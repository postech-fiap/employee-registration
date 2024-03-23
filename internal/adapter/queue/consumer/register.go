package consumer

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/postech-fiap/employee-registration/internal/adapter/queue/consumer/dto"
	"github.com/postech-fiap/employee-registration/internal/adapter/queue/consumer/mapper"
	"github.com/postech-fiap/employee-registration/internal/core/port"
	amqp "github.com/rabbitmq/amqp091-go"
)

const newRegistry = "new-registry"

type registerQueueConsumer struct {
	channel         *amqp.Channel
	registerUseCase port.RegisterUseCaseInterface
}

func NewRegisterQueueConsumer(channel *amqp.Channel, registerUseCase port.RegisterUseCaseInterface) *registerQueueConsumer {
	r := registerQueueConsumer{
		channel:         channel,
		registerUseCase: registerUseCase,
	}

	r.createQueue()

	return &r
}

func (r *registerQueueConsumer) Listen() {
	go func() {
		msgs, err := r.channel.Consume(newRegistry, "", false, false, false, false, nil)
		if err != nil {
			panic(err)
		}

		for message := range msgs {
			var registerDto dto.NewRegisterMessage
			err := json.Unmarshal(message.Body, &registerDto)
			if err != nil {
				fmt.Println(err)
				continue
			}

			validate := validator.New()
			err = validate.Struct(registerDto)
			if err != nil {
				fmt.Println(err, message.MessageId, string(message.Body))
				continue
			}

			newRegister := mapper.MapNewRegisterMessageToDomain(&registerDto)

			err = r.registerUseCase.Insert(newRegister)
			if err != nil {
				fmt.Println(err)
				continue
			}

			err = message.Ack(false)
			if err != nil {
				fmt.Println(err)
				continue
			}
		}
	}()
}

func (o *registerQueueConsumer) createQueue() {
	_, err := o.channel.QueueDeclare(newRegistry, false, false, false, false, nil)
	if err != nil {
		panic(err)
	}
}
