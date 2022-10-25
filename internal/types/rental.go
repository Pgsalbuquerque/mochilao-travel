package types

type Rental struct {
	Fields   Fields   `json:"fields"`
	Geometry Geometry `json:"geometry"`
}

type Fields struct {
	City        string  `json:"city"`
	Country     string  `json:"country"`
	Name        string  `json:"name"`
	Summary     string  `json:"summary"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	Street      string  `json:"street"`
}

type Geometry struct {
	Coordinates []float64 `json:"coordinates"`
}

type RentalWithEmail struct {
	Fields           Fields   `json:"fields"`
	DestinationEmail string   `json:"destinationEmail"`
	Geometry         Geometry `json:"geometry"`
}
