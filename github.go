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
	*withGitHubHostname
}

func NewGitHubService(hostname string) *GitHubService {
	return &GitHubService{
		Service:            &service{},
		withGitHubHostname: &withGitHubHostname{hostname},
	}
}

func (e GitHubService) String() string {
	return fmt.Sprintf("https://%s/", e.Hostname())
}

func (e GitHubService) RelDecodeString(input string) (Entity, error) {
	return gitHubRelDecodeString(e.Hostname(), "", "", input, false)
}

func (e GitHubService) Equals(other Entity) bool {
	if typed, valid := other.(*GitHubService); valid {
		return e.Hostname() == typed.Hostname()
	}
	return false
}

func (e GitHubService) Contains(other Entity) bool {
	switch other.(type) {
	case *GitHubRepo, *GitHubOwner, *GitHubMilestone, *GitHubIssueOrPullRequest, *GitHubIssue, *GitHubPullRequest:
		if typed, valid := other.(hasWithGitHubHostname); valid {
			return e.Hostname() == typed.Hostname()
		}
	}
	return false
}

//
// GitHubIssue
//

type GitHubIssue struct {
	Issue
	*withGitHubID
}

func NewGitHubIssue(hostname, owner, repo, id string) *GitHubIssue {
	return &GitHubIssue{
		Issue:        &issue{},
		withGitHubID: &withGitHubID{hostname, owner, repo, id},
	}
}

func (e GitHubIssue) String() string {
	return fmt.Sprintf("https://%s/%s/%s/issues/%s", e.Hostname(), e.Owner(), e.Repo(), e.ID())
}

func (e GitHubIssue) RelDecodeString(input string) (Entity, error) {
	return gitHubRelDecodeString(e.Hostname(), e.Owner(), e.Repo(), input, false)
}

func (e GitHubIssue) Equals(other Entity) bool {
	switch other.(type) {
	case *GitHubIssueOrPullRequest, *GitHubIssue, *GitHubPullRequest:
		if typed, valid := other.(hasWithGitHubID); valid {
			return e.Hostname() == typed.Hostname() &&
				e.Owner() == typed.Owner() &&
				e.Repo() == typed.Repo() &&
				e.ID() == typed.ID()
		}
	}
	return false
}

func (e GitHubIssue) Contains(other Entity) bool {
	return false
}

//
// GitHubMilestone
//

type GitHubMilestone struct {
	Milestone
	*withGitHubID
}

func NewGitHubMilestone(hostname, owner, repo, id string) *GitHubMilestone {
	return &GitHubMilestone{
		Milestone:    &milestone{},
		withGitHubID: &withGitHubID{hostname, owner, repo, id},
	}
}

func (e GitHubMilestone) String() string {
	return fmt.Sprintf("https://%s/%s/%s/milestone/%s", e.Hostname(), e.Owner(), e.Repo(), e.ID())
}

func (e GitHubMilestone) RelDecodeString(input string) (Entity, error) {
	return gitHubRelDecodeString(e.Hostname(), e.Owner(), e.Repo(), input, false)
}

func (e GitHubMilestone) Equals(other Entity) bool {
	if typed, valid := other.(*GitHubMilestone); valid {
		return e.Hostname() == typed.Hostname() &&
			e.Owner() == typed.Owner() &&
			e.Repo() == typed.Repo() &&
			e.ID() == typed.ID()
	}
	return false
}

func (e GitHubMilestone) Contains(other Entity) bool {
	return false
}

//
// GitHubPullRequest
//

type GitHubPullRequest struct {
	MergeRequest
	*withGitHubID
}

func NewGitHubPullRequest(hostname, owner, repo, id string) *GitHubPullRequest {
	return &GitHubPullRequest{
		MergeRequest: &mergeRequest{},
		withGitHubID: &withGitHubID{hostname, owner, repo, id},
	}
}

func (e GitHubPullRequest) String() string {
	return fmt.Sprintf("https://%s/%s/%s/pull/%s", e.Hostname(), e.Owner(), e.Repo(), e.ID())
}

func (e GitHubPullRequest) RelDecodeString(input string) (Entity, error) {
	return gitHubRelDecodeString(e.Hostname(), e.Owner(), e.Repo(), input, false)
}

