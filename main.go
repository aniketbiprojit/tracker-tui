package main

import (
	"tracker/cmd"
	"tracker/cmd/tui"
)

func main() {
	executor := cmd.Execute()

	if tui.GetModel().Render {
		if executor.HandleInit != nil {
			tui.InitTea(executor.HandleInit)
		}
	}
}
