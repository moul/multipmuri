package multipmuri

import "net/url"

func DecodeString(input string) (Entity, error) {
	u, err := url.Parse(input)
	if err != nil {
		return nil, err
	}
	switch u.Scheme {
	case string(GitHubProvider):
		return GitHubService{}.RelDecodeString(input[len(u.Scheme)+3:])
	}

	return GitHubService{}.RelDecodeString(input) // fallback on GitHub (for now)
}
