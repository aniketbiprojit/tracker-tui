/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"tracker/cmd/tui"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

// createProjectCmd represents the createProject command
var createProjectCmd = &cobra.Command{
	Use:   "create-project project-name",
	Short: "Create a new project",
	Long: `A project is what is tracked in time tracker. It is base for all data.
Usage: tracker create-project project-name
	`,
	Args: cobra.RangeArgs(0, 1),
	Run: func(cmd *cobra.Command, args []string) {

		ti := textinput.New()
		ti.Placeholder = "Pikachu"
		ti.Focus()
		ti.CharLimit = 156
		ti.Width = 20

		m := tui.GetModel()
		if len(args) == 0 {

			m.Render = true
			m.HandleUpdate = func(msg tea.Msg) tea.Cmd {

				var returnCmd tea.Cmd
				switch msg := msg.(type) {
				case tea.KeyMsg:
					switch msg.Type {
					case tea.KeyEnter:
						return tea.Quit
					}
				}

				// We handle errors just like any other message
				ti, returnCmd = ti.Update(msg)

				return returnCmd
			}

			m.HandleView = func() string {
				s := fmt.Sprintf("%s%s\n", "Please enter project name you want to add?\n", ti.View())
				return s
			}
		}

		// projectName := args[0]

		// m.AddProject(projectName)
		m.Render = true
	},
}

func init() {
	rootHolder.rootCmd.AddCommand(createProjectCmd)
	rootHolder.CurrentCommand = "create-project"
	rootHolder.HandleInit = func() tea.Cmd {
		return textinput.Blink
	}
}
