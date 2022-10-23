package repository

import (
	"github.com/abdukhashimov/go_api/internal/config"
	"github.com/abdukhashimov/go_api/internal/core/ports"
	"github.com/abdukhashimov/go_api/internal/core/repository/mongodb"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repositories struct {
	Messages ports.MessagesStore
}

func NewRepositories(cfg *config.Config) *Repositories {
	db := mongo.Database{}

	return &Repositories{
		Messages: mongodb.NewMessagesRepo(&db),
	}
}
