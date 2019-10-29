package multipmuri

import "fmt"

func ExampleNewGitHubIssue() {
	entity := NewGitHubIssue("", "moul", "depviz", "42")
	fmt.Println("entity")
	fmt.Println(" ", entity.String())
	fmt.Println(" ", entity.Kind())
	fmt.Println(" ", entity.Provider())

	relatives := []string{
		"@moul",
		"#4242",
		"moul2/depviz2#43",
		"moul/depviz#42",
		"github.com/moul2/depviz2#42",
		"https://github.com/moul2/depviz2#42",
		"https://example.com/a/b#42",
		"https://gitlab.com/moul/depviz/issues/42",
	}
	fmt.Println("relationships")
	for _, name := range relatives {
		attrs := ""
		rel, err := entity.RelDecodeString(name)
		if err != nil {
			fmt.Printf("  %-42s -> error: %v\n", name, err)
			continue
		}
		if rel.Equals(entity) {
			attrs += " (equals)"
		}
		if entity.Contains(rel) {
			attrs += " (contains)"
		}
		if rel.Contains(entity) {
			attrs += " (is contained)"
		}
		fmt.Printf("  %-42s -> %s%s\n", name, rel.String(), attrs)
	}
	fmt.Println("repo:", entity.RepoEntity().String())
	fmt.Println("owner:", entity.OwnerEntity().String())
	fmt.Println("complex relationship:",
		entity.Owner().
			Service().
			Owner("test1").
			Repo("test2").
			Issue("42").
			Repo().
			Service().
			Owner("test3").
			Repo("test4").
			Milestone("42").
			String())
	// Output:
	// entity
	//   https://github.com/moul/depviz/issues/42
	//   issue
	//   github
	// relationships
	//   @moul                                      -> https://github.com/moul (is contained)
	//   #4242                                      -> https://github.com/moul/depviz/issues/4242
	//   moul2/depviz2#43                           -> https://github.com/moul2/depviz2/issues/43
	//   moul/depviz#42                             -> https://github.com/moul/depviz/issues/42 (equals)
	//   github.com/moul2/depviz2#42                -> https://github.com/moul2/depviz2/issues/42
	//   https://github.com/moul2/depviz2#42        -> https://github.com/moul2/depviz2/issues/42
	//   https://example.com/a/b#42                 -> error: ambiguous uri "https://example.com/a/b#42"
	//   https://gitlab.com/moul/depviz/issues/42   -> https://gitlab.com/moul/depviz/issues/42
	// repo: https://github.com/moul/depviz
	// owner: https://github.com/moul
	// complex relationship: https://github.com/test3/test4/milestone/42
}

func ExampleNewGitHubService() {
	entity := NewGitHubService("github.com")
	fmt.Println("entity")
	fmt.Println(" ", entity.String())
	fmt.Println(" ", entity.Kind())
	fmt.Println(" ", entity.Provider())

	relatives := []string{
		"https://github.com",
		"github.com",
		"github.com/moul",
		"@moul",
		"github.com/moul/depviz",
		"moul/depviz",
		"moul/depviz/milestone/1",
		"moul/depviz#1",
		"github.com/moul/depviz/issues/2",
		"github.com/moul/depviz/pull/1",
		"https://github.com/moul/depviz/issues/1",
		"https://github.com/moul/depviz#1",
		"github://moul/depviz#1",
		"github://github.com/moul/depviz#1",
		"github://https://github.com/moul/depviz#1",
	}
	fmt.Println("relationships")
	for _, name := range relatives {
		rel, err := entity.RelDecodeString(name)
		if err != nil {
			fmt.Printf("  %-42s -> error: %v\n", name, err)
			continue
		}
		fmt.Printf("  %-42s -> %-43s %s\n", name, rel.String(), rel.Kind())
	}
	// Output:
	// entity
	//   https://github.com/
	//   service
	//   github
	// relationships
	//   https://github.com                         -> https://github.com/                         service
	//   github.com                                 -> https://github.com/                         service
	//   github.com/moul                            -> https://github.com/moul                     user-or-organization
	//   @moul                                      -> https://github.com/moul                     user-or-organization
	//   github.com/moul/depviz                     -> https://github.com/moul/depviz              project
	//   moul/depviz                                -> https://github.com/moul/depviz              project
	//   moul/depviz/milestone/1                    -> https://github.com/moul/depviz/milestone/1  milestone
	//   moul/depviz#1                              -> https://github.com/moul/depviz/issues/1     issue-or-merge-request
	//   github.com/moul/depviz/issues/2            -> https://github.com/moul/depviz/issues/2     issue
	//   github.com/moul/depviz/pull/1              -> https://github.com/moul/depviz/issues/1     merge-request
	//   https://github.com/moul/depviz/issues/1    -> https://github.com/moul/depviz/issues/1     issue
	//   https://github.com/moul/depviz#1           -> https://github.com/moul/depviz/issues/1     issue-or-merge-request
	//   github://moul/depviz#1                     -> https://github.com/moul/depviz/issues/1     issue-or-merge-request
	//   github://github.com/moul/depviz#1          -> https://github.com/moul/depviz/issues/1     issue-or-merge-request
	//   github://https://github.com/moul/depviz#1  -> https://github.com/moul/depviz/issues/1     issue-or-merge-request
}

