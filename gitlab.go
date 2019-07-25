package multipmuri

import (
	"fmt"
	"net/url"
	"strings"
)

//
// GitLabService
//

type GitLabService struct {
	Service
	*withGitLabCommon
}

func NewGitLabService(hostname string) *GitLabService {
	return &GitLabService{
		Service:          &service{},
		withGitLabCommon: &withGitLabCommon{hostname},
	}
}

func (e GitLabService) Canonical() string {
	return fmt.Sprintf("https://%s/", e.Hostname())
}

func (e GitLabService) RelDecodeString(input string) (Entity, error) {
	return gitLabRelDecodeString(e.Hostname(), "", "", input, false)
}

//
// GitLabIssue
//

type GitLabIssue struct {
	Issue
	*withGitLabCommon
	*withGitLabOwner
	*withGitLabRepo
	*withGitLabID
}

func NewGitLabIssue(hostname, owner, repo, id string) *GitLabIssue {
	return &GitLabIssue{
		Issue:            &issue{},
		withGitLabCommon: &withGitLabCommon{hostname},
		withGitLabOwner:  &withGitLabOwner{owner},
		withGitLabRepo:   &withGitLabRepo{repo},
		withGitLabID:     &withGitLabID{id},
	}
}

func (e GitLabIssue) Canonical() string {
	return fmt.Sprintf("https://%s/%s/%s/issues/%s", e.Hostname(), e.Owner(), e.Repo(), e.ID())
}

func (e GitLabIssue) RelDecodeString(input string) (Entity, error) {
	return gitLabRelDecodeString(e.Hostname(), e.Owner(), e.Repo(), input, false)
}

//
// GitLabMilestone
//

type GitLabMilestone struct {
	Milestone
	*withGitLabCommon
	*withGitLabOwner
	*withGitLabRepo
	*withGitLabID
}

func NewGitLabMilestone(hostname, owner, repo, id string) *GitLabMilestone {
	return &GitLabMilestone{
		Milestone:        &milestone{},
		withGitLabCommon: &withGitLabCommon{hostname},
		withGitLabOwner:  &withGitLabOwner{owner},
		withGitLabRepo:   &withGitLabRepo{repo},
		withGitLabID:     &withGitLabID{id},
	}
}

func (e GitLabMilestone) Canonical() string {
	return fmt.Sprintf("https://%s/%s/%s/-/milestones/%s", e.Hostname(), e.Owner(), e.Repo(), e.ID())
}

func (e GitLabMilestone) RelDecodeString(input string) (Entity, error) {
	return gitLabRelDecodeString(e.Hostname(), e.Owner(), e.Repo(), input, false)
}

//
// GitLabMergeRequest
//

type GitLabMergeRequest struct {
	MergeRequest
	*withGitLabCommon
	*withGitLabOwner
	*withGitLabRepo
	*withGitLabID
}

func NewGitLabMergeRequest(hostname, owner, repo, id string) *GitLabMergeRequest {
	return &GitLabMergeRequest{
		MergeRequest:     &mergeRequest{},
		withGitLabCommon: &withGitLabCommon{hostname},
		withGitLabOwner:  &withGitLabOwner{owner},
		withGitLabRepo:   &withGitLabRepo{repo},
		withGitLabID:     &withGitLabID{id},
	}
}

func (e GitLabMergeRequest) Canonical() string {
	return fmt.Sprintf("https://%s/%s/%s/merge_requests/%s", e.Hostname(), e.Owner(), e.Repo(), e.ID())
}

func (e GitLabMergeRequest) RelDecodeString(input string) (Entity, error) {
	return gitLabRelDecodeString(e.Hostname(), e.Owner(), e.Repo(), input, false)
}

//
// GitLabUserOrOrganization
//

type GitLabUserOrOrganization struct {
	UserOrOrganization
	*withGitLabCommon
	*withGitLabOwner
}

func NewGitLabUserOrOrganization(hostname, owner string) *GitLabUserOrOrganization {
	return &GitLabUserOrOrganization{
		UserOrOrganization: &userOrOrganization{},
		withGitLabCommon:   &withGitLabCommon{hostname},
		withGitLabOwner:    &withGitLabOwner{owner},
	}
}

func (e GitLabUserOrOrganization) Canonical() string {
	return fmt.Sprintf("https://%s/%s", e.Hostname(), e.Owner())
}

func (e GitLabUserOrOrganization) RelDecodeString(input string) (Entity, error) {
	return gitLabRelDecodeString(e.Hostname(), e.Owner(), "", input, false)
}

//
// GitLabUserOrOrganization
//

type GitLabOrganizationOrRepo struct {
	OrganizationOrProject
	*withGitLabCommon
	*withGitLabOwner
	*withGitLabRepo
}

