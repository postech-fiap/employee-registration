package publisher

import (
	"context"
	"encoding/json"
	"github.com/postech-fiap/employee-registration/internal/adapter/queue/publisher/mapper"
	"github.com/postech-fiap/employee-registration/internal/core/domain"
	amqp "github.com/rabbitmq/amqp091-go"
)

const newRegistryQueueName string = "new-registry"

type registerQueuePublisher struct {
	channel *amqp.Channel
}

func NewRegisterQueuePublisher(channel *amqp.Channel) *registerQueuePublisher {
	r := registerQueuePublisher{
		channel: channel,
	}

	r.createQueue()

	return &r
}

func (o *registerQueuePublisher) PublishRegistry(register domain.Register) error {
	dto := mapper.DomainToRegisterNewMessage(register)

	body, err := json.Marshal(dto)
	if err != nil {
		return err
	}

	return o.channel.PublishWithContext(context.Background(),
		"",
		newRegistryQueueName,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})
}

func (o *registerQueuePublisher) createQueue() {
	_, err := o.channel.QueueDeclare(newRegistryQueueName, false, false, false, false, nil)
	if err != nil {
		panic(err)
	}
}
