package main

import (
	"topupservice/internal/api"
	"topupservice/internal/component"
	"topupservice/internal/config"
	"topupservice/internal/repository"
	"topupservice/internal/service"

	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	cnf := config.Get()
	dbConnection := component.GetDatabaseConnection(cnf)

	userRepository := repository.User(dbConnection)

	userService := service.User(userRepository)

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000, http://localhost:5173, http://localhost:8081, http://localhost:8080", // Membolehkan domain tertentu
		AllowMethods: "GET,POST,PUT,DELETE,PATCH",                                                                  // Metode HTTP yang diizinkan
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",                                                // Header yang diizinkan
	}))

	app.Use(logger.New(logger.Config{
		Format:     "${time} ${method} ${url} ${status} - ${latency} ${bytesSent}\n",
		TimeFormat: "02-Jan-2006 15:04:05",
		TimeZone:   "Asia/Jakarta",
	}))

	// Tentukan konfigurasi Swagger
	cfg := swagger.Config{
		BasePath: "/",
		FilePath: "./docs/swagger.json",
		Path:     "/swagger",
		Title:    "Swagger API Docs",
	}

	// Gunakan middleware Swagger
	app.Use(swagger.New(cfg))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Test Ping")
	})
	apiRoutes := app.Group("api")

	api.User(apiRoutes.(*fiber.Group), userService)
	// app.Get("/", func(c *fiber.Ctx) error {
	// 	return c.SendString("Hello, World!")
	// })

	app.Listen("localhost" + ":" + "8080")
}
