package repositories

import (
	"encoding/json"

	"github.com/wptest/configs"
	"github.com/wptest/internal/models"
	"github.com/wptest/pkg/kafka"
)

type IMessageRepository interface {
	Publish(request models.Device) (err error)
}
type MessageRepository struct {
	kafka  kafka.Kafka
	config *configs.Config
}

func NewMessageRepository(k kafka.Kafka, cfg *configs.Config) *MessageRepository {
	return &MessageRepository{
		kafka:  k,
		config: cfg,
	}
}

func (mr *MessageRepository) Publish(request models.Device) (err error) {
	msg, _ := json.Marshal(request)
	err = mr.kafka.SendMessage(mr.config.Kafka.MessagingConsumer.Topic, []byte{}, msg)
	return
}
