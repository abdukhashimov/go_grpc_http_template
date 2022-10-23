package ports

import (
	"context"

	v1 "github.com/abdukhashimov/go_api/internal/common/gen/go/user/service/v1"
)

type MessagesService interface {
	Create(ctx context.Context, input *v1.CreateMessageRequest) (*v1.Message, error)
}

type MessagesStore interface {
	Create(ctx context.Context, input *v1.CreateMessageRequest) (string, error)
}
