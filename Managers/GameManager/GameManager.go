package GameManager

import (
	"github.com/xanzy/go-gitlab"
	"github.com/0x-1/GitGame/Models"
	"github.com/0x-1/GitGame/Managers/GitLabManager"
	"log"
)

//Function to test wether projectName and accessToken are working
func M_TestGame(gitLabURL string, projectName string,accessToken string) (error) {
	return nil
}

//Actual Function to get all game relevant data
func M_GetGame(gitLabURL string, projectName string, accessToken string) (error, Models.GitLabData) {
	git := gitlab.NewClient(nil, accessToken)
	git.SetBaseURL(gitLabURL)

	//User
	user, _, err := git.Users.CurrentUser()
	if(err != nil) {
		return err, Models.GitLabData{}
	}

	//Project
	project, err := GitLabManager.M_GetProjectByName(git, projectName)
	if(err != nil) {
		return err, Models.GitLabData{}
	}
	log.Println("UserID of AccessToken is: ", user.ID, "and ProjectID is: ", project.ID)

	//Contributors
	contributors, err := GitLabManager.M_GetProjectContributors(git, project.ID)
	if(err != nil) {
		return err, Models.GitLabData{}
	}

	//Issues
	issues, err := GitLabManager.M_GetProjectIssues(git, project.ID)
	if(err != nil) {
		return err, Models.GitLabData{}
	}

	//Project Members
	members, err := GitLabManager.M_GetProjectMembers(git, project.ID)
	if(err != nil) {
		return err, Models.GitLabData{}
	}

	//Events
	events, err := GitLabManager.M_GetProjectEvents(git, project.ID)
	if(err != nil) {
		return err, Models.GitLabData{}
	}

	if(err != nil) {
		return err, Models.GitLabData{}
	}

	var response Models.GitLabData
	response.Contributors = contributors
	response.Project = project
	response.Issues = issues
	response.Events = events
	response.Members = members
	return nil, response

}

func M_SaveGame( game Models.GitLabData, accessToken string) (error) {
	return nil
	//git := gitlab.NewClient(nil, accessToken)
	//git.SetBaseURL("https://inf-git.fh-rosenheim.de/")
	//err := GitLabManager.M_CreateEditWikiPage(git, game.Project.ID)
}
