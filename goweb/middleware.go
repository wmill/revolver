package main

import (
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

func startRequireJWT(app *fiber.App) {
	config := GetConfig()
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: config.JwtSecret},
	}))
}