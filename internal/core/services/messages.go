package services

import (
	"context"

	v1 "github.com/abdukhashimov/go_api/internal/common/gen/go/user/service/v1"
	"github.com/abdukhashimov/go_api/internal/config"
	"github.com/abdukhashimov/go_api/internal/core/ports"
)

var _ ports.MessagesService = (*MessagesService)(nil)

type MessagesService struct {
	repo ports.MessagesStore
}

func NewMessagesService(repo ports.MessagesStore, cfg *config.Config) *MessagesService {
	return &MessagesService{
		repo: repo,
	}
}

func (m *MessagesService) Create(ctx context.Context, input *v1.CreateMessageRequest) (*v1.Message, error) {

	return &v1.Message{
		Name: "name",
	}, nil
}
