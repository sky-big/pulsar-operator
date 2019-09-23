package v1alpha1

// BookieSpec defines the desired state of Bookie
type BookieSpec struct {
	// Image is the  container image. default is apachepulsar/pulsar-all:latest
	Image ContainerImage `json:"image"`

	// Labels specifies the labels to attach to pods the operator creates for
	// the bookie cluster.
	Labels map[string]string `json:"labels,omitempty"`

	// Size (DEPRECATED) is the expected size of the bookie cluster. This
	// has been replaced with "Replicas"
	//
	// The valid range of size is from 1 to 7.
	Size int32 `json:"size"`

	// Pod defines the policy to create pod for the bookie cluster.
	//
	// Updating the Pod does not take effect on any existing pods.
	Pod PodPolicy `json:"pod,omitempty"`
}

func (b *BookieSpec) SetDefault(cluster *PulsarCluster) bool {
	return false
}
