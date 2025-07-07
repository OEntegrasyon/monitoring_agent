package main

import (
	"log"

	"github.com/streadway/amqp"
)

type Sender struct {
	conn *amqp.Connection
	ch   *amqp.Channel
	q    amqp.Queue
}

func NewSender(url string) *Sender {
	conn, err := amqp.Dial(url)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open channel: %v", err)
	}

	q, err := ch.QueueDeclare("sysmon", true, false, false, false, nil)
	if err != nil {
		log.Fatalf("Queue declare error: %v", err)
	}

	return &Sender{conn, ch, q}
}

func (s *Sender) Send(body string) {
	err := s.ch.Publish("", s.q.Name, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(body),
	})
	if err != nil {
		log.Printf("Publish error: %v", err)
	}
}

func (s *Sender) Close() {
	s.ch.Close()
	s.conn.Close()
}
