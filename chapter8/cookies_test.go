package chapter8_test

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/jhseoeo/Golang-DDD/chapter8"
	"github.com/jhseoeo/Golang-DDD/chapter8/chapter8/mocks"
)

func Test_CookiePurchases(t *testing.T) {
	t.Run(`Given a user tries to purchase a cookie and we have them in stock,
    when they tap their card, they get charged and then receive an email receipt a few moments later.`,
		func(t *testing.T) {
			ctrl := gomock.NewController(t)
			e := mocks.NewMockEmailSender(ctrl)
			c := mocks.NewMockCardCharger(ctrl)
			s := mocks.NewMockCookieStockChecker(ctrl)
			ctx := context.Background()

			const cookiesToBuy = 5
			const totalExpectedCost = 250
			const cardToken = "some-token__"
			const email = "example@gmail.com"

			cs, err := chapter8.NewCookieService(e, c, s)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			gomock.InOrder(
				s.EXPECT().AmountInStock(ctx).Times(1).Return(cookiesToBuy),
				c.EXPECT().ChargeCard(ctx, cardToken, totalExpectedCost).Times(1).Return(nil),
				e.EXPECT().SendEmailReceipt(ctx, email).Times(1).Return(nil),
			)
			err = cs.PurchaseCookies(ctx, cookiesToBuy, cardToken, email)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
		})

	t.Run(`Given a user tries to purchase a cookie and we don't have any in stock, we return an error to the cashier
      so they can apologize to the customer.`,
		func(t *testing.T) {
			ctrl := gomock.NewController(t)
			e := mocks.NewMockEmailSender(ctrl)
			c := mocks.NewMockCardCharger(ctrl)
			s := mocks.NewMockCookieStockChecker(ctrl)
			ctx := context.Background()

			const cookiesToBuy = 5
			const cardToken = "some-token__"
			const email = "example@gmail.com"

			cs, err := chapter8.NewCookieService(e, c, s)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			gomock.InOrder(
				s.EXPECT().AmountInStock(ctx).Times(1).Return(0),
			)
			err = cs.PurchaseCookies(ctx, cookiesToBuy, cardToken, email)
			if err == nil {
				t.Fatalf("expected error, got nil")
			}
		})

	t.Run(`Given a user tries to purchase a cookie, we have them in stock, but their card gets declined, we return
   an error to the cashier so that we can ban the customer from the store.`,
		func(t *testing.T) {
			ctrl := gomock.NewController(t)
			e := mocks.NewMockEmailSender(ctrl)
			c := mocks.NewMockCardCharger(ctrl)
			s := mocks.NewMockCookieStockChecker(ctrl)
			ctx := context.Background()

			const cookiesToBuy = 5
			const totalExpectedCost = 250
			const cardToken = "some-token__"
			const email = "example@gmail.com"

			cs, err := chapter8.NewCookieService(e, c, s)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			gomock.InOrder(
				s.EXPECT().AmountInStock(ctx).Times(1).Return(cookiesToBuy),
				c.EXPECT().ChargeCard(ctx, cardToken, totalExpectedCost).Times(1).Return(errors.New("your card was declined")),
			)
			err = cs.PurchaseCookies(ctx, cookiesToBuy, cardToken, email)
			if err == nil {
				t.Fatalf("expected error, got nil")
			}
			if err.Error() != "your card was declined" {
				t.Fatalf("expected error, got %v", err)
			}
		})

	t.Run(`Given a user purchases a cookie and we have them in stock, their card is charged successfully but we
   fail to send an email, we return a message to the cashier so they know can notify the customer that they will not
   get an e-mail, but the transaction is still considered done.`,
		func(t *testing.T) {
			ctrl := gomock.NewController(t)
			e := mocks.NewMockEmailSender(ctrl)
			c := mocks.NewMockCardCharger(ctrl)
			s := mocks.NewMockCookieStockChecker(ctrl)
			ctx := context.Background()

			const cookiesToBuy = 5
			const totalExpectedCost = 250
			const cardToken = "some-token__"
			const email = "example@gmail.com"

			cs, err := chapter8.NewCookieService(e, c, s)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			gomock.InOrder(
				s.EXPECT().AmountInStock(ctx).Times(1).Return(cookiesToBuy),
				c.EXPECT().ChargeCard(ctx, cardToken, totalExpectedCost).Times(1).Return(nil),
				e.EXPECT().SendEmailReceipt(ctx, email).Times(1).Return(errors.New("failed to send email")),
			)
			err = cs.PurchaseCookies(ctx, cookiesToBuy, cardToken, email)
			if err == nil {
				t.Fatalf("expected error, got nil")
			}
			if err.Error() != "failed to send email" {
				t.Fatalf("expected error, got %v", err)
			}
		})

	t.Run(`Given someone wants to purchase more cookies than we have in stock we only charge them for the ones we do have`,
		func(t *testing.T) {
			ctrl := gomock.NewController(t)
			e := mocks.NewMockEmailSender(ctrl)
			c := mocks.NewMockCardCharger(ctrl)
			s := mocks.NewMockCookieStockChecker(ctrl)
			ctx := context.Background()

			const cookiesToBuy = 5
			const inStock = 3
			const totalExpectedCost = 150
			const cardToken = "some-token__"
			const email = "example@gmail.com"

			cs, err := chapter8.NewCookieService(e, c, s)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			gomock.InOrder(
				s.EXPECT().AmountInStock(ctx).Times(1).Return(inStock),
				c.EXPECT().ChargeCard(ctx, cardToken, totalExpectedCost).Times(1).Return(nil),
				e.EXPECT().SendEmailReceipt(ctx, email).Times(1).Return(nil),
			)
			err = cs.PurchaseCookies(ctx, cookiesToBuy, cardToken, email)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
		})

	t.Run(`Given card token is not 12 characters, it returns an error`, func(t *testing.T) {
		ctrl := gomock.NewController(t)
		e := mocks.NewMockEmailSender(ctrl)
		c := mocks.NewMockCardCharger(ctrl)
		s := mocks.NewMockCookieStockChecker(ctrl)
		ctx := context.Background()

		const cookiesToBuy = 5
		const cardToken = "token-too-long"
		const email = "example@gmail.com"

		cs, err := chapter8.NewCookieService(e, c, s)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		gomock.InOrder(
			s.EXPECT().AmountInStock(ctx).Times(1).Return(cookiesToBuy),
		)
		err = cs.PurchaseCookies(ctx, cookiesToBuy, cardToken, email)
		if err == nil {
			t.Fatalf("expected error, got nil")
		}
		if err.Error() != "invalid card token" {
			t.Fatalf("expected error, got %v", err)
		}
	})

	t.Run(`Given email is not a valid format, it returns an error`, func(t *testing.T) {
		ctrl := gomock.NewController(t)
		e := mocks.NewMockEmailSender(ctrl)
		c := mocks.NewMockCardCharger(ctrl)
		s := mocks.NewMockCookieStockChecker(ctrl)
		ctx := context.Background()

		const cookiesToBuy = 5
		const cardToken = "some-token__"
		const email = "examplegmail.com"

		cs, err := chapter8.NewCookieService(e, c, s)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		gomock.InOrder(
			s.EXPECT().AmountInStock(ctx).Times(1).Return(cookiesToBuy),
		)
		err = cs.PurchaseCookies(ctx, cookiesToBuy, cardToken, email)
		if err == nil {
			t.Fatalf("expected error, got nil")
		}
		if err.Error() != "invalid email format" {
			t.Fatalf("expected error, got %v", err)
		}
	})

	t.Run(`Given email's suffix is not a one of gmail, yahoo, msn, it returns an error`, func(t *testing.T) {
		ctrl := gomock.NewController(t)
		e := mocks.NewMockEmailSender(ctrl)
		c := mocks.NewMockCardCharger(ctrl)
		s := mocks.NewMockCookieStockChecker(ctrl)
		ctx := context.Background()

		const cookiesToBuy = 5
		const cardToken = "some-token__"
		const email = "example@naver.com"

		cs, err := chapter8.NewCookieService(e, c, s)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		gomock.InOrder(
			s.EXPECT().AmountInStock(ctx).Times(1).Return(cookiesToBuy),
		)
		err = cs.PurchaseCookies(ctx, cookiesToBuy, cardToken, email)
		if err == nil {
			t.Fatalf("expected error, got nil")
		}
		if err.Error() != "only gmail, yahoo and msn are supported" {
			t.Fatalf("expected error, got %v", err)
		}
	})
}
