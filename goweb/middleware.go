package main

import (
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func initMiddleware(app *fiber.App) {
	app.Use(cors.New())
}

func startRequireJWT(app *fiber.App) {
	config := GetConfig()
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: config.JwtSecret},
	}))
}