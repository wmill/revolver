package main

import (
	"log"
	"os"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

func installMiddleware(app *fiber.App) {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		log.Fatal("JWT_SECRET is not set")
	
	}
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: secret},
	}))
}