package database

import (
	"context"
	"mochilao-travel/internal/types"
	"time"

	"go.mongodb.org/mongo-driver/bson"
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

func (t *TravelsClient) FindTravel(email string) (*types.Travel, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 40*time.Second)
	defer cancel()

	cur := t.travelsCollection.FindOne(ctx, bson.M{"email": email})
	if cur.Err() != nil {
		return nil, cur.Err()
	}

	var travel types.Travel
	err := cur.Decode(&travel)
	if err != nil {
		return nil, err
	}

	return &travel, nil

}

func (t *TravelsClient) FindTenant(rental types.Rental) (*types.Travel, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 40*time.Second)
	defer cancel()

	conditions := []bson.M{
		{"firstlocation": rental.Fields.City},
		{"secondlocation": rental.Fields.City},
		{"thirdlocation": rental.Fields.City},
	}

	cur := t.travelsCollection.FindOne(ctx, bson.M{"$or": conditions})
	if cur.Err() != nil {
		return nil, cur.Err()
	}

	var travel types.Travel
	err := cur.Decode(&travel)
	if err != nil {
		return nil, err
	}

	return &travel, nil

}
