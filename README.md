# Pulsar Operator

## Overview

Pulsar Operator is to manage Pulsar service instances deployed on the Kubernetes cluster.
It is built using the [Operator SDK](https://github.com/operator-framework/operator-sdk), which is part of the [Operator Framework](https://github.com/operator-framework/).

## Quick Start

### Deploy Pulsar Operator

1. Clone the project on your Kubernetes cluster master node:
```
$ git clone https://github.com/sky-big/pulsar-operator.git
$ cd pulsar-operator
```

2. To deploy the Pulsar Operator on your Kubernetes cluster, please run the following script:
```
$ make install
```

3. Use command ```kubectl get pods``` to check Pulsar Operator deploy status like:
```
$ kubectl get pods
NAME                                      READY   STATUS    RESTARTS   AGE
pulsar-operator-564b5d75d-jllzk           1/1     Running   0          108s
```

Now you can use the CRDs provide by Pulsar Operator to deploy your Pulsar Cluster.

### Define Your Pulsar Cluster

1. Check the file ```pulsar_v1alpha1_pulsarcluster_cr.yaml```in the deploy/crd directory, for example:
```
apiVersion: pulsar.apache.org/v1alpha1
kind: PulsarCluster
metadata:
  name: example-pulsarcluster
spec:
  bookie:
    size: 3
  broker:
    size: 3
  proxy:
    size: 3
```

which defines pulsar cluster zookeeper, bookkeeper, broker, proxy components configuration

2. If you need pulsar dashboard, prometheus, grafana need configuration, for example:
```
apiVersion: pulsar.apache.org/v1alpha1
kind: PulsarCluster
metadata:
  name: example-pulsarcluster
spec:
  bookie:
    size: 3
  broker:
    size: 3
  proxy:
    size: 3
  monitor:
    isActive: true              // true/false: active monitor
    dashboardPort: 30001        // pulsar dashboard expose port on kubernetes
    prometheusPort: 30002       // prometheus expose port on kubernetes
    grafanaPort: 30003          // grafana expose port on kubernetes
```

### Create Your Pulsar Cluster

1. Deploy the pulsar cluster by running:
```
$ kubectl create -f deploy/crds/pulsar_v1alpha1_pulsarcluster_cr.yaml
```

2. Use command ```kubectl get pods```to check example Pulsar Cluster status like:
```
$ kubectl get pods
NAME                                                              READY   STATUS      RESTARTS   AGE
example-pulsarcluster-bookie-autorecovery-deployment-6775f2sbxj   1/1     Running     0          67s
example-pulsarcluster-bookie-autorecovery-deployment-6775fdqhmb   1/1     Running     0          67s
example-pulsarcluster-bookie-autorecovery-deployment-6775ftnftd   1/1     Running     0          67s
example-pulsarcluster-bookie-statefulset-0                        1/1     Running     0          68s
example-pulsarcluster-bookie-statefulset-1                        1/1     Running     0          55s
example-pulsarcluster-bookie-statefulset-2                        1/1     Running     0          42s
example-pulsarcluster-broker-deployment-5bb58577b4-4tr4l          1/1     Running     0          67s
example-pulsarcluster-broker-deployment-5bb58577b4-6vzhm          1/1     Running     0          67s
example-pulsarcluster-broker-deployment-5bb58577b4-mzphh          1/1     Running     0          67s
example-pulsarcluster-init-cluster-metadata-job-98rsd             0/1     Completed   0          80s
example-pulsarcluster-proxy-deployment-6555968487-7df5l           1/1     Running     0          67s
example-pulsarcluster-proxy-deployment-6555968487-cfxl7           1/1     Running     0          67s
example-pulsarcluster-proxy-deployment-6555968487-cxhc6           1/1     Running     0          67s
example-pulsarcluster-zookeeper-statefulset-0                     1/1     Running     0          2m2s
example-pulsarcluster-zookeeper-statefulset-1                     1/1     Running     0          109s
example-pulsarcluster-zookeeper-statefulset-2                     1/1     Running     0          95s
```

### Storage of Pulsar Cluster

1. EmptyDir Volume(Default) For Test


2. Local Volume(Next TODO)

## Horizontal Scale Pulsar Cluster

### Scale Pulsar Proxy

1. If you want to enlarge your proxy component. Modify your CR file pulsar_v1alpha1_pulsarcluster_cr.yaml, increase the field size to the number you want, for example, from ```size: 3``` to ```size: 5```
```
apiVersion: pulsar.apache.org/v1alpha1
kind: PulsarCluster
metadata:
  name: example-pulsarcluster
spec:
  proxy:
    size: 5
```

2. After configuring the size fields, simply run:
```
$ kubectl apply -f deploy/crds/pulsar_v1alpha1_pulsarcluster_cr.yaml
```

### Scale Pulsar Broker

1. If you want to enlarge your broker component. Modify your CR file pulsar_v1alpha1_pulsarcluster_cr.yaml, increase the field size to the number you want, for example, from ```size: 3``` to ```size: 5```
```
apiVersion: pulsar.apache.org/v1alpha1
kind: PulsarCluster
metadata:
  name: example-pulsarcluster
spec:
  broker:
    size: 5
```

2. After configuring the size fields, simply run:
```
$ kubectl apply -f deploy/crds/pulsar_v1alpha1_pulsarcluster_cr.yaml
```

### Scale Pulsar Bookie

1. If you want to enlarge your bookie component. Modify your CR file pulsar_v1alpha1_pulsarcluster_cr.yaml, increase the field size to the number you want, for example, from ```size: 3``` to ```size: 5```
```
apiVersion: pulsar.apache.org/v1alpha1
kind: PulsarCluster
metadata:
  name: example-pulsarcluster
spec:
  bookie:
    size: 5
```

2. After configuring the size fields, simply run:
```
$ kubectl apply -f deploy/crds/pulsar_v1alpha1_pulsarcluster_cr.yaml
```

## Local Test On Kubernetes

## Start Local Test

1. Install kubernetes cluster

2. Install golang on kubernetes master node

3. Install operator sdk[Install Operator SDK](https://github.com/operator-framework/operator-sdk/blob/master/doc/user/install-operator-sdk.md) on kubernetes master node

4. Clone project to the kubernetes master node:
```
$ git clone https://github.com/sky-big/pulsar-operator.git
$ cd pulsar-operator
```

5. Execute script on kubernetes master node:
```
$ make start-local
```

## Stop Local Test
1. Execute script on kubernetes master node:
```
$ make stop-local
```

## Compile Pulsar Operator
```
$ make build
```

## Build Pulsar Operator Image
```
$ make image
```

## Generate CRD Code And Project Vendor Code
```
$ make generate
```
