package travel

import "mochilao-travel/internal/types"

type Travel struct {
	travelDb TravelDB
}

type TravelDB interface {
	InsertTravel(types.Travel) (*types.Travel, error)
	FindTravel(email string) (*types.Travel, error)
}

func NewTravel(travelDb TravelDB) *Travel {
	return &Travel{
		travelDb: travelDb,
	}
}

func (t *Travel) CreateTravel(firstLocation, secondLocation, thirdLocation, email string) (*types.Travel, error) {
	travel := types.Travel{
		FirstLocation:  firstLocation,
		SecondLocation: secondLocation,
		ThirdLocation:  thirdLocation,
		FirstRental:    types.Rental{},
		SecondRental:   types.Rental{},
		ThirdRental:    types.Rental{},
		Email:          email,
	}

	result, err := t.travelDb.InsertTravel(travel)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (t *Travel) FindTravel(email string) (*types.Travel, error) {
	travel, err := t.travelDb.FindTravel(email)
	if err != nil {
		return nil, err
	}

	return travel, nil
}
