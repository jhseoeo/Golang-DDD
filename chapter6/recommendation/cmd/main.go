package main

import "github.com/gofiber/fiber/v2"

const ServiceAddr = ":3030"
const PartnershipsServiceUrl = "http://localhost:3031"

func main() {
	app := fiber.New()
	_, err := InitializeRecommendationHandler(PartnershipsServiceUrl, app)
	if err != nil {
		panic(err)
	}

	err = app.Listen(ServiceAddr)
	if err != nil {
		panic(err)
	}
}
