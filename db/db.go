package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

func ConnectToDatabase() *sql.DB {

	url := "libsql://tracker-aniketbiprojit.aws-eu-west-1.turso.io?authToken=eyJhbGciOiJFZERTQSIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE3MzUzNzE0NzQsImlkIjoiODZkZTk1OGItMmU0My00Y2M4LThkZTAtYzNiNmI3MDFmNTUzIn0.Kc9T9iqJhejdHA-V602NuZNINmEOFXlFwQeEP7GlOp3ZtroQkfEboZA2TedOaoNarD5j8h98jPXkHcM8yEkJBA"

	db, err := sql.Open("libsql", url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open db %s: %s", url, err)
		os.Exit(1)
	}
	return db
}
