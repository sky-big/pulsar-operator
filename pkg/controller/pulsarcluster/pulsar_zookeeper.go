package pulsarcluster

import (
	"context"

	pulsarv1alpha1 "github.com/sky-big/pulsar-operator/pkg/apis/pulsar/v1alpha1"
	"github.com/sky-big/pulsar-operator/pkg/resource/zookeeper"

	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/api/core/v1"
	"k8s.io/api/policy/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"

	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
)

// Reconcile For All Resource About Zookeeper
func (r *ReconcilePulsarCluster) reconcileZookeeper(c *pulsarv1alpha1.PulsarCluster) error {
	r.log.Info("[Start] Reconciling PulsarCluster Zookeeper")
	for _, fun := range []reconcileFunc{
		r.reconcileZookeeperConfigMap,
		r.reconcileZookeeperStatefulSet,
		r.reconcileZookeeperService,
		r.reconcileZookeeperPodDisruptionBudget,
	} {
		if err := fun(c); err != nil {
			r.log.Error(err, "Reconciling PulsarCluster Zookeeper Error", c)
			return err
		}
	}
	return nil
}

func (r *ReconcilePulsarCluster) reconcileZookeeperConfigMap(c *pulsarv1alpha1.PulsarCluster) (err error) {
	cmCreate := zookeeper.MakeConfigMap(c)

	if err = controllerutil.SetControllerReference(c, cmCreate, r.scheme); err != nil {
		return err
	}

	cmCur := &v1.ConfigMap{}
	err = r.client.Get(context.TODO(), types.NamespacedName{
		Name:      cmCreate.Name,
		Namespace: cmCreate.Namespace,
	}, cmCur)
	if err != nil && errors.IsNotFound(err) {
		err = r.client.Create(context.TODO(), cmCreate)
		if err == nil {
			r.log.Info("Create Pulsar Zookeeper Config Map Success",
				"ConfigMap.Namespace", c.Namespace,
				"ConfigMap.Name", cmCreate.GetName())
		}
	}
	return
}

func (r *ReconcilePulsarCluster) reconcileZookeeperStatefulSet(c *pulsarv1alpha1.PulsarCluster) (err error) {
	ssCreate := zookeeper.MakeStatefulSet(c)

	if err = controllerutil.SetControllerReference(c, ssCreate, r.scheme); err != nil {
		return err
	}

	ssCur := &appsv1.StatefulSet{}
	err = r.client.Get(context.TODO(), types.NamespacedName{
		Name:      ssCreate.Name,
		Namespace: ssCreate.Namespace,
	}, ssCur)
	if err != nil && errors.IsNotFound(err) {
		err = r.client.Create(context.TODO(), ssCreate)
		if err == nil {
			r.log.Info("Create Pulsar Zookeeper StatefulSet Success",
				"StatefulSet.Namespace", c.Namespace,
				"StatefulSet.Name", ssCreate.GetName())
		}
	} else if err != nil {
		return err
	} else {
		if c.Spec.Zookeeper.Size != *ssCur.Spec.Replicas {
			old := *ssCur.Spec.Replicas
			ssCur.Spec.Replicas = &c.Spec.Zookeeper.Size
			err = r.client.Update(context.TODO(), ssCur)
			if err == nil {
				r.log.Info("Scale Pulsar Zookeeper StatefulSet Success",
					"OldSize", old,
					"NewSize", c.Spec.Zookeeper.Size)
			}
		}
	}
	return
}

func (r *ReconcilePulsarCluster) reconcileZookeeperService(c *pulsarv1alpha1.PulsarCluster) (err error) {
	sCreate := zookeeper.MakeService(c)

	if err = controllerutil.SetControllerReference(c, sCreate, r.scheme); err != nil {
		return err
	}

	sCur := &v1.Service{}
	err = r.client.Get(context.TODO(), types.NamespacedName{
		Name:      sCreate.Name,
		Namespace: sCreate.Namespace,
	}, sCur)
	if err != nil && errors.IsNotFound(err) {
		err = r.client.Create(context.TODO(), sCreate)
		if err == nil {
			r.log.Info("Create Pulsar Zookeeper Service Success",
				"Service.Namespace", c.Namespace,
				"Service.Name", sCreate.GetName())
		}
	}
	return
}

func (r *ReconcilePulsarCluster) reconcileZookeeperPodDisruptionBudget(c *pulsarv1alpha1.PulsarCluster) (err error) {
	pdb := zookeeper.MakePodDisruptionBudget(c)

	if err = controllerutil.SetControllerReference(c, pdb, r.scheme); err != nil {
		return err
	}

	pdbCur := &v1beta1.PodDisruptionBudget{}
	err = r.client.Get(context.TODO(), types.NamespacedName{
		Name:      pdb.Name,
		Namespace: pdb.Namespace,
	}, pdbCur)
	if err != nil && errors.IsNotFound(err) {
		err = r.client.Create(context.TODO(), pdb)
		if err == nil {
			r.log.Info("Create Pulsar Zookeeper PodDisruptionBudget Success",
				"PodDisruptionBudget.Namespace", c.Namespace,
				"PodDisruptionBudget.Name", pdb.GetName())
		}
	}
	return
}
