package blogrenderer

import (
	"bytes"
	"embed"
	"fmt"
	"html/template"
	"io"

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
	postVm, err := newPostVM(post)
	if err != nil {
		return err
	}

	if err := r.templ.ExecuteTemplate(w, "blog.gohtml", postVm); err != nil {
		return err
	}

	return nil
}

func (r PostRenderer) RenderIndex(w io.Writer, posts []Post) error {
	return r.templ.ExecuteTemplate(w, "index.gohtml", posts)
}

type postViewModel struct {
	Post
	HTMLBody template.HTML
}

func newPostVM(p Post) (*postViewModel, error) {
	var bodyMdToHtml bytes.Buffer
	if err := goldmark.Convert([]byte(p.Body), &bodyMdToHtml); err != nil {
		return nil, fmt.Errorf("error on converting body to html: %q", err.Error())
	}

	vm := &postViewModel{p, template.HTML(bodyMdToHtml.String())}

	return vm, nil
}
