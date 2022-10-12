package repository

import (
	"context"

	"github.com/caiomp87/bots-grpc-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var BotCollection IBot

type IBot interface {
	Create(ctx context.Context, bot *models.Bot) error
	FindAll(ctx context.Context, filter bson.M, opts ...*options.FindOptions) ([]*models.Bot, error)
}

type BotHelper struct {
	collection *mongo.Collection
}

func InitBot(collection *mongo.Collection) IBot {
	return &BotHelper{
		collection,
	}
}
