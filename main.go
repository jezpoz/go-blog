package main

import (
	"errors"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	// Create a new engine
	engine := html.New("./views", ".html")

	// Or from an embedded system
	// See github.com/gofiber/embed for examples
	// engine := html.NewFileSystem(http.Dir("./views", ".html"))

	// Pass the engine to the Views
	app := fiber.New(fiber.Config{
		Views: engine,
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
			}
			err = ctx.Status(code).Render(fmt.Sprintf("%d", code), fiber.Map{
				"Error": err.Error(),
			}, "layouts/main")
			if err != nil {
				return ctx.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
			}
			return nil
		},
	})

	app.Static("/", "./static")

	app.Get("/", func(c *fiber.Ctx) error {
		// Render index
		return c.Render("index", fiber.Map{
			"Title": "Hello, World!",
		}, "layouts/main")
	})

	app.Get("/about", func(c *fiber.Ctx) error {
		// Render index within layouts/main
		return c.Render("about", fiber.Map{
			"Title": "Hello, About!",
		}, "layouts/main")
	})

	app.Get("", func(c *fiber.Ctx) error {
		return c.Render("404", fiber.Map{}, "layouts/main")
	})

	log.Fatal(app.Listen(":3000"))
}
