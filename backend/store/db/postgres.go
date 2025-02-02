package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

func NewDefaultDBConfig() DBConfig {
	return DBConfig{
		Host:     "127.0.0.1",
		Port:     "6432",
		User:     "postgres",
		Password: "postgres",
		DBName:   "postgres",
	}
}

func NewPostgresDB(config DBConfig) (*sql.DB, error) {

	db, err := sql.Open(
		"postgres",
		fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Shanghai",
			config.Host,
			config.Port,
			config.User,
			config.Password,
			config.DBName,
		),
	)

	return db, err
}
