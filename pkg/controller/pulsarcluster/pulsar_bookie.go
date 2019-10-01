package pulsarcluster

import (
	"context"

	pulsarv1alpha1 "github.com/sky-big/pulsar-operator/pkg/apis/pulsar/v1alpha1"
	"github.com/sky-big/pulsar-operator/pkg/components/bookie"
	"github.com/sky-big/pulsar-operator/pkg/components/bookie/autorecovery"

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

	r.log.Info("Reconciling PulsarCluster Bookie")
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
			r.log.Info("Create Pulsar Bookie Config Map Success",
				"ConfigMap.Namespace", c.Namespace,
				"ConfigMap.Name", cmCreate.GetName())
		}
	}
	return
}

func (r *ReconcilePulsarCluster) reconcileBookieStatefulSet(c *pulsarv1alpha1.PulsarCluster) (err error) {
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
			r.log.Info("Create Pulsar Bookie AutoRecovery Deployment Success",
				"Deployment.Namespace", c.Namespace,
				"Deployment.Name", dmCreate.GetName())
		}
	}

	r.log.Info("Bookie AutoRecovery Node Num Info",
		"Replicas", dmCur.Status.Replicas,
		"ReadyNum", dmCur.Status.ReadyReplicas,
	)
	return
}
