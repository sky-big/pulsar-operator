package v1alpha1

// Manager defines the desired state of Manager
// +k8s:openapi-gen=true
type Manager struct {
	// Is enable pulsar cluster manager flag.
	Enable bool `json:"enable,omitempty"`

	// Image is the  container image. default is apachepulsar/pulsar-all:latest
	Image ContainerImage `json:"image,omitempty"`

	// Labels specifies the labels to attach to pods the operator creates for
	// the broker cluster.
	Labels map[string]string `json:"labels,omitempty"`

	// User Name
	UserName string `json:"userName,omitempty"`

	// User Password
	UserPassword string `json:"userPassword,omitempty"`

	// Host (DEPRECATED) is the expected host of the pulsar manager.
	Host string `json:"host,omitempty"`

	// Ingress additional annotation
	Annotations map[string]string `json:"annotations,omitempty"`

	// NodePort (DEPRECATED) is the expected port of the pulsar manager.
	NodePort int32 `json:"nodePort,omitempty"`
}

func (m *Manager) SetDefault(cluster *PulsarCluster) bool {
	changed := false

	if m.Image.SetDefault(cluster, ManagerComponent) {
		changed = true
	}

	if m.UserName == "" {
		m.UserName = ManagerDefaultUserName
		changed = true
	}

	if m.UserPassword == "" {
		m.UserPassword = ManagerDefaultUserPassword
		changed = true
	}
	return changed
}
