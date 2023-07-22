package store

import (
	"context"
	"github.com/google/uuid"
	coffeeco "github.com/jhseoeo/Golang-DDD/chapter5/internal"
)

type Store struct {
	ID             uuid.UUID
	Location       string
	ProductForSale coffeeco.Product
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s Service) GetStoreSpecificDiscount(ctx context.Context, storeId uuid.UUID) (float32, error) {
	dis, err := s.repo.GetStoreDiscount(ctx, storeId)
	if err != nil {
		return 0, err
	}
	return float32(dis), nil
}
