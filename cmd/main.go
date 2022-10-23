package main

import (
	"context"
	"mochilao-travel/internal/config"
	"mochilao-travel/internal/database"
	"mochilao-travel/internal/grpc"
	"mochilao-travel/internal/travel"
)

func main() {
	connectionMongoDB, err := RegisterMongo()
	if err != nil {
		return
	}
	travelDB := database.NewTravelsClient(connectionMongoDB)
	travel := travel.NewTravel(travelDB)
	go grpc.StartGrpcServer(travel)
	grpc.StartGrpcGateway()
}

func RegisterMongo() (mongo *database.Mongo, err error) {
	mongo = &database.Mongo{ConnectionString: config.Get().MongoConnectionString, DatabaseName: config.Get().DBName}
	// Make connection
	err = mongo.Connect(context.Background())

	return
}
