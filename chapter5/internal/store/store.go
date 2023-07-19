package store

import (
	"github.com/google/uuid"
	coffeeco "github.com/jhseoeo/Golang-DDD/chapter5/internal"
)

type Store struct {
	ID             uuid.UUID
	Location       string
	ProductForSale coffeeco.Product
}
