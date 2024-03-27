// Copyright 2016 The fer Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"log"
	"time"

	"github.com/converged-computing/gomq/pkg/config"
	"github.com/converged-computing/gomq/pkg/gomq"
)

type sampler struct {
	cfg   config.Device
	datac chan gomq.Msg
}

func (dev *sampler) Configure(cfg config.Device) error {
	dev.cfg = cfg
	return nil
}

func (dev *sampler) Init(ctl gomq.Controller) error {
	datac, err := ctl.Chan("data", 0)
	if err != nil {
		return err
	}

	dev.datac = datac
	return nil
}

func (dev *sampler) Run(ctl gomq.Controller) error {
	for {
		select {
		case dev.datac <- gomq.Msg{Data: []byte("HELLO")}:
			ctl.Printf("sent 'HELLO' (%v)\n", time.Now())
		case <-ctl.Done():
			return nil
		}
	}
}

func main() {
	err := gomq.Main(&sampler{})
	if err != nil {
		log.Fatal(err)
	}
}
