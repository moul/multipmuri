package multipmuri

//
// Entity
//

type Entity interface {
	WithKind
	WithProvider
	Canonical() string
	RelDecodeString(string) (Entity, error)
}

type WithKind interface {
	Kind() Kind
}

type WithProvider interface {
	Provider() Provider
}

//
// Enums
//

type Provider string

const (
	UnknownProvider Provider = "unknown-provider"
	GitHubProvider  Provider = "github"
	GitLabProvider  Provider = "gitlab"
	JiraProvider    Provider = "jira"
	TrelloProvider  Provider = "trello"
)

type Kind string

const (
	UnknownKind               Kind = "unknown-kind"
	IssueKind                 Kind = "issue"
	MergeRequestKind          Kind = "merge-request"
	ProviderKind              Kind = "provider"
	UserOrOrganizationKind    Kind = "user-or-organization"
	OrganizationOrProjectKind Kind = "organization-or-project"
	ServiceKind               Kind = "service"
	MilestoneKind             Kind = "milestone"
	IssueOrMergeRequestKind   Kind = "issue-or-merge-request"
	UserKind                  Kind = "user"
	ProjectKind               Kind = "project"
)

//
// Issue
//

type Issue interface {
	WithKind
	IsIssue()
}

type issue struct{}

func (issue) IsIssue()   {}
func (issue) Kind() Kind { return IssueKind }

//
// OrganizationOrProject
//

type OrganizationOrProject interface {
	WithKind
	IsOrganizationOrProject()
}

type organizationOrProject struct{}

func (organizationOrProject) IsOrganizationOrProject() {}
func (organizationOrProject) Kind() Kind               { return OrganizationOrProjectKind }

//
// IssueOrMergeRequest
//

type IssueOrMergeRequest interface {
	WithKind
	IsIssueOrMergeRequest()
}

type issueOrMergeRequest struct{}

func (issueOrMergeRequest) IsIssueOrMergeRequest() {}
func (issueOrMergeRequest) Kind() Kind             { return IssueOrMergeRequestKind }

//
// Milestone
//

type Milestone interface {
	WithKind
	IsMilestone()
}

type milestone struct{}

func (milestone) IsMilestone() {}
func (milestone) Kind() Kind   { return MilestoneKind }

//
// Project
//

type Project interface {
	WithKind
	IsProject()
}

type project struct{}

func (project) IsProject() {}
func (project) Kind() Kind { return ProjectKind }

//
// MergeRequest
//

type MergeRequest interface {
	WithKind
	IsMergeRequest()
}

type mergeRequest struct{}

func (mergeRequest) IsMergeRequest() {}
func (mergeRequest) Kind() Kind      { return MergeRequestKind }

//
// Service
//

type Service interface {
	WithKind
	IsService()
}

type service struct{}

func (service) IsService() {}
func (service) Kind() Kind { return ServiceKind }

//
// UserOrOrganization
//

type UserOrOrganization interface {
	WithKind
	IsUserOrOrganization()
}

type userOrOrganization struct{}

func (userOrOrganization) IsUserOrOrganization() {}
func (userOrOrganization) Kind() Kind            { return UserOrOrganizationKind }
