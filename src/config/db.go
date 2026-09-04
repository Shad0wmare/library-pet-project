package config

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func OpenConnection() (*sqlx.DB, error) {
	Db, err := sqlx.Open("postgres", "user=postgres dbname=pet1 password=postgres sslmode=disable")
	return Db, err
}
