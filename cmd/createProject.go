/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"tracker/cmd/tui"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

const (
	defaultWidth = 80

	//default colors
	purple    = `#7e2fcc`
	darkGrey  = `#353C3B`
	lightTeal = `#03DAC5`
	darkTeal  = `#01A299`
	white     = `#e5e5e5`
	red       = `#FF3333`
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

		TextInputStyle := lipgloss.NewStyle().Foreground(lipgloss.Color(white)).
			Padding(0, 1).
			Margin(1, 0)

		CursorStyle := lipgloss.NewStyle().Foreground(lipgloss.Color(lightTeal))

		ti := textinput.New()
		ti.Placeholder = "Project name here"
		ti.Focus()
		ti.CharLimit = 156
		ti.Width = 20
		ti.TextStyle = TextInputStyle
		ti.CursorStyle = CursorStyle

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

				ti, returnCmd = ti.Update(msg)

				return returnCmd
			}

			BorderStyle := lipgloss.NewStyle().
				Padding(0, 1, 0, 1).
				Width(defaultWidth).
				BorderForeground(lipgloss.AdaptiveColor{Light: darkTeal, Dark: lightTeal}).
				Border(lipgloss.ThickBorder())

			TitleStyle := lipgloss.NewStyle().Bold(true).
				// Width(defaultWidth - 4).
				Align(lipgloss.Center)
			m.HandleView = func() string {

				title := TitleStyle.Render("Please enter project name you want to add?\n")
				inputView := ti.View()
				verticalView := lipgloss.JoinVertical(lipgloss.Left, title, inputView)

				border := BorderStyle.Render(verticalView)

				return border + "\nPress q to exit"
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
