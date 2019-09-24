package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
)

// PodPolicy defines the common pod configuration for Pods, including when used
// in deployments, stateful-sets, etc.
// +k8s:openapi-gen=true
type PodPolicy struct {
	// Labels specifies the labels to attach to pods the operator creates for
	// the pulsar cluster.
	Labels map[string]string `json:"labels,omitempty"`

	// NodeSelector specifies a map of key-value pairs. For the pod to be
	// eligible to run on a node, the node must have each of the indicated
	// key-value pairs as labels.
	NodeSelector map[string]string `json:"nodeSelector,omitempty"`

	// The scheduling constraints on pods.
	Affinity *corev1.Affinity `json:"affinity,omitempty"`

	// Resources is the resource requirements for the container.
	// This field cannot be updated once the cluster is created.
	Resources corev1.ResourceRequirements `json:"resources,omitempty"`

	// Tolerations specifies the pod's tolerations.
	Tolerations []corev1.Toleration `json:"tolerations,omitempty"`

	// List of environment variables to set in the container.
	// This field cannot be updated.
	Env []corev1.EnvVar `json:"env,omitempty"`

	// Annotations specifies the annotations to attach to pods the operator
	// creates.
	Annotations map[string]string `json:"annotations,omitempty"`

	// SecurityContext specifies the security context for the entire pod
	// More info: https://kubernetes.io/docs/tasks/configure-pod-container/security-context
	SecurityContext *corev1.PodSecurityContext `json:"securityContext,omitempty"`

	// TerminationGracePeriodSeconds is the amount of time that kubernetes will
	// give for a pod instance to shutdown normally.
	// The default value is 1800.
	TerminationGracePeriodSeconds int64 `json:"terminationGracePeriodSeconds"`
}

func (p *PodPolicy) SetDefault(cluster *PulsarCluster, component string) bool {
	changed := false

	if p.Labels == nil {
		p.Labels = make(map[string]string)
		changed = true
	}

	if p.NodeSelector == nil {
		p.NodeSelector = make(map[string]string)
		changed = true
	}

	if p.Tolerations == nil {
		p.Tolerations = make([]corev1.Toleration, 0)
		changed = true
	}

	if p.Env == nil {
		p.Env = make([]corev1.EnvVar, 0)
		changed = true
	}

	if p.Annotations == nil {
		p.Annotations = make(map[string]string)
		changed = true
	}
	return changed
}
