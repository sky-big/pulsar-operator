package pulsarcluster

import (
	pulsarv1alpha1 "github.com/sky-big/pulsar-operator/pkg/apis/pulsar/v1alpha1"
)

func (r *ReconcilePulsarCluster) reconcileBookie(c *pulsarv1alpha1.PulsarCluster) error {
	if c.Status.Phase == pulsarv1alpha1.PulsarClusterInitingPhase {
		return nil
	}
	return nil
}