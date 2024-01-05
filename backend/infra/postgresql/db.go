package postgresql

import (
	"database/sql"
	"fmt"
	"os"
)

type Config struct {
	Addr     string
	User     string
	Password string
	DB       string
}

func New() (*sql.DB, error) {
	var err error
	connStr := fmt.Sprintf("user=%s dbname=%s password=%s host=postgres sslmode=disable", os.Getenv("DB_USER"), os.Getenv("DB_NAME"), os.Getenv("DB_PASSWORD"))
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	return db, nil
}
