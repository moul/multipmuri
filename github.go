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

func NewGitHubIssue(hostname, ownerID, repoID, id string) *GitHubIssue {
	return &GitHubIssue{
		Issue:        &issue{},
		withGitHubID: &withGitHubID{hostname, ownerID, repoID, id},
	}
}

func (e GitHubIssue) String() string {
	return fmt.Sprintf("https://%s/%s/%s/issues/%s", e.Hostname(), e.OwnerID(), e.RepoID(), e.ID())
}

func (e GitHubIssue) RelDecodeString(input string) (Entity, error) {
	return gitHubRelDecodeString(e.Hostname(), e.OwnerID(), e.RepoID(), input, false)
}

func (e GitHubIssue) Equals(other Entity) bool {
	switch other.(type) {
	case *GitHubIssueOrPullRequest, *GitHubIssue, *GitHubPullRequest:
		if typed, valid := other.(hasWithGitHubID); valid {
			return e.Hostname() == typed.Hostname() &&
				e.OwnerID() == typed.OwnerID() &&
				e.RepoID() == typed.RepoID() &&
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

func NewGitHubMilestone(hostname, ownerID, repoID, id string) *GitHubMilestone {
	return &GitHubMilestone{
		Milestone:    &milestone{},
		withGitHubID: &withGitHubID{hostname, ownerID, repoID, id},
	}
}

func (e GitHubMilestone) String() string {
	return fmt.Sprintf("https://%s/%s/%s/milestone/%s", e.Hostname(), e.OwnerID(), e.RepoID(), e.ID())
}

func (e GitHubMilestone) RelDecodeString(input string) (Entity, error) {
	return gitHubRelDecodeString(e.Hostname(), e.OwnerID(), e.RepoID(), input, false)
}

func (e GitHubMilestone) Equals(other Entity) bool {
	if typed, valid := other.(*GitHubMilestone); valid {
		return e.Hostname() == typed.Hostname() &&
			e.OwnerID() == typed.OwnerID() &&
			e.RepoID() == typed.RepoID() &&
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

func NewGitHubPullRequest(hostname, ownerID, repoID, id string) *GitHubPullRequest {
	return &GitHubPullRequest{
		MergeRequest: &mergeRequest{},
		withGitHubID: &withGitHubID{hostname, ownerID, repoID, id},
	}
}

func (e GitHubPullRequest) String() string {
	return fmt.Sprintf("https://%s/%s/%s/pull/%s", e.Hostname(), e.OwnerID(), e.RepoID(), e.ID())
}

func (e GitHubPullRequest) RelDecodeString(input string) (Entity, error) {
	return gitHubRelDecodeString(e.Hostname(), e.OwnerID(), e.RepoID(), input, false)
}

func (e GitHubPullRequest) Equals(other Entity) bool {
	switch other.(type) {
	case *GitHubIssueOrPullRequest, *GitHubIssue, *GitHubPullRequest:
		if typed, valid := other.(hasWithGitHubID); valid {
			return e.Hostname() == typed.Hostname() &&
				e.OwnerID() == typed.OwnerID() &&
				e.RepoID() == typed.RepoID() &&
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

func NewGitHubIssueOrPullRequest(hostname, ownerID, repoID, id string) *GitHubIssueOrPullRequest {
	return &GitHubIssueOrPullRequest{
		IssueOrMergeRequest: &issueOrMergeRequest{},
		withGitHubID:        &withGitHubID{hostname, ownerID, repoID, id},
	}
}

func (e GitHubIssueOrPullRequest) String() string {
	return fmt.Sprintf("https://%s/%s/%s/issues/%s", e.Hostname(), e.OwnerID(), e.RepoID(), e.ID())
}

func (e GitHubIssueOrPullRequest) RelDecodeString(input string) (Entity, error) {
	return gitHubRelDecodeString(e.Hostname(), e.OwnerID(), e.RepoID(), input, false)
}

func (e GitHubIssueOrPullRequest) Equals(other Entity) bool {
	switch other.(type) {
	case *GitHubIssueOrPullRequest, *GitHubIssue, *GitHubPullRequest:
		if typed, valid := other.(hasWithGitHubID); valid {
			return e.Hostname() == typed.Hostname() &&
				e.OwnerID() == typed.OwnerID() &&
				e.RepoID() == typed.RepoID() &&
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

func NewGitHubOwner(hostname, ownerID string) *GitHubOwner {
	return &GitHubOwner{
		UserOrOrganization: &userOrOrganization{},
		withGitHubOwner:    &withGitHubOwner{hostname, ownerID},
	}
}

func (e GitHubOwner) String() string {
	return fmt.Sprintf("https://%s/%s", e.Hostname(), e.OwnerID())
}

func (e GitHubOwner) RelDecodeString(input string) (Entity, error) {
	return gitHubRelDecodeString(e.Hostname(), e.OwnerID(), "", input, false)
}

func (e GitHubOwner) Equals(other Entity) bool {
	if typed, valid := other.(*GitHubOwner); valid {
		return e.Hostname() == typed.Hostname() &&
			e.OwnerID() == typed.OwnerID()
	}
	return false
}

func (e GitHubOwner) Contains(other Entity) bool {
	switch other.(type) {
	case *GitHubRepo, *GitHubMilestone, *GitHubIssueOrPullRequest, *GitHubIssue, *GitHubPullRequest:
		if typed, valid := other.(hasWithGitHubOwner); valid {
			return e.Hostname() == typed.Hostname() &&
				e.OwnerID() == typed.OwnerID()
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

func NewGitHubRepo(hostname, ownerID, repoID string) *GitHubRepo {
	return &GitHubRepo{
		Project:        &project{},
		withGitHubRepo: &withGitHubRepo{hostname, ownerID, repoID},
	}
}

func (e GitHubRepo) String() string {
	return fmt.Sprintf("https://%s/%s/%s", e.Hostname(), e.OwnerID(), e.RepoID())
}

func (e GitHubRepo) RelDecodeString(input string) (Entity, error) {
	return gitHubRelDecodeString(e.Hostname(), e.OwnerID(), e.RepoID(), input, false)
}

func (e GitHubRepo) Equals(other Entity) bool {
	if typed, valid := other.(*GitHubMilestone); valid {
		return e.Hostname() == typed.Hostname() &&
			e.OwnerID() == typed.OwnerID() &&
			e.RepoID() == typed.RepoID()
	}
	return false
}

func (e GitHubRepo) Contains(other Entity) bool {
	switch other.(type) {
	case *GitHubMilestone, *GitHubIssueOrPullRequest, *GitHubIssue, *GitHubPullRequest:
		if typed, valid := other.(hasWithGitHubRepo); valid {
			return e.Hostname() == typed.Hostname() &&
				e.OwnerID() == typed.OwnerID() &&
				e.RepoID() == typed.RepoID()
		}
	}
	return false
}

//
// GitHubCommon
//

// githubHostname

type hasWithGitHubHostname interface {
	Hostname() string
}

type withGitHubHostname struct{ hostname string }

func (e *withGitHubHostname) Provider() Provider      { return GitHubProvider }
func (e *withGitHubHostname) Hostname() string        { return githubHostname(e.hostname) }
func (e *withGitHubHostname) Service() *GitHubService { return NewGitHubService(e.hostname) }
func (e *withGitHubHostname) ServiceEntity() Entity   { return e.Service() }
func (e *withGitHubHostname) Owner(ownerID string) *GitHubOwner {
	return NewGitHubOwner(e.hostname, ownerID)
}
func (e *withGitHubHostname) OwnerEntity(ownerID string) Entity { return e.Owner(ownerID) }

//githubOwner

type hasWithGitHubOwner interface {
	Hostname() string
	OwnerID() string
}

type withGitHubOwner struct{ hostname, ownerID string }

func (e *withGitHubOwner) Provider() Provider      { return GitHubProvider }
func (e *withGitHubOwner) Hostname() string        { return githubHostname(e.hostname) }
func (e *withGitHubOwner) Service() *GitHubService { return NewGitHubService(e.hostname) }
func (e *withGitHubOwner) ServiceEntity() Entity   { return e.Service() }
func (e *withGitHubOwner) OwnerID() string         { return e.ownerID }
func (e *withGitHubOwner) Owner() *GitHubOwner     { return NewGitHubOwner(e.hostname, e.ownerID) }
func (e *withGitHubOwner) OwnerEntity() Entity     { return e.Owner() }
func (e *withGitHubOwner) Repo(repoID string) *GitHubRepo {
	return NewGitHubRepo(e.hostname, e.ownerID, repoID)
}
func (e *withGitHubOwner) RepoEntity(repoID string) Entity { return e.Repo(repoID) }

// githubRepo

type hasWithGitHubRepo interface {
	Hostname() string
	OwnerID() string
	RepoID() string
}

type withGitHubRepo struct{ hostname, ownerID, repoID string }

func (e *withGitHubRepo) Provider() Provider      { return GitHubProvider }
func (e *withGitHubRepo) Hostname() string        { return githubHostname(e.hostname) }
func (e *withGitHubRepo) OwnerID() string         { return e.ownerID }
func (e *withGitHubRepo) RepoID() string          { return e.repoID }
func (e *withGitHubRepo) Service() *GitHubService { return NewGitHubService(e.hostname) }
func (e *withGitHubRepo) ServiceEntity() Entity   { return e.Service() }
func (e *withGitHubRepo) Owner() *GitHubOwner     { return NewGitHubOwner(e.hostname, e.ownerID) }
func (e *withGitHubRepo) OwnerEntity() Entity     { return e.Owner() }
func (e *withGitHubRepo) Repo() *GitHubRepo       { return NewGitHubRepo(e.hostname, e.ownerID, e.repoID) }
func (e *withGitHubRepo) RepoEntity() Entity      { return e.Repo() }
func (e *withGitHubRepo) Issue(id string) *GitHubIssue {
	return NewGitHubIssue(e.hostname, e.ownerID, e.repoID, id)
}
func (e *withGitHubRepo) IssueEntity(id string) Entity { return e.Issue(id) }
func (e *withGitHubRepo) Milestone(id string) *GitHubMilestone {
	return NewGitHubMilestone(e.hostname, e.ownerID, e.repoID, id)
}
func (e *withGitHubRepo) MilestoneEntity(id string) Entity { return e.Milestone(id) }

// githubID (issue, milestone, PR, ...))

type hasWithGitHubID interface {
	Hostname() string
	OwnerID() string
	RepoID() string
	ID() string
}

type withGitHubID struct{ hostname, ownerID, repoID, id string }

func (e *withGitHubID) Provider() Provider      { return GitHubProvider }
func (e *withGitHubID) Hostname() string        { return githubHostname(e.hostname) }
func (e *withGitHubID) OwnerID() string         { return e.ownerID }
func (e *withGitHubID) RepoID() string          { return e.repoID }
func (e *withGitHubID) ID() string              { return e.id }
func (e *withGitHubID) Service() *GitHubService { return NewGitHubService(e.hostname) }
func (e *withGitHubID) ServiceEntity() Entity   { return e.Service() }
func (e *withGitHubID) Owner() *GitHubOwner     { return NewGitHubOwner(e.hostname, e.ownerID) }
func (e *withGitHubID) OwnerEntity() Entity     { return e.Owner() }
func (e *withGitHubID) Repo() *GitHubRepo       { return NewGitHubRepo(e.hostname, e.ownerID, e.repoID) }
func (e *withGitHubID) RepoEntity() Entity      { return e.Repo() }

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
