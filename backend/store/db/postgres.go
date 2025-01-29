package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

const Dsn = "host=14.103.71.103 port=6432 user=aimemos password=aimemos%2025 dbname=aimemos sslmode=disable TimeZone=Asia/Shanghai"

func NewPostgresDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", Dsn)

	return db, err
}
