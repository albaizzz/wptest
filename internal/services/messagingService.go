package services

import (
	"github.com/wptest/internal/models"
	"github.com/wptest/internal/repositories"
)

type IMessageService interface {
	CreateKessage(device models.Device) (err error)
}

type MessageService struct {
	messageRepo repositories.MessageRepository
}

// CreateMessage this service for created message and send
func (ms *MessageService) CreateMessage(param models.Device) (err error) {
	// store message
	err = ms.messageRepo.Publish(param)
	if err != nil {
		return err
	}

	return
}
