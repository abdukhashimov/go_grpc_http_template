package mongodb

import (
	"context"

	v1 "github.com/abdukhashimov/go_api/internal/common/gen/go/user/service/v1"
	"github.com/abdukhashimov/go_api/internal/core/ports"
	"github.com/abdukhashimov/go_api/internal/pkg/logger"
	"go.mongodb.org/mongo-driver/mongo"
)

var _ ports.MessagesStore = (*MessagesRepo)(nil)

type MessagesRepo struct {
	db *mongo.Collection
}

func NewMessagesRepo(db *mongo.Database) *MessagesRepo {
	return &MessagesRepo{
		db: db.Collection(messagesCollection),
	}
}

func (m *MessagesRepo) Create(ctx context.Context, input *v1.CreateMessageRequest) (string, error) {
	logger.Log.Warn("not implemented method")
	return "", nil
}
