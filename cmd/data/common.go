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

func InitDB() (*database.Queries, error) {
	var queries *database.Queries
	ctx := context.Background()
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		return nil, err
	}
	defer db.Close()
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
	return queries, nil

}
