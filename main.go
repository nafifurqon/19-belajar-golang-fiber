package main

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {
	fiber := fiber.New(fiber.Config{
		IdleTimeout:  time.Second * 5,
		ReadTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 5,
	})

	err := fiber.Listen("localhost:3000")
	if err != nil {
		panic(err)
	}
}
