package main

import (
	"log"
	"time"

	"github.com/converged-computing/gomq/pkg/gomq"

	"github.com/converged-computing/gomq/pkg/config"
)

type processor struct {
	cfg config.Device

	// input and output data message channels
	idatac chan gomq.Msg
	odatac chan gomq.Msg
}

func (dev *processor) Configure(cfg config.Device) error {
	dev.cfg = cfg
	return nil
}

func (dev *processor) Init(ctl gomq.Controller) error {
	idatac, err := ctl.Chan("data1", 0)
	if err != nil {
		return err
	}

	odatac, err := ctl.Chan("data2", 0)
	if err != nil {
		return err
	}

	dev.idatac = idatac
	dev.odatac = odatac
	return nil
}

func (dev *processor) Run(ctl gomq.Controller) error {
	for {
		select {
		case data := <-dev.idatac:
			// ctl.Printf("received: %q\n", string(data.Data))
			out := append([]byte(nil), data.Data...)
			out = append(out, []byte(" (modified by "+dev.cfg.Name()+")"+time.Now().String())...)
			dev.odatac <- gomq.Msg{Data: out}
		case <-ctl.Done():
			return nil
		}
	}
}

func main() {
	err := gomq.Main(&processor{})
	if err != nil {
		log.Fatal(err)
	}
}
