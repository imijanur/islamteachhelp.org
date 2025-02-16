package handlers

import (
	"html/template"
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/imijanur/islamteachhelps.com/internal/models"
)

func renderTemplate(c fiber.Ctx, pageName string, page models.Page) error {
	tmpl := template.New("base")

	// Parse the base layout and partial templates
	tmpl, err := tmpl.ParseFiles(
		"internal/templates/layouts/base.html",
		"internal/templates/partials/header.html",
		"internal/templates/partials/navbar.html",
		"internal/templates/partials/footer.html",
	)
	if err != nil {
		// Handle the error
		log.Println(err)
		return err
	}

	// Parse the specific page template
	tmpl, err = tmpl.ParseFiles("internal/templates/pages/" + pageName + ".html")
	if err != nil {
		// Handle the error
		log.Println(err)
		return err
	}

	// Set the Content-Type header
	c.Set("Content-Type", "text/html")

	// Execute the template
	return tmpl.ExecuteTemplate(c.Response().BodyWriter(), "base", page)
}