func ExampleNewGitHubService_Enterprise() {
	entity := NewGitHubService("ge.company.com")
	fmt.Println("entity")
	fmt.Println(" ", entity.String())
	fmt.Println(" ", entity.Kind())
	fmt.Println(" ", entity.Provider())

	relatives := []string{
		"https://github.com",
		"github.com",
		"github.com/moul",
		"@moul",
		"github.com/moul/depviz",
		"moul/depviz",
		"moul/depviz/milestone/1",
		"moul/depviz#1",
		"github.com/moul/depviz/issues/2",
		"github.com/moul/depviz/pull/1",
		"https://github.com/moul/depviz/issues/1",
		"https://github.com/moul/depviz#1",
		"github://moul/depviz#1",
		"github://github.com/moul/depviz#1",
		"github://https://github.com/moul/depviz#1",
	}
	fmt.Println("relationships")
	for _, name := range relatives {
		rel, err := entity.RelDecodeString(name)
		if err != nil {
			fmt.Printf("  %-42s -> error: %v\n", name, err)
			continue
		}
		fmt.Printf("  %-42s -> %-43s %s\n", name, rel.String(), rel.Kind())
	}
	// Output:
	// entity
	//   https://ge.company.com/
	//   service
	//   github
	// relationships
	//   https://github.com                         -> https://github.com/                         service
	//   github.com                                 -> https://github.com/                         service
	//   github.com/moul                            -> https://github.com/moul                     user-or-organization
	//   @moul                                      -> https://ge.company.com/moul                 user-or-organization
	//   github.com/moul/depviz                     -> https://github.com/moul/depviz              project
	//   moul/depviz                                -> https://ge.company.com/moul/depviz          project
	//   moul/depviz/milestone/1                    -> https://ge.company.com/moul/depviz/milestone/1 milestone
	//   moul/depviz#1                              -> https://ge.company.com/moul/depviz/issues/1 issue-or-merge-request
	//   github.com/moul/depviz/issues/2            -> https://github.com/moul/depviz/issues/2     issue
	//   github.com/moul/depviz/pull/1              -> https://github.com/moul/depviz/issues/1     merge-request
	//   https://github.com/moul/depviz/issues/1    -> https://github.com/moul/depviz/issues/1     issue
	//   https://github.com/moul/depviz#1           -> https://github.com/moul/depviz/issues/1     issue-or-merge-request
	//   github://moul/depviz#1                     -> https://github.com/moul/depviz/issues/1     issue-or-merge-request
	//   github://github.com/moul/depviz#1          -> https://github.com/moul/depviz/issues/1     issue-or-merge-request
	//   github://https://github.com/moul/depviz#1  -> https://github.com/moul/depviz/issues/1     issue-or-merge-request
}
