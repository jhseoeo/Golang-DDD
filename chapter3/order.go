package chapter3

import "github.com/google/uuid"

type item struct {
	name string
}

type Order struct {
	items       []item
	taxAmount   Money
	discount    Money
	paymentCard uuid.UUID
	customerId  uuid.UUID
}
