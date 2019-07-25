package multipmuri

import (
	"fmt"
	"net/url"
	"strings"
)

//
// GitHubService
//

type GitHubService struct {
	Service
	*withGitHubCommon
}

func NewGitHubService(hostname string) *GitHubService {
	return &GitHubService{
		Service:          &service{},
		withGitHubCommon: &withGitHubCommon{hostname},
	}
}

func (e GitHubService) Canonical() string {
	return fmt.Sprintf("https://%s/", e.Hostname())
}

func (e GitHubService) RelDecodeString(input string) (Entity, error) {
	return gitHubRelDecodeString(e.Hostname(), "", "", input, false)
}

//
// GitHubIssue
//

type GitHubIssue struct {
	Issue
	*withGitHubCommon
	*withGitHubOwner
	*withGitHubRepo
	*withGitHubID
}

func NewGitHubIssue(hostname, owner, repo, id string) *GitHubIssue {
	return &GitHubIssue{
		Issue:            &issue{},
		withGitHubCommon: &withGitHubCommon{hostname},
		withGitHubOwner:  &withGitHubOwner{owner},
		withGitHubRepo:   &withGitHubRepo{repo},
		withGitHubID:     &withGitHubID{id},
	}
}

func (e GitHubIssue) Canonical() string {
	return fmt.Sprintf("https://%s/%s/%s/issues/%s", e.Hostname(), e.Owner(), e.Repo(), e.ID())
}

func (e GitHubIssue) RelDecodeString(input string) (Entity, error) {
	return gitHubRelDecodeString(e.Hostname(), e.Owner(), e.Repo(), input, false)
}

//
// GitHubMilestone
//

type GitHubMilestone struct {
	Milestone
	*withGitHubCommon
	*withGitHubOwner
	*withGitHubRepo
	*withGitHubID
}

func NewGitHubMilestone(hostname, owner, repo, id string) *GitHubMilestone {
	return &GitHubMilestone{
		Milestone:        &milestone{},
		withGitHubCommon: &withGitHubCommon{hostname},
		withGitHubOwner:  &withGitHubOwner{owner},
		withGitHubRepo:   &withGitHubRepo{repo},
		withGitHubID:     &withGitHubID{id},
	}
}

func (e GitHubMilestone) Canonical() string {
	return fmt.Sprintf("https://%s/%s/%s/milestone/%s", e.Hostname(), e.Owner(), e.Repo(), e.ID())
}

func (e GitHubMilestone) RelDecodeString(input string) (Entity, error) {
	return gitHubRelDecodeString(e.Hostname(), e.Owner(), e.Repo(), input, false)
}

//
// GitHubPullRequest
//

type GitHubPullRequest struct {
	MergeRequest
	*withGitHubCommon
	*withGitHubOwner
	*withGitHubRepo
	*withGitHubID
}

func NewGitHubPullRequest(hostname, owner, repo, id string) *GitHubPullRequest {
	return &GitHubPullRequest{
		MergeRequest:     &mergeRequest{},
		withGitHubCommon: &withGitHubCommon{hostname},
		withGitHubOwner:  &withGitHubOwner{owner},
		withGitHubRepo:   &withGitHubRepo{repo},
		withGitHubID:     &withGitHubID{id},
	}
}

func (e GitHubPullRequest) Canonical() string {
	return fmt.Sprintf("https://%s/%s/%s/pull/%s", e.Hostname(), e.Owner(), e.Repo(), e.ID())
}

func (e GitHubPullRequest) RelDecodeString(input string) (Entity, error) {
	return gitHubRelDecodeString(e.Hostname(), e.Owner(), e.Repo(), input, false)
}

//
// GitHubIssueOrPullRequest
//

type GitHubIssueOrPullRequest struct {
	IssueOrMergeRequest
	*withGitHubCommon
	*withGitHubOwner
	*withGitHubRepo
	*withGitHubID
}

func NewGitHubIssueOrPullRequest(hostname, owner, repo, id string) *GitHubIssueOrPullRequest {
	return &GitHubIssueOrPullRequest{
		IssueOrMergeRequest: &issueOrMergeRequest{},
		withGitHubCommon:    &withGitHubCommon{hostname},
		withGitHubOwner:     &withGitHubOwner{owner},
		withGitHubRepo:      &withGitHubRepo{repo},
		withGitHubID:        &withGitHubID{id},
	}
}

func (e GitHubIssueOrPullRequest) Canonical() string {
	return fmt.Sprintf("https://%s/%s/%s/issues/%s", e.Hostname(), e.Owner(), e.Repo(), e.ID())
}

func (e GitHubIssueOrPullRequest) RelDecodeString(input string) (Entity, error) {
	return gitHubRelDecodeString(e.Hostname(), e.Owner(), e.Repo(), input, false)
}

