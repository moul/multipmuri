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
	*withGitLabHostname
}

func NewGitLabService(hostname string) *GitLabService {
	return &GitLabService{
		Service:            &service{},
		withGitLabHostname: &withGitLabHostname{hostname},
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
	*withGitLabID
}

func NewGitLabIssue(hostname, owner, repo, id string) *GitLabIssue {
	return &GitLabIssue{
		Issue:        &issue{},
		withGitLabID: &withGitLabID{hostname, owner, repo, id},
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
	*withGitLabID
}

func NewGitLabMilestone(hostname, owner, repo, id string) *GitLabMilestone {
	return &GitLabMilestone{
		Milestone:    &milestone{},
		withGitLabID: &withGitLabID{hostname, owner, repo, id},
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
	*withGitLabID
}

func NewGitLabMergeRequest(hostname, owner, repo, id string) *GitLabMergeRequest {
	return &GitLabMergeRequest{
		MergeRequest: &mergeRequest{},
		withGitLabID: &withGitLabID{hostname, owner, repo, id},
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
	*withGitLabOwner
}

func NewGitLabUserOrOrganization(hostname, owner string) *GitLabUserOrOrganization {
	return &GitLabUserOrOrganization{
		UserOrOrganization: &userOrOrganization{},
		withGitLabOwner:    &withGitLabOwner{hostname, owner},
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
	*withGitLabRepo
}

func NewGitLabOrganizationOrRepo(hostname, owner, repo string) *GitLabOrganizationOrRepo {
	return &GitLabOrganizationOrRepo{
		OrganizationOrProject: &organizationOrProject{},
		withGitLabRepo:        &withGitLabRepo{hostname, owner, repo},
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
	*withGitLabRepo
}

func NewGitLabRepo(hostname, owner, repo string) *GitLabRepo {
	return &GitLabRepo{
		Project:        &project{},
		withGitLabRepo: &withGitLabRepo{hostname, owner, repo},
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

type withGitLabHostname struct{ hostname string }

func (e *withGitLabHostname) Provider() Provider            { return GitLabProvider }
func (e *withGitLabHostname) Hostname() string              { return gitlabHostname(e.hostname) }
func (e *withGitLabHostname) ServiceEntity() *GitLabService { return NewGitLabService(e.hostname) }

type withGitLabOwner struct{ hostname, owner string }

func (e *withGitLabOwner) Provider() Provider            { return GitLabProvider }
func (e *withGitLabOwner) Hostname() string              { return gitlabHostname(e.hostname) }
func (e *withGitLabOwner) Owner() string                 { return e.owner }
func (e *withGitLabOwner) ServiceEntity() *GitLabService { return NewGitLabService(e.hostname) }
func (e *withGitLabOwner) RepoEntity(repo string) *GitLabRepo {
	return NewGitLabRepo(e.hostname, e.owner, repo)
}

type withGitLabRepo struct{ hostname, owner, repo string }

func (e *withGitLabRepo) Provider() Provider            { return GitLabProvider }
func (e *withGitLabRepo) Hostname() string              { return gitlabHostname(e.hostname) }
func (e *withGitLabRepo) Owner() string                 { return e.owner }
func (e *withGitLabRepo) Repo() string                  { return e.repo }
func (e *withGitLabRepo) ServiceEntity() *GitLabService { return NewGitLabService(e.hostname) }
func (e *withGitLabRepo) RepoEntity() *GitLabRepo       { return NewGitLabRepo(e.hostname, e.owner, e.repo) }

type withGitLabID struct{ hostname, owner, repo, id string }

func (e *withGitLabID) Provider() Provider      { return GitLabProvider }
func (e *withGitLabID) Hostname() string        { return gitlabHostname(e.hostname) }
func (e *withGitLabID) Owner() string           { return e.owner }
func (e *withGitLabID) Repo() string            { return e.repo }
func (e *withGitLabID) ID() string              { return e.id }
func (e *withGitLabID) RepoEntity() *GitLabRepo { return NewGitLabRepo(e.hostname, e.owner, e.repo) }

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

func gitlabHostname(input string) string {
	if input == "" {
		return "gitlab.com"
	}
	return input
}
