package grpcservice

import (
	"context"
	gen "mochilao-travel/internal/grpc/gen/go"
)

type Travels interface{}

type TravelServer struct {
	gen.TravelServer
	travels Travels
}

func NewTravelServer(travels Travels) *TravelServer {
	return &TravelServer{
		travels: travels,
	}
}

func (server *TravelServer) PostTravel(ctx context.Context, request *gen.TravelRequest) (*gen.TravelResponse, error) {

	return &gen.TravelResponse{}, nil

}

func (server *TravelServer) GetTravel(ctx context.Context, request *gen.GetTravelRequest) (*gen.TravelResponse, error) {

	return &gen.TravelResponse{}, nil

}
