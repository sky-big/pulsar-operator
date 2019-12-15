package pulsarcluster

import (
	"context"

	pulsarv1alpha1 "github.com/sky-big/pulsar-operator/pkg/apis/pulsar/v1alpha1"
	"github.com/sky-big/pulsar-operator/pkg/pulsar/components/zookeeper"

	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/api/core/v1"
	"k8s.io/api/policy/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"

	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
)

// Reconcile For All Resource About Zookeeper
func (r *ReconcilePulsarCluster) reconcileZookeeper(c *pulsarv1alpha1.PulsarCluster) error {
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

	cmCur := &v1.ConfigMap{}
	err = r.client.Get(context.TODO(), types.NamespacedName{
		Name:      cmCreate.Name,
		Namespace: cmCreate.Namespace,
	}, cmCur)
	if err != nil && errors.IsNotFound(err) {
		if err = controllerutil.SetControllerReference(c, cmCreate, r.scheme); err != nil {
			return err
		}

		if err = r.client.Create(context.TODO(), cmCreate); err == nil {
			r.log.Info("Create pulsar zookeeper config map success",
				"ConfigMap.Namespace", c.Namespace,
				"ConfigMap.Name", cmCreate.GetName())
		}
	}
	return
}

func (r *ReconcilePulsarCluster) reconcileZookeeperStatefulSet(c *pulsarv1alpha1.PulsarCluster) (err error) {
	ssCreate := zookeeper.MakeStatefulSet(c)

	ssCur := &appsv1.StatefulSet{}
	err = r.client.Get(context.TODO(), types.NamespacedName{
		Name:      ssCreate.Name,
		Namespace: ssCreate.Namespace,
	}, ssCur)
	if err != nil && errors.IsNotFound(err) {
		if err = controllerutil.SetControllerReference(c, ssCreate, r.scheme); err != nil {
			return err
		}

		if err = r.client.Create(context.TODO(), ssCreate); err == nil {
			r.log.Info("Create pulsar zookeeper statefulSet success",
				"StatefulSet.Namespace", c.Namespace,
				"StatefulSet.Name", ssCreate.GetName())
		}
	} else if err != nil {
		return err
	} else {
		if c.Spec.Zookeeper.Size != *ssCur.Spec.Replicas {
			old := *ssCur.Spec.Replicas
			ssCur.Spec.Replicas = &c.Spec.Zookeeper.Size
			if err = r.client.Update(context.TODO(), ssCur); err == nil {
				r.log.Info("Scale pulsar zookeeper statefulSet success",
					"OldSize", old,
					"NewSize", c.Spec.Zookeeper.Size)
			}
		}
	}

	r.log.Info("Zookeeper node num info",
		"Replicas", ssCur.Status.Replicas,
		"ReadyNum", ssCur.Status.ReadyReplicas,
		"CurrentNum", ssCur.Status.CurrentReplicas,
	)
	return
}

func (r *ReconcilePulsarCluster) reconcileZookeeperService(c *pulsarv1alpha1.PulsarCluster) (err error) {
	sCreate := zookeeper.MakeService(c)

	sCur := &v1.Service{}
	err = r.client.Get(context.TODO(), types.NamespacedName{
		Name:      sCreate.Name,
		Namespace: sCreate.Namespace,
	}, sCur)
	if err != nil && errors.IsNotFound(err) {
		if err = controllerutil.SetControllerReference(c, sCreate, r.scheme); err != nil {
			return err
		}

		if err = r.client.Create(context.TODO(), sCreate); err == nil {
			r.log.Info("Create pulsar zookeeper service success",
				"Service.Namespace", c.Namespace,
				"Service.Name", sCreate.GetName())
		}
	}
	return
}

func (r *ReconcilePulsarCluster) reconcileZookeeperPodDisruptionBudget(c *pulsarv1alpha1.PulsarCluster) (err error) {
	pdb := zookeeper.MakePodDisruptionBudget(c)

	pdbCur := &v1beta1.PodDisruptionBudget{}
	err = r.client.Get(context.TODO(), types.NamespacedName{
		Name:      pdb.Name,
		Namespace: pdb.Namespace,
	}, pdbCur)
	if err != nil && errors.IsNotFound(err) {
		if err = controllerutil.SetControllerReference(c, pdb, r.scheme); err != nil {
			return err
		}

		if err = r.client.Create(context.TODO(), pdb); err == nil {
			r.log.Info("Create pulsar zookeeper podDisruptionBudget success",
				"PodDisruptionBudget.Namespace", c.Namespace,
				"PodDisruptionBudget.Name", pdb.GetName())
		}
	}
	return
}

func (r *ReconcilePulsarCluster) isZookeeperRunning(c *pulsarv1alpha1.PulsarCluster) bool {
	ss := &appsv1.StatefulSet{}
	err := r.client.Get(context.TODO(), types.NamespacedName{
		Name:      zookeeper.MakeStatefulSetName(c),
		Namespace: c.Namespace,
	}, ss)
	if err == nil {
		return ss.Status.ReadyReplicas == c.Spec.Zookeeper.Size
	}
	return false
}
