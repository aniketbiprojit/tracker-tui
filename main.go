package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type TrackingData struct {
	startedAt int
	endedAt   int
}

type Tracking struct {
	active *string
	tracks map[string][]TrackingData
}

func DefaultStyle() lipgloss.Style {
	return lipgloss.NewStyle()
}

type Project struct {
	client string
	name   string
}

type Model struct {
	cursor   int
	width    int
	height   int
	style    *lipgloss.Style
	projects []Project
	tracking Tracking
	err      error
}

func initialTracking() Tracking {
	return Tracking{
		active: nil,
		tracks: make(map[string][]TrackingData),
	}
}

func (t Model) Init() tea.Cmd {
	return nil
}

func (t Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return t, tea.Quit
		}
	case tea.WindowSizeMsg:
		t.width = msg.Width
		t.height = msg.Height
	}

	return t, nil
}

func (t Model) View() string {

	if t.width == 0 {
		return "loading..."
	}

	s := "Options:"

	s += "\n1. New Tracking."
	s += "\n2. List Tracking."
	s += "\n3. Start Tracking."

	s += "\nPress q to quit."
	return s
}

func main() {
	style := DefaultStyle()
	p := tea.NewProgram(Model{
		cursor:   0,
		width:    0,
		height:   0,
		style:    &style,
		projects: make([]Project, 0),
	})
	if _, err := p.Run(); err != nil {
		fmt.Printf("Damn gurrrrrl, what did you do. %v", err)
		os.Exit(420)
	}
}
