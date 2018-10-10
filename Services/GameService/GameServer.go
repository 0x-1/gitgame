package GameService

import (
	"github.com/gin-gonic/gin"

	"github.com/0x-1/GitGame/Models"
	"net/http"
	"github.com/xanzy/go-gitlab"
	"log"
	"github.com/0x-1/GitGame/Managers/GameManager"
)

func M_InitGameService(engine *gin.Engine) {
		engine.POST("/requestGame", m_OnRequestGame)
}

func m_OnRequestGame(context *gin.Context) {
	//claims := jwt.ExtractClaims(context)
	//userID := claims["id"].(string)
	var request Models.RequestGame
	err := context.BindJSON(&request)
	if(err == nil) {
		git := gitlab.NewClient(nil, request.V_AccessToken)
		git.SetBaseURL("https://inf-git.fh-rosenheim.de/")
		user, _, err := git.Users.CurrentUser()
		if(err == nil) {
			project, err := GameManager.M_GetProjectByName(git,request.V_ProjectName)
			if(err==nil) {
				log.Println("UserID of AccessToken is: ", user.ID, "and ProjectID is: ",project.ID)
				var response Models.ResponseGame
				context.JSON(http.StatusOK, response)
			} else {
				log.Println("could not get project: ",request.V_ProjectName, " with token: ",request.V_AccessToken)
				context.AbortWithStatus(http.StatusBadRequest)
			}
		} else {
			log.Println("could not get current user: ",request.V_AccessToken)
			context.AbortWithStatus(http.StatusBadRequest)
		}
	} else {
		log.Println("could not bind request: ",request.V_AccessToken)
		context.AbortWithStatus(http.StatusBadRequest)
	}

}