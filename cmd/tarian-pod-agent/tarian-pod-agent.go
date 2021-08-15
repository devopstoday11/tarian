package main

import (
	"bufio"
	"crypto/tls"
	"crypto/x509"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/devopstoday11/tarian/pkg/logger"
	"github.com/devopstoday11/tarian/pkg/podagent"
	"github.com/devopstoday11/tarian/pkg/tarianpb"
	cli "github.com/urfave/cli/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	defaultClusterAgentPort = "50052"
	defaultClusterAgentHost = ""
)

// nolint: gochecknoglobals
var (
	version = "dev"
	commit  = "main"
)

func main() {
	app := getCliApp()
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func getCliApp() *cli.App {
	return &cli.App{
		Name:    "Tarian Pod Agent",
		Usage:   "The Tarian pod agent is the component which runs as a sidecar to monitor your main container.",
		Version: version + " (" + commit + ")",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "log-level",
				Usage: "Log level: debug, info, warn, error",
				Value: "info",
			},
			&cli.StringFlag{
				Name:  "log-encoding",
				Usage: "log-encoding: json, console",
				Value: "console",
			},
		},
		Action: func(ctx *cli.Context) error {
			return ctx.App.Command("run").Run(ctx)
		},
		Commands: []*cli.Command{
			{
				Name:  "run",
				Usage: "Run the pod agent",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "host",
						Usage: "Host address of the cluster agent to communicate with",
						Value: defaultClusterAgentHost,
					},
					&cli.StringFlag{
						Name:  "port",
						Usage: "Host port of the cluster agent to communicate with",
						Value: defaultClusterAgentPort,
					},
					&cli.BoolFlag{
						Name:  "tls-enabled",
						Usage: "If enabled, it will communicate with the cluster agent using TLS",
						Value: false,
					},
					&cli.BoolFlag{
						Name:  "tls-insecure-skip-verify",
						Usage: "If set to true, it will skip cluster agent's certificate chain and hostname verification",
						Value: true,
					},
					&cli.StringFlag{
						Name:  "pod-labels-file",
						Usage: "File path containing pod labels. This is intended to be a file from Kubernetes DownwardAPIVolumeFile",
						Value: "",
					},
					&cli.StringFlag{
						Name:  "pod-name",
						Usage: "Pod name where it is running. This is intended to be set from Downward API",
						Value: "",
					},
					&cli.StringFlag{
						Name:  "pod-uid",
						Usage: "Pod UID where it is running. This is intended to be set from Downward API",
						Value: "",
					},
					&cli.StringFlag{
						Name:  "namespace",
						Usage: "Kubernetes namespace where it is running",
						Value: "tarian-system",
					},
				},
				Action: run,
			},
		},
	}
}

func run(c *cli.Context) error {
	logger := logger.GetLogger(c.String("log-level"), c.String("log-encoding"))
	podagent.SetLogger(logger)
	logger.Infow("tarian-pod-agent is running")

	dialOpts := []grpc.DialOption{}
	if c.Bool("tls-enabled") {
		// TODO: handle err
		certPool, _ := x509.SystemCertPool()
		tlsConfig := &tls.Config{ServerName: "", RootCAs: certPool}

		tlsConfig.InsecureSkipVerify = c.Bool("tls-insecure-skip-verify")
		dialOpts = append(dialOpts, grpc.WithTransportCredentials(credentials.NewTLS(tlsConfig)))
	} else {
		dialOpts = append(dialOpts, grpc.WithInsecure())
	}

	agent := podagent.NewPodAgent(c.String("host")+":"+c.String("port"), dialOpts)

	podLabelsFile := c.String("pod-labels-file")
	if podLabelsFile != "" {
		podLabels, err := readLabelsFromFile(podLabelsFile)

		if err != nil {
			logger.Errorw("failed reading pod-labels-file", "err", err)
		}

		agent.SetPodLabels(podLabels)
	}

	podName := c.String("pod-name")
	if podName != "" {
		agent.SetPodName(podName)
	}

	podUid := c.String("pod-uid")
	if podUid != "" {
		agent.SetPodUid(podUid)
	}

	namespace := c.String("namespace")
	if namespace != "" {
		agent.SetNamespace(namespace)
	}

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

	go func() {
		sig := <-sigCh
		logger.Infow("got sigterm signal, attempting graceful shutdown", "signal", sig)

		agent.GracefulStop()
	}()

	agent.Run()
	logger.Info("tarian-pod-agent shutdown gracefully")

	return nil
}

func readLabelsFromFile(path string) ([]*tarianpb.Label, error) {
	labels := []*tarianpb.Label{}

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		idx := strings.Index(line, "=")

		if idx < 0 {
			continue
		}

		key := line[:idx]
		value := strings.Trim(line[idx+1:], "\"")

		labels = append(labels, &tarianpb.Label{Key: key, Value: value})
	}

	return labels, nil
}
