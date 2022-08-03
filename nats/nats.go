package main

import (
	"log"
	"runtime"
	"time"

	"github.com/nats-io/nats.go"
)

type NATS struct {
	Address string
	Subject string
}

func NewNATSSubject(subject string) *NATS {
	return &NATS{Subject: subject}
}

func NewNATSAddress(address string) *NATS {
	return &NATS{Address: address}
}

type Nats interface {
	Subscribe() error
	Publish() error
}

func (n *NATS) Subscribe() error {
	nc, err := nats.Connect(n.Address)
	if err != nil {
		return err
	}
	nc.Subscribe(n.Subject, func(msg *nats.Msg) {
		log.Printf("received message '%s\n", string(msg.Data)+"'")
	})

	runtime.Goexit()
	return nil
}

func (n *NATS) Publish(data []byte) error {
	nc, err := nats.Connect(n.Address)
	if err != nil {
		return err
	}
	defer nc.Close()
	nc.Publish(n.Subject, data)

	time.Sleep(30 * time.Second)
	return nil
}
