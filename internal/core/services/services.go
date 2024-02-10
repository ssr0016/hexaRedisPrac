package services

import (
	"github.com/google/uuid"

	"messenger/internal/core/domain"
	"messenger/internal/core/ports"
)

type MessengerService struct {
	repo ports.MessengerRepository
}

func NewMessengerService(repo ports.MessengerRepository) *MessengerService {
	return &MessengerService{repo: repo}
}

func (m *MessengerService) SaveMessage(message domain.Message) error {
	message.ID = uuid.New().String()
	return m.repo.SaveMessage(message)
}

func (m *MessengerService) ReadMessage(id string) (*domain.Message, error) {
	return m.repo.ReadMessage(id)
}

func (m *MessengerService) ReadMessages() ([]*domain.Message, error) {
	return m.repo.ReadMessages()
}
