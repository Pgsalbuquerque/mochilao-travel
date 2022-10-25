package rabbit

import (
	"context"
	"fmt"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMq struct {
	Connection       *amqp.Connection
	Channel          *amqp.Channel
	NewRentalQueue   amqp.Queue
	RentalFoundQueue amqp.Queue
}

func ConnectRabbit() (*RabbitMq, error) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	q, err := ch.QueueDeclare(
		"NEW_RENTAL", // name
		true,         // durable
		false,        // delete when unused
		false,        // exclusive
		false,        // no-wait
		nil,          // arguments
	)
	if err != nil {
		return nil, err
	}

	q2, err := ch.QueueDeclare(
		"RENTAL_FOUND", // name
		true,           // durable
		false,          // delete when unused
		false,          // exclusive
		false,          // no-wait
		nil,            // arguments
	)
	if err != nil {
		return nil, err
	}

	return &RabbitMq{
		Connection:       conn,
		Channel:          ch,
		NewRentalQueue:   q,
		RentalFoundQueue: q2,
	}, nil
}

func (rabbit *RabbitMq) Publish(body []byte) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	fmt.Println("Um rental foi enviado pra fila")
	err := rabbit.Channel.PublishWithContext(ctx,
		"",                           // exchange
		rabbit.RentalFoundQueue.Name, // routing key
		false,                        // mandatory
		false,                        // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})
	if err != nil {
		return err
	}
	return nil
}
