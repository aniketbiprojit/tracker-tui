package main

import (
	"fmt"
	"os"
	"tracker/cmd"
	"tracker/cmd/tui"
	"tracker/db"
)

func main() {

	db := db.ConnectToDatabase()

	rows, err := db.Query("SELECT 1")
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to execute query: %v\n", err)
		os.Exit(1)
	} else {
		fmt.Println("Connected to database.")
	}
	defer rows.Close()

	executor := cmd.Execute()

	if tui.GetModel().Render {
		if executor.HandleInit != nil {
			tui.InitTea(executor.HandleInit)
		}
	}

	defer db.Close()
}
