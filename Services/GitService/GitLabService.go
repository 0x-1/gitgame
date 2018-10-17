package GitService

import (
	"github.com/gin-gonic/gin"
	"github.com/0x-1/GitGame/Models"
	"net/http"
	"github.com/xanzy/go-gitlab"
	"log"
	"github.com/0x-1/GitGame/Managers/DataManager"
)

func M_InitGameService(engine *gin.Engine) {
	engine.POST("/requestGame", m_OnRequestGame)

	engine.GET("/gitgame/:projectName/:accessToken", m_OnRequestGame)
}

//TODO RETURN OK AFTER SANITY CHECK ONLY (token ok?, project name ok?) and do the rest in a go routine!

func m_OnRequestGame(context *gin.Context) {
	//claims := jwt.ExtractClaims(context)
	//userID := claims["id"].(string)
	projectName := context.Param("projectName")
	accessToken := context.Param("accessToken")
	//var request Models.RequestGame
	//err := context.BindJSON(&request)
	git := gitlab.NewClient(nil, accessToken)
	git.SetBaseURL("https://inf-git.fh-rosenheim.de/")
	user, _, err := git.Users.CurrentUser()
	if (err == nil) {
		project, err := DataManager.M_GetProjectByName(git, projectName)
		if (err == nil) {
			log.Println("UserID of AccessToken is: ", user.ID, "and ProjectID is: ", project.ID)
			contributors, err := DataManager.M_GetProjectContributors(git, project.ID)
			issues, err2 := DataManager.M_GetProjectIssues(git, project.ID)
			events, err3 := DataManager.M_GetProjectEvents(git, project.ID)
			err4 := DataManager.M_CreateEditWikiPage(git, project.ID)
			var response Models.ResponseGame
			if(err == nil && err2==nil && err3==nil && err4==nil) {
				response.Contributors = contributors
				response.Project = project
				response.Issues = issues
				response.Events = events
				context.JSON(http.StatusOK, response)
			} else {
				context.AbortWithStatus(http.StatusBadRequest)
			}

		} else {
			log.Println("could not get project: ", projectName, " with token: ", accessToken)
			context.AbortWithStatus(http.StatusBadRequest)
		}
	} else {
		log.Println("could not get current user: ", accessToken)
		context.AbortWithStatus(http.StatusBadRequest)
	}

}
