package bookie

import (
	"fmt"

	pulsarv1alpha1 "github.com/sky-big/pulsar-operator/pkg/apis/pulsar/v1alpha1"

	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func makeVolumeClaimTemplates(c *pulsarv1alpha1.PulsarCluster) []v1.PersistentVolumeClaim {
	return []v1.PersistentVolumeClaim{
		makeJournalDataVolumeClaimTemplate(c),
		makeLedgersDataVolumeClaimTemplate(c),
	}
}

func makeJournalDataVolumeClaimTemplate(c *pulsarv1alpha1.PulsarCluster) v1.PersistentVolumeClaim {
	return v1.PersistentVolumeClaim{
		ObjectMeta: metav1.ObjectMeta{
			Name:      makeJournalDataVolumeClaimName(c),
			Namespace: c.Namespace,
		},
		Spec: makeJournalDataVolumeClaimSpec(c),
	}
}

func makeJournalDataVolumeClaimName(c *pulsarv1alpha1.PulsarCluster) string {
	return fmt.Sprintf("%s-journal-disk", c.GetName())
}

func makeJournalDataVolumeClaimSpec(c *pulsarv1alpha1.PulsarCluster) v1.PersistentVolumeClaimSpec {
	storageName := "test"
	return v1.PersistentVolumeClaimSpec{
		AccessModes:      []v1.PersistentVolumeAccessMode{v1.ReadWriteOnce},
		Resources:        v1.ResourceRequirements{Requests: v1.ResourceList{v1.ResourceStorage: resource.MustParse("100Gi")}},
		StorageClassName: &storageName,
	}
}

func makeLedgersDataVolumeClaimTemplate(c *pulsarv1alpha1.PulsarCluster) v1.PersistentVolumeClaim {
	return v1.PersistentVolumeClaim{
		ObjectMeta: metav1.ObjectMeta{
			Name:      makeLedgersDataVolumeClaimName(c),
			Namespace: c.Namespace,
		},
		Spec: makeLedgersDataVolumeClaimSpec(c),
	}
}

func makeLedgersDataVolumeClaimName(c *pulsarv1alpha1.PulsarCluster) string {
	return fmt.Sprintf("%s-ledgers-disk", c.GetName())
}

func makeLedgersDataVolumeClaimSpec(c *pulsarv1alpha1.PulsarCluster) v1.PersistentVolumeClaimSpec {
	storageName := "test"
	return v1.PersistentVolumeClaimSpec{
		AccessModes:      []v1.PersistentVolumeAccessMode{v1.ReadWriteOnce},
		Resources:        v1.ResourceRequirements{Requests: v1.ResourceList{v1.ResourceStorage: resource.MustParse("100Gi")}},
		StorageClassName: &storageName,
	}
}
