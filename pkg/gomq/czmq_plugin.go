//go:build czmq
// +build czmq

package gomq

import (
	_ "github.com/converged-computing/gomq/pkg/mq/czmq" // load C-bindings zeromq plugin
)
