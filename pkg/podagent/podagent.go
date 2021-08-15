package podagent

import (
	"context"
	"crypto/sha256"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"regexp"
	"sync"
	"time"

	"github.com/devopstoday11/tarian/pkg/tarianpb"
	psutil "github.com/shirou/gopsutil/process"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var logger *zap.SugaredLogger

func init() {
	l, err := zap.NewProduction()

	if err != nil {
		panic("Can not create logger")
	}

	logger = l.Sugar()
}

func SetLogger(l *zap.SugaredLogger) {
	logger = l
}

type PodAgent struct {
	clusterAgentAddress string
	grpcConn            *grpc.ClientConn
	grpcDialOpts        []grpc.DialOption
	configClient        tarianpb.ConfigClient
	eventClient         tarianpb.EventClient
	podName             string
	podUid              string
	podLabels           []*tarianpb.Label
	namespace           string

	constraints     []*tarianpb.Constraint
	constraintsLock sync.RWMutex

	cancelFunc context.CancelFunc
	cancelCtx  context.Context
}

func NewPodAgent(clusterAgentAddress string, grpcDialOpts []grpc.DialOption) *PodAgent {
	ctx, cancel := context.WithCancel(context.Background())

	return &PodAgent{
		clusterAgentAddress: clusterAgentAddress,
		cancelCtx:           ctx,
		cancelFunc:          cancel,
		grpcDialOpts:        grpcDialOpts,
	}
}

func (p *PodAgent) SetPodLabels(labels []*tarianpb.Label) {
	p.podLabels = labels
}

func (p *PodAgent) SetPodName(name string) {
	p.podName = name
}

func (p *PodAgent) SetPodUid(uid string) {
	p.podUid = uid
}

func (p *PodAgent) SetNamespace(namespace string) {
	p.namespace = namespace
}

func (p *PodAgent) Dial() {
	var err error
	p.grpcConn, err = grpc.Dial(p.clusterAgentAddress, p.grpcDialOpts...)
	p.configClient = tarianpb.NewConfigClient(p.grpcConn)
	p.eventClient = tarianpb.NewEventClient(p.grpcConn)

	if err != nil {
		logger.Fatalw("couldn't connect to the cluster agent", "err", err)
	}
}

func (p *PodAgent) GracefulStop() {
	p.cancelFunc()

	if p.grpcConn != nil {
		p.grpcConn.Close()
	}
}

func (p *PodAgent) Run() {
	p.Dial()
	defer p.grpcConn.Close()

	wg := sync.WaitGroup{}
	wg.Add(3)

	go func() {
		p.loopSyncConstraints(p.cancelCtx)
		wg.Done()
	}()

	go func() {
		p.loopValidateProcesses(p.cancelCtx)
		wg.Done()
	}()

	go func() {
		p.loopValidateFileChecksums(p.cancelCtx)
		wg.Done()
	}()

	wg.Wait()
}

func (p *PodAgent) SetConstraints(constraints []*tarianpb.Constraint) {
	p.constraintsLock.Lock()
	defer p.constraintsLock.Unlock()

	p.constraints = constraints
}

func (p *PodAgent) GetConstraints() []*tarianpb.Constraint {
	return p.constraints
}

func (p *PodAgent) loopSyncConstraints(ctx context.Context) error {
	for {
		p.SyncConstraints()

		select {
		case <-time.After(3 * time.Second):
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

func (p *PodAgent) SyncConstraints() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	token, _ := ioutil.ReadFile("/run/secrets/kubernetes.io/serviceaccount/token")
	ctx = metadata.AppendToOutgoingContext(ctx, "Authorization", "Bearer "+string(token))
	logger.Infow(">>> using token", "token", string(token))

	r, err := p.configClient.GetConstraints(ctx, &tarianpb.GetConstraintsRequest{Namespace: p.namespace, Labels: p.podLabels})

	if err != nil {
		logger.Errorw("error while getting constraints from the cluster agent", "err", err)
	}

	logger.Debugw("received constraints from the cluster agent", "constraint", r.GetConstraints())
	cancel()

	p.SetConstraints(r.GetConstraints())
}

func (p *PodAgent) loopValidateProcesses(ctx context.Context) error {
	for {
		ps, _ := psutil.Processes()
		processes := NewProcessesFromPsutil(ps)

		violations := p.ValidateProcesses(processes)

		for _, violation := range violations {
			name := violation.GetName()

			logger.Debugw("found process that violates regex", "process", name)
		}

		if len(violations) > 0 {
			p.ReportViolationsToClusterAgent(violations)
		}

		select {
		case <-time.After(3 * time.Second):
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

func (p *PodAgent) loopValidateFileChecksums(ctx context.Context) error {
	for {
		violatedFiles := p.ValidateFileChecksums()

		for _, violation := range violatedFiles {
			logger.Debugw("found a file that violates checksum", "file", violation.name, "actual", violation.actualSha256Sum, "expected", violation.expectedSha256Sum)
		}

		if len(violatedFiles) > 0 {
			p.ReportViolatedFilesToClusterAgent(violatedFiles)
		}

		select {
		case <-time.After(3 * time.Second):
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

func (p *PodAgent) ReportViolationsToClusterAgent(processes map[int32]*Process) {
	violatingProcesses := make([]*tarianpb.Process, len(processes))

	i := 0
	for _, p := range processes {
		violatingProcesses[i] = &tarianpb.Process{Pid: p.GetPid(), Name: p.GetName()}
		i++
	}

	req := &tarianpb.IngestEventRequest{
		Event: &tarianpb.Event{
			Type:            tarianpb.EventTypeViolation,
			ClientTimestamp: timestamppb.Now(),
			Targets: []*tarianpb.Target{
				{
					Pod: &tarianpb.Pod{
						Uid:       p.podUid,
						Name:      p.podName,
						Namespace: p.namespace,
						Labels:    p.podLabels,
					},
					ViolatingProcesses: violatingProcesses,
				},
			},
		},
	}

	response, err := p.eventClient.IngestEvent(context.Background(), req)

	if err != nil {
		logger.Errorw("error while reporting violation events", "err", err)
	} else {
		logger.Debugw("ingest event response", "response", response)
	}
}

func (p *PodAgent) ValidateProcesses(processes []*Process) map[int32]*Process {
	p.constraintsLock.RLock()

	// map[pid]*process
	violations := make(map[int32]*Process)
	allowedProcesses := make(map[int32]*Process)

	for _, constraint := range p.constraints {
		if constraint.GetAllowedProcesses() == nil {
			continue
		}

		for _, allowedProcess := range constraint.GetAllowedProcesses() {
			if allowedProcess.GetRegex() == "" {
				continue
			}

			rgx, err := regexp.Compile(allowedProcess.GetRegex())

			if err != nil {
				logger.Errorw("can not compile regex", "err", err)
				continue
			}

			logger.Debugw("looking for running processes that violate regex", "expr", rgx.String())

			for _, process := range processes {
				name := process.GetName()

				if err != nil {
					logger.Errorw("can not read process name", "err", err)
					continue
				}

				if !rgx.MatchString(name) {
					violations[process.GetPid()] = process
				} else {
					allowedProcesses[process.GetPid()] = process
				}
			}
		}
	}

	p.constraintsLock.RUnlock()

	for pid := range allowedProcesses {
		delete(violations, pid)
	}

	return violations
}

type violatedFile struct {
	name              string
	expectedSha256Sum string
	actualSha256Sum   string
}

func (p *PodAgent) ValidateFileChecksums() map[string]*violatedFile {
	p.constraintsLock.RLock()

	// Copy constraints to a local var to not block SyncConstraints() because this function can run quite long
	constraints := make([]*tarianpb.Constraint, len(p.constraints))
	copy(constraints, p.constraints)

	p.constraintsLock.RUnlock()

	violatedFiles := make(map[string]*violatedFile)
	allowedFiles := make(map[string]struct{})

	for _, constraint := range constraints {
		if constraint.GetAllowedFiles() == nil {
			continue
		}

		for _, allowedFile := range constraint.GetAllowedFiles() {
			if allowedFile.GetName() == "" || allowedFile.GetSha256Sum() == "" {
				continue
			}

			logger.Debugw("validating file sha256 checksum", "file", allowedFile.GetName(), "allowedSha256Sum", allowedFile.GetSha256Sum())

			f, err := os.Open(allowedFile.GetName())
			if err != nil {
				logger.Errorw("can not open file to check the sha256 checksum", "file", allowedFile.GetName(), "err", err)
			}

			s256 := sha256.New()
			if _, err := io.Copy(s256, f); err != nil {
				logger.Errorw("can not read file to check the sha256 checksum", "file", allowedFile.GetName(), "err", err)
			}

			actualSha256Sum := fmt.Sprintf("%x", s256.Sum(nil))

			if actualSha256Sum == allowedFile.GetSha256Sum() {
				allowedFiles[allowedFile.GetName()] = struct{}{}
			} else {
				violated := &violatedFile{
					name:              allowedFile.GetName(),
					actualSha256Sum:   actualSha256Sum,
					expectedSha256Sum: allowedFile.GetSha256Sum(),
				}

				violatedFiles[allowedFile.GetName()] = violated
			}

			f.Close()
		}
	}

	for name := range allowedFiles {
		delete(violatedFiles, name)
	}

	return violatedFiles
}

func (p *PodAgent) ReportViolatedFilesToClusterAgent(violatedFiles map[string]*violatedFile) {
	vf := make([]*tarianpb.ViolatedFile, len(violatedFiles))

	i := 0
	for _, f := range violatedFiles {
		vf[i] = &tarianpb.ViolatedFile{Name: f.name, ActualSha256Sum: f.actualSha256Sum, ExpectedSha256Sum: f.expectedSha256Sum}
		i++
	}

	req := &tarianpb.IngestEventRequest{
		Event: &tarianpb.Event{
			Type:            tarianpb.EventTypeViolation,
			ClientTimestamp: timestamppb.Now(),
			Targets: []*tarianpb.Target{
				{
					Pod: &tarianpb.Pod{
						Uid:       p.podUid,
						Name:      p.podName,
						Namespace: p.namespace,
						Labels:    p.podLabels,
					},
					ViolatedFiles: vf,
				},
			},
		},
	}

	response, err := p.eventClient.IngestEvent(context.Background(), req)

	if err != nil {
		logger.Errorw("error while reporting violation events", "err", err)
	} else {
		logger.Debugw("ingest event response", "response", response)
	}
}
