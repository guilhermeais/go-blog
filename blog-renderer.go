package blogrenderer

import (
	"embed"
	"io"
	"text/template"
)

type Post struct {
	Title, Body, Description string
	Tags                     []string
}

var (
	//go:embed "templates/*"
	postTemplate embed.FS
)

func Render(w io.Writer, post Post) error {
	templ, err := template.ParseFS(postTemplate, "templates/*.gohtml")
	if err != nil {
		return err
	}

	if err := templ.Execute(w, post); err != nil {
		return err
	}

	return nil
}
