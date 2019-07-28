package pmbodyparser

import (
	"fmt"

	"moul.io/multipmuri"
)

func ExampleParseString() {
	body := `
This PR fixes a lot of things and implement plenty new features.

Addresses #42
Depends on: github.com/moul/depviz#42
Blocks #45
Block: #46
fixes: #58
FIX github.com/moul/depviz#1337

Signed-off-by: Super Developer <super.dev@gmail.com>
`
	relationships, errs := RelParseString(
		multipmuri.NewGitHubIssue("github.com", "moul", "depviz", "1"),
		body,
	)
	if len(errs) > 0 {
		panic(errs)
	}
	for _, relationship := range relationships {
		fmt.Println(relationship)
	}
	// Output:
	// addresses https://github.com/moul/depviz/issues/42
	// blocks https://github.com/moul/depviz/issues/45
	// blocks https://github.com/moul/depviz/issues/46
	// depends-on https://github.com/moul/depviz/issues/42
	// fixes https://github.com/moul/depviz/issues/1337
	// fixes https://github.com/moul/depviz/issues/58
}
