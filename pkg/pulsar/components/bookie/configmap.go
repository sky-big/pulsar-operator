package bookie

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
			"PULSAR_MEM":                        PulsarMemData,
			"dbStorage_writeCacheMaxSizeMb":     DbStorage_writeCacheMaxSizeMb,
			"dbStorage_readAheadCacheMaxSizeMb": DbStorage_readAheadCacheMaxSizeMb,
			"zkServers":                         zookeeper.MakeServiceName(c),
			"statsProviderClass":                StatsProviderClass,
		},
	}
}

func MakeConfigMapName(c *pulsarv1alpha1.PulsarCluster) string {
	return fmt.Sprintf("%s-bookie-configmap", c.GetName())
}
