package handlers

import (
	"github.com/gofiber/fiber/v3"
	"github.com/imijanur/islamteachhelps.com/internal/models"
)

func HomeHandler(c fiber.Ctx) error {
	return renderTemplate(c, "index", models.Page{
		Title: "Home Page",
	})
}

func TestHandler(c fiber.Ctx) error {
	return renderTemplate(c, "test", models.Page{
		Title: "Test Page",
	})
}
func AboutHandler(c fiber.Ctx) error {
	return renderTemplate(c, "about", models.Page{
		Title: "ABOUT &#8211; Donation",
	})
}
func ActivitiesHandler(c fiber.Ctx) error {
	return renderTemplate(c, "activities", models.Page{
		Title: "ACTIVITIES &#8211; Donation",
	})
}
func AuthorHandler(c fiber.Ctx) error {
	return renderTemplate(c, "author", models.Page{
		Title: "Author Page",
	})
}
func CategoryHandler(c fiber.Ctx) error {
	return renderTemplate(c, "category", models.Page{
		Title: "Uncategorized &#8211; Donation",
	})
}
func CommentsHandler(c fiber.Ctx) error {
	return renderTemplate(c, "test", models.Page{
		Title: "Test Page",
	})
}
func ContactHandler(c fiber.Ctx) error {
	return renderTemplate(c, "contact", models.Page{
		Title: "CONTACT &#8211; Donation",
	})
}
func DonationHandler(c fiber.Ctx) error {
	return renderTemplate(c, "donation", models.Page{
		Title: "Donation &#8211; Donation",
	})
}
func FeedHandler(c fiber.Ctx) error {
	return renderTemplate(c, "feed", models.Page{
		Title: "Feed &#8211; Donation",
	})
}
func GalleryHandler(c fiber.Ctx) error {
	return renderTemplate(c, "gallery", models.Page{
		Title: "GALLERY &#8211; Donation",
	})
}
func MadrashaHandler(c fiber.Ctx) error {
	return renderTemplate(c, "madrasha", models.Page{
		Title: "Madrasha &#8211; Donation",
	})
}
func SamplePageHandler(c fiber.Ctx) error {
	return renderTemplate(c, "sample-page", models.Page{
		Title: "Sample Page &#8211; Donation",
	})
}
func South24ParganasDistrictHandler(c fiber.Ctx) error {
	return renderTemplate(c, "south-24-parganas-district", models.Page{
		Title: "South 24 Parganas District &#8211; Donation",
	})
}

func NotFoundHandler(c fiber.Ctx) error {
	return renderTemplate(c, "notfound", models.Page{
		Title: "Not Found",
	})
}
