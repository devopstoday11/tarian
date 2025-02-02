package webhookserver

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"go.uber.org/zap"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/utils/pointer"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// +kubebuilder:webhook:path=/inject-pod-agent,mutating=true,sideEffects=none,failurePolicy=ignore,groups="",resources=pods,verbs=create,versions=v1,admissionReviewVersions=v1,name=pod-agent.k8s.tarian.dev

type PodAgentInjector struct {
	Client  client.Client
	decoder *admission.Decoder
	config  PodAgentContainerConfig
	logger  *zap.SugaredLogger
}

type PodAgentContainerConfig struct {
	Name        string
	Image       string
	LogEncoding string
	Host        string
	Port        string
}

const (
	ThreatScanAnnotation              = "pod-agent.k8s.tarian.dev/threat-scan"
	RegisterAnnotation                = "pod-agent.k8s.tarian.dev/register"
	FileValidationIntervalAnnotation  = "pod-agent.k8s.tarian.dev/file-validation-interval"
	RegisterFileIgnorePathsAnnotation = "pod-agent.k8s.tarian.dev/register-file-ignore-paths"
)

// podAnnotator adds an annotation to every incoming pods.
func (p *PodAgentInjector) Handle(ctx context.Context, req admission.Request) admission.Response {
	p.logger.Debugw("handling a webhook request")

	pod := &corev1.Pod{}

	err := p.decoder.Decode(req, pod)
	if err != nil {
		p.logger.Errorw("error while decoding webhook request payload", "err", err)
		return admission.Errored(http.StatusBadRequest, err)
	}

	if pod.Annotations == nil {
		p.logger.Debugw("not injecting container because no annotation is found", "pod_name", pod.GetObjectMeta().GetName())
		return admission.Allowed("no annotation found")
	}

	_, threatScanAnnotationPresent := pod.Annotations[ThreatScanAnnotation]
	registerAnnotationValue, registerAnnotationPresent := pod.Annotations[RegisterAnnotation]
	registerFileIgnorePathsAnnotationValue, registerFileIgnorePathsAnnotationPresent := pod.Annotations[RegisterFileIgnorePathsAnnotation]

	if !threatScanAnnotationPresent && !registerAnnotationPresent {
		p.logger.Debugw("not injecting container because no tarian annotation is found", "pod_name", pod.GetObjectMeta().GetName())
		return admission.Allowed("annotation " + ThreatScanAnnotation + " or " + RegisterAnnotation + " not found")
	}

	for _, c := range pod.Spec.Containers {
		if c.Name == p.config.Name {
			return admission.Allowed("container with name " + p.config.Name + " already exists")
		}
	}

	podInfoPath := "/etc/podinfo"
	// mount all volumes into pod agent
	volumeMounts := []corev1.VolumeMount{{Name: "podinfo", MountPath: podInfoPath}}
	mountNamesAdded := make(map[string]struct{})
	for _, c := range pod.Spec.Containers {
		for _, vm := range c.VolumeMounts {
			if _, found := mountNamesAdded[vm.Name]; found {
				continue
			}

			volumeMounts = append(volumeMounts, vm)
			mountNamesAdded[vm.Name] = struct{}{}
		}
	}

	podAgentCommand := "threat-scan"
	if registerAnnotationPresent {
		podAgentCommand = "register"
	}

	podAgentArgs := []string{
		"--log-encoding=" + p.config.LogEncoding,
		podAgentCommand,
		"--host=" + p.config.Host,
		"--port=" + p.config.Port,
		"--namespace=$(NAMESPACE)",
		"--pod-name=$(POD_NAME)",
		"--pod-uid=$(POD_UID)",
		"--pod-labels-file=/etc/podinfo/labels",
	}

	if registerAnnotationPresent {
		podAgentArgs = append(podAgentArgs, "--register-rules="+registerAnnotationValue)

		mountPathsAdded := make(map[string]struct{})
		mountPaths := []string{}
		for _, vm := range volumeMounts {
			if _, found := mountPathsAdded[vm.MountPath]; found {
				continue
			}

			mountNamesAdded[vm.MountPath] = struct{}{}

			if vm.MountPath != podInfoPath && vm.MountPath != "/var/run/secrets/kubernetes.io/serviceaccount" {
				mountPaths = append(mountPaths, vm.MountPath)
			}
		}

		podAgentArgs = append(podAgentArgs, "--register-file-paths="+strings.Join(mountPaths, ","))
	}

	if registerFileIgnorePathsAnnotationPresent {
		podAgentArgs = append(podAgentArgs, "--register-file-ignore-paths="+registerFileIgnorePathsAnnotationValue)
	}

	if fileValidationInterval, ok := pod.Annotations[FileValidationIntervalAnnotation]; ok {
		podAgentArgs = append(podAgentArgs, "--file-validation-interval="+fileValidationInterval)
	}

	podAgentContainer := corev1.Container{
		Name:  p.config.Name,
		Image: p.config.Image,
		Env: []corev1.EnvVar{
			{
				Name: "NAMESPACE", ValueFrom: &corev1.EnvVarSource{FieldRef: &corev1.ObjectFieldSelector{FieldPath: "metadata.namespace"}},
			},
			{
				Name: "POD_NAME", ValueFrom: &corev1.EnvVarSource{FieldRef: &corev1.ObjectFieldSelector{FieldPath: "metadata.name"}},
			},
			{
				Name: "POD_UID", ValueFrom: &corev1.EnvVarSource{FieldRef: &corev1.ObjectFieldSelector{FieldPath: "metadata.uid"}},
			},
		},
		Args:         podAgentArgs,
		VolumeMounts: volumeMounts,
	}
	pod.Spec.Containers = append(pod.Spec.Containers, podAgentContainer)
	pod.Spec.ShareProcessNamespace = pointer.BoolPtr(true)

	podInfoVolume := corev1.Volume{
		Name: "podinfo",
		VolumeSource: corev1.VolumeSource{
			DownwardAPI: &corev1.DownwardAPIVolumeSource{
				Items: []corev1.DownwardAPIVolumeFile{{Path: "labels", FieldRef: &corev1.ObjectFieldSelector{FieldPath: "metadata.labels"}}},
			},
		},
	}
	pod.Spec.Volumes = append(pod.Spec.Volumes, podInfoVolume)

	marshaledPod, err := json.Marshal(pod)
	if err != nil {
		p.logger.Errorw("error while marshalling pod into json", "pod_name", pod.GetObjectMeta().GetName(), "err", err)
		return admission.Errored(http.StatusInternalServerError, err)
	}

	p.logger.Infow("responding webhook with a sidecar container", "pod_name", pod.GetObjectMeta().GetName())
	return admission.PatchResponseFromRaw(req.Object.Raw, marshaledPod)
}

// InjectDecoder injects the decoder.
func (p *PodAgentInjector) InjectDecoder(d *admission.Decoder) error {
	p.decoder = d
	return nil
}
