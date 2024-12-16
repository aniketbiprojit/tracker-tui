package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type TrackingData struct {
	startedAt int
	endedAt   int
}

type Tracking struct {
	active *string
	tracks map[string]TrackingData
	cursor int
}

func initialTracking() Tracking {
	return Tracking{
		active: nil,
		tracks: make(map[string]TrackingData),
		cursor: 0,
	}
}

func (t Tracking) Init() tea.Cmd {
	return nil
}

func (t Tracking) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return t, tea.Quit
		}
	}
	return t, nil
}

func (t Tracking) View() string {
	s := ""

	s += "\nPress q to quit."
	return s
}

func main() {
	p := tea.NewProgram(initialTracking())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Damn gurrrrrl, what did you do. %v", err)
	}
}
