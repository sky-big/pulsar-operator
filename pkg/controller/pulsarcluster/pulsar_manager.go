package pulsarcluster

import (
	"context"

	pulsarv1alpha1 "github.com/sky-big/pulsar-operator/pkg/apis/pulsar/v1alpha1"
	"github.com/sky-big/pulsar-operator/pkg/pulsar/components/manager"

	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	"k8s.io/api/extensions/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"

	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
)

func (r *ReconcilePulsarCluster) reconcileManager(c *pulsarv1alpha1.PulsarCluster) error {
	if c.Status.Phase == pulsarv1alpha1.PulsarClusterInitingPhase {
		return nil
	}
	if !c.Spec.Manager.Enable {
		return nil
	}

	for _, fun := range []reconcileFunc{
		r.reconcileManagerDeployment,
		r.reconcileManagerService,
	} {
		if err := fun(c); err != nil {
			r.log.Error(err, "Reconciling PulsarCluster Manager Error", c)
			return err
		}
	}

	if c.Spec.Manager.Host != "" {
		if err := r.reconcileManagerIngress(c); err != nil {
			r.log.Error(err, "Reconciling PulsarCluster Manager Ingress Error", c)
			return err
		}
	}
	return nil
}

func (r *ReconcilePulsarCluster) reconcileManagerDeployment(c *pulsarv1alpha1.PulsarCluster) (err error) {
	dCreate := manager.MakeDeployment(c)

	dCur := &appsv1.Deployment{}
	err = r.client.Get(context.TODO(), types.NamespacedName{
		Name:      dCreate.Name,
		Namespace: dCreate.Namespace,
	}, dCur)
	if err != nil && errors.IsNotFound(err) {
		if err = controllerutil.SetControllerReference(c, dCreate, r.scheme); err != nil {
			return err
		}

		if err = r.client.Create(context.TODO(), dCreate); err == nil {
			r.log.Info("Create pulsar manager deployment success",
				"Deployment.Namespace", c.Namespace,
				"Deployment.Name", dCreate.GetName())
		}
	}
	return
}

func (r *ReconcilePulsarCluster) reconcileManagerService(c *pulsarv1alpha1.PulsarCluster) (err error) {
	sCreate := manager.MakeService(c)

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
			r.log.Info("Create Pulsar Manager Service Success",
				"Service.Namespace", c.Namespace,
				"Service.Name", sCreate.GetName())
		}
	}
	return
}

func (r *ReconcilePulsarCluster) reconcileManagerIngress(c *pulsarv1alpha1.PulsarCluster) (err error) {
	inCreate := manager.MakeIngress(c)

	inCur := &v1beta1.Ingress{}
	err = r.client.Get(context.TODO(), types.NamespacedName{
		Name:      inCreate.Name,
		Namespace: inCreate.Namespace,
	}, inCur)
	if err != nil && errors.IsNotFound(err) {
		if err = controllerutil.SetControllerReference(c, inCreate, r.scheme); err != nil {
			return err
		}

		if err = r.client.Create(context.TODO(), inCreate); err == nil {
			r.log.Info("Create pulsar manager ingress success",
				"Ingress.Namespace", c.Namespace,
				"Ingress.Name", inCreate.GetName())
		}
	}
	return
}
