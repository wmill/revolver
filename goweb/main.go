package main

import (
	"context"
	"fmt"
	"log"
	"os"
	user_v1 "proto/gen/go/user/v1"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

func main() {
	app := fiber.New()
	err := godotenv.Load()

	if err != nil {
    log.Print("Error loading .env file")
  }

	LoadConfig()
	
	fmt.Println("loaded env var:" + os.Getenv("USER_SERV_GRPC"))

	api := app.Group("/api")
	v1 := api.Group("/v1")
	
	v1.Post("/login", func(c *fiber.Ctx) error {
		email := c.FormValue("email")
		password := c.FormValue("password")
		token, err := loginUser(email, password)
		
		if err != nil {
			st, ok := status.FromError(err)
			if ok && st.Code() == codes.PermissionDenied {
				return c.Status(403).JSON(fiber.Map{"error": "Invalid email or password"})
			} else {
				return c.Status(500).JSON(fiber.Map{"error": err.Error()})
			}
			
		}
		return c.JSON(fiber.Map{"token": token})
	})


	startRequireJWT(app)

	// routes after this middleware require a valid jwt



	app.Listen(":" + GetConfig().Port)
}


func loginUser(email string, password string) (string, error) {
	config := GetConfig()

	conn, err := grpc.Dial(config.UserServGrpc, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return "", err
	}
	defer conn.Close()
	c := user_v1.NewUserServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.PasswordLogin(ctx, &user_v1.PasswordLoginRequest{Email: email, Password: password})
	if err != nil {
		return "", err
	}

	claims := jwt.MapClaims{
		"email": r.GetEmail(),
		"admin": r.GetAdmin(),
		"userId": r.GetUserId(),
		"exp": time.Now().Add(config.JwtExpiration).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(config.JwtSecret))
	if err != nil {
			return "", err
	}


	return t, nil
}