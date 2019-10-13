package pulsarcluster

import (
	"context"

	pulsarv1alpha1 "github.com/sky-big/pulsar-operator/pkg/apis/pulsar/v1alpha1"
	"github.com/sky-big/pulsar-operator/pkg/components/monitor/dashboard"
	"github.com/sky-big/pulsar-operator/pkg/components/monitor/grafana"
	"github.com/sky-big/pulsar-operator/pkg/components/monitor/prometheus"

	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"

	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
)

func (r *ReconcilePulsarCluster) reconcileMonitor(c *pulsarv1alpha1.PulsarCluster) error {
	if c.Spec.Monitor.IsActive != pulsarv1alpha1.MonitorActived {
		return nil
	}

	for _, fun := range []reconcileFunc{
		r.reconcileMonitorDashboard,
		r.reconcileMonitorPrometheus,
		r.reconcileMonitorGrafana,
	} {
		if err := fun(c); err != nil {
			r.log.Error(err, "Reconciling PulsarCluster Monitor Error", c)
			return err
		}
	}
	return nil
}

func (r *ReconcilePulsarCluster) reconcileMonitorDashboard(c *pulsarv1alpha1.PulsarCluster) error {
	if c.Spec.Monitor.DashboardPort == 0 {
		return nil
	}

	for _, fun := range []reconcileFunc{
		r.reconcileMonitorDashboardDeployment,
		r.reconcileMonitorDashboardService,
	} {
		if err := fun(c); err != nil {
			return err
		}
	}
	return nil
}

func (r *ReconcilePulsarCluster) reconcileMonitorDashboardDeployment(c *pulsarv1alpha1.PulsarCluster) (err error) {
	dmCreate := dashboard.MakeDeployment(c)

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
			r.log.Info("Create pulsar monitor dashboard deployment success",
				"Deployment.Namespace", c.Namespace,
				"Deployment.Name", dmCreate.GetName())
		}
	}
	return
}

func (r *ReconcilePulsarCluster) reconcileMonitorDashboardService(c *pulsarv1alpha1.PulsarCluster) (err error) {
	sCreate := dashboard.MakeService(c)

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
			r.log.Info("Create pulsar monitor dashboard service success",
				"Service.Namespace", c.Namespace,
				"Service.Name", sCreate.GetName())
		}
	}
	return
}

func (r *ReconcilePulsarCluster) reconcileMonitorPrometheus(c *pulsarv1alpha1.PulsarCluster) error {
	if c.Spec.Monitor.PrometheusPort == 0 {
		return nil
	}

	for _, fun := range []reconcileFunc{
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
	if c.Spec.Monitor.GrafanaPort == 0 {
		return nil
	}

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
