package multipmuri

type DecodedMultipmuri interface {
	Canonical() string
	Provider() Provider
	Kind() Kind
	RelDecodeString(string) (DecodedMultipmuri, error)
}

type Issue interface {
	DecodedMultipmuri
	IsIssue() bool
}

type MergeRequest interface {
	DecodedMultipmuri
	IsMergeRequest() bool
}

type Service interface {
	DecodedMultipmuri
	IsService() bool
}

type UserOrOrganization interface {
	DecodedMultipmuri
	IsUserOrOrganization() bool
}

type Provider string

const (
	UnknownProvider Provider = "unknown-provider"
	GitHubProvider  Provider = "github"
	GitLabProvider  Provider = "gitlab"
	//GitHubEnterpriseProvider Provider = "github-enterprise"
	//JiraProvider             Provider = "jira"
	//TrelloProvider           Provider = "trello"
)

type Kind string

const (
	UnknownKind             Kind = "unknown-kind"
	IssueKind               Kind = "issue"
	MergeRequestKind        Kind = "merge-request"
	ProviderKind            Kind = "provider"
	UserOrOrganizationKind  Kind = "user-or-organization"
	ServiceKind             Kind = "service"
	ProjectKind             Kind = "project"
	MilestoneKind           Kind = "milestone"
	IssueOrMergeRequestKind Kind = "issue-or-merge-request"
	//UserKind                Kind = "user"
	//ProjectKind             Kind = "project"
)
