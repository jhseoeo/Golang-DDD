package membership

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	coffeeco "github.com/jhseoeo/Golang-DDD/chapter5/internal"
	"github.com/jhseoeo/Golang-DDD/chapter5/internal/store"
)

type CoffeeBux struct {
	ID                                    uuid.UUID
	store                                 store.Store
	coffeeLover                           coffeeco.CoffeeLover
	FreeDrinksAvailable                   int
	RemainingDrinkPurchasesUntilFreeDrink int
}

func (c *CoffeeBux) AddStamp() {
	if c.RemainingDrinkPurchasesUntilFreeDrink == 1 {
		c.RemainingDrinkPurchasesUntilFreeDrink = 10
		c.FreeDrinksAvailable += 1
	} else {
		c.RemainingDrinkPurchasesUntilFreeDrink--
	}
}

func (c *CoffeeBux) Pay(ctx context.Context, products []coffeeco.Product) error {
	lp := len(products)
	if lp == 0 {
		return errors.New("nothing to buy")
	}

	if c.FreeDrinksAvailable < lp {
		return fmt.Errorf("not enough free drinks available, %d requestsed, %d available", lp, c.FreeDrinksAvailable)
	}

	c.FreeDrinksAvailable -= lp
	return nil
}
