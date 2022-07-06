package main

import (
	"demo"
	"fmt"
)

func main() {
	context, _ := demo.InitialNatServer()
	_, err := context.Publish("demo.data", []byte("Hello"))
	if err != nil {
		fmt.Printf("Can not publish a new message to NAT server")
	}
	fmt.Printf("Done")
}
