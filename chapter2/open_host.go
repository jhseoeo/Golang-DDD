package chapter2

import (
	"context"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type UserHandler interface {
	IsUserPaymentActive(ctx context.Context, userID string) bool
}

type UserActiveResponse struct {
	IsActive bool
}

func router(u UserHandler, app *fiber.App) {
	app.Get("/user/{userID}/payment/active", func(ctx *fiber.Ctx) error {
		uID := ctx.Query("userID")
		if uID == "" {
			return ctx.SendStatus(http.StatusBadRequest)
		}
		isActive := u.IsUserPaymentActive(ctx.UserContext(), uID)

		b, err := json.Marshal(UserActiveResponse{IsActive: isActive})
		if err != nil {
			return ctx.SendStatus(http.StatusInternalServerError)
		}

		return ctx.Send(b)
	})
}
