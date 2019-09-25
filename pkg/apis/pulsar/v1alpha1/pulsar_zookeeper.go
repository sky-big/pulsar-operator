package v1alpha1

const (
	// zookeeper cluster default num is 3
	ZookeeperClusterDefaultNodeNum = 3
)

// Zookeeper defines the desired state of Zookeeper
// +k8s:openapi-gen=true
type Zookeeper struct {
	// Image is the  container image. default is apachepulsar/pulsar-all:latest
	Image ContainerImage `json:"image"`

	// Labels specifies the labels to attach to pods the operator creates for
	// the zookeeper cluster.
	Labels map[string]string `json:"labels,omitempty"`

	// Size (DEPRECATED) is the expected size of the zookeeper cluster. This
	// has been replaced with "Replicas"
	//
	// The valid range of size is from 1 to 7.
	Size int32 `json:"size"`

	// Pod defines the policy to create pod for the zookeeper cluster.
	//
	// Updating the Pod does not take effect on any existing pods.
	Pod PodPolicy `json:"pod,omitempty"`
}

func (s *Zookeeper) SetDefault(cluster *PulsarCluster) bool {
	changed := false

	if s.Image.SetDefault(cluster, ZookeeperComponent) {
		changed = true
	}

	if s.Labels == nil {
		s.Labels = make(map[string]string)
		changed = true
	}

	if s.Size == 0 {
		s.Size = ZookeeperClusterDefaultNodeNum
		changed = true
	}

	if s.Pod.SetDefault(cluster, ZookeeperComponent) {
		changed = true
	}

	return changed
}
