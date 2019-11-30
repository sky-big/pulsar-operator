package v1alpha1

import (
	"fmt"

	corev1 "k8s.io/api/core/v1"
)

// ContainerImage defines the fields needed for a Docker repository image. The
// format here matches the predominant format used in Helm charts.
// +k8s:openapi-gen=true
type ContainerImage struct {
	Repository string            `json:"repository,omitempty"`
	Tag        string            `json:"tag,omitempty"`
	PullPolicy corev1.PullPolicy `json:"pullPolicy,omitempty"`
}

func (c *ContainerImage) SetDefault(cluster *PulsarCluster, component string) bool {
	changed := false
	switch component {
	case ZookeeperComponent:
		if cluster.Spec.Zookeeper.Image.Repository == "" {
			cluster.Spec.Zookeeper.Image.Repository = DefaultAllPulsarContainerRepository
			changed = true
		}
		if cluster.Spec.Zookeeper.Image.Tag == "" {
			cluster.Spec.Zookeeper.Image.Tag = DefaultAllPulsarContainerVersion
			changed = true
		}
		if cluster.Spec.Zookeeper.Image.PullPolicy == "" {
			cluster.Spec.Zookeeper.Image.PullPolicy = DefaultContainerPolicy
			changed = true
		}

	case BrokerComponent:
		if cluster.Spec.Broker.Image.Repository == "" {
			cluster.Spec.Broker.Image.Repository = DefaultAllPulsarContainerRepository
			changed = true
		}
		if cluster.Spec.Broker.Image.Tag == "" {
			cluster.Spec.Broker.Image.Tag = DefaultAllPulsarContainerVersion
			changed = true
		}
		if cluster.Spec.Broker.Image.PullPolicy == "" {
			cluster.Spec.Broker.Image.PullPolicy = DefaultContainerPolicy
			changed = true
		}

	case BookieComponent:
		if cluster.Spec.Bookie.Image.Repository == "" {
			cluster.Spec.Bookie.Image.Repository = DefaultAllPulsarContainerRepository
			changed = true
		}
		if cluster.Spec.Bookie.Image.Tag == "" {
			cluster.Spec.Bookie.Image.Tag = DefaultAllPulsarContainerVersion
			changed = true
		}
		if cluster.Spec.Bookie.Image.PullPolicy == "" {
			cluster.Spec.Bookie.Image.PullPolicy = DefaultContainerPolicy
			changed = true
		}

	case ProxyComponent:
		if cluster.Spec.Proxy.Image.Repository == "" {
			cluster.Spec.Proxy.Image.Repository = DefaultPulsarContainerRepository
			changed = true
		}
		if cluster.Spec.Proxy.Image.Tag == "" {
			cluster.Spec.Proxy.Image.Tag = DefaultPulsarContainerVersion
			changed = true
		}
		if cluster.Spec.Proxy.Image.PullPolicy == "" {
			cluster.Spec.Proxy.Image.PullPolicy = DefaultContainerPolicy
			changed = true
		}

	case ManagerComponent:
		if cluster.Spec.Manager.Image.Repository == "" {
			cluster.Spec.Manager.Image.Repository = DefaultPulsarManagerContainerRepository
			changed = true
		}
		if cluster.Spec.Manager.Image.Tag == "" {
			cluster.Spec.Manager.Image.Tag = DefaultPulsarManagerContainerVersion
			changed = true
		}
		if cluster.Spec.Manager.Image.PullPolicy == "" {
			cluster.Spec.Manager.Image.PullPolicy = DefaultContainerPolicy
			changed = true
		}

	}
	return changed
}

func (c ContainerImage) GenerateImage() string {
	return fmt.Sprintf("%s:%s", c.Repository, c.Tag)
}
