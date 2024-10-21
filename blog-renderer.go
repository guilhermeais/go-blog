package blogrenderer

import (
	"io"
	"text/template"
)

type Post struct {
	Title, Body, Description string
	Tags                     []string
}

const (
	POST_TEMPLATE = `<h1>{{.Title}}</h1><p>{{.Description}}</p>Tags: <ul>{{range .Tags}}<li>{{.}}</li>{{end}}</ul>`
)

func Render(w io.Writer, post Post) error {
	templ, err := template.New("blog").Parse(POST_TEMPLATE)
	if err != nil {
		return err
	}

	if err := templ.Execute(w, post); err != nil {
		return err
	}

	return nil
}
