package main

import (
	"context"
	"log"
	"mochilao-travel/internal/config"
	"mochilao-travel/internal/database"
	"mochilao-travel/internal/grpc"
	"mochilao-travel/internal/jobs"
	"mochilao-travel/internal/rabbit"
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
	rabbit, err := rabbit.ConnectRabbit()
	if err != nil {
		log.Fatal(err)
		return
	}
	go jobs.FoundTenant(rabbit, travelDB)
	grpc.StartGrpcGateway()
}

func RegisterMongo() (mongo *database.Mongo, err error) {
	mongo = &database.Mongo{ConnectionString: config.Get().MongoConnectionString, DatabaseName: config.Get().DBName}
	// Make connection
	err = mongo.Connect(context.Background())

	return
}