func NewGitLabOrganizationOrRepo(hostname, owner, repo string) *GitLabOrganizationOrRepo {
	return &GitLabOrganizationOrRepo{
		OrganizationOrProject: &organizationOrProject{},
		withGitLabCommon:      &withGitLabCommon{hostname},
		withGitLabOwner:       &withGitLabOwner{owner},
		withGitLabRepo:        &withGitLabRepo{repo},
	}
}

func (e GitLabOrganizationOrRepo) Canonical() string {
	return fmt.Sprintf("https://%s/%s/%s", e.Hostname(), e.Owner(), e.Repo())
}

func (e GitLabOrganizationOrRepo) RelDecodeString(input string) (Entity, error) {
	return gitLabRelDecodeString(e.Hostname(), e.Owner(), e.Repo(), input, false)
}

//
// GitLabRepo
//

type GitLabRepo struct {
	Project
	*withGitLabCommon
	*withGitLabOwner
	*withGitLabRepo
}

func NewGitLabRepo(hostname, owner, repo string) *GitLabRepo {
	return &GitLabRepo{
		Project:          &project{},
		withGitLabCommon: &withGitLabCommon{hostname},
		withGitLabOwner:  &withGitLabOwner{owner},
		withGitLabRepo:   &withGitLabRepo{repo},
	}
}

func (e GitLabRepo) Canonical() string {
	return fmt.Sprintf("https://%s/%s/%s", e.Hostname(), e.Owner(), e.Repo())
}

func (e GitLabRepo) RelDecodeString(input string) (Entity, error) {
	return gitLabRelDecodeString(e.Hostname(), e.Owner(), e.Repo(), input, false)
}

//
// GitLabCommon
//

type withGitLabCommon struct{ hostname string }

func (e *withGitLabCommon) Provider() Provider { return GitLabProvider }
func (e *withGitLabCommon) Hostname() string {
	if e.hostname == "" {
		return "gitlab.com"
	}
	return e.hostname
}

type withGitLabOwner struct{ owner string }

func (e *withGitLabOwner) Owner() string { return e.owner }

type withGitLabRepo struct{ repo string }

func (e *withGitLabRepo) Repo() string { return e.repo }

type withGitLabID struct{ id string }

func (e *withGitLabID) ID() string { return e.id }

//
// Helpers
//

func gitLabRelDecodeString(hostname, owner, repo, input string, force bool) (Entity, error) {
	if hostname == "" {
		hostname = "gitlab.com"
	}
	u, err := url.Parse(input)
	if err != nil {
		return nil, err
	}
	if isProviderScheme(u.Scheme) { // gitlab://, gitlab://, ...
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
	if owner != "" && repo != "" && strings.HasPrefix(u.Path, "!") { // !42 from a repo
		return NewGitLabMergeRequest(hostname, owner, repo, u.Path[1:]), nil
	}
	if owner != "" && repo != "" && u.Path == "" && u.Fragment != "" { // #42 from a repo
		return NewGitLabIssue(hostname, owner, repo, u.Fragment), nil
	}
	if u.Path == "" && u.Fragment == "" {
		return NewGitLabService(hostname), nil
	}
	if strings.Contains(u.Path, "!") {
		parts := strings.Split(u.Path, "!")
		u.Path = fmt.Sprintf("%s/merge_requests/%s", parts[0], parts[1])
	}
	if u.Path != "" && u.Fragment != "" { // user/repo#42
		u.Path += "/issues/" + u.Fragment
	}
	parts := strings.Split(u.Path, "/")
	lenParts := len(parts)
	switch lenParts {
	case 1: // user or org
		if u.Host != "" && parts[0][0] != '@' {
			return NewGitLabUserOrOrganization(hostname, parts[0]), nil
		}
		if parts[0][0] == '@' {
			return NewGitLabUserOrOrganization(hostname, parts[0][1:]), nil
		}
	case 2:
		// org or rep
		return NewGitLabOrganizationOrRepo(hostname, parts[0], parts[1]), nil
	case 0:
		panic("should not happen")
	default: // more than 2
		switch {
		case parts[lenParts-2] == "issues":
			return NewGitLabIssue(
				hostname,
				strings.Join(parts[:lenParts-3], "/"),
				parts[lenParts-3],
				parts[lenParts-1],
			), nil
		case parts[lenParts-2] == "merge_requests":
			return NewGitLabMergeRequest(
				hostname,
				strings.Join(parts[:lenParts-3], "/"),
				parts[lenParts-3],
				parts[lenParts-1],
			), nil
		case parts[lenParts-2] == "milestones" && parts[lenParts-3] == "-":
			return NewGitLabMilestone(
				hostname,
				strings.Join(parts[:lenParts-4], "/"),
				parts[lenParts-4],
				parts[lenParts-1],
			), nil
		default:
			return NewGitLabRepo(
				hostname,
				strings.Join(parts[:lenParts-1], "/"),
				parts[lenParts-1],
			), nil
		}
	}

	return nil, fmt.Errorf("failed to parse %q", input)
}
