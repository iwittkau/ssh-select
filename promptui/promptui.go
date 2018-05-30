package promptui

import (
	"strings"

	"github.com/iwittkau/ssh-select"
	"github.com/manifoldco/promptui"
)

var (
	Template = promptui.SelectTemplates{
		Label: "{{ . }}:",
		//Active:   " {{ .Name | bgCyan }} ({{ .Username | red }}{{  \"@\" | red }}{{ .IpAddress | red }})",
		Active: "[{{.Index}}] {{ .Name | bgCyan }}",
		//Inactive: " {{ .Name | cyan }} ({{ .Username | red }}{{  \"@\" | red }}{{ .IpAddress | red }})",
		Inactive: "[{{.Index}}] {{ .Name | cyan }}",
		Selected: "[{{.Index}}] {{ .Name | cyan }} ({{ .Username | red }}{{  \"@\" | red }}{{ .IpAddress | red }})",
		Details: `
  {{ "Name:" | faint }}	{{ .Name }}
  {{ "IP Address:" | faint }}	{{ .IpAddress }}
  {{ "Username:" | faint }}	{{ .Username }}`,
	}
)

type Frontend struct {
	prompt promptui.Select
}

func New(config *sshselect.Configuration) *Frontend {

	searcher := func(input string, index int) bool {
		server := config.Servers[index]
		name := strings.Replace(strings.ToLower(server.Name), " ", "", -1)
		input = strings.Replace(strings.ToLower(input), " ", "", -1)

		return strings.Contains(name, input)
	}

	prompt := promptui.Select{
		Label:     "Select a host to connect to",
		Items:     config.Servers,
		Templates: &Template,
		Size:      10,
		Searcher:  searcher,
	}

	return &Frontend{prompt}
}

func (f *Frontend) Draw() (*int, error) {
	i, _, err := f.prompt.Run()
	return &i, err
}
