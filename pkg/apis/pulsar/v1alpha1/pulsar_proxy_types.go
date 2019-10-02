package v1alpha1

// Proxy defines the desired state of Proxy
// +k8s:openapi-gen=true
type Proxy struct {
	// Image is the  container image. default is apachepulsar/pulsar:latest
	Image ContainerImage `json:"image,omitempty"`

	// Labels specifies the labels to attach to pods the operator creates for
	// the proxy cluster.
	Labels map[string]string `json:"labels,omitempty"`

	// Size (DEPRECATED) is the expected size of the proxy cluster.
	Size int32 `json:"size,omitempty"`

	// Pod defines the policy to create pod for the proxy cluster.
	//
	// Updating the pod does not take effect on any existing pods.
	Pod PodPolicy `json:"pod,omitempty"`
}

func (b *Proxy) SetDefault(cluster *PulsarCluster) bool {
	changed := false

	if b.Image.SetDefault(cluster, ProxyComponent) {
		changed = true
	}

	if b.Size == 0 {
		b.Size = ProxyClusterDefaultNodeNum
		changed = true
	}

	if b.Pod.SetDefault(cluster, ProxyComponent) {
		changed = true
	}
	return changed
}
