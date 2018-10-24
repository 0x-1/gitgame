package Models

import "github.com/xanzy/go-gitlab"

type GitLabData struct {
	Contributors []gitlab.Contributor

	Issues []gitlab.Issue
	Events []gitlab.ContributionEvent

	Project gitlab.Project
	Members []gitlab.ProjectMember
}