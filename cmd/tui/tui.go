package tui

import (
	"errors"
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type TrackedData struct {
	started int
	ended   int
}

type Project struct {
	Name string
	Id   int
}

type model struct {
	command       string
	activeProject int
	trackedData   map[int][]TrackedData
	projects      []Project
	Err           error
	Render        bool
	ViewString    string
}

func (m model) AddProject(projectName string) {

	found := false
	for _, val := range m.projects {
		if val.Name == projectName {
			found = true
			break
		}
	}

	if found {
		m.Err = errors.New("Project already exists.")
		return
	}

	Id := len(m.projects) + 1

	m.projects = append(m.projects, Project{
		Id:   Id,
		Name: projectName,
	})

	return
}

var modelData = model{}

func init() {
	modelData = model{
		command:       "",
		activeProject: -1,
		trackedData:   make(map[int][]TrackedData),
		projects:      make([]Project, 0),
		Err:           nil,
		Render:        false,
		ViewString:    "",
	}
	modelData.AddProject("some-project")
}

func GetModel() *model {
	return &modelData
}

func (m model) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m model) View() string {
	s := m.ViewString

	s += "\nPress q to quit..."

	return s
}

func InitTea() {
	p := tea.NewProgram(modelData)
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
