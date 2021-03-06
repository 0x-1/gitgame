package GitLabManager

import (
	"encoding/base64"
	"github.com/0x-1/GitGame/Managers/InterpreterManager"
	"github.com/0x-1/GitGame/Models"
	"github.com/pkg/errors"
	"github.com/xanzy/go-gitlab"
	"strings"
	"log"
)

func M_TestGitLabData(gitLabURL string, projectName string, nameSpace, accessToken string)(error) {
	git := gitlab.NewClient(nil, accessToken)
	git.SetBaseURL(gitLabURL)

	//Test Token
	_, _, err := git.Users.CurrentUser()
	if err != nil {
		return errors.New("permission error, please check token")
	}

	project, err := M_GetProjectByName(git, projectName, nameSpace)
	if err != nil{
		return errors.New("could not find project: "+projectName+" in namespace: "+nameSpace)
	}
	if(len(project.DefaultBranch)<=0) {
		return errors.New("there is no default branch")
	}

	configFileContent, err := m_GetFileContent(git, project.ID, project.DefaultBranch)
	if err != nil {
		return errors.New("there is no .gitgame file in the default branch or no permission")
	}

	err = InterpreterManager.M_TestInterpret(configFileContent)
	if(err != nil) {
		log.Println("err")
		return err
	}

	return nil
}

func M_GetGitLabData(gitLabURL string, projectName string, nameSpace, accessToken string) (Models.GitLabData, error) {
	git := gitlab.NewClient(nil, accessToken)
	git.SetBaseURL(gitLabURL)

	//User
	/*user, _, err := git.Users.CurrentUser()
	if err != nil {
		return Models.GitLabData{}, err
	}*/

	//Project
	project, err := M_GetProjectByName(git, projectName, nameSpace)
	if err != nil{
		return Models.GitLabData{}, err
	}

	if(len(project.DefaultBranch)<=0) {
		return Models.GitLabData{}, errors.New("there is no default branch")
	}

	//Contributors
	/*contributors, err := m_GetProjectContributors(git, project.ID)
	if err != nil {
		return Models.GitLabData{}, err
	}*/

	//Issues
	/*issues, err := m_GetProjectIssues(git, project.ID)
	if err != nil {
		return Models.GitLabData{}, err
	}*/

	//Project Members
	members, err := m_GetProjectMembers(git, project.ID)
	if err != nil {
		return Models.GitLabData{}, err
	}

	//Events
	projectEvents, err := m_GetProjectEvents(git, project.ID)
	if err != nil {
		return Models.GitLabData{}, err
	}

	//Config File
	configFileContent, err := m_GetFileContent(git, project.ID, project.DefaultBranch)
	if err != nil {
		return Models.GitLabData{}, err
	}

	//PipelineList
	pipelineList, err := m_GetPipelineList(git, project.ID)
	if(err != nil) {
		return Models.GitLabData{}, err
	}

	var gitLabData Models.GitLabData
	//gitLabData.Contributors = contributors
	gitLabData.Project = project
	//gitLabData.Issues = issues
	gitLabData.ProjectEvents = projectEvents
	gitLabData.Members = members
	gitLabData.ConfigFileContent = configFileContent
	gitLabData.PipelineList = pipelineList
	return gitLabData, nil

}


func M_GetProjectByName(git *gitlab.Client ,projectName string, nameSpace string)(gitlab.Project, error) {
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
		if(strings.EqualFold(element.Namespace.Name, nameSpace) && strings.EqualFold(element.Name, projectName)) {
			return *element, nil
		}
	}
	return gitlab.Project{}, errors.New("not found")
}

//viable, can filter userid
func m_GetProjectIssues(git *gitlab.Client ,projectID int)([]gitlab.Issue, error) {



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

func m_GetProjectContributors(git *gitlab.Client, projectID int)([]gitlab.Contributor, error) {

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

func m_GetProjectCommits(git *gitlab.Client, projectID int)([]gitlab.Commit, error) {

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

func m_GetProjectEvents(git *gitlab.Client, projectID int) ([]gitlab.ContributionEvent, error) {

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
				allEvents = append(allEvents, *element)
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

func m_GetProjectWikiEvents(git *gitlab.Client, projectID int) ([]gitlab.ContributionEvent, error) {


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
				allEvents = append(allEvents, *element)
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

func m_GetProjectMembers(git *gitlab.Client, projectID int)([]gitlab.ProjectMember, error) {
	opt := &gitlab.ListProjectMembersOptions{

	}
	var allMembers []gitlab.ProjectMember
	members, _, err := git.ProjectMembers.ListProjectMembers(projectID, opt)

	if(err == nil) {
		for _, member := range members {
			allMembers = append(allMembers, *member)
		}
		return allMembers, nil
	} else {
		return nil, err
	}


}

func m_GetFileContent(git *gitlab.Client, projectID int, branch string)(string, error) {
	opt := &gitlab.GetFileOptions{
		Ref: gitlab.String(branch),
	}

	file, _, err := git.RepositoryFiles.GetFile(projectID, ".gitgame", opt)
	if(err==nil) {
		//Readable Bytes
		bytes, err := base64.StdEncoding.DecodeString(file.Content)
		if(err != nil) {
			return "", err
		}
		return string(bytes), nil
	} else {
		return "", err
	}
}

func m_GetPipelineList(git *gitlab.Client, projectID int)(gitlab.PipelineList, error) {
	opt := &gitlab.ListProjectPipelinesOptions{

	}

	list, _, err := git.Pipelines.ListProjectPipelines(projectID, opt)
	if(err != nil) {
		return nil, err
	}

	return list, nil
}

func M_CreateEditWikiPage(git *gitlab.Client, projectID int, content string) (error){

	optEdit := &gitlab.EditWikiPageOptions{
		Content:gitlab.String(content),
		Title:gitlab.String("#gitGame-Result"),
	}



	optCreate := &gitlab.CreateWikiPageOptions{
		Title:gitlab.String("#gitGame-Result"),
		Content:gitlab.String(content),
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
