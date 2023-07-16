package chapter4

import "errors"

type CheckOutService struct {
	shoppingCart *ShoppingCart
}

func NewCheckOutService(shoppingCart *ShoppingCart) *CheckOutService {
	return &CheckOutService{shoppingCart: shoppingCart}
}

func (c *CheckOutService) AddProductToCart(p *Product) error {
	if c.shoppingCart.IsFull {
		return errors.New("cannot add product to full cart")
	}
	if p.CanBeBought() {
		c.shoppingCart.Products = append(c.shoppingCart.Products, *p)
		return nil
	}
	if c.shoppingCart.MaxCartSize == len(c.shoppingCart.Products) {
		c.shoppingCart.IsFull = true
	}

	return nil
}
