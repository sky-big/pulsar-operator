package pulsarcluster

import (
	"context"

	pulsarv1alpha1 "github.com/sky-big/pulsar-operator/pkg/apis/pulsar/v1alpha1"
	"github.com/sky-big/pulsar-operator/pkg/pulsar/components/broker"

	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"

	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
)

func (r *ReconcilePulsarCluster) reconcileBroker(c *pulsarv1alpha1.PulsarCluster) error {
	if c.Status.Phase == pulsarv1alpha1.PulsarClusterInitingPhase {
		return nil
	}

	for _, fun := range []reconcileFunc{
		r.reconcileBrokerConfigMap,
		r.reconcileBrokerDeployment,
		r.reconcileBrokerService,
	} {
		if err := fun(c); err != nil {
			r.log.Error(err, "Reconciling PulsarCluster Broker Error", c)
			return err
		}
	}
	return nil
}

func (r *ReconcilePulsarCluster) reconcileBrokerConfigMap(c *pulsarv1alpha1.PulsarCluster) (err error) {
	cmCreate := broker.MakeConfigMap(c)

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
			r.log.Info("Create pulsar broker config map success",
				"ConfigMap.Namespace", c.Namespace,
				"ConfigMap.Name", cmCreate.GetName())
		}
	}
	return
}

func (r *ReconcilePulsarCluster) reconcileBrokerDeployment(c *pulsarv1alpha1.PulsarCluster) (err error) {
	dmCreate := broker.MakeDeployment(c)

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
			r.log.Info("Create pulsar broker deployment success",
				"Deployment.Namespace", c.Namespace,
				"Deployment.Name", dmCreate.GetName())
		}
	} else if err != nil {
		return err
	} else {
		if c.Spec.Broker.Size != *dmCur.Spec.Replicas {
			old := *dmCur.Spec.Replicas
			dmCur.Spec.Replicas = &c.Spec.Broker.Size
			if err = r.client.Update(context.TODO(), dmCur); err == nil {
				r.log.Info("Scale pulsar broker deployment success",
					"OldSize", old,
					"NewSize", c.Spec.Broker.Size)
			}
		}
	}
	return
}

func (r *ReconcilePulsarCluster) reconcileBrokerService(c *pulsarv1alpha1.PulsarCluster) (err error) {
	sCreate := broker.MakeService(c)

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
			r.log.Info("Create pulsar broker service success",
				"Service.Namespace", c.Namespace,
				"Service.Name", sCreate.GetName())
		}
	}
	return
}

func (r *ReconcilePulsarCluster) isBrokerRunning(c *pulsarv1alpha1.PulsarCluster) bool {
	dm := broker.MakeDeployment(c)

	dmCur := &appsv1.Deployment{}
	err := r.client.Get(context.TODO(), types.NamespacedName{
		Name:      dm.Name,
		Namespace: dm.Namespace,
	}, dmCur)
	return err == nil && dmCur.Status.ReadyReplicas == c.Spec.Broker.Size
}
