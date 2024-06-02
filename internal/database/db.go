package database

import (
    "github.com/jmoiron/sqlx"
    "github.com/lib/pq"
)

func Connect(dsn string) (*sqlx.DB, error) {
    return sqlx.Connect("postgres", dsn)
}
