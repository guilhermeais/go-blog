package blogrenderer_test

import (
	"bytes"
	blogrenderer "go-blog"
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

	t.Run("it converts a single post into HTML", func(t *testing.T) {
		buf := bytes.Buffer{}
		err := blogrenderer.Render(&buf, aPost)

		if err != nil {
			t.Fatal(err)
		}

		approvals.VerifyString(t, buf.String())
	})
}
