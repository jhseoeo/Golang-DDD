package chapter4

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
)

type BookingRepository interface {
	SaveBooking(ctx context.Context, booking Booking) error
	DeleteBooking(ctx context.Context, booking Booking) error
}

type PostgresRepository struct {
	connPool *pgx.Conn
}

func NewPostgresRepository(ctx context.Context, dbConnString string) (*PostgresRepository, error) {
	conn, err := pgx.Connect(ctx, dbConnString)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	defer conn.Close(ctx)
	return &PostgresRepository{connPool: conn}, nil
}

func (p PostgresRepository) SaveBooking(ctx context.Context, booking Booking) error {
	_, err := p.connPool.Exec(
		ctx,
		"INSERT INTO bookings (id, from, to, hair_dresser_id) VALUES ($1, $2, $3, $4)",
		booking.id.String(),
		booking.from.String(),
		booking.to.String(),
		booking.hairDresserId.String(),
	)

	if err != nil {
		return fmt.Errorf("failed to save booking: %w", err)
	}
	return nil
}

func (p PostgresRepository) DeleteBooking(ctx context.Context, booking Booking) error {
	_, err := p.connPool.Exec(
		ctx,
		"DELETE FROM bookings WHERE id = $1",
		booking.id,
	)

	if err != nil {
		return fmt.Errorf("failed to delete booking: %w", err)
	}
	return nil
}
