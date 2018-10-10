package GameManager

import (
	"github.com/xanzy/go-gitlab"
	"strings"
	"github.com/pkg/errors"
)
func M_GetProjectByName(git *gitlab.Client ,projectName string)(gitlab.Project, error) {
	opt := gitlab.ListProjectsOptions{

		Membership:gitlab.Bool(true),
		MinAccessLevel:gitlab.AccessLevel(30),
		ListOptions: gitlab.ListOptions{
			PerPage: 10,
			Page:    1,
		},
	}

	var allProjects []*gitlab.Project
	for {
		projects, resp, err := git.Projects.ListProjects(&opt)
		if(err == nil) {
			allProjects = append(allProjects, projects...)

			if opt.Page >= resp.TotalPages {
				break
			}

			opt.Page = resp.NextPage
		} else {
			return gitlab.Project{}, err
		}
	}

	for _, element := range allProjects {
		if(strings.EqualFold(element.Name, projectName)) {
			return *element, nil
		}
	}
	return gitlab.Project{}, errors.New("not found")
}

//viable, can filter userid
func M_GetProjectIssues(git *gitlab.Client ,projectID int)([]*gitlab.Issue, error) {



	opt := &gitlab.ListProjectIssuesOptions{
		ListOptions: gitlab.ListOptions{
			PerPage: 10,
			Page:    1,
		},
	}

	var allIssues []*gitlab.Issue
	for {
		issues, resp, err := git.Issues.ListProjectIssues(projectID, opt)
		if(err == nil) {
			allIssues = append(allIssues, issues...)

			if opt.Page >= resp.TotalPages {
				break
			}

			opt.Page = resp.NextPage
		} else {
			return nil, err
		}
	}
	return allIssues, nil
}

func M_GetProjectContributors(git *gitlab.Client, projectID int)([]*gitlab.Contributor, error) {

	opt := &gitlab.ListContributorsOptions{
		PerPage:10,
		Page:1,
	}

	var allContributors []*gitlab.Contributor
	for {
		contributors, resp, err := git.Repositories.Contributors(projectID, opt)
		if(err == nil) {
			allContributors = append(allContributors, contributors...)

			if opt.Page >= resp.TotalPages {
				break
			}

			opt.Page = resp.NextPage
		} else {
			return nil, err
		}
	}
	return allContributors, nil
}

func M_GetProjectCommits(git *gitlab.Client, projectID int)([]*gitlab.Commit, error) {

	opt := &gitlab.ListCommitsOptions{
		ListOptions: gitlab.ListOptions{
			PerPage: 10,
			Page:    1,
		},
	}

	var allCommits []*gitlab.Commit
	for {
		commits, resp, err := git.Commits.ListCommits(projectID, opt)
		if(err == nil) {
			allCommits = append(allCommits, commits...)

			if opt.Page >= resp.TotalPages {
				break
			}

			opt.Page = resp.NextPage
		} else {
			return nil, err
		}
	}
	return allCommits, nil
}

//not usefull, only one user possible
func M_GetProjectEvents(git *gitlab.Client, projectID int)([]*gitlab.ContributionEvent, error) {

	opt := &gitlab.ListContributionEventsOptions{
		ListOptions: gitlab.ListOptions{
			PerPage: 10,
			Page:    1,
		},
	}

	var allEvents []*gitlab.ContributionEvent
	for {
		events, resp, err := git.Events.ListProjectVisibleEvents(projectID, opt)
		if (err == nil) {
			for _, element := range events {
				if (element.ProjectID == projectID) {
					allEvents = append(allEvents, element)
				}
			}

			if opt.Page >= resp.TotalPages {
				break
			}

			opt.Page = resp.NextPage
		} else {
			return nil, err
		}
	}

	return allEvents, nil
}
