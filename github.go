package multipmuri

import (
	"fmt"
	"net/url"
	"strings"
)

// FIXME: reduce redundant code

type githubCommon struct{}

type GitHubIssue struct {
	githubCommon
	Issue
	Owner string
	Repo  string
	ID    string
}

type GitHubMilestone struct {
	githubCommon
	Issue
	Owner string
	Repo  string
	ID    string
}

type GitHubPullRequest struct {
	githubCommon
	MergeRequest
	Owner string
	Repo  string
	ID    string
}

type GitHubIssueOrPullRequest struct {
	githubCommon
	MergeRequest
	Owner string
	Repo  string
	ID    string
}

type GitHubService struct {
	githubCommon
	Service
}

type GitHubUserOrOrganization struct {
	githubCommon
	UserOrOrganization
	Owner string
}

type GitHubRepo struct {
	githubCommon
	UserOrOrganization
	Owner string
	Repo  string
}

func (e GitHubRepo) Canonical() string {
	return fmt.Sprintf("https://github.com/%s/%s", e.Owner, e.Repo)
}
func (e GitHubRepo) Provider() Provider { return GitHubProvider }
func (e GitHubRepo) Kind() Kind         { return ProjectKind }
func (e GitHubRepo) RelDecodeString(input string) (DecodedMultipmuri, error) {
	return githubRelDecodeString(e.githubCommon, input)
}

func (e GitHubMilestone) Canonical() string {
	return fmt.Sprintf("https://github.com/%s/%s/milestone/%s", e.Owner, e.Repo, e.ID)
}
func (e GitHubMilestone) Provider() Provider { return GitHubProvider }
func (e GitHubMilestone) Kind() Kind         { return MilestoneKind }
func (e GitHubMilestone) RelDecodeString(input string) (DecodedMultipmuri, error) {
	return githubRelDecodeString(e.githubCommon, input)
}

func (e GitHubIssue) Canonical() string {
	return fmt.Sprintf("https://github.com/%s/%s/issues/%s", e.Owner, e.Repo, e.ID)
}
func (e GitHubIssue) Provider() Provider { return GitHubProvider }
func (e GitHubIssue) Kind() Kind         { return IssueKind }
func (e GitHubIssue) RelDecodeString(input string) (DecodedMultipmuri, error) {
	return githubRelDecodeString(e.githubCommon, input)
}

func (e GitHubIssueOrPullRequest) Canonical() string {
	return fmt.Sprintf("https://github.com/%s/%s/issues/%s", e.Owner, e.Repo, e.ID)
}
func (e GitHubIssueOrPullRequest) Provider() Provider { return GitHubProvider }
func (e GitHubIssueOrPullRequest) Kind() Kind         { return IssueOrMergeRequestKind }
func (e GitHubIssueOrPullRequest) RelDecodeString(input string) (DecodedMultipmuri, error) {
	return githubRelDecodeString(e.githubCommon, input)
}

func (e GitHubPullRequest) Canonical() string {
	return fmt.Sprintf("https://github.com/%s/%s/pull/%s", e.Owner, e.Repo, e.ID)
}
func (e GitHubPullRequest) Provider() Provider { return GitHubProvider }
func (e GitHubPullRequest) Kind() Kind         { return MergeRequestKind }
func (e GitHubPullRequest) RelDecodeString(input string) (DecodedMultipmuri, error) {
	return githubRelDecodeString(e.githubCommon, input)
}

func (e GitHubUserOrOrganization) Canonical() string  { return "https://github.com/" + e.Owner }
func (e GitHubUserOrOrganization) Provider() Provider { return GitHubProvider }
func (e GitHubUserOrOrganization) Kind() Kind         { return UserOrOrganizationKind }
func (e GitHubUserOrOrganization) RelDecodeString(input string) (DecodedMultipmuri, error) {
	return githubRelDecodeString(e.githubCommon, input)
}

func (e GitHubService) Canonical() string  { return "https://github.com" }
func (e GitHubService) Provider() Provider { return GitHubProvider }
func (e GitHubService) Kind() Kind         { return ServiceKind }
func (e GitHubService) RelDecodeString(input string) (DecodedMultipmuri, error) {
	return githubRelDecodeString(e.githubCommon, input)
}

func githubRelDecodeString(common githubCommon, input string) (DecodedMultipmuri, error) {
	// sanitization
	u, err := url.Parse(input)
	if err != nil {
		return nil, err
	}
	if u.Host == "" && strings.HasPrefix(u.Path, "github.com") {
		u.Path = u.Path[10:]
	}
	if len(u.Path) > 0 && u.Path[0] == '/' {
		u.Path = u.Path[1:]
	}
	if u.Path == "" {
		return &GitHubService{}, nil
	}
	if u.Fragment != "" {
		u.Path += "/issue-or-pull-request/" + u.Fragment
	}

	// mapping
	parts := strings.Split(u.Path, "/")
	switch len(parts) {
	case 1:
		if parts[0][0] == '@' {
			parts[0] = parts[0][1:]
		}
		return &GitHubUserOrOrganization{Owner: parts[0]}, nil
	case 2:
		return &GitHubRepo{Owner: parts[0], Repo: parts[1]}, nil
	case 4:
		switch parts[2] {
		case "issues":
			return &GitHubIssue{Owner: parts[0], Repo: parts[1], ID: parts[3]}, nil
		case "milestone":
			return &GitHubMilestone{Owner: parts[0], Repo: parts[1], ID: parts[3]}, nil
		case "pull":
			return &GitHubPullRequest{Owner: parts[0], Repo: parts[1], ID: parts[3]}, nil
		case "issue-or-pull-request":
			return &GitHubIssueOrPullRequest{Owner: parts[0], Repo: parts[1], ID: parts[3]}, nil
		}
	}

	return nil, fmt.Errorf("failed to parse %q", input)
}
