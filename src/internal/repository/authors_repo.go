package repository

import (
	"main/internal/model"

	"github.com/jmoiron/sqlx"
)

type AuthorRepository interface {
	GetAll() ([]model.Authors, error)
	GetById(id string) ([]model.Authors, error)
	Create(author *model.CreateAuthor) error
	Replace(id string, author *model.UpdateAuthor) (int64, error)
	Update(id string, author *model.UpdateAuthor) (int64, error)
	Delete(id string) (int64, error)
}

type authorRepository struct {
	db *sqlx.DB
}

func NewAuthorRepository(db *sqlx.DB) AuthorRepository {
	return &authorRepository{db: db}
}

const (
	queryGetAll  = `SELECT * FROM authors`
	queryGetByID = `SELECT * FROM authors WHERE author_id = $1`
	queryCreate  = `INSERT INTO authors(first_name, last_name) VALUES($1, $2)`
	queryReplace = `UPDATE authors SET first_name = $1, last_name = $2 WHERE author_id = $3`
	queryUpdate  = `UPDATE authors SET first_name = COALESCE($1, first_name), last_name = COALESCE($2, last_name) WHERE author_id = $3`
	queryDelete  = `DELETE FROM authors WHERE author_id = $1`
)

func (r *authorRepository) GetAll() ([]model.Authors, error) {
	var authors []model.Authors
	err := r.db.Select(&authors, queryGetAll)
	return authors, err
}

func (r *authorRepository) GetById(id string) ([]model.Authors, error) {
	var author []model.Authors
	err := r.db.Select(&author, queryGetByID, id)
	return author, err
}

func (r *authorRepository) Create(req *model.CreateAuthor) error {
	_, err := r.db.Exec(queryCreate, req.FirstName, req.LastName)
	return err
}

func (r *authorRepository) Replace(id string, req *model.UpdateAuthor) (int64, error) {
	res, err := r.db.Exec(queryReplace, req.FirstName, req.LastName, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func (r *authorRepository) Update(id string, req *model.UpdateAuthor) (int64, error) {
	res, err := r.db.Exec(queryUpdate, req.FirstName, req.LastName, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func (r *authorRepository) Delete(id string) (int64, error) {
	res, err := r.db.Exec(queryDelete, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}
