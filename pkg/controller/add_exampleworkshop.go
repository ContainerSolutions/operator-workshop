package controller

import (
	"github.com/ContainerSolutions/example-operator/pkg/controller/exampleworkshop"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, exampleworkshop.Add)
}
