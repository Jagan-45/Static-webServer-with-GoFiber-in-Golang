package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World")
	})

	app.Static("/hello", "./static")
	/*app.Post("/hello", func(c *fiber.Ctx) error {
		uname := c.Params("username")
		pswd := c.Params("password")
		return c.SendString(uname + " " + pswd)
	})*/
	err := app.Listen(":3000")
	if err != nil {
		log.Fatal("Error")
	}
}
