/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"tracker/cmd/tui"

	"github.com/spf13/cobra"
)

// createProjectCmd represents the createProject command
var createProjectCmd = &cobra.Command{
	Use:   "create-project project-name",
	Short: "Create a new project",
	Long: `A project is what is tracked in time tracker. It is base for all data.
Usage: tracker create-project project-name
	`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		projectName := args[0]
		fmt.Printf("creating project: %s\n", projectName)

		m := tui.GetModel()

		m.AddProject(projectName)
	},
}

func init() {
	rootCmd.AddCommand(createProjectCmd)
}
