package pulsarcluster

import (
	"context"

	pulsarv1alpha1 "github.com/sky-big/pulsar-operator/pkg/apis/pulsar/v1alpha1"
	"github.com/sky-big/pulsar-operator/pkg/pulsar/components/monitor/grafana"
	"github.com/sky-big/pulsar-operator/pkg/pulsar/components/monitor/ingress"
	"github.com/sky-big/pulsar-operator/pkg/pulsar/components/monitor/prometheus"

	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	"k8s.io/api/extensions/v1beta1"
	rbacv1 "k8s.io/api/rbac/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"

	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
)

func (r *ReconcilePulsarCluster) reconcileMonitor(c *pulsarv1alpha1.PulsarCluster) error {
	if c.Status.Phase == pulsarv1alpha1.PulsarClusterInitingPhase {
		return nil
	}
	if !c.Spec.Monitor.Enable {
		return nil
	}

	for _, fun := range []reconcileFunc{
		r.reconcileMonitorPrometheus,
		r.reconcileMonitorGrafana,
	} {
		if err := fun(c); err != nil {
			r.log.Error(err, "Reconciling PulsarCluster Monitor Error", c)
			return err
		}
	}

	if c.Spec.Monitor.Ingress.Enable &&
		(c.Spec.Monitor.Grafana.Host != "" ||
			c.Spec.Monitor.Prometheus.Host != "") {
		if err := r.reconcileMonitorIngress(c); err != nil {
			r.log.Error(err, "Reconciling PulsarCluster Monitor Ingress Error", c)
			return err
		}
	}
	return nil
}

func (r *ReconcilePulsarCluster) reconcileMonitorPrometheus(c *pulsarv1alpha1.PulsarCluster) error {
	for _, fun := range []reconcileFunc{
		r.reconcileMonitorPrometheusRBAC,
		r.reconcileMonitorPrometheusConfigMap,
		r.reconcileMonitorPrometheusDeployment,
		r.reconcileMonitorPrometheusService,
	} {
		if err := fun(c); err != nil {
			return err
		}
	}
	return nil
}

func (r *ReconcilePulsarCluster) reconcileMonitorPrometheusRBAC(c *pulsarv1alpha1.PulsarCluster) (err error) {
	// cluster role
	crCreate := prometheus.MakeClusterRole(c)
	crCur := &rbacv1.ClusterRole{}
	err = r.client.Get(context.TODO(), types.NamespacedName{
		Name: crCreate.Name,
	}, crCur)
	if err != nil && errors.IsNotFound(err) {
		if err = controllerutil.SetControllerReference(c, crCreate, r.scheme); err != nil {
			return err
		}

		if err = r.client.Create(context.TODO(), crCreate); err == nil {
			r.log.Info("Create pulsar monitor prometheus cluster role success",
				"ClusterRole.Name", crCreate.GetName())
		}
	}

	// service account
	saCreate := prometheus.MakeServiceAccount(c)
	saCur := &v1.ServiceAccount{}
	err = r.client.Get(context.TODO(), types.NamespacedName{
		Name:      saCreate.Name,
		Namespace: c.GetNamespace(),
	}, saCur)
	if err != nil && errors.IsNotFound(err) {
		if err = controllerutil.SetControllerReference(c, saCreate, r.scheme); err != nil {
			return err
		}

		if err = r.client.Create(context.TODO(), saCreate); err == nil {
			r.log.Info("Create pulsar monitor prometheus service account success",
				"ServiceAccount.Name", saCreate.GetName())
		}
	}

	// cluster role and service account binding
	rbCreate := prometheus.MakeClusterRoleBinding(c)
	rbCur := &rbacv1.ClusterRoleBinding{}
	err = r.client.Get(context.TODO(), types.NamespacedName{
		Name: rbCreate.Name,
	}, rbCur)
	if err != nil && errors.IsNotFound(err) {
		if err = controllerutil.SetControllerReference(c, rbCreate, r.scheme); err != nil {
			return err
		}

		if err = r.client.Create(context.TODO(), rbCreate); err == nil {
			r.log.Info("Create pulsar monitor prometheus cluster role binding success",
				"ClusterRoleBinding.Name", rbCreate.GetName())
		}
	}
	return
}

func (r *ReconcilePulsarCluster) reconcileMonitorPrometheusConfigMap(c *pulsarv1alpha1.PulsarCluster) (err error) {
	cmCreate := prometheus.MakeConfigMap(c)

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
			r.log.Info("Create pulsar monitor prometheus config map success",
				"ConfigMap.Namespace", c.Namespace,
				"ConfigMap.Name", cmCreate.GetName())
		}
	}
	return
}

func (r *ReconcilePulsarCluster) reconcileMonitorPrometheusDeployment(c *pulsarv1alpha1.PulsarCluster) (err error) {
	dmCreate := prometheus.MakeDeployment(c)

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
			r.log.Info("Create pulsar monitor prometheus deployment success",
				"Deployment.Namespace", c.Namespace,
				"Deployment.Name", dmCreate.GetName())
		}
	}
	return
}

func (r *ReconcilePulsarCluster) reconcileMonitorPrometheusService(c *pulsarv1alpha1.PulsarCluster) (err error) {
	sCreate := prometheus.MakeService(c)

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
			r.log.Info("Create pulsar monitor prometheus service success",
				"Service.Namespace", c.Namespace,
				"Service.Name", sCreate.GetName())
		}
	}
	return
}

func (r *ReconcilePulsarCluster) reconcileMonitorGrafana(c *pulsarv1alpha1.PulsarCluster) error {
	for _, fun := range []reconcileFunc{
		r.reconcileMonitorGrafanaDeployment,
		r.reconcileMonitorGrafanaService,
	} {
		if err := fun(c); err != nil {
			return err
		}
	}
	return nil
}

func (r *ReconcilePulsarCluster) reconcileMonitorGrafanaDeployment(c *pulsarv1alpha1.PulsarCluster) (err error) {
	dmCreate := grafana.MakeDeployment(c)

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
			r.log.Info("Create pulsar monitor grafana deployment success",
				"Deployment.Namespace", c.Namespace,
				"Deployment.Name", dmCreate.GetName())
		}
	}
	return
}

func (r *ReconcilePulsarCluster) reconcileMonitorGrafanaService(c *pulsarv1alpha1.PulsarCluster) (err error) {
	sCreate := grafana.MakeService(c)

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
			r.log.Info("Create pulsar monitor grafana service success",
				"Service.Namespace", c.Namespace,
				"Service.Name", sCreate.GetName())
		}
	}
	return
}

func (r *ReconcilePulsarCluster) reconcileMonitorIngress(c *pulsarv1alpha1.PulsarCluster) (err error) {
	inCreate := ingress.MakeIngress(c)

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
			r.log.Info("Create pulsar monitor ingress success",
				"Ingress.Namespace", c.Namespace,
				"Ingress.Name", inCreate.GetName())
		}
	}
	return
}
