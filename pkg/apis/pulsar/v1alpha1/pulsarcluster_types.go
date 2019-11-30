package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// PulsarClusterSpec defines the desired state of PulsarCluster
// +k8s:openapi-gen=true
type PulsarClusterSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html

	// Zookeeper defines the desired state of Zookeeper
	Zookeeper Zookeeper `json:"zookeeper,omitempty"`

	// Bookie defines the desired state of Bookie
	Bookie Bookie `json:"bookie,omitempty"`

	// Broker defines the desired state of Broker
	Broker Broker `json:"broker,omitempty"`

	// Proxy defines the desired state of Proxy
	Proxy Proxy `json:"proxy,omitempty"`

	// Monitor defines the desired state of Monitor
	Monitor Monitor `json:"monitor,omitempty"`

	// Manager defines the desired state of Manager
	Manager Manager `json:"manager,omitempty"`
}

func (s *PulsarClusterSpec) SetDefault(cluster *PulsarCluster) bool {
	changed := false

	if s.Zookeeper.SetDefault(cluster) {
		changed = true
	}

	if s.Bookie.SetDefault(cluster) {
		changed = true
	}

	if s.Broker.SetDefault(cluster) {
		changed = true
	}

	if s.Proxy.SetDefault(cluster) {
		changed = true
	}

	if s.Monitor.SetDefault(cluster) {
		changed = true
	}

	if s.Manager.SetDefault(cluster) {
		changed = true
	}
	return changed
}

// PulsarClusterStatus defines the observed state of PulsarCluster
// +k8s:openapi-gen=true
type PulsarClusterStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html

	// Pulsar cluster phase
	Phase string `json:"phase,omitempty"`
}

func (s *PulsarClusterStatus) SetDefault(cluster *PulsarCluster) bool {
	changed := false

	if s.Phase == "" {
		s.Phase = PulsarClusterInitingPhase
		changed = true
	}
	return changed
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// PulsarCluster is the Schema for the pulsarclusters API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
type PulsarCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PulsarClusterSpec   `json:"spec,omitempty"`
	Status PulsarClusterStatus `json:"status,omitempty"`
}

func (c *PulsarCluster) SpecSetDefault() bool {
	return c.Spec.SetDefault(c)
}

func (c *PulsarCluster) StatusSetDefault() bool {
	return c.Status.SetDefault(c)
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// PulsarClusterList contains a list of PulsarCluster
type PulsarClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PulsarCluster `json:"items"`
}

func init() {
	SchemeBuilder.Register(&PulsarCluster{}, &PulsarClusterList{})
}
