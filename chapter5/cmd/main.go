package main

import (
	"context"
	"github.com/jhseoeo/Golang-DDD/chapter5/internal/payment"
	"github.com/jhseoeo/Golang-DDD/chapter5/internal/purchase"
	"github.com/jhseoeo/Golang-DDD/chapter5/internal/store"
	"log"
)

const stripeTestApiKey = "qweqwe"
const mongoConnString = "mongodb://localhost:27017"

func wireUp() *purchase.Service {
	ctx := context.Background()

	csvc, err := payment.NewStripeService(stripeTestApiKey)
	if err != nil {
		log.Fatal(err)
	}

	prepo, err := purchase.NewMongoRepo(ctx, mongoConnString)
	if err != nil {
		log.Fatal(err)
	}

	srepo, err := store.NewMongoRepo(ctx, mongoConnString)
	if err != nil {
		log.Fatal(err)
	}

	sSvc := store.NewService(srepo)
	svc := purchase.NewService(csvc, prepo, sSvc)

	return svc
}

func main() {
	_ = wireUp()

	log.Println("Service is running...")
}
