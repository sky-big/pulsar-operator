#!/usr/bin/env bash

export WORK_DIR=$(cd `dirname $0`; pwd)

# delete one example pulsar cluster
kubectl delete -f ${WORK_DIR}/crds/pulsar_v1alpha1_pulsarcluster_cr.yaml

# delete pulsar cluster operator account and role
kubectl delete -f ${WORK_DIR}/role.yaml
kubectl delete -f ${WORK_DIR}/role_binding.yaml
kubectl delete -f ${WORK_DIR}/service_account.yaml

# delete pulsar cluster crd
kubectl delete -f ${WORK_DIR}/crds/pulsar_v1alpha1_pulsarcluster_crd.yaml
