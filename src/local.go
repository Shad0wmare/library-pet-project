package main

import (
	"log"
	"main/config"

	"github.com/gofiber/fiber/v3"
)

func main() {
	db, err := config.OpenConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	app := fiber.New()
	app.Get("/authors", GetAuthors(db))
	app.Get("/author/:id", GetAuthor(db))
	app.Post("/authors", AddAuthor(db))
	app.Delete("/author/:id", DeleteAuthor(db))
	app.Put("author/:id", ReplaceAuthorInfo(db))
	app.Patch("author/:id", ReplaceAuthorInfo(db))

	log.Fatal(app.Listen(":3000"))
}
