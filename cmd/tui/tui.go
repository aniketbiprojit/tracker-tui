package tui

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
}

func (m model) AddProject(projectName string) int {

	Id := len(m.projects) + 1
	m.projects = append(m.projects, Project{
		Id:   Id,
		Name: projectName,
	})

	return Id
}

var modelData = model{}

func init() {
	modelData = model{
		command:       "",
		activeProject: -1,
		trackedData:   make(map[int][]TrackedData),
		projects:      make([]Project, 0),
	}
}

func GetModel() *model {
	return &modelData
}
