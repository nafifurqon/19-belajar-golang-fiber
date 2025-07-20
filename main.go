package main

import "github.com/gofiber/fiber/v2"

func main() {
	fiber := fiber.New()

	err := fiber.Listen("localhost:3000")
	if err != nil {
		panic(err)
	}
}
