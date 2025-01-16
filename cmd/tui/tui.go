package tui

import (
	"errors"
	"fmt"
	"os"
	"tracker/db"

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

type HandleUpdateFunc func(msg tea.Msg) tea.Cmd
type HandleInitFunc func() tea.Cmd
type HandleViewFunc func() string
type model struct {
	command       string
	activeProject int
	trackedData   map[int][]TrackedData
	projects      []Project
	Err           error
	ErrPrinted    bool
	Render        bool
	ViewString    string
	HandleInit    HandleInitFunc
	HandleUpdate  HandleUpdateFunc
	HandleView    HandleViewFunc
}

func (m model) AddProject(projectName string) (err error) {

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
	db.AddProjectToDatabase(projectName)

	fmt.Printf("Project `%s` added with id %d\n", projectName, Id)

	return nil
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
		HandleUpdate:  nil,
		ErrPrinted:    false,
	}
	modelData.projects = append(modelData.projects, Project{
		Id:   1,
		Name: "project",
	})

}

func GetModel() *model {
	return &modelData
}

func (m model) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	if m.HandleInit != nil {
		return m.HandleInit()
	}
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	if m.ErrPrinted {

		fmt.Println("Printed ")
		return m, tea.Quit
	}
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	if m.HandleUpdate != nil {
		returnCmd := m.HandleUpdate(msg)
		return m, returnCmd
	}

	return m, cmd
}

func (m model) View() string {

	if m.Err != nil {
		// m.ErrPrinted = true

		fmt.Println("Error: ", m.Err)

		return "Error: " + m.Err.Error() + "\n\n\n"
	}

	if m.HandleView != nil {
		return m.HandleView()
	}

	s := m.ViewString

	s += "\nPress ctrl+c to quit..."

	return s
}

func InitTea(handleInit HandleInitFunc) {
	modelData.HandleInit = handleInit
	p := tea.NewProgram(modelData)
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
