package main

import (
	"log"
	"main/internal/db"
	"main/internal/handler"
	"main/internal/repository"
	"main/internal/service"

	"github.com/gofiber/fiber/v3"
)

func main() {
	database, err := db.NewPostgresConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	authorRepo := repository.NewAuthorRepository(database)
	authorService := service.NewAuthorService(authorRepo)
	authorHandler := handler.NewAuthorHandler(authorService)

	app := fiber.New()
	library := app.Group("/library")

	authorsGroup := library.Group("/authors")
	authorsGroup.Get("/", authorHandler.GetAuthors)
	authorsGroup.Get("/:id", authorHandler.GetAuthor)
	authorsGroup.Post("/", authorHandler.AddAuthor)
	authorsGroup.Put("/:id", authorHandler.ReplaceAuthorInfo)
	authorsGroup.Patch("/:id", authorHandler.EditAuthorInfo)
	authorsGroup.Delete("/:id", authorHandler.DeleteAuthor)

	log.Fatal(app.Listen(":3000"))
}
