package chapter8

import (
	"context"
	"errors"
)

type EmailSender interface {
	SendEmailReceipt(ctx context.Context, email string) error
}

type CardCharger interface {
	ChargeCard(ctx context.Context, cardToken string, amountInCent int) error
}

type CookieStockChecker interface {
	AmountInStock(ctx context.Context) int
}

type CookieService struct {
	emailSender        EmailSender
	cardCharger        CardCharger
	cookieStockChecker CookieStockChecker
}

func NewCookieService(e EmailSender, c CardCharger, a CookieStockChecker) (*CookieService, error) {
	return &CookieService{
		emailSender:        e,
		cardCharger:        c,
		cookieStockChecker: a,
	}, nil
}

func (c *CookieService) PurchaseCookies(ctx context.Context, amountOfCookies int, cardToken string, email string) error {
	const priceOfCookie = 50

	cookiesInStock := c.cookieStockChecker.AmountInStock(ctx)
	if cookiesInStock == 0 {
		return errors.New("no cookies in stock")
	}
	if cookiesInStock < amountOfCookies {
		amountOfCookies = cookiesInStock
	}

	err := ValidateList(
		NewCardTokenValidator(cardToken),
		NewEmailValidator(email),
		NewEmailSuffixValidator(email),
	)
	if err != nil {
		return err
	}

	cost := priceOfCookie * amountOfCookies
	err = c.cardCharger.ChargeCard(ctx, cardToken, cost)
	if err != nil {
		return err
	}

	err = c.emailSender.SendEmailReceipt(ctx, email)
	if err != nil {
		return err
	}

	return nil
}
