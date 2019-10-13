#!/usr/bin/env bash

export WORK_DIR=$(cd `dirname $0`; pwd)

# create pulsar cluster crd
kubectl create -f ${WORK_DIR}/crds/pulsar_v1alpha1_pulsarcluster_crd.yaml

# create pulsar cluster operator account and role
kubectl create -f ${WORK_DIR}/role.yaml
kubectl create -f ${WORK_DIR}/role_binding.yaml
kubectl create -f ${WORK_DIR}/service_account.yaml

# install pulsar cluster operator
kubectl create -f ${WORK_DIR}/operator.yaml

# create one example pulsar cluster
kubectl create -f ${WORK_DIR}/crds/pulsar_v1alpha1_pulsarcluster_cr.yaml
