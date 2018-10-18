package GitLabManager

import (
	"github.com/xanzy/go-gitlab"
	"strings"
	"github.com/pkg/errors"
	"log"
	"encoding/base64"
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
func M_GetProjectIssues(git *gitlab.Client ,projectID int)([]gitlab.Issue, error) {



	opt := &gitlab.ListProjectIssuesOptions{
		ListOptions: gitlab.ListOptions{
			PerPage: 10,
			Page:    1,
		},
	}

	var allIssues []gitlab.Issue
	for {
		issues, resp, err := git.Issues.ListProjectIssues(projectID, opt)
		if(err == nil) {
			for _, element := range issues {
				allIssues = append(allIssues, *element)
			}

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

func M_GetProjectContributors(git *gitlab.Client, projectID int)([]gitlab.Contributor, error) {

	opt := &gitlab.ListContributorsOptions{
		PerPage:10,
		Page:1,
	}

	var allContributors []gitlab.Contributor
	for {
		contributors, resp, err := git.Repositories.Contributors(projectID, opt)
		if(err == nil) {
			for _, element := range contributors {
				allContributors = append(allContributors, *element)
			}


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

func M_GetProjectCommits(git *gitlab.Client, projectID int)([]gitlab.Commit, error) {

	opt := &gitlab.ListCommitsOptions{
		ListOptions: gitlab.ListOptions{
			PerPage: 10,
			Page:    1,
		},
	}

	var allCommits []gitlab.Commit
	for {
		commits, resp, err := git.Commits.ListCommits(projectID, opt)
		if(err == nil) {
			for _, element := range commits {
				allCommits = append(allCommits, *element)
			}

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
func M_GetProjectEvents(git *gitlab.Client, projectID int)([]gitlab.ContributionEvent, error) {

	opt := &gitlab.ListContributionEventsOptions{
		ListOptions: gitlab.ListOptions{
			PerPage: 10,
			Page:    1,
		},
	}

	var allEvents []gitlab.ContributionEvent
	for {
		events, resp, err := git.Events.ListProjectVisibleEvents(projectID, opt)
		if (err == nil) {
			for _, element := range events {
				if (element.ProjectID == projectID) {
					for _, element := range events {
						allEvents = append(allEvents, *element)
					}
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

func M_GetFile(git *gitlab.Client, projectID int, branch string)(content string, err error) {
	opt := &gitlab.GetFileOptions{
		Ref: gitlab.String(branch),
	}

	file, _, err := git.RepositoryFiles.GetFile(projectID, ".gitattributes", opt)
	if(err==nil) {
		bytes, err := base64.StdEncoding.DecodeString(file.Content)
		if(err==nil) {
			temp := strings.Split(string(bytes),"\n")
			for _, element := range temp {
				log.Println(element)
			}
		} else {
			return "", err
		}
		return file.Content, nil
	} else {
		return "", err
	}
}

func M_CreateEditWikiPage(git *gitlab.Client, projectID int) (error){
	optEdit := &gitlab.EditWikiPageOptions{
		Content:gitlab.String("yey 2"),
		Title:gitlab.String("#gitGame-Result"),
	}

	optCreate := &gitlab.CreateWikiPageOptions{
		Title:gitlab.String("#gitGame-Result"),
		Content:gitlab.String("yey"),
	}

	//optList := &gitlab.ListWikisOptions{
	//
	//}
	//
	//wikies, _, err := git.Wikis.ListWikis(projectID, optList)
	//for _, element := range wikies {
	//	log.Println(element.Slug)
	//}

	_, _ ,err := git.Wikis.GetWikiPage(projectID, "#gitGame-Result")
	if(err == nil) {
		_, _ ,err := git.Wikis.EditWikiPage(projectID, "#gitGame-Result", optEdit)
		if(err==nil) {
			return nil
		} else {
			return err
		}
	} else {
		_,_,err := git.Wikis.CreateWikiPage(projectID, optCreate)
		if(err==nil) {
			return nil
		} else {
			return err
		}
	}



}