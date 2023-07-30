//go:build wireinject
// +build wireinject

package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/wire"
	"github.com/jhseoeo/Golang-DDD/chapter6/recommendation/internal/recommendation"
	"net/http"
)

func newHttpClient() (*http.Client, error) {
	return &http.Client{}, nil
}

func InitializeRecommendationHandler(url string, app *fiber.App) (*recommendation.Handler, error) {

	wire.Build(
		wire.Bind(new(recommendation.AvailabilityGetter), new(*recommendation.PartnershipAdapter)),
		newHttpClient,
		recommendation.NewPartnershipAdapter,
		recommendation.NewService,
		recommendation.NewHandler,
	)

	return nil, nil
}
