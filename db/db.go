package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

func ConnectToDatabase() *sql.DB {

	url := fmt.Sprintf("libsql://tracker-aniketbiprojit.aws-eu-west-1.turso.io?authToken=%s", os.Getenv("TURSO_TOKEN"))

	db, err := sql.Open("libsql", url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open db %s: %s", url, err)
		os.Exit(1)
	}
	return db
}
