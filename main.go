package main

import (
	"fmt"
	"os"
	"tracker/cmd"
	"tracker/cmd/db"
	"tracker/cmd/tui"
)

func main() {

	db := db.ConnectToDatabase()

	rows, err := db.Query("SELECT 1")
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to execute query: %v\n", err)
		os.Exit(1)
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
