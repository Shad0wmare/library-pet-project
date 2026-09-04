package main

import (
	"database/sql"
	"main/model"

	"github.com/gofiber/fiber/v3"
	"github.com/jmoiron/sqlx"
)

const (
	QueryGetAuthors        = `SELECT * FROM authors`
	QueryGetAuthor         = `SELECT * FROM authors WHERE author_id = $1`
	QueryAddAuthor         = `INSERT INTO authors(first_name, last_name) VALUES($1, $2)`
	QueryReplaceAuthorInfo = `UPDATE authors SET first_name = $1, last_name = $2 WHERE author_id = $3`
	QueryEditAuthorInfo    = `UPDATE authors SET first_name = COALESCE($1, first_name), last_name = COALESCE($2, last_name) WHERE author_id = $3`
	QueryDeleteAuthor      = `DELETE FROM authors WHERE author_id = $1`
)

func GetAuthors(db *sqlx.DB) fiber.Handler {
	return func(c fiber.Ctx) error {
		var authors []model.Authors
		err := db.Select(&authors, QueryGetAuthors)
		if err != nil {
			return err
		}
		return c.JSON(authors)
	}
}

func GetAuthor(db *sqlx.DB) fiber.Handler {
	return func(c fiber.Ctx) error {
		var author []model.Authors
		id := c.Params("id")
		err := db.Select(&author, QueryGetAuthor, id)
		if err != nil {
			return err
		}
		if len(author) == 0 {
			return c.SendString("Author not found")
		}
		return c.JSON(author)
	}
}

func AddAuthor(db *sqlx.DB) fiber.Handler {
	return func(c fiber.Ctx) error {
		req := new(model.Authors)
		err := c.Bind().Body(req)
		if err != nil {
			return err
		}
		_, err = db.Exec(QueryAddAuthor, req.FirstName, req.LastName)
		if err != nil {
			return err
		}
		return c.SendString("Author added successfully")
	}
}

func ReplaceAuthorInfo(db *sqlx.DB) fiber.Handler {
	return func(c fiber.Ctx) error {
		id := c.Params("id")
		req := new(model.Authors)
		err := c.Bind().Body(req)
		if err != nil {
			return err
		}

		var res sql.Result
		if c.Method() == "PUT" {
			if req.LastName == nil || *req.LastName == "" {
				return c.SendString("LastName must not be empty")
			}
			res, err = db.Exec(QueryReplaceAuthorInfo, req.FirstName, req.LastName, id)
			if err != nil {
				return err
			}
			isEdited, _ := res.RowsAffected()
			if isEdited == 0 {
				return c.SendString("Author not found")
			}
			return c.SendString("Author info replaced successfully")
		}
		if c.Method() == "PATCH" {
			res, err = db.Exec(QueryEditAuthorInfo, req.FirstName, req.LastName, id)
			if err != nil {
				return err
			}
			isEdited, _ := res.RowsAffected()
			if isEdited == 0 {
				return c.SendString("Author not found")
			}
			return c.SendString("Author info edited successfully")
		}
		return nil
	}
}

func DeleteAuthor(db *sqlx.DB) fiber.Handler {
	return func(c fiber.Ctx) error {
		id := c.Params("id")
		res, err := db.Exec(QueryDeleteAuthor, id)
		if err != nil {
			return err
		}

		isDeleted, _ := res.RowsAffected()
		if isDeleted == 0 {
			return c.SendString("Author not found")
		}
		return c.SendString("Author deleted successfully")
	}
}
