package recommendation

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"time"
)

type Handler struct {
	svc *Service
	app *fiber.App
}

func NewHandler(svc *Service, app *fiber.App) (*Handler, error) {
	if svc == nil {
		return nil, errors.New("service must not be empty")
	}
	if app == nil {
		return nil, errors.New("fiber app must not be nil")
	}

	h := &Handler{svc: svc, app: app}

	app.Get("/recommendation", h.getRecommendation)

	return h, nil
}

type GetRecommendationResponse struct {
	HotelName  string `json:"hotelName"`
	TotalConst struct {
		Cost     int64  `json:"cost"`
		Currency string `json:"currency"`
	} `json:"totalCost"`
}

func (h Handler) getRecommendation(ctx *fiber.Ctx) error {
	location := ctx.Query("location")
	_from := ctx.Query("from")
	_to := ctx.Query("to")
	_budget := ctx.Query("budget")
	if location == "" || _from == "" || _to == "" || _budget == "" {
		return ctx.Status(400).JSON(fiber.Map{
			"error": "missing required query parameter",
		})
	}

	const expectedFormat = "2006-01-02"
	from, err := time.Parse(expectedFormat, _from)
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"error": "invalid from date",
		})
	}
	to, err := time.Parse(expectedFormat, _to)
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"error": "invalid to date",
		})
	}
	budget, err := strconv.ParseInt(_budget, 10, 64)
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"error": "invalid budget",
		})
	}

	res, err := h.svc.Get(ctx.UserContext(), from, to, location, Money(budget))
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(200).JSON(GetRecommendationResponse{
		HotelName: res.HotelName,
		TotalConst: struct {
			Cost     int64  `json:"cost"`
			Currency string `json:"currency"`
		}{
			Cost:     int64(res.TripPrice),
			Currency: "USD",
		},
	})
}
