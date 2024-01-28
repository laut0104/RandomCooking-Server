package postgresql

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"cloud.google.com/go/cloudsqlconn"
	"cloud.google.com/go/cloudsqlconn/postgres/pgxv4"
)

type Config struct {
	Addr     string
	User     string
	Password string
	DB       string
}

func New() (*sql.DB, error) {
	var err error
	if os.Getenv("ENV") == "PROD" {
		credentialFilePath := "./credentials.json"
		cleanup, err := pgxv4.RegisterDriver("cloudsql-postgres", cloudsqlconn.WithIAMAuthN(), cloudsqlconn.WithCredentialsFile(credentialFilePath))
		if err != nil {
			log.Println(err)
			return nil, err
		}
		// // call cleanup when you're done with the database connection
		defer cleanup()
	}

	connStr := fmt.Sprintf("user=%s dbname=%s password=%s host=%s sslmode=disable", os.Getenv("DB_USER"), os.Getenv("DB_NAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"))
	db, err := sql.Open(os.Getenv("DB_TYPE"), connStr)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	log.Println(db)

	return db, nil
}
