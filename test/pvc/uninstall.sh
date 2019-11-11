#!/usr/bin/env bash

export WORK_DIR=$(cd `dirname $0`; cd ../..; pwd)

# delete one example pulsar cluster
kubectl delete -f ${WORK_DIR}/test/pvc/pulsar_v1alpha1_pulsarcluster_cr.yaml
kubectl delete -f ${WORK_DIR}/test/admin/admin.yaml

# delete pulsar cluster operator account and role
kubectl delete -f ${WORK_DIR}/deploy/rbac/all_namespace_rbac.yaml

# delete pulsar cluster crd
kubectl delete -f ${WORK_DIR}/deploy/crds/pulsar.apache.org_pulsarclusters_crd.yaml

# delete pvc
kubectl delete pvc journal-disk-volume-pvc-example-pulsarcluster-bookie-statefulset-0
kubectl delete pvc journal-disk-volume-pvc-example-pulsarcluster-bookie-statefulset-1
kubectl delete pvc journal-disk-volume-pvc-example-pulsarcluster-bookie-statefulset-2
kubectl delete pvc ledgers-disk-volume-pvc-example-pulsarcluster-bookie-statefulset-0
kubectl delete pvc ledgers-disk-volume-pvc-example-pulsarcluster-bookie-statefulset-1
kubectl delete pvc ledgers-disk-volume-pvc-example-pulsarcluster-bookie-statefulset-2
