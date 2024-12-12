package data

import (
	"context"
	"database/sql"
	"gh_static_portfolio/cmd/data/database"
	"log"

	_ "embed"
)

//go:embed database/schema.sql
var DDL string

func InitDB() (*database.Queries, *sql.DB, error) {
	var queries *database.Queries
	ctx := context.Background()
	db, err := sql.Open("sqlite3", "course_manager.db")
	if err != nil {
		return nil, nil, err
	}
	if _, err := db.ExecContext(ctx, DDL); err != nil {
		log.Fatal(err)
	}
	// Enable foreign keys
	_, err = db.ExecContext(ctx, "PRAGMA foreign_keys = ON;")
	if err != nil {
		log.Fatal("Failed to enable foreign keys:", err)
	}

	// Check if foreign keys are enabled
	var foreignKeysEnabled int
	err = db.QueryRow("PRAGMA foreign_keys;").Scan(&foreignKeysEnabled)
	if err != nil {
		log.Fatal("Failed to check foreign_keys status:", err)
	}

	log.Println("Foreign keys enabled:", foreignKeysEnabled)
	queries = database.New(db)
	return queries, db, nil

}
