package repository

import (
	"github.com/abdukhashimov/go_api/internal/config"
	"github.com/abdukhashimov/go_api/internal/core/ports"
	"github.com/abdukhashimov/go_api/internal/core/repository/mongodb"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repasitories struct {
	Messages ports.MessagesStore
}

func NewRepositories(cfg *config.Config) *Repasitories {
	db := mongo.Database{}

	return &Repasitories{
		Messages: mongodb.NewMessagesRepo(&db),
	}
}