func (e GitHubPullRequest) Equals(other Entity) bool {
	switch other.(type) {
	case *GitHubIssueOrPullRequest, *GitHubIssue, *GitHubPullRequest:
		if typed, valid := other.(hasWithGitHubID); valid {
			return e.Hostname() == typed.Hostname() &&
				e.Owner() == typed.Owner() &&
				e.Repo() == typed.Repo() &&
				e.ID() == typed.ID()
		}
	}
	return false
}

func (e GitHubPullRequest) Contains(other Entity) bool {
	return false
}

//
// GitHubIssueOrPullRequest
//

type GitHubIssueOrPullRequest struct {
	IssueOrMergeRequest
	*withGitHubID
}

func NewGitHubIssueOrPullRequest(hostname, owner, repo, id string) *GitHubIssueOrPullRequest {
	return &GitHubIssueOrPullRequest{
		IssueOrMergeRequest: &issueOrMergeRequest{},
		withGitHubID:        &withGitHubID{hostname, owner, repo, id},
	}
}

func (e GitHubIssueOrPullRequest) String() string {
	return fmt.Sprintf("https://%s/%s/%s/issues/%s", e.Hostname(), e.Owner(), e.Repo(), e.ID())
}

func (e GitHubIssueOrPullRequest) RelDecodeString(input string) (Entity, error) {
	return gitHubRelDecodeString(e.Hostname(), e.Owner(), e.Repo(), input, false)
}

func (e GitHubIssueOrPullRequest) Equals(other Entity) bool {
	switch other.(type) {
	case *GitHubIssueOrPullRequest, *GitHubIssue, *GitHubPullRequest:
		if typed, valid := other.(hasWithGitHubID); valid {
			return e.Hostname() == typed.Hostname() &&
				e.Owner() == typed.Owner() &&
				e.Repo() == typed.Repo() &&
				e.ID() == typed.ID()
		}
	}
	return false
}

func (e GitHubIssueOrPullRequest) Contains(other Entity) bool {
	return false
}

//
// GitHubOwner
//

type GitHubOwner struct {
	UserOrOrganization
	*withGitHubOwner
}

func NewGitHubOwner(hostname, owner string) *GitHubOwner {
	return &GitHubOwner{
		UserOrOrganization: &userOrOrganization{},
		withGitHubOwner:    &withGitHubOwner{hostname, owner},
	}
}

func (e GitHubOwner) String() string {
	return fmt.Sprintf("https://%s/%s", e.Hostname(), e.Owner())
}

func (e GitHubOwner) RelDecodeString(input string) (Entity, error) {
	return gitHubRelDecodeString(e.Hostname(), e.Owner(), "", input, false)
}

func (e GitHubOwner) Equals(other Entity) bool {
	if typed, valid := other.(*GitHubOwner); valid {
		return e.Hostname() == typed.Hostname() &&
			e.Owner() == typed.Owner()
	}
	return false
}

func (e GitHubOwner) Contains(other Entity) bool {
	switch other.(type) {
	case *GitHubRepo, *GitHubMilestone, *GitHubIssueOrPullRequest, *GitHubIssue, *GitHubPullRequest:
		if typed, valid := other.(hasWithGitHubOwner); valid {
			return e.Hostname() == typed.Hostname() &&
				e.Owner() == typed.Owner()
		}
	}
	return false
}

//
// GitHubRepo
//

type GitHubRepo struct {
	Project
	*withGitHubRepo
}

func NewGitHubRepo(hostname, owner, repo string) *GitHubRepo {
	return &GitHubRepo{
		Project:        &project{},
		withGitHubRepo: &withGitHubRepo{hostname, owner, repo},
	}
}

func (e GitHubRepo) String() string {
	return fmt.Sprintf("https://%s/%s/%s", e.Hostname(), e.Owner(), e.Repo())
}

func (e GitHubRepo) RelDecodeString(input string) (Entity, error) {
	return gitHubRelDecodeString(e.Hostname(), e.Owner(), e.Repo(), input, false)
}

