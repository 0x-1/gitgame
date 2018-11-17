package Models

import "github.com/xanzy/go-gitlab"

type GitLabData struct {
	//Contributors []gitlab.Contributor
	ConfigFileContent string
	//Issues []gitlab.Issue
	ProjectEvents []gitlab.ContributionEvent
	PipelineList gitlab.PipelineList
	Project gitlab.Project
	Members []gitlab.ProjectMember
}