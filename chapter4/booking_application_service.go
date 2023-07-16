package chapter4

import (
	"context"
	"errors"
	"fmt"
	"github.com/jhseoeo/Golang-DDD/chapter2"
)

type accountKey = int

const accountCtxKey = accountKey(1)

type BookingDomainService interface {
	CreateBooking(ctx context.Context, booking Booking) error
}

type BookingAppService struct {
	bookingRepo          BookingRepository
	bookingDomainService BookingDomainService
}

func NewBookingAppService(bookingRepo BookingRepository, bookingDomainService BookingDomainService) *BookingAppService {
	return &BookingAppService{
		bookingRepo:          bookingRepo,
		bookingDomainService: bookingDomainService,
	}
}

func (b *BookingAppService) CreateBooking(ctx context.Context, booking Booking) error {
	u, ok := ctx.Value(accountCtxKey).(*chapter2.Customer)
	if !ok {
		return errors.New("invalid customer")
	}
	if u.UserID() != booking.userId.String() {
		return errors.New("cannot create booking for other users")
	}

	err := b.bookingDomainService.CreateBooking(ctx, booking)
	if err != nil {
		return fmt.Errorf("could not create booking: %w", err)
	}
	err = b.bookingRepo.SaveBooking(ctx, booking)
	if err != nil {
		return fmt.Errorf("could not save bookign: %w", err)
	}

	return nil
}
