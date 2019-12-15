package pulsarcluster

import (
	"context"

	pulsarv1alpha1 "github.com/sky-big/pulsar-operator/pkg/apis/pulsar/v1alpha1"
	"github.com/sky-big/pulsar-operator/pkg/pulsar/metadata"

	"github.com/go-logr/logr"

	appsv1 "k8s.io/api/apps/v1"
	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/api/policy/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"

	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

type reconcileFunc func(cluster *pulsarv1alpha1.PulsarCluster) error

var log = logf.Log.WithName("controller_pulsarcluster")

/**
* USER ACTION REQUIRED: This is a scaffold file intended for the user to modify with their own Controller
* business logic.  Delete these comments after modifying this file.*
 */

// Add creates a new PulsarCluster Controller and adds it to the Manager. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func Add(mgr manager.Manager) error {
	return add(mgr, newReconciler(mgr))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	return &ReconcilePulsarCluster{client: mgr.GetClient(), scheme: mgr.GetScheme()}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Create a new controller
	c, err := controller.New("pulsarcluster-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to primary resource PulsarCluster
	err = c.Watch(&source.Kind{Type: &pulsarv1alpha1.PulsarCluster{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	// Watch Config Map
	err = c.Watch(&source.Kind{Type: &corev1.ConfigMap{}}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    &pulsarv1alpha1.PulsarCluster{},
	})
	if err != nil {
		return err
	}

	// Watch StatefulSet
	err = c.Watch(&source.Kind{Type: &appsv1.StatefulSet{}}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    &pulsarv1alpha1.PulsarCluster{},
	})
	if err != nil {
		return err
	}

	// Watch Service
	err = c.Watch(&source.Kind{Type: &corev1.Service{}}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    &pulsarv1alpha1.PulsarCluster{},
	})
	if err != nil {
		return err
	}

	// Watch Pod
	err = c.Watch(&source.Kind{Type: &corev1.Pod{}}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    &pulsarv1alpha1.PulsarCluster{},
	})

	// Watch PodDisruptionBudget
	err = c.Watch(&source.Kind{Type: &v1beta1.PodDisruptionBudget{}}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    &pulsarv1alpha1.PulsarCluster{},
	})

	// Watch Job
	err = c.Watch(&source.Kind{Type: &batchv1.Job{}}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    &pulsarv1alpha1.PulsarCluster{},
	})

	return nil
}

// blank assignment to verify that ReconcilePulsarCluster implements reconcile.Reconciler
var _ reconcile.Reconciler = &ReconcilePulsarCluster{}

// ReconcilePulsarCluster reconciles a PulsarCluster object
type ReconcilePulsarCluster struct {
	// This client, initialized using mgr.Client() above, is a split client
	// that reads objects from the cache and writes to the apiserver
	client client.Client
	scheme *runtime.Scheme
	log    logr.Logger
}

// Reconcile reads that state of the cluster for a PulsarCluster object and makes changes based on the state read
// and what is in the PulsarCluster.Spec
// Note:
// The Controller will requeue the Request to be processed again if the returned error is non-nil or
// Result.Requeue is true, otherwise upon completion it will remove the work from the queue-proxy.
func (r *ReconcilePulsarCluster) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	r.log = log.WithValues("Request.Namespace", request.Namespace, "Request.Name", request.Name)
	r.log.Info("[Start] Reconciling PulsarCluster")
	defer func() {
		r.log.Info("[End] Reconciling PulsarCluster")
	}()

	// Fetch the PulsarCluster instance
	instance := &pulsarv1alpha1.PulsarCluster{}
	err := r.client.Get(context.TODO(), request.NamespacedName, instance)
	if err != nil {
		if errors.IsNotFound(err) {
			// Request object not found, could have been deleted after reconcile request.
			// Owned objects are automatically garbage collected. For additional cleanup logic use finalizers.
			// Return and don't requeue
			return reconcile.Result{}, nil
		}
		// Error reading the object - requeue the request.
		return reconcile.Result{}, err
	}

	// Set pulsar cluster resource spec default
	changed := instance.SpecSetDefault()
	if changed {
		r.log.Info("Setting spec default settings for pulsar-cluster")
		if err := r.client.Update(context.TODO(), instance); err != nil {
			return reconcile.Result{}, err
		}
		return reconcile.Result{Requeue: true}, nil
	}

	// Set pulsar cluster resource status default
	changed = instance.StatusSetDefault()
	if changed {
		r.log.Info("Setting status default settings for pulsar-cluster")
		if err := r.client.Status().Update(context.TODO(), instance); err != nil {
			return reconcile.Result{}, err
		}
		return reconcile.Result{Requeue: true}, nil
	}

	// Reconcile all pulsar cluster child resource
	for _, fun := range []reconcileFunc{
		r.reconcileZookeeper,
		r.reconcileBookie,
		r.reconcileBroker,
		r.reconcileProxy,
		r.reconcilePulsarCluster,
		r.reconcileMonitor,
		r.reconcileManager,
	} {
		if err = fun(instance); err != nil {
			return reconcile.Result{}, err
		}
	}

	return reconcile.Result{}, nil
}

// Reconcile pulsarCluster resource
func (r *ReconcilePulsarCluster) reconcilePulsarCluster(c *pulsarv1alpha1.PulsarCluster) error {
	if err := r.reconcileInitPulsarClusterMetaData(c); err != nil {
		return err
	}

	if err := r.reconcilePulsarClusterPhase(c); err != nil {
		return err
	}
	return nil
}

// Init pulsar metaData
func (r *ReconcilePulsarCluster) reconcileInitPulsarClusterMetaData(c *pulsarv1alpha1.PulsarCluster) (err error) {
	if c.Status.Phase == pulsarv1alpha1.PulsarClusterInitingPhase && r.isZookeeperRunning(c) {
		job := metadata.MakeInitClusterMetaDataJob(c)

		jobCur := &batchv1.Job{}
		err = r.client.Get(context.TODO(), types.NamespacedName{
			Name:      job.Name,
			Namespace: job.Namespace,
		}, jobCur)
		if err != nil && errors.IsNotFound(err) {
			if err = controllerutil.SetControllerReference(c, job, r.scheme); err != nil {
				return err
			}

			if err = r.client.Create(context.TODO(), job); err == nil {
				r.log.Info("Start init pulsar cluster metaData job",
					"Job.Namespace", job.Namespace,
					"Job.Name", job.Name)
			}

		} else if err == nil && jobCur.Status.Succeeded == 1 {
			// Init pulsar cluster success
			c.Status.Phase = pulsarv1alpha1.PulsarClusterLaunchingPhase
			if err = r.client.Status().Update(context.TODO(), c); err == nil {
				r.log.Info("Init pulsar cluster metaData success",
					"PulsarCluster.Namespace", c.Namespace,
					"PulsarCluster.Name", c.Name)
			}
		}
	}
	return
}

func (r *ReconcilePulsarCluster) reconcilePulsarClusterPhase(c *pulsarv1alpha1.PulsarCluster) (err error) {
	if c.Status.Phase == pulsarv1alpha1.PulsarClusterLaunchingPhase &&
		r.isZookeeperRunning(c) &&
		r.isBookieRunning(c) &&
		r.isBrokerRunning(c) &&
		r.isProxyRunning(c) {
		c.Status.Phase = pulsarv1alpha1.PulsarClusterRunningPhase
		if err = r.client.Status().Update(context.TODO(), c); err == nil {
			r.log.Info("start pulsar cluster success",
				"PulsarCluster.Namespace", c.Namespace,
				"PulsarCluster.Name", c.Name)
		}
	}
	return
}