//
// GitHubUserOrOrganization
//

type GitHubUserOrOrganization struct {
	UserOrOrganization
	*withGitHubCommon
	*withGitHubOwner
}

func NewGitHubUserOrOrganization(hostname, owner string) *GitHubUserOrOrganization {
	return &GitHubUserOrOrganization{
		UserOrOrganization: &userOrOrganization{},
		withGitHubCommon:   &withGitHubCommon{hostname},
		withGitHubOwner:    &withGitHubOwner{owner},
	}
}

func (e GitHubUserOrOrganization) Canonical() string {
	return fmt.Sprintf("https://%s/%s", e.Hostname(), e.Owner())
}

func (e GitHubUserOrOrganization) RelDecodeString(input string) (Entity, error) {
	return gitHubRelDecodeString(e.Hostname(), e.Owner(), "", input, false)
}

//
// GitHubRepo
//

type GitHubRepo struct {
	Project
	*withGitHubCommon
	*withGitHubOwner
	*withGitHubRepo
}

func NewGitHubRepo(hostname, owner, repo string) *GitHubRepo {
	return &GitHubRepo{
		Project:          &project{},
		withGitHubCommon: &withGitHubCommon{hostname},
		withGitHubOwner:  &withGitHubOwner{owner},
		withGitHubRepo:   &withGitHubRepo{repo},
	}
}

func (e GitHubRepo) Canonical() string {
	return fmt.Sprintf("https://%s/%s/%s", e.Hostname(), e.Owner(), e.Repo())
}

func (e GitHubRepo) RelDecodeString(input string) (Entity, error) {
	return gitHubRelDecodeString(e.Hostname(), e.Owner(), e.Repo(), input, false)
}

//
// GitHubCommon
//

type withGitHubCommon struct{ hostname string }

func (e *withGitHubCommon) Provider() Provider { return GitHubProvider }
func (e *withGitHubCommon) Hostname() string {
	if e.hostname == "" {
		return "github.com"
	}
	return e.hostname
}

type withGitHubOwner struct{ owner string }

func (e *withGitHubOwner) Owner() string { return e.owner }

type withGitHubRepo struct{ repo string }

func (e *withGitHubRepo) Repo() string { return e.repo }

type withGitHubID struct{ id string }

func (e *withGitHubID) ID() string { return e.id }

//
// Helpers
//

func gitHubRelDecodeString(hostname, owner, repo, input string, force bool) (Entity, error) {
	if hostname == "" {
		hostname = "github.com"
	}
	u, err := url.Parse(input)
	if err != nil {
		return nil, err
	}
	if isProviderScheme(u.Scheme) { // github://, gitlab://, ...
		return DecodeString(input)
	}
	u.Path = strings.Trim(u.Path, "/")
	if u.Host == "" && len(u.Path) > 0 { // domain.com/a/b
		u.Host = getHostname(u.Path)
		if u.Host != "" {
			u.Path = u.Path[len(u.Host):]
			u.Path = strings.Trim(u.Path, "/")
		}
	}
	if u.Host != "" && u.Host != hostname && !force {
		return DecodeString(input)
	}
	if owner != "" && repo != "" && u.Path == "" && u.Fragment != "" { // #42 from a repo
		return NewGitHubIssueOrPullRequest(hostname, owner, repo, u.Fragment), nil
	}
	if u.Path == "" && u.Fragment == "" {
		return NewGitHubService(hostname), nil
	}
	if u.Path != "" && u.Fragment != "" { // user/repo#42
		u.Path += "/issue-or-pull-request/" + u.Fragment
	}
	parts := strings.Split(u.Path, "/")
	switch len(parts) {
	case 1:
		if u.Host != "" && parts[0][0] != '@' {
			return NewGitHubUserOrOrganization(hostname, parts[0]), nil
		}
		if parts[0][0] == '@' {
			return NewGitHubUserOrOrganization(hostname, parts[0][1:]), nil
		}
	case 2:
		// FIXME: if starting with @ -> it's a team
		return NewGitHubRepo(hostname, parts[0], parts[1]), nil
	case 4:
		switch parts[2] {
		case "issues":
			return NewGitHubIssue(hostname, parts[0], parts[1], parts[3]), nil
		case "milestone":
			return NewGitHubMilestone(hostname, parts[0], parts[1], parts[3]), nil
		case "pull":
			return NewGitHubPullRequest(hostname, parts[0], parts[1], parts[3]), nil
		case "issue-or-pull-request":
			return NewGitHubIssueOrPullRequest(hostname, parts[0], parts[1], parts[3]), nil
		}
	}

	return nil, fmt.Errorf("failed to parse %q", input)
}
