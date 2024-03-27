package main

import (
	"fmt"
	"log"
	"time"

	"github.com/converged-computing/gomq/pkg/config"
	"github.com/converged-computing/gomq/pkg/gomq"
)

type client struct {
	cfg   config.Device
	datac chan gomq.Msg
}

func (dev *client) Configure(cfg config.Device) error {
	dev.cfg = cfg
	return nil
}

func (dev *client) Init(ctl gomq.Controller) error {
	datac, err := ctl.Chan("data", 0)
	if err != nil {
		return err
	}

	dev.datac = datac
	return nil
}

func (dev *client) Run(ctl gomq.Controller) error {

	// Give some delay for other worker to start (just a hack)
	fmt.Println("Client has started.")
	time.Sleep(10 * time.Second)

	// Run for 10 iterations
	for i := 0; i < 10; i++ {

		// Send a message to the channel
		message := fmt.Sprintf("Hello from iteration %d", i)
		fmt.Printf("Sending message %s\n", message)
		dev.datac <- gomq.Msg{Data: []byte(message)}

		// Wait for a response
		notReceived := false
		for notReceived {
			select {

			// Receive the response, and break from the loop
			case data := <-dev.datac:
				ctl.Printf("received reply: %q\n", string(data.Data))
				notReceived = false

			case <-ctl.Done():
				return nil
			}
		}
		time.Sleep(2 * time.Second)
	}
	return nil
}

func main() {
	err := gomq.Main(&client{})
	if err != nil {
		log.Fatal(err)
	}
}
