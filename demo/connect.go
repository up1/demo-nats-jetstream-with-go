package demo

import (
	"github.com/nats-io/nats.go"
)

func InitialNatServer() (nats.JetStreamContext, error) {
	n, err := nats.Connect("nats://127.0.0.1:4222")
	if err != nil {
		return nil, err
	}
	context, err := n.JetStream(nats.PublishAsyncMaxPending(256))
	if err != nil {
		return nil, err
	}

	s, err := context.StreamInfo("demo")
	if s == nil {
		// Create new stream
		_, err := context.AddStream(&nats.StreamConfig{
			Name:     "demo",
			Subjects: []string{"demo.*"},
		})
		if err != nil {
			return nil, err
		}
	}

	return context, nil
}
