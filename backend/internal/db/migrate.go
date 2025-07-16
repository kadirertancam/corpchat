package db

import "database/sql"

func Migrate(db *sql.DB) error {
	schema := `
CREATE EXTENSION IF NOT EXISTS pgcrypto;
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username TEXT UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT now()
);`
	_, err := db.Exec(schema)
	return err
}