#!/usr/bin/env bash

export WORK_DIR=$(cd `dirname $0`; pwd)

# create pulsar cluster crd
kubectl create -f ${WORK_DIR}/crds/pulsar.apache.org_pulsarclusters_crd.yaml

# create pulsar cluster operator account and role
kubectl create -f ${WORK_DIR}/rbac/all_namespace_rbac.yaml

# install pulsar cluster operator
kubectl create -f ${WORK_DIR}/operator.yaml
