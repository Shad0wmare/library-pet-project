package handler

import (
	"main/internal/model"
	"main/internal/service"
	"net/http"

	"github.com/gofiber/fiber/v3"
)

type AuthorHandler struct {
	service service.AuthorService
}

func NewAuthorHandler(service service.AuthorService) *AuthorHandler {
	return &AuthorHandler{service: service}
}

func (h *AuthorHandler) GetAuthors(c fiber.Ctx) error {
	authors, err := h.service.GetAll()
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString(err.Error())
	}
	return c.JSON(authors)
}

func (h *AuthorHandler) GetAuthor(c fiber.Ctx) error {
	id := c.Params("id")
	authors, err := h.service.GetByID(id)
	if err != nil {
		if err.Error() == "author not found" {
			return c.Status(http.StatusNotFound).SendString("Author not found")
		}
		return c.Status(http.StatusInternalServerError).SendString(err.Error())
	}
	return c.JSON(authors)
}

func (h *AuthorHandler) AddAuthor(c fiber.Ctx) error {
	var req model.CreateAuthor
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(http.StatusBadRequest).SendString("Invalid request body")
	}

	if err := h.service.Create(&req); err != nil {
		if err.Error() == "last name is required" {
			return c.Status(http.StatusBadRequest).SendString(err.Error())
		}
		return c.Status(http.StatusInternalServerError).SendString(err.Error())
	}
	return c.Status(http.StatusCreated).SendString("Author added successfully")
}

func (h *AuthorHandler) ReplaceAuthorInfo(c fiber.Ctx) error {
	id := c.Params("id")
	var req model.UpdateAuthor
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(http.StatusBadRequest).SendString("Invalid request body")
	}

	if err := h.service.Replace(id, &req); err != nil {
		if err.Error() == "author not found" {
			return c.Status(http.StatusNotFound).SendString(err.Error())
		}
		if err.Error() == "last name is required for replace" {
			return c.Status(http.StatusBadRequest).SendString(err.Error())
		}
		return c.Status(http.StatusInternalServerError).SendString(err.Error())
	}
	return c.SendString("Author info replaced successfully")
}

func (h *AuthorHandler) EditAuthorInfo(c fiber.Ctx) error {
	id := c.Params("id")
	var req model.UpdateAuthor
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(http.StatusBadRequest).SendString("Invalid request body")
	}

	if err := h.service.Update(id, &req); err != nil {
		if err.Error() == "author not found" {
			return c.Status(http.StatusNotFound).SendString(err.Error())
		}
		return c.Status(http.StatusInternalServerError).SendString(err.Error())
	}
	return c.SendString("Author info edited successfully")
}

func (h *AuthorHandler) DeleteAuthor(c fiber.Ctx) error {
	id := c.Params("id")
	if err := h.service.Delete(id); err != nil {
		if err.Error() == "author not found" {
			return c.Status(http.StatusNotFound).SendString(err.Error())
		}
		return c.Status(http.StatusInternalServerError).SendString(err.Error())
	}
	return c.SendString("Author deleted successfully")
}
