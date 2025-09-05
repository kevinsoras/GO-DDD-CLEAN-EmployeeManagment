package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

// TestPostgresConnection tries to connect and prints a message if successful
func TestPostgresConnection(dsn string) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("failed to open db: %v", err)
	}
	defer db.Close()
	if err := db.Ping(); err != nil {
		log.Fatalf("failed to ping db: %v", err)
	}
	log.Println("✅ Conexión a PostgreSQL exitosa!")
}

// NewPostgresConnection returns a *sql.DB for PostgreSQL
func NewPostgresConnection(dsn string) *sql.DB {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("failed to open db: %v", err)
	}
	if err := db.Ping(); err != nil {
		log.Fatalf("failed to ping db: %v", err)
	}
	return db
}
