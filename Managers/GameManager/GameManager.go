package GameManager

import (
	"github.com/xanzy/go-gitlab"
	"github.com/0x-1/GitGame/Models"
	"github.com/0x-1/GitGame/Managers/GitLabManager"
	"log"
)

//Function to test wether projectName and accessToken are working
func M_TestGame(projectName string,accessToken string) (error) {
	return nil
}

//Actual Function to get all game relevant data
func M_GetGame(projectName string,accessToken string) (error, Models.ResponseGame) {
	git := gitlab.NewClient(nil, accessToken)
	git.SetBaseURL("https://inf-git.fh-rosenheim.de/")

	//User
	user, _, err := git.Users.CurrentUser()
	if(err != nil) {
		return err, Models.ResponseGame{}
	}

	//Project
	project, err := GitLabManager.M_GetProjectByName(git, projectName)
	if(err != nil) {
		return err, Models.ResponseGame{}
	}
	log.Println("UserID of AccessToken is: ", user.ID, "and ProjectID is: ", project.ID)

	//Contributors
	contributors, err := GitLabManager.M_GetProjectContributors(git, project.ID)
	if(err != nil) {
		return err, Models.ResponseGame{}
	}

	//Issues
	issues, err := GitLabManager.M_GetProjectIssues(git, project.ID)
	if(err != nil) {
		return err, Models.ResponseGame{}
	}

	//Events
	events, err := GitLabManager.M_GetProjectEvents(git, project.ID)
	if(err != nil) {
		return err, Models.ResponseGame{}
	}

	if(err != nil) {
		return err, Models.ResponseGame{}
	}

	var response Models.ResponseGame
	response.Contributors = contributors
	response.Project = project
	response.Issues = issues
	response.Events = events
	return nil, response

}

func M_SaveGame( game Models.ResponseGame, accessToken string) (error) {
	git := gitlab.NewClient(nil, accessToken)
	git.SetBaseURL("https://inf-git.fh-rosenheim.de/")
	err := GitLabManager.M_CreateEditWikiPage(git, game.Project.ID)
}
