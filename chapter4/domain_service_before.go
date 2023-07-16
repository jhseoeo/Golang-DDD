package chapter4

import "github.com/google/uuid"

type Product struct {
	Id             uuid.UUID
	InStock        bool
	InSomeonesCart bool
}

func (p Product) CanBeBought() bool {
	return p.InStock && !p.InSomeonesCart
}

type ShoppingCart struct {
	Id          uuid.UUID
	Products    []Product
	IsFull      bool
	MaxCartSize int
}

func (s *ShoppingCart) AddToCard(p Product) bool {
	if s.IsFull {
		return false
	}
	if p.CanBeBought() {
		s.Products = append(s.Products, p)
		return true
	}
	if s.MaxCartSize == len(s.Products) {
		s.IsFull = true
	}
	return true
}
