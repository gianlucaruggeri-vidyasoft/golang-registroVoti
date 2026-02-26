package main

import (
	"log"

	"github.com/gofiber/fiber/v3"

	"goApp/src/internal/middleware"
)

func main() {
	app := middleware.Init()

	log.Fatal(app.Listen(":3000", fiber.ListenConfig{
		DisableStartupMessage: true,
	}))
}