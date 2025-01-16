package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

func ConnectToDatabase() *sql.DB {

	url := fmt.Sprintf("%s?authToken=%s", os.Getenv("TURSO_URL"), os.Getenv("TURSO_TOKEN"))

	db, err := sql.Open("libsql", url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open db %s: %s", url, err)
		os.Exit(1)
	}

	_, err = db.Exec("SELECT 1;")
	if err != nil {
		fmt.Fprintf(os.Stderr, "DB Conn failed %s", err)
		os.Exit(1)
	}

	return db
}
