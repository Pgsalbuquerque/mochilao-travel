package travel

type Travel struct {
	travelDb TravelDB
}

type TravelDB interface {
}

func NewTravel(travelDb TravelDB) *Travel {
	return &Travel{
		travelDb: travelDb,
	}
}
