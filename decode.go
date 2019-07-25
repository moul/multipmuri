package multipmuri

import (
	"fmt"
	"net/url"
	"strings"
)

func isProviderScheme(scheme string) bool {
	switch scheme {
	case string(GitHubProvider),
		string(TrelloProvider),
		string(JiraProvider),
		string(GitLabProvider):
		return true
	}
	return false
}

func DecodeString(input string) (Entity, error) {
	u, err := url.Parse(input)
	if err != nil {
		return nil, err
	}

	if isProviderScheme(u.Scheme) {
		input = input[len(u.Scheme)+3:]
		switch u.Scheme {
		case string(GitHubProvider):
			return gitHubRelDecodeString(getHostname(input), "", "", input, true)
			//case string(GitLabProvider):
			//case string(JiraProvider):
			//case string(TrelloProvider):
		}
	}

	if u.Scheme == "" && u.Host == "" && u.Path != "" { // github.com/x/x
		u.Host = strings.Split(u.Path, "/")[0]
		// u.Path = u.Path[len(u.Host)+1:]
	}

	switch u.Scheme {
	case "", "https", "http":
		switch u.Host {
		case "github.com":
			return gitHubRelDecodeString("", "", "", input, true)
			// case "gitlab.com":
			// case "jira.com", "atlassian.com":
			// case "trello.com":
		}
	}

	return nil, fmt.Errorf("ambiguous uri %q", input)
}
