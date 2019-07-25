package multipmuri

import "fmt"

func ExampleNewGitHubIssue() {
	issue := NewGitHubIssue("moul", "depviz", "42")
	fmt.Println(issue.Canonical())
	fmt.Println(issue.Kind())
	fmt.Println(issue.Provider())
	// Output:
	// https://github.com/moul/depviz/issues/42
	// issue
	// github
}

func ExampleNewGitHubService() {
	issue := NewGitHubService()
	fmt.Println(issue.Canonical())
	fmt.Println(issue.Kind())
	fmt.Println(issue.Provider())
	// Output:
	// https://github.com/
	// service
	// github
}
