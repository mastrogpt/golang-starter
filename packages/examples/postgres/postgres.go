// --param POSTGRES_URL $POSTGRES_URL

package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// Main function for the action
func Postgres(obj map[string]any) map[string]any {

	url, ok := obj["POSTGRES_URL"].(string)
	if !ok {
		return send500Error(nil, 400, "POSTGRES_URL not found")
	}

	db, err := sql.Open("postgres", url)
	if err != nil {
		return send500Error(err, 500, "Connection to Postgres failed")
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
		return send500Error(err, 500, "Table creation failed")
	}

	_, err = db.Exec("INSERT INTO nuvolaris_table(message) VALUES($1)", "Nuvolaris Postgres is up and running!")
	if err != nil {
		return send500Error(err, 500, "Insert test failed")
	}

	var message string
	err = db.QueryRow("SELECT message FROM nuvolaris_table").Scan(&message)
	if err != nil {
		return send500Error(err, 500, "Query execution test failed")
	}

	response := make(map[string]any)
	response["body"] = message
	return response
}

func send500Error(err any, status int, message string) map[string]any {
	fmt.Println(err)
	response := make(map[string]any)

	response["body"] = message
	response["statusCode"] = status
	return response

}
