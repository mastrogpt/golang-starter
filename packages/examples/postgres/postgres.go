// --param POSTGRES_URL $POSTGRES_URL

package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// Main function for the action
func Main(obj map[string]any) map[string]any {
	return postgres(obj)
}

// actual logic for the action
func postgres(obj map[string]any) map[string]any {
	url, ok := obj["POSTGRES_URL"].(string)
	if !ok {
		return sendError(nil, 400, "POSTGRES_URL not found")
	}

	db, err := sql.Open("postgres", url)
	if err != nil {
		return sendError(err, 500, "Connection to Postgres failed")
	}
	defer db.Close()

	_, err = db.Exec(`
    CREATE EXTENSION IF NOT EXISTS "pgcrypto";
    CREATE TABLE IF NOT EXISTS nuvolaris_table (
      id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
      message varchar(100) NOT NULL
    );
  `)
	if err != nil {
		return sendError(err, 500, "Table creation failed")
	}

	_, err = db.Exec("INSERT INTO nuvolaris_table(message) VALUES($1)", "Nuvolaris Postgres is up and running!")
	if err != nil {
		return sendError(err, 500, "Insert test failed")
	}

	var message string
	err = db.QueryRow("SELECT message FROM nuvolaris_table").Scan(&message)
	if err != nil {
		return sendError(err, 500, "Query execution test failed")
	}

	return map[string]any{
		"body": message,
	}
}

func sendError(err any, status int, message string) map[string]any {
	fmt.Println(err)
	return map[string]any{
		"body":       message,
		"statusCode": status,
	}
}
