package jobs

import (
	"bytes"
	"encoding/json"
	"log"
	"mochilao-travel/internal/rabbit"
	"mochilao-travel/internal/types"
)

type TravelDB interface {
	FindTenant(types.Rental) (*types.Travel, error)
}

func FoundTenant(rabbit *rabbit.RabbitMq, travelDB TravelDB) {
	msgs, err := rabbit.Channel.Consume(rabbit.NewRentalQueue.Name, "", true, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
		return
	}

	var forever chan struct{}

	go func() {
		for d := range msgs {
			var rental types.Rental
			if err := json.Unmarshal(d.Body, &rental); err != nil {
				log.Fatal(err)
				panic(err)
			}

			travel, err := travelDB.FindTenant(rental)

			if err != nil && err.Error() != "mongo: no documents in result" {
				break
			}
			if err == nil {
				rentalWithEmail := types.RentalWithEmail{
					Fields:           rental.Fields,
					DestinationEmail: travel.Email,
					Geometry:         rental.Geometry,
				}
				reqBodyBytes := new(bytes.Buffer)
				json.NewEncoder(reqBodyBytes).Encode(rentalWithEmail)
				rabbit.Publish(reqBodyBytes.Bytes())
			}
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
