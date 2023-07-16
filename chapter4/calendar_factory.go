package chapter4

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

type Booking struct {
	id            uuid.UUID
	userId        uuid.UUID
	from          time.Time
	to            time.Time
	hairDresserId uuid.UUID
}

func CreateBooking(from time.Time, to time.Time, userId uuid.UUID, hairDresserId uuid.UUID) (*Booking, error) {
	closingTime, _ := time.Parse(time.Kitchen, "17:00pm")

	if from.After(closingTime) {
		return nil, errors.New("no appointments after closing time")
	} else {
		return &Booking{
			hairDresserId: hairDresserId,
			id:            uuid.New(),
			userId:        userId,
			from:          from,
			to:            to,
		}, nil
	}
}
