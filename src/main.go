package main

import (
	"log"

	"github.com/Pradeep-kumar-py/fiber-mongodb/src/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Post("/submit", func(c *fiber.Ctx) error {
		// Parse JSON from request body
		var data struct {
			Name     string `json:"name"`
			Email    string `json:"email"`
			Password string `json:"password"`
		}

		// Parse JSON into struct (like destructuring)
		if err := c.BodyParser(&data); err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error": "Cannot parse JSON",
			})
		}

		// Use the fields directly
		return c.JSON(fiber.Map{
			"message": "Hello, " + data.Name + "!",
			"email":   data.Email,
			"name":    data.Name,
			// Note: Never return password in response!
		})

	})
	
	routes.UserRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
