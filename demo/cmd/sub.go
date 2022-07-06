package main

import (
	"demo"
	"fmt"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {
	context, _ := demo.InitialNatServer()
	_, err := context.Subscribe("demo.data", receiveMessage)
	if err != nil {
		fmt.Printf("Can not subscribe from NAT server")
	}
	fmt.Println("Done")
	time.Sleep( 5 * time.Second)
}

func receiveMessage(m *nats.Msg) {
	err := m.Ack()
	if err != nil {
		fmt.Errorf("Can not send ack to NAT server")
	}
	fmt.Printf("Got data =%s\n", m.Data)
}
