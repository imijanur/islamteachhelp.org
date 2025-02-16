package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/recover"
	"github.com/gofiber/fiber/v3/middleware/static"
	"github.com/imijanur/islamteachhelps.com/internal/handlers"
)

func main() {
	// engine := mustache.New("./internal/templates", ".mustache")
	app := fiber.New()

	// Middleware
	app.Use(recover.New())

	// Routes
	app.Get("/", handlers.HomeHandler)
	app.Get("/test", handlers.TestHandler)
	app.Get("/about", handlers.AboutHandler)
	app.Get("/activities", handlers.ActivitiesHandler)
	app.Get("/author", handlers.AuthorHandler)
	app.Get("/category", handlers.CategoryHandler)
	app.Get("/comments", handlers.CommentsHandler)
	app.Get("/contact", handlers.ContactHandler)
	app.Get("/donation", handlers.DonationHandler)
	app.Get("/feed", handlers.FeedHandler)
	app.Get("/gallery", handlers.GalleryHandler)
	app.Get("/madrasha", handlers.MadrashaHandler)
	app.Get("/balance-sheet", handlers.BalanceSheetHandler)
	app.Get("/sample-page", handlers.SamplePageHandler)
	app.Get("/south-24-parganas-district", handlers.South24ParganasDistrictHandler)
	app.Use("/static", static.New("./static"))

	// 404 Handler
	app.Use(func(c fiber.Ctx) error {
		return handlers.NotFoundHandler(c)
	})

	// Start server
	go func() {
		if err := app.Listen(":8080"); err != nil {
			log.Fatalf("Could not listen on :8080: %v\n", err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit
	log.Println("Shutting down server...")

	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := app.Shutdown(); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exiting")
}
