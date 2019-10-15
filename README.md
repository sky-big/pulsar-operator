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

4. Now you can use the CRDs provide by Pulsar Operator to deploy your Pulsar Cluster, Start one example pulsar cluster:
```
$ kubectl create -f deploy/crds/pulsar_v1alpha1_pulsarcluster_cr.yaml
```

5. Use command ```kubectl get pods```to check example Pulsar Cluster status like:
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
