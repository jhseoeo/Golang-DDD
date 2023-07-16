package chapter3

import "github.com/google/uuid"

type SomeEntity struct {
	id uuid.UUID
}

func NewSomeEntity() *SomeEntity {
	return &SomeEntity{
		id: uuid.New(),
	}
}
