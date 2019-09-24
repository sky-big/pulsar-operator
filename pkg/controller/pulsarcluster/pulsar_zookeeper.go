package pulsarcluster

import (
	pulsarv1alpha1 "github.com/sky-big/pulsar-operator/pkg/apis/pulsar/v1alpha1"
)

// Reconcile For All Resource About Zookeeper
func (r *ReconcilePulsarCluster) reconcileZookeeper(c *pulsarv1alpha1.PulsarCluster) error {
	for _, fun := range []reconcileFunc{
		r.reconcileZookeeperConfigMap,
		r.reconcileZookeeperStatefulSet,
		r.reconcileZookeeperService,
	} {
		if err := fun(c); err != nil {
			return err
		}
	}
	return nil
}

func (r *ReconcilePulsarCluster) reconcileZookeeperConfigMap(c *pulsarv1alpha1.PulsarCluster) error {
	return nil
}

func (r *ReconcilePulsarCluster) reconcileZookeeperStatefulSet(c *pulsarv1alpha1.PulsarCluster) error {
	return nil
}

func (r *ReconcilePulsarCluster) reconcileZookeeperService(c *pulsarv1alpha1.PulsarCluster) error {
	return nil
}