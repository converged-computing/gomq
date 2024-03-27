package main

import (
	"fmt"
	"log"

	"github.com/converged-computing/gomq/pkg/config"
	"github.com/converged-computing/gomq/pkg/gomq"
)

type server struct {
	cfg   config.Device
	datac chan gomq.Msg
}

func (dev *server) Configure(cfg config.Device) error {
	dev.cfg = cfg
	return nil
}

func (dev *server) Init(ctl gomq.Controller) error {
	datac, err := ctl.Chan("data", 0)
	if err != nil {
		return err
	}

	dev.datac = datac
	return nil
}

func (dev *server) Run(ctl gomq.Controller) error {
	fmt.Println("Server has started.")

	for {
		select {
		case data := <-dev.datac:
			// This is receiving the message and replying
			ctl.Printf("received: %q\n", string(data.Data))
			dev.datac <- gomq.Msg{Data: []byte("WORLD")}
		case <-ctl.Done():
			return nil
		default:
			fmt.Println("Case not covererd")
		}
	}
}

func main() {
	err := gomq.Main(&server{})
	if err != nil {
		log.Fatal(err)
	}
}
