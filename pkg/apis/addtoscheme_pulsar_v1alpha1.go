package apis

import (
	"github.com/sky-big/pulsar-operator/pkg/apis/pulsar/v1alpha1"
)

func init() {
	// Register the types with the Scheme so the resources can map objects to GroupVersionKinds and back
	AddToSchemes = append(AddToSchemes, v1alpha1.SchemeBuilder.AddToScheme)
}
