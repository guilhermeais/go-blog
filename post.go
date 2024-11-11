package blogrenderer

import "strings"

type Post struct {
	Title, Body, Description string
	Tags                     []string
}

func (p Post) SanitisedTitle() string {
	return strings.ToLower(strings.ReplaceAll(p.Title, " ", "-"))
}
