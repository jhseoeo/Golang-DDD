package chapter4

import (
	"errors"
	"fmt"
	"log"
)

type Car interface {
	BeepBeep()
}

type BMW struct {
	heatedSeatSubscriptionEnabled bool
}

func (B BMW) BeepBeep() {
	// TODO: implement me
	panic("implement me")
}

type Tesla struct {
	autopilotEnabled bool
}

func (T Tesla) BeepBeep() {
	// TODO: implement me
	panic("implement me")
}

func BuildCar(carType string) (Car, error) {
	switch carType {
	case "bmw":
		return BMW{heatedSeatSubscriptionEnabled: true}, nil
	case "tesla":
		return Tesla{autopilotEnabled: true}, nil
	default:
		return nil, errors.New("unknown car type")
	}
}

func main() {
	myCar, err := BuildCar("tesla")
	if err != nil {
		log.Fatal(err)
	}
	// TODO: do something with myCar
	fmt.Printf("%v is my car\n", myCar)
}
