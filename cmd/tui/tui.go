package tui

import "errors"

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
	}
	modelData.AddProject("some-project")
}

func GetModel() *model {
	return &modelData
}
