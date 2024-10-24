package blogrenderer

import (
	"bytes"
	"embed"
	"fmt"
	"io"
	"text/template"

	"github.com/yuin/goldmark"
)

var (
	//go:embed "templates/*"
	postTemplate embed.FS
)

type PostRenderer struct {
	templ *template.Template
}

func NewPostRenderer() (*PostRenderer, error) {
	templ, err := template.ParseFS(postTemplate, "templates/*.gohtml")
	if err != nil {
		return nil, err
	}

	return &PostRenderer{templ}, nil
}

func (r PostRenderer) Render(w io.Writer, post Post) error {
	var bodyMdToHtml bytes.Buffer
	if err := goldmark.Convert([]byte(post.Body), &bodyMdToHtml); err != nil {

		return fmt.Errorf("error on converting body to html: %q", err.Error())
	}

	post.Body = bodyMdToHtml.String()

	if err := r.templ.ExecuteTemplate(w, "blog.gohtml", post); err != nil {
		return err
	}

	return nil
}
