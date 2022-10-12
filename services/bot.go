package services

import (
	"context"
	"time"

	"github.com/caiomp87/bots-grpc-api/models"
	"github.com/caiomp87/bots-grpc-api/pb"
	"github.com/caiomp87/bots-grpc-api/repository"
	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *Server) CreateBot(ctx context.Context, request *pb.CreateBotRequest) (*pb.CreateBotResponse, error) {
	err := repository.BotCollection.Create(ctx, &models.Bot{
		Name:      request.Name,
		Strategy:  request.Strategy,
		UserID:    request.UserID,
		Squad:     request.Squad,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	if err != nil {
		return nil, err
	}

	return &pb.CreateBotResponse{
		Message: "bot created successfully",
	}, nil
}

func (s *Server) FindBots(ctx context.Context, request *pb.Empty) (*pb.FindBotsResponse, error) {
	bots, err := repository.BotCollection.FindAll(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	var botstoReturn []*pb.Bot
	for _, bot := range bots {
		botstoReturn = append(botstoReturn, &pb.Bot{
			UserID:    bot.UserID,
			Squad:     bot.Squad,
			Strategy:  bot.Strategy,
			Name:      bot.Name,
			CreatedAt: timestamppb.New(bot.CreatedAt),
			UpdatedAt: timestamppb.New(bot.UpdatedAt),
		})
	}

	return &pb.FindBotsResponse{
		Bots: botstoReturn,
	}, nil
}
