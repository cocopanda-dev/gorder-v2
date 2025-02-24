package broker

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
)

func Connect(user, password, host, port string) (*amqp.Channel, func() error, error) {
	address := fmt.Sprintf("amqp://%s:%s@%s:%s", user, password, host, port)
	conn, err := amqp.Dial(address)
	if err != nil {
		return nil, func() error { return nil }, err
	}
	ch, err := conn.Channel()
	if err != nil {
		return nil, func() error { return nil }, err
	}

	ch.ExchangeDeclare(EventOrderCreated, "direct")
}
