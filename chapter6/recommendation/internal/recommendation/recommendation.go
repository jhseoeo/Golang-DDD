package recommendation

import (
	"context"
	"errors"
	"fmt"
	"math"
	"time"
)

type Money = int

type Recommendation struct {
	TripStart time.Time
	TripEnd   time.Time
	HotelName string
	Location  string
	TripPrice Money
}

type Option struct {
	HotelName     string
	Location      string
	PricePerNight Money
}

type AvailabilityGetter interface {
	GetAvailability(ctx context.Context, tripStart time.Time, tripEnd time.Time, location string) ([]Option, error)
}

type Service struct {
	availabilityGetter AvailabilityGetter
}

func NewService(availabilityGetter AvailabilityGetter) (*Service, error) {
	if availabilityGetter == nil {
		return nil, errors.New("availabilityGetter must not be nil")
	}

	return &Service{availabilityGetter: availabilityGetter}, nil
}

func (svc *Service) Get(ctx context.Context, tripStart time.Time, tripEnd time.Time, location string, budget Money) (*Recommendation, error) {
	switch {
	case tripStart.IsZero():
		return nil, errors.New("tripStart must not be zero")
	case tripEnd.IsZero():
		return nil, errors.New("tripEnd must not be zero")
	case location == "":
		return nil, errors.New("location must not be empty")
	}

	opts, err := svc.availabilityGetter.GetAvailability(ctx, tripStart, tripEnd, location)
	if err != nil {
		return nil, fmt.Errorf("error getting availability: %w", err)
	}

	tripDuration := math.Round(tripEnd.Sub(tripStart).Hours() / 24)
	var lowestPrice = math.MaxFloat64
	var cheapestTrip *Option
	for _, opt := range opts {
		price := float64(opt.PricePerNight) * tripDuration
		if price > float64(budget) {
			continue
		}
		if price < lowestPrice {
			lowestPrice = price
			cheapestTrip = &opt
		}
	}

	if cheapestTrip == nil {
		return nil, errors.New("no trips available")
	}

	return &Recommendation{
		TripStart: tripStart,
		TripEnd:   tripEnd,
		HotelName: cheapestTrip.HotelName,
		Location:  cheapestTrip.Location,
		TripPrice: Money(lowestPrice),
	}, nil
}
