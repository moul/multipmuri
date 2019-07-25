package multipmuri

import (
	"fmt"
	"net/url"
	"strings"
)

func DecodeString(input string) (Entity, error) {
	u, err := url.Parse(input)
	if err != nil {
		return nil, err
	}
	switch u.Scheme {
	case string(GitHubProvider):
		return gitHubRelDecodeString("", "", input[len(u.Scheme)+3:], true)
	}

	if u.Scheme == "" && u.Host == "" && u.Path != "" {
		u.Host = strings.Split(u.Path, "/")[0]
		// u.Path = u.Path[len(u.Host)+1:]
	}

	if u.Scheme == "" || u.Scheme == "https" || u.Scheme == "http" {
		switch u.Host {
		case "github.com":
			return gitHubRelDecodeString("", "", input, true)
		}
	}

	return nil, fmt.Errorf("ambiguous uri %q", input)
}
