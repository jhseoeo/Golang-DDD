package main

import (
	"context"
	"github.com/google/uuid"
	coffeeco "github.com/jhseoeo/Golang-DDD/chapter5/internal"
	"github.com/jhseoeo/Golang-DDD/chapter5/internal/payment"
	"github.com/jhseoeo/Golang-DDD/chapter5/internal/purchase"
	"github.com/jhseoeo/Golang-DDD/chapter5/internal/store"
	"testing"
)

func Test_main(t *testing.T) {
	var cardToken = "asdasd"

	t.Log("TestMain")

	storeId := uuid.New()

	svc := wireUp()

	pur := &purchase.Purchase{
		CardToken: &cardToken,
		Store: store.Store{
			ID: storeId,
		},
		ProductsToPurchase: []coffeeco.Product{{
			ItemName:  "item1",
			BasePrice: 10,
		}},
		PaymentMeans: payment.MEANS_CARD,
	}

	err := svc.CompletePurchase(context.Background(), storeId, pur, nil)
	if err != nil {
		t.Error(err)
	}
}
