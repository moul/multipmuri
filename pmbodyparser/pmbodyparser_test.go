package pmbodyparser

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"moul.io/multipmuri"
)

func ExampleRelParseString() {
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

func ExampleParseString() {
	rels, errs := ParseString("Depends on github.com/moul/depviz#1")
	if len(errs) > 0 {
		panic(errs)
	}
	for _, rel := range rels {
		fmt.Println(rel)
	}
	// Output:
	// depends-on https://github.com/moul/depviz/issues/1
}

func TestRelParseString(t *testing.T) {
	Convey("Testing RelParseString", t, func() {
		// simple
		rels, errs := RelParseString(
			multipmuri.NewGitHubIssue("", "moul", "depviz", "1"),
			"Depends on #2",
		)
		So(len(errs), ShouldEqual, 0)
		So(len(rels), ShouldEqual, 1)
		So(rels[0].Kind, ShouldEqual, DependsOn)
		So(rels[0].Target.String(), ShouldEqual, "https://github.com/moul/depviz/issues/2")

		// multiple
		rels, errs = RelParseString(
			multipmuri.NewGitHubIssue("", "moul", "depviz", "1"),
			"Depends on #2\nDepends on #3",
		)
		So(len(errs), ShouldEqual, 0)
		So(len(rels), ShouldEqual, 2)
		So(rels[0].Kind, ShouldEqual, DependsOn)
		So(rels[0].Target.String(), ShouldEqual, "https://github.com/moul/depviz/issues/2")
		So(rels[1].Kind, ShouldEqual, DependsOn)
		So(rels[1].Target.String(), ShouldEqual, "https://github.com/moul/depviz/issues/3")

		// spaces
		rels, errs = RelParseString(
			multipmuri.NewGitHubIssue("", "moul", "depviz", "1"),
			" Depends on #2 \n Depends on #3 \n\n ",
		)
		So(len(errs), ShouldEqual, 0)
		So(len(rels), ShouldEqual, 2)
		So(rels[0].Kind, ShouldEqual, DependsOn)
		So(rels[0].Target.String(), ShouldEqual, "https://github.com/moul/depviz/issues/2")
		So(rels[1].Kind, ShouldEqual, DependsOn)
		So(rels[1].Target.String(), ShouldEqual, "https://github.com/moul/depviz/issues/3")
	})
}
