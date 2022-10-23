package types

type Travel struct {
	FirstLocation  string `json:"first_location"`
	SecondLocation string `json:"second_location"`
	ThirdLocation  string `json:"third_location"`
	FirstRental    Rental `json:"first_rental"`
	SecondRental   Rental `json:"second_rental"`
	ThirdRental    Rental `json:"third_rental"`
	Email          string `json:"email"`
}

type Rental struct {
	City        string  `json:"city"`
	Country     string  `json:"country"`
	Name        string  `json:"name"`
	Summary     string  `json:"summary"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	Street      string  `json:"street"`
}
