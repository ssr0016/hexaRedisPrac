package repository

import (
	"errors"
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"messenger/internal/core/domain"
)

type MessengerPostgresRepository struct {
	db *gorm.DB
}

func NewMessengerPostgresRepository() *MessengerPostgresRepository {
	host := "localhost"
	port := "5432"
	user := "postgres"
	password := "sercret"
	dbname := "postgres"

	con := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", host, port, user, dbname, password)

	db, err := gorm.Open("postgres", con)
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	db.AutoMigrate(&domain.Message{})

	return &MessengerPostgresRepository{
		db: db,
	}
}

func (m *MessengerPostgresRepository) SaveMessage(message domain.Message) error {
	req := m.db.Create(&message)
	if req.RowsAffected == 0 {
		return errors.New(fmt.Sprintf("messages not saved: %v", req.Error))
	}

	return nil
}

func (m *MessengerPostgresRepository) ReadMessage(id string) (*domain.Message, error) {
	message := &domain.Message{}
	req := m.db.First(&message, "id = ?", id)
	if req.RowsAffected == 0 {
		return nil, errors.New("message not found")
	}

	return message, nil
}

func (m *MessengerPostgresRepository) ReadMessages() ([]*domain.Message, error) {
	var messages []*domain.Message
	req := m.db.Find(&messages)
	if req.Error != nil {
		return nil, errors.New(fmt.Sprintf("message not found: %v", req.Error))
	}

	return messages, nil
}
