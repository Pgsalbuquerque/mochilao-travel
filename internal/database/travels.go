package database

import "go.mongodb.org/mongo-driver/mongo"

type TravelsClient struct {
	travelsCollection *mongo.Collection
}

func NewTravelsClient(mongoConn *Mongo) *TravelsClient {
	return &TravelsClient{
		travelsCollection: mongoConn.Collection(TravelsCollection),
	}
}
