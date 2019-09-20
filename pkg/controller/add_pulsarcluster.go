package controller

import (
	"github.com/sky-big/pulsar-operator/pkg/controller/pulsarcluster"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, pulsarcluster.Add)
}
