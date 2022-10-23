package services

import (
	"github.com/abdukhashimov/go_api/internal/config"
	"github.com/abdukhashimov/go_api/internal/core/ports"
	"github.com/abdukhashimov/go_api/internal/core/repository"
)

type Services struct {
	Messages ports.MessagesService
}

func NewServices(repos *repository.Repositories, cfg *config.Config) *Services {
	messagesService := NewMessagesService(repos.Messages, cfg)

	return &Services{
		Messages: messagesService,
	}
}
