package main

import (
	"tracker/cmd"
	"tracker/cmd/tui"
)

func main() {
	cmd.Execute()

	if tui.GetModel().Render {
		tui.InitTea()
	}
}
