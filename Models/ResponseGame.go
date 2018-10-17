package Models

import "github.com/xanzy/go-gitlab"

type ResponseGame struct {
	Contributors []gitlab.Contributor
	Project gitlab.Project
	Issues []gitlab.Issue
	Events []gitlab.ContributionEvent
}