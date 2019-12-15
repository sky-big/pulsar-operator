package broker

import (
	"fmt"

	pulsarv1alpha1 "github.com/sky-big/pulsar-operator/pkg/apis/pulsar/v1alpha1"
	"github.com/sky-big/pulsar-operator/pkg/pulsar/components/zookeeper"

	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func MakeConfigMap(c *pulsarv1alpha1.PulsarCluster) *v1.ConfigMap {
	return &v1.ConfigMap{
		TypeMeta: metav1.TypeMeta{
			Kind:       "ConfigMap",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      MakeConfigMapName(c),
			Namespace: c.Namespace,
		},
		Data: map[string]string{
			"PULSAR_MEM":                       PulsarMemData,
			"zookeeperServers":                 zookeeper.MakeServiceName(c),
			"configurationStoreServers":        zookeeper.MakeServiceName(c),
			"clusterName":                      c.GetName(),
			"managedLedgerDefaultEnsembleSize": ManagedLedgerDefaultEnsembleSize,
			"managedLedgerDefaultWriteQuorum":  ManagedLedgerDefaultWriteQuorum,
			"managedLedgerDefaultAckQuorum":    ManagedLedgerDefaultAckQuorum,
			"functionsWorkerEnabled":           FunctionsWorkerEnabled,
			"PF_pulsarFunctionsCluster":        c.GetName(),
		},
	}
}

func MakeConfigMapName(c *pulsarv1alpha1.PulsarCluster) string {
	return fmt.Sprintf("%s-broker-configmap", c.GetName())
}
