package main

import (
	"context"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"time"
)

func testMQConn() {
	conn, err := amqp.Dial("amqp://guest:guest@10.188.163.68:80/")
	if err != nil {
		fmt.Printf("err=%v\n", err)
		return
	} else {
		fmt.Printf("connection created.")
		defer conn.Close()
	}
}

func send() error {
	conn, err := amqp.Dial("amqp://guest:guest@10.188.163.68:5672/")
	if err != nil {
		fmt.Printf("Failed to create connection: %v\n", err)
		return err
	}
	ch, err := conn.Channel()
	if err != nil {
		fmt.Printf("Failed to create channel: %v\n", err)
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	msgDeliver := make(chan amqp.Return, 1)
	ch.NotifyReturn(msgDeliver)

	imm := false
	err = ch.PublishWithContext(ctx,
		"my.topic",
		"svc.auth",
		true, // true - undeliverable when no queue is bounded
		imm,  // true - if no consumer connect to queue
		amqp.Publishing{
			// DeliveryMode: amqp.Transient,
			// DeliveryMode: amqp.Persistent,
			// Timestamp:    time.Now(),
			ContentType: "text/plain",
			Body:        []byte(fmt.Sprintf("hello, world, immediate=%v", imm)),
		})
	if err != nil {
		fmt.Printf("sending error: %v\n", err)
		return err
	} else {
		fmt.Printf("message sent, immediate=%v\n", imm)
	}

	ch.Close()
	conn.Close()

	select {
	case d := <-msgDeliver:
		fmt.Printf("code: %d, text: %s, exch:%s, routing key: %s\n",
			d.ReplyCode, d.ReplyText, d.Exchange, d.RoutingKey)
	case <-time.After(15 * time.Second):
		fmt.Printf("running timeout")
	}

	return nil
}

// func recv() error {
// 	conn, err := amqp.Dial("amqp://guest:guest@10.188.163.68:5672/")
// 	if err != nil {
// 		fmt.Printf("Failed to create connection: %v\n", err)
// 		return err
// 	}
// 	ch, err := conn.Channel()
// 	if err != nil {
// 		fmt.Printf("Failed to create channel: %v\n", err)
// 		return err
// 	}
// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()
//
// 	q, err := ch.QueueDeclare("svc.auth",
// 		false,
// 		false,
// 		false,
// 		false,
// 		nil)
// 	if err != nil {
// 		fmt.Printf("Failed to declare queue: %v\n", err)
// 		return err
// 	}
//
// }

func main() {
	send()
}
