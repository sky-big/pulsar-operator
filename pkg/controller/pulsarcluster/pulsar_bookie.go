package pulsarcluster

import (
	"context"

	pulsarv1alpha1 "github.com/sky-big/pulsar-operator/pkg/apis/pulsar/v1alpha1"
	"github.com/sky-big/pulsar-operator/pkg/pulsar/components/bookie"
	"github.com/sky-big/pulsar-operator/pkg/pulsar/components/bookie/autorecovery"

	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"

	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
)

func (r *ReconcilePulsarCluster) reconcileBookie(c *pulsarv1alpha1.PulsarCluster) error {
	if c.Status.Phase == pulsarv1alpha1.PulsarClusterInitingPhase {
		return nil
	}

	for _, fun := range []reconcileFunc{
		r.reconcileBookieConfigMap,
		r.reconcileBookieStatefulSet,
		r.reconcileBookieService,
		r.reconcileBookieAutoRecoveryDeployment,
	} {
		if err := fun(c); err != nil {
			r.log.Error(err, "Reconciling PulsarCluster Bookie Error", c)
			return err
		}
	}
	return nil
}

func (r *ReconcilePulsarCluster) reconcileBookieConfigMap(c *pulsarv1alpha1.PulsarCluster) (err error) {
	cmCreate := bookie.MakeConfigMap(c)

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
			r.log.Info("Create pulsar bookie config map success",
				"ConfigMap.Namespace", c.Namespace,
				"ConfigMap.Name", cmCreate.GetName())
		}
	}
	return
}

func (r *ReconcilePulsarCluster) reconcileBookieStatefulSet(c *pulsarv1alpha1.PulsarCluster) (err error) {
	ssCreate := bookie.MakeStatefulSet(c)

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
			r.log.Info("Create pulsar bookie statefulSet success",
				"StatefulSet.Namespace", c.Namespace,
				"StatefulSet.Name", ssCreate.GetName())
		}
	} else if err != nil {
		return err
	} else {
		if c.Spec.Bookie.Size != *ssCur.Spec.Replicas {
			old := *ssCur.Spec.Replicas
			ssCur.Spec.Replicas = &c.Spec.Bookie.Size
			if err = r.client.Update(context.TODO(), ssCur); err == nil {
				r.log.Info("Scale pulsar bookie statefulSet success",
					"OldSize", old,
					"NewSize", c.Spec.Bookie.Size)
			}
		}
	}

	r.log.Info("Bookie node num info",
		"Replicas", ssCur.Status.Replicas,
		"ReadyNum", ssCur.Status.ReadyReplicas,
		"CurrentNum", ssCur.Status.CurrentReplicas,
	)
	return
}

func (r *ReconcilePulsarCluster) reconcileBookieService(c *pulsarv1alpha1.PulsarCluster) (err error) {
	sCreate := bookie.MakeService(c)

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
			r.log.Info("Create Pulsar Bookie Service Success",
				"Service.Namespace", c.Namespace,
				"Service.Name", sCreate.GetName())
		}
	}
	return
}

func (r *ReconcilePulsarCluster) reconcileBookieAutoRecoveryDeployment(c *pulsarv1alpha1.PulsarCluster) (err error) {
	dmCreate := autorecovery.MakeDeployment(c)

	dmCur := &appsv1.Deployment{}
	err = r.client.Get(context.TODO(), types.NamespacedName{
		Name:      dmCreate.Name,
		Namespace: dmCreate.Namespace,
	}, dmCur)
	if err != nil && errors.IsNotFound(err) {
		if err = controllerutil.SetControllerReference(c, dmCreate, r.scheme); err != nil {
			return err
		}

		if err = r.client.Create(context.TODO(), dmCreate); err == nil {
			r.log.Info("Create pulsar bookie autoRecovery deployment success",
				"Deployment.Namespace", c.Namespace,
				"Deployment.Name", dmCreate.GetName())
		}
	}

	r.log.Info("Bookie pulsar autoRecovery node num info",
		"Replicas", dmCur.Status.Replicas,
		"ReadyNum", dmCur.Status.ReadyReplicas,
	)
	return
}

func (r *ReconcilePulsarCluster) isBookieRunning(c *pulsarv1alpha1.PulsarCluster) bool {
	ss := &appsv1.StatefulSet{}
	err := r.client.Get(context.TODO(), types.NamespacedName{
		Name:      bookie.MakeStatefulSetName(c),
		Namespace: c.Namespace,
	}, ss)
	if err == nil {
		return ss.Status.ReadyReplicas == c.Spec.Bookie.Size
	}
	return false
}
