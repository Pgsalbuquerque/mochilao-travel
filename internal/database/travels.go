package database

import (
	"context"
	"mochilao-travel/internal/types"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type TravelsClient struct {
	travelsCollection *mongo.Collection
}

func NewTravelsClient(mongoConn *Mongo) *TravelsClient {
	return &TravelsClient{
		travelsCollection: mongoConn.Collection(TravelsCollection),
	}
}

func (t *TravelsClient) InsertTravel(travel types.Travel) (*types.Travel, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 40*time.Second)
	defer cancel()
	_, err := t.travelsCollection.InsertOne(ctx, travel)
	if err != nil {
		return nil, err
	}

	return &travel, nil
}
