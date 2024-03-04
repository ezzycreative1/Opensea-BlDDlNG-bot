// main.go
package main

// @title Your GoFiber API
// @version 1.0
// @description This is a sample API using GoFiber and Swagger.
// @termsOfService https://example.com/terms
// @contact.name API Support
// @contact.url https://www.example.com/support
// @contact.email support@example.com
// @license.name MIT
// @license.url https://opensource.org/licenses/MIT
// @host localhost:3000
// @BasePath /v1
// @schemes http https
import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	//"github.com/swaggo/fiber-swagger"
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

	// Route untuk dokumentasi Swagger
	// app.Get("/swagger/*", fiberSwagger.Handler)

	// Variabel untuk menyimpan token
	var authToken string

	// Unauthenticated route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, unauthenticated!")
	})

	// JWT Authentication Middleware
	app.Use(func(c *fiber.Ctx) error {
		if authToken == "" {
			// Token belum di-generate, lanjutkan ke middleware atau route selanjutnya
			return c.Next()
		}

		// Mengambil token dari header Authorization
		tokenHeader := c.Get("Authorization")

		// Memeriksa apakah header Authorization kosong atau tidak memiliki format yang benar
		if tokenHeader == "" || len(tokenHeader) < 7 || tokenHeader[:7] != "Bearer " {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized: Invalid Authorization header format",
			})
		}

		// Mengambil token dari string setelah "Bearer "
		token := tokenHeader[7:]

		// Membuat struktur untuk menyimpan klaim JWT
		claims := &Claims{}

		// Parsing dan verifikasi token JWT
		parsedToken, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})

		// Memeriksa kesalahan saat parsing token
		if err != nil {
			fmt.Println("Error parsing token:", err)
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized: Error parsing token",
			})
		}

		// Memeriksa apakah token valid
		if !parsedToken.Valid {
			fmt.Println("Invalid token:", parsedToken)
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized: Invalid token",
			})
		}

		// Jika token valid, lanjutkan ke middleware atau route selanjutnya
		return c.Next()
	})

	// Authenticated route
	app.Get("/auth", func(c *fiber.Ctx) error {
		return c.SendString("Hello, authenticated!")
	})

	// Login endpoint
	app.Post("/login", func(c *fiber.Ctx) error {
		username := c.FormValue("username")
		passwd := c.FormValue("password")

		// Perform authentication logic (contoh sederhana)
		if username == "user123" && passwd == "password123" {
			// Jika autentikasi berhasil, hasilkan token
			token, err := generateToken(username)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"message": "Internal Server Error",
				})
			}

			// Sertakan token dalam respons JSON
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
