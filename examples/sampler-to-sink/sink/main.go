// Copyright 2016 The fer Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"log"

	"github.com/converged-computing/gomq/pkg/config"
	"github.com/converged-computing/gomq/pkg/gomq"
)

type sink struct {
	cfg   config.Device
	datac chan gomq.Msg
}

func (dev *sink) Configure(cfg config.Device) error {
	dev.cfg = cfg
	return nil
}

func (dev *sink) Init(ctl gomq.Controller) error {
	datac, err := ctl.Chan("data", 0)
	if err != nil {
		return err
	}

	dev.datac = datac
	return nil
}

func (dev *sink) Run(ctl gomq.Controller) error {
	for {
		select {
		case data := <-dev.datac:
			ctl.Printf("received: %q\n", string(data.Data))
		case <-ctl.Done():
			return nil
		}
	}
}

func main() {
	err := gomq.Main(&sink{})
	if err != nil {
		log.Fatal(err)
	}
}
