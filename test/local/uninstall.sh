#!/usr/bin/env bash

export WORK_DIR=$(cd `dirname $0`; cd ../..; pwd)

# delete one example pulsar cluster
kubectl delete -f ${WORK_DIR}/deploy/crds/pulsar_v1alpha1_pulsarcluster_cr.yaml

# delete pulsar cluster operator account and role
kubectl delete -f ${WORK_DIR}/deploy/role.yaml
kubectl delete -f ${WORK_DIR}/deploy/role_binding.yaml
kubectl delete -f ${WORK_DIR}/deploy/service_account.yaml

# delete pulsar cluster crd
kubectl delete -f ${WORK_DIR}/deploy/crds/pulsar.apache.org_pulsarclusters_crd.yaml