func (e GitHubRepo) Equals(other Entity) bool {
	if typed, valid := other.(*GitHubMilestone); valid {
		return e.Hostname() == typed.Hostname() &&
			e.Owner() == typed.Owner() &&
			e.Repo() == typed.Repo()
	}
	return false
}

func (e GitHubRepo) Contains(other Entity) bool {
	switch other.(type) {
	case *GitHubMilestone, *GitHubIssueOrPullRequest, *GitHubIssue, *GitHubPullRequest:
		if typed, valid := other.(hasWithGitHubRepo); valid {
			return e.Hostname() == typed.Hostname() &&
				e.Owner() == typed.Owner() &&
				e.Repo() == typed.Repo()
		}
	}
	return false
}

//
// GitHubCommon
//

type hasWithGitHubHostname interface {
	Hostname() string
}

type withGitHubHostname struct{ hostname string }

func (e *withGitHubHostname) Provider() Provider { return GitHubProvider }
func (e *withGitHubHostname) Hostname() string   { return githubHostname(e.hostname) }
func (e *withGitHubHostname) OwnerEntity(owner string) *GitHubOwner {
	return NewGitHubOwner(e.hostname, owner)
}

type hasWithGitHubOwner interface {
	Hostname() string
	Owner() string
}

type withGitHubOwner struct{ hostname, owner string }

func (e *withGitHubOwner) Provider() Provider            { return GitHubProvider }
func (e *withGitHubOwner) Hostname() string              { return githubHostname(e.hostname) }
func (e *withGitHubOwner) Owner() string                 { return e.owner }
func (e *withGitHubOwner) ServiceEntity() *GitHubService { return NewGitHubService(e.hostname) }
func (e *withGitHubOwner) RepoEntity(repo string) *GitHubRepo {
	return NewGitHubRepo(e.hostname, e.owner, repo)
}

type hasWithGitHubRepo interface {
	Hostname() string
	Owner() string
	Repo() string
}

type withGitHubRepo struct{ hostname, owner, repo string }

func (e *withGitHubRepo) Provider() Provider            { return GitHubProvider }
func (e *withGitHubRepo) Hostname() string              { return githubHostname(e.hostname) }
func (e *withGitHubRepo) Owner() string                 { return e.owner }
func (e *withGitHubRepo) Repo() string                  { return e.repo }
func (e *withGitHubRepo) ServiceEntity() *GitHubService { return NewGitHubService(e.hostname) }
func (e *withGitHubRepo) OwnerEntity() *GitHubOwner     { return NewGitHubOwner(e.hostname, e.owner) }
func (e *withGitHubRepo) RepoEntity() *GitHubRepo {
	return NewGitHubRepo(e.hostname, e.owner, e.repo)
}
func (e *withGitHubRepo) IssueEntity(id string) *GitHubIssue {
	return NewGitHubIssue(e.hostname, e.owner, e.repo, id)
}
func (e *withGitHubRepo) MilestoneEntity(id string) *GitHubMilestone {
	return NewGitHubMilestone(e.hostname, e.owner, e.repo, id)
}

type hasWithGitHubID interface {
	Hostname() string
	Owner() string
	Repo() string
	ID() string
}

type withGitHubID struct{ hostname, owner, repo, id string }

func (e *withGitHubID) Provider() Provider            { return GitHubProvider }
func (e *withGitHubID) Hostname() string              { return githubHostname(e.hostname) }
func (e *withGitHubID) Owner() string                 { return e.owner }
func (e *withGitHubID) Repo() string                  { return e.repo }
func (e *withGitHubID) ID() string                    { return e.id }
func (e *withGitHubID) ServiceEntity() *GitHubService { return NewGitHubService(e.hostname) }
func (e *withGitHubID) OwnerEntity() *GitHubOwner     { return NewGitHubOwner(e.hostname, e.owner) }
func (e *withGitHubID) RepoEntity() *GitHubRepo       { return NewGitHubRepo(e.hostname, e.owner, e.repo) }

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
			return NewGitHubOwner(hostname, parts[0]), nil
		}
		if parts[0][0] == '@' {
			return NewGitHubOwner(hostname, parts[0][1:]), nil
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

func githubHostname(input string) string {
	if input == "" {
		return "github.com"
	}
	return input
}
