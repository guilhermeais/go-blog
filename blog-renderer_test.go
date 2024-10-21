package blogrenderer_test

import (
	"bytes"
	blogrenderer "go-blog"
	"testing"
)

func TestRender(t *testing.T) {
	aPost := blogrenderer.Post{
		Title:       "hello, world!",
		Body:        "this is my first post",
		Description: "This is a description",
		Tags:        []string{"dev", "go", "tdd"},
	}

	t.Run("it converts a single psot into HTML", func(t *testing.T) {
		buf := bytes.Buffer{}
		err := blogrenderer.Render(&buf, aPost)

		if err != nil {
			t.Fatal(err)
		}

		want := `<h1>hello, world!</h1><p>This is a description</p>Tags: <ul><li>dev</li><li>go</li><li>tdd</li></ul>`

		got := buf.String()

		if got != want {
			t.Errorf("got '%s', want '%s'", got, want)
		}
	})
}
