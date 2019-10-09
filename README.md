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
$ 
```

3. Use command ```kubectl get pods``` to check Pulsar Operator deploy status like:
```
$ kubectl get pods
NAME                                      READY   STATUS    RESTARTS   AGE
pulsar-operator-564b5d75d-jllzk           1/1     Running   0          108s
```

Now you can use the CRDs provide by Pulsar Operator to deploy your Pulsar Cluster.
