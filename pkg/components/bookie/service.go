package bookie

import (
	"fmt"

	pulsarv1alpha1 "github.com/sky-big/pulsar-operator/pkg/apis/pulsar/v1alpha1"
)

func MakeBookieService(c *pulsarv1alpha1.PulsarCluster) string {
	return fmt.Sprintf("%s-bookie-service", c.GetName())
}
