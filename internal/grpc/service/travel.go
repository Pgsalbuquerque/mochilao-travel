package grpcservice

import (
	"context"
	gen "mochilao-travel/internal/grpc/gen/go"
	"mochilao-travel/internal/types"
)

type Travels interface {
	CreateTravel(firstLocation, secondLocation, thirdLocation, email string) (*types.Travel, error)
	FindTravel(email string) (*types.Travel, error)
}

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
	firstLocation := request.GetFirstLocation()
	secondLocation := request.GetSecondLocation()
	thirdLocation := request.GetThirdLocation()
	email := request.GetEmail()

	travel, err := server.travels.CreateTravel(firstLocation, secondLocation, thirdLocation, email)
	if err != nil {
		return nil, err
	}
	return &gen.TravelResponse{
		FirstLocation:  travel.FirstLocation,
		SecondLocation: travel.SecondLocation,
		ThirdLocation:  travel.ThirdLocation,
		Email:          travel.Email,
	}, nil

}

func (server *TravelServer) GetTravel(ctx context.Context, request *gen.GetTravelRequest) (*gen.TravelResponse, error) {

	email := request.GetEmail()

	travel, err := server.travels.FindTravel(email)
	if err != nil {
		return nil, err
	}

	return &gen.TravelResponse{
		FirstLocation:  travel.FirstLocation,
		SecondLocation: travel.SecondLocation,
		ThirdLocation:  travel.ThirdLocation,
		Email:          travel.Email,
	}, nil

}
