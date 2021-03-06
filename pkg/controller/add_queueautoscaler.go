package controller

import (
	"github.com/evalsocket/oldmonk/pkg/controller/queueautoscaler"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, queueautoscaler.Add)
}
