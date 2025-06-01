package postgres

import (
	"fmt"
	"github.com/g-stro/tech-task/pkg/database"
	_ "github.com/lib/pq"
	"os"
)

type DB struct {
	conn *database.Connection
}

// NewPostgresDB creates a database connection returned as a DB struct
func NewPostgresDB() (*DB, error) {
	dsn := getDSN()

	// create connection
	conn, err := database.NewConnection("postgres", dsn)
	if err != nil {
		return nil, err
	}

	// configs
	conn.DB.SetMaxOpenConns(25)
	conn.DB.SetMaxIdleConns(5)

	return &DB{conn: conn}, nil
}

func getDSN() string {
	return fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=%s timezone=%s",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSL_MODE"),
		os.Getenv("DB_TIMEZONE"))
}

func (db *DB) Close() {
	if db.conn != nil {
		db.conn.Close()
	}
}
