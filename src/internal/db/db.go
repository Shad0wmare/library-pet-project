package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewPostgresConnection() (*sqlx.DB, error) {
	Db, err := sqlx.Open("postgres", "user=postgres dbname=pet1 password=postgres sslmode=disable")
	return Db, err
}
