package main

import (
	"log"
	"time"

	"github.com/TanishkBansode/rtc-api/database"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.InitDB("../data")
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	app.Post("/comments/:videoid?", func(c *fiber.Ctx) error {
		addComment(c)
		return nil
	})
	app.Get("/comments/:videoid?", func(c *fiber.Ctx) error {
		getComments(c)
		return nil
	})

	if err := app.Listen(":3000"); err != nil {
		log.Fatal(err)
	}
}

func addComment(c *fiber.Ctx) error {
	videoId := c.Params("videoId")
	commentText := c.FormValue("comment")

	err := database.AddComment(videoId, commentText)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to load comment"})
	}

	// Fetch updated comments after adding the new one
	getComments(c)
	return nil
}

func getComments(c *fiber.Ctx) error {
	videoId := c.Params("videoId")
	comments, err := database.GetComments(videoId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to load comments"})
	}

	var commentsJSON []fiber.Map
	for _, comment := range comments {
		parsedDate, err := time.Parse(time.RFC3339, comment["createdAt"])
		if err != nil {
			log.Println("Error parsing date:", err)
			parsedDate = time.Now() // Fallback if parsing fails
		}
		formattedDate := parsedDate.Format("2 Jan 2006")

		commentsJSON = append(commentsJSON, fiber.Map{
			"text":      comment["text"],
			"createdAt": formattedDate,
		})
	}

	// Return the comments as JSON
	return c.Status(fiber.StatusOK).JSON(commentsJSON)
}
