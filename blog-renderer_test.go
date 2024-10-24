package blogrenderer_test

import (
	"bytes"
	blogrenderer "go-blog"
	"io"
	"testing"

	approvals "github.com/approvals/go-approval-tests"
)

func TestRender(t *testing.T) {
	aPost := blogrenderer.Post{
		Title:       "hello, world!",
		Body:        "this is my first post",
		Description: "This is a description",
		Tags:        []string{"dev", "go", "tdd"},
	}

	postRenderer, err := blogrenderer.NewPostRenderer()
	if err != nil {
		t.Fatal(err)
	}

	t.Run("it converts a single post into HTML", func(t *testing.T) {
		buf := bytes.Buffer{}

		if err := postRenderer.Render(&buf, aPost); err != nil {
			t.Fatal(err)
		}

		approvals.VerifyString(t, buf.String())
	})

	t.Run("it converts a single post into HTML with markdown on body", func(t *testing.T) {
		aPostWithMarkdownBody := blogrenderer.Post{
			Title:       "hello, world!",
			Body:        "## Hello, word\nthis is a **hello, wold** text using **markdown**",
			Description: "This is a description",
			Tags:        []string{"dev", "go", "tdd"},
		}
		buf := bytes.Buffer{}

		if err := postRenderer.Render(&buf, aPostWithMarkdownBody); err != nil {
			t.Fatal(err)
		}

		approvals.VerifyString(t, buf.String())
	})

	t.Run("it handles malicious post", func(t *testing.T) {
		aPostWithMarkdownBody := blogrenderer.Post{
			Title: "hello, world!",
			Body: `# Teste de XSS

Aqui está um link malicioso:

[XSS](javascript:alert('XSS'))

E aqui está um script embutido:

<script>alert('XSS');</script>`,
			Description: "This is a description",
			Tags:        []string{"dev", "go", "tdd"},
		}
		buf := bytes.Buffer{}

		if err := postRenderer.Render(&buf, aPostWithMarkdownBody); err != nil {
			t.Fatal(err)
		}

		approvals.VerifyString(t, buf.String())
	})
}

func BenchmarkRender(b *testing.B) {
	postRenderer, _ := blogrenderer.NewPostRenderer()

	aPost := blogrenderer.Post{
		Title:       "hello, world!",
		Body:        "this is my first post",
		Description: "This is a description",
		Tags:        []string{"dev", "go", "tdd"},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		postRenderer.Render(io.Discard, aPost)
	}
}
