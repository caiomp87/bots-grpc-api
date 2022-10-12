package repository

import (
	"context"

	"github.com/caiomp87/bots-grpc-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (botHelper *BotHelper) Create(ctx context.Context, bot *models.Bot) error {
	_, err := botHelper.collection.InsertOne(ctx, bot)
	if err != nil {
		return err
	}

	return err
}
func (botHelper *BotHelper) FindAll(ctx context.Context, filter bson.M, opts ...*options.FindOptions) ([]*models.Bot, error) {
	cursor, err := botHelper.collection.Find(ctx, filter, opts...)
	if err != nil {
		return nil, err
	}

	var bots []*models.Bot
	for cursor.Next(context.Background()) {
		var bot *models.Bot
		if err := cursor.Decode(&bot); err != nil {
			return nil, err
		}

		bots = append(bots, bot)
	}

	return bots, nil
}
