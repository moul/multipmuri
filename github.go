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

func NewGitHubService() *GitHubService {
	return &GitHubService{
		Service:          &service{},
		withGitHubCommon: &withGitHubCommon{},
	}
}

func (e GitHubService) Canonical() string {
	return "https://github.com/"
}

func (e GitHubService) RelDecodeString(input string) (Entity, error) {
	return gitHubRelDecodeString("", "", input, false)
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

func NewGitHubIssue(owner, repo, id string) *GitHubIssue {
	return &GitHubIssue{
		Issue:            &issue{},
		withGitHubCommon: &withGitHubCommon{},
		withGitHubOwner:  &withGitHubOwner{owner},
		withGitHubRepo:   &withGitHubRepo{repo},
		withGitHubID:     &withGitHubID{id},
	}
}

func (e GitHubIssue) Canonical() string {
	return fmt.Sprintf("https://github.com/%s/%s/issues/%s", e.Owner(), e.Repo(), e.ID())
}

func (e GitHubIssue) RelDecodeString(input string) (Entity, error) {
	return gitHubRelDecodeString(e.Owner(), e.Repo(), input, false)
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

func NewGitHubMilestone(owner, repo, id string) *GitHubMilestone {
	return &GitHubMilestone{
		Milestone:        &milestone{},
		withGitHubCommon: &withGitHubCommon{},
		withGitHubOwner:  &withGitHubOwner{owner},
		withGitHubRepo:   &withGitHubRepo{repo},
		withGitHubID:     &withGitHubID{id},
	}
}

func (e GitHubMilestone) Canonical() string {
	return fmt.Sprintf("https://github.com/%s/%s/milestone/%s", e.Owner(), e.Repo(), e.ID())
}

func (e GitHubMilestone) RelDecodeString(input string) (Entity, error) {
	return gitHubRelDecodeString(e.Owner(), e.Repo(), input, false)
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

func NewGitHubPullRequest(owner, repo, id string) *GitHubPullRequest {
	return &GitHubPullRequest{
		MergeRequest:     &mergeRequest{},
		withGitHubCommon: &withGitHubCommon{},
		withGitHubOwner:  &withGitHubOwner{owner},
		withGitHubRepo:   &withGitHubRepo{repo},
		withGitHubID:     &withGitHubID{id},
	}
}

func (e GitHubPullRequest) Canonical() string {
	return fmt.Sprintf("https://github.com/%s/%s/pull/%s", e.Owner(), e.Repo(), e.ID())
}

func (e GitHubPullRequest) RelDecodeString(input string) (Entity, error) {
	return gitHubRelDecodeString(e.Owner(), e.Repo(), input, false)
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

func NewGitHubIssueOrPullRequest(owner, repo, id string) *GitHubIssueOrPullRequest {
	return &GitHubIssueOrPullRequest{
		IssueOrMergeRequest: &issueOrMergeRequest{},
		withGitHubCommon:    &withGitHubCommon{},
		withGitHubOwner:     &withGitHubOwner{owner},
		withGitHubRepo:      &withGitHubRepo{repo},
		withGitHubID:        &withGitHubID{id},
	}
}

func (e GitHubIssueOrPullRequest) Canonical() string {
	return fmt.Sprintf("https://github.com/%s/%s/issues/%s", e.Owner(), e.Repo(), e.ID())
}

func (e GitHubIssueOrPullRequest) RelDecodeString(input string) (Entity, error) {
	return gitHubRelDecodeString(e.Owner(), e.Repo(), input, false)
}

//
// GitHubUserOrOrganization
//

type GitHubUserOrOrganization struct {
	UserOrOrganization
	*withGitHubCommon
	*withGitHubOwner
}

func NewGitHubUserOrOrganization(owner string) *GitHubUserOrOrganization {
	return &GitHubUserOrOrganization{
		UserOrOrganization: &userOrOrganization{},
		withGitHubCommon:   &withGitHubCommon{},
		withGitHubOwner:    &withGitHubOwner{owner},
	}
}

func (e GitHubUserOrOrganization) Canonical() string {
	return fmt.Sprintf("https://github.com/%s", e.Owner())
}

func (e GitHubUserOrOrganization) RelDecodeString(input string) (Entity, error) {
	return gitHubRelDecodeString(e.Owner(), "", input, false)
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

func NewGitHubRepo(owner, repo string) *GitHubRepo {
	return &GitHubRepo{
		Project:          &project{},
		withGitHubCommon: &withGitHubCommon{},
		withGitHubOwner:  &withGitHubOwner{owner},
		withGitHubRepo:   &withGitHubRepo{repo},
	}
}

func (e GitHubRepo) Canonical() string {
	return fmt.Sprintf("https://github.com/%s/%s", e.Owner(), e.Repo())
}

func (e GitHubRepo) RelDecodeString(input string) (Entity, error) {
	return gitHubRelDecodeString(e.Owner(), e.Repo(), input, false)
}

//
// GitHubCommon
//

type withGitHubCommon struct{}

func (e *withGitHubCommon) Provider() Provider { return GitHubProvider }

type withGitHubOwner struct{ owner string }

func (e *withGitHubOwner) Owner() string { return e.owner }

type withGitHubRepo struct{ repo string }

func (e *withGitHubRepo) Repo() string { return e.repo }

type withGitHubID struct{ id string }

func (e *withGitHubID) ID() string { return e.id }

//
// Helpers
//

func gitHubRelDecodeString(owner, repo, input string, force bool) (Entity, error) {
	u, err := url.Parse(input)
	if err != nil {
		return nil, err
	}
	if u.Host == "" && strings.HasPrefix(u.Path, "github.com") {
		u.Path = u.Path[10:]
		u.Host = "github.com"
	}
	if u.Scheme != "" && u.Scheme != "https" && !force {
		return DecodeString(input)
	}
	if u.Host != "" && u.Host != "github.com" && !force {
		return DecodeString(input)
	}
	if len(u.Path) > 0 && u.Path[0] == '/' {
		u.Path = u.Path[1:]
	}
	if owner != "" && repo != "" && u.Path == "" && u.Fragment != "" {
		return NewGitHubIssueOrPullRequest(owner, repo, u.Fragment), nil
	}
	if u.Path == "" && u.Fragment == "" {
		return NewGitHubService(), nil
	}
	if u.Path != "" && u.Fragment != "" {
		u.Path += "/issue-or-pull-request/" + u.Fragment
	}

	parts := strings.Split(u.Path, "/")
	switch len(parts) {
	case 1:
		if u.Host != "" && parts[0][0] != '@' {
			return NewGitHubUserOrOrganization(parts[0]), nil
		}
		if parts[0][0] == '@' {
			return NewGitHubUserOrOrganization(parts[0][1:]), nil
		}
	case 2:
		// FIXME: if starting with @ -> it's a team
		return NewGitHubRepo(parts[0], parts[1]), nil
	case 4:
		switch parts[2] {
		case "issues":
			return NewGitHubIssue(parts[0], parts[1], parts[3]), nil
		case "milestone":
			return NewGitHubMilestone(parts[0], parts[1], parts[3]), nil
		case "pull":
			return NewGitHubPullRequest(parts[0], parts[1], parts[3]), nil
		case "issue-or-pull-request":
			return NewGitHubIssueOrPullRequest(parts[0], parts[1], parts[3]), nil
		}
	}

	return nil, fmt.Errorf("failed to parse %q", input)
}
