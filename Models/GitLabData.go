package Models

import "github.com/xanzy/go-gitlab"

type GitLabData struct {
	//Contributors []gitlab.Contributor
	ConfigFileContent string
	//Issues []gitlab.Issue
	Events []gitlab.ContributionEvent

	Project gitlab.Project
	Members []gitlab.ProjectMember
}