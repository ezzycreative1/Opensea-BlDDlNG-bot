// main.go
package main

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

var jwtSecret = []byte("secret_key")

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func generateToken(username string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func main() {
	app := fiber.New()

	// Middleware
	app.Use(cors.New())
	app.Use(compress.New())
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(requestid.New())
	app.Use(limiter.New())

	// Unauthenticated route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, unauthenticated!")
	})

	// JWT Authentication Middleware
	app.Use(func(c *fiber.Ctx) error {
		token := c.Get("Authorization")[7:]
		claims := &Claims{}

		_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})

		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		}

		return c.Next()
	})

	// Authenticated route
	app.Get("/auth", func(c *fiber.Ctx) error {
		return c.SendString("Hello, authenticated!")
	})

	// Login endpoint
	app.Post("/login", func(c *fiber.Ctx) error {
		

		// Perform authentication logic (contoh sederhana)
		if username == "user123" && passwd == "password123" {
			// Jika autentikasi berhasil, hasilkan token
			token, err := generateToken(username)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"message": "Internal Server Error",
				})
			}

			return c.JSON(fiber.Map{
				"token": token,
			})
		} else {
			// Jika autentikasi gagal
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Authentication failed",
			})
		}
	})

	// Start the Fiber app
	app.Listen(":3000")
}
