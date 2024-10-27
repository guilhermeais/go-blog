package blogrenderer

import (
	"bytes"
	"embed"
	"fmt"
	"io"
	"strings"
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

func (r PostRenderer) RenderIndex(w io.Writer, posts []Post) error {
	indexTemplate := `<ol>{{range .}}<li><a href="/post/{{titleToSlug .Title}}">{{.Title}}</a></li>{{end}}</ol>`
	templ, err := template.New("index").Funcs(template.FuncMap{
		"titleToSlug": func(title string) string {
			return strings.ToLower(strings.ReplaceAll(title, " ", "-"))
		},
	}).Parse(indexTemplate)

	if err != nil {
		return err
	}

	if err := templ.Execute(w, posts); err != nil {
		return err
	}

	return nil
}
