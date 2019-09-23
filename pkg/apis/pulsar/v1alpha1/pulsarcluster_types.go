package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	// Zookeeper Pod Type
	ZookeeperPodType = "zookeeper"

	// Broker Pod Type
	BrokerPodType = "broker"

	// Bookie Pod Type
	BookiePodType = "bookie"
)

const (
	// DefaultZkContainerRepository is the default docker repo for the container
	DefaultContainerRepository = "apachepulsar/pulsar-all"

	// DefaultZkContainerVersion is the default tag used for for the container
	DefaultContainerVersion = "latest"

	// DefaultZkContainerPolicy is the default container pull policy used
	DefaultContainerPolicy = "Always"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// PulsarClusterSpec defines the desired state of PulsarCluster
// +k8s:openapi-gen=true
type PulsarClusterSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html

	// ZookeeperSpec defines the desired state of Zookeeper
	ZookeeperSpec ZookeeperSpec `json:"zookeeper,omitempty"`

	// BookieSpec defines the desired state of Bookie
	BookieSpec BookieSpec `json:"bookie,omitempty"`

	// BrokerSpec defines the desired state of Broker
	BrokerSpec BrokerSpec `json:"broker,omitempty"`
}

func (s *PulsarClusterSpec) SetDefault(cluster *PulsarCluster) bool {
	zookeeperChanged := s.ZookeeperSpec.SetDefault(cluster)

	bookieChanged := s.BookieSpec.SetDefault(cluster)

	brokerChanged := s.BrokerSpec.SetDefault(cluster)

	return zookeeperChanged || bookieChanged || brokerChanged
}

// PulsarClusterStatus defines the observed state of PulsarCluster
// +k8s:openapi-gen=true
type PulsarClusterStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html

	// Pulsar Cluster Phase
	Phase string `json:"phase,omitempty"`
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

func (c *PulsarCluster) SetDefault() bool {
	return c.Spec.SetDefault(c)
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
