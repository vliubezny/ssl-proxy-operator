package controller

import (
	"github.com/vliubezny/ssl-proxy-operator/pkg/controller/sslproxy"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, sslproxy.Add)
}
