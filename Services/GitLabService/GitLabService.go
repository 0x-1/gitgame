package GitLabService

import (
	"github.com/0x-1/GitGame/Managers/GitLabManager"
	"github.com/0x-1/GitGame/Managers/InterpreterManager"
	"github.com/0x-1/GitGame/Managers/OutputManager"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"github.com/0x-1/GitGame/Managers/CryptManager"
	"encoding/base64"
)

func M_InitGitLabService(engine *gin.Engine) {
	//engine.POST("/requestGame", m_OnRequestGame)

	engine.GET("/gitgame/:namespace/:project", m_OnStartGame)
	engine.GET("/test", m_OnTest)
}

func m_OnTest(context *gin.Context) {
	/*err, gitData := GameManager.M_GetGame("https://inf-git.fh-rosenheim.de/", "unity", "S8WXCdS2yrJwhSZ_C-oH")
	if(err != nil) {
		log.Println(err)
		return
	}

	var list []gitlab.ContributionEvent
	for _,element := range gitData.Events {
		if(element.ActionName=="pushed to") {
			list = append(list, element)
		}

	}*/
	context.JSON(http.StatusOK, "pong")
}

//TODO RETURN OK AFTER SANITY CHECK ONLY (token ok?, project name ok?) and do the rest in a go routine!

func m_OnStartGame(context *gin.Context) {
	namespace := context.Param("namespace")
	projectName := context.Param("project")

	gitLabURL := context.DefaultQuery("url", "https://gitlab.com")
	accessToken := context.Query("token")
	cryptedToken := context.Query("ctoken")

	if(len(projectName) <= 0 || len(namespace) <= 0 || (len(accessToken) <= 0 && len(cryptedToken)<= 0)) {
		context.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if len(cryptedToken) > 0 {
		bytes, err := base64.URLEncoding.DecodeString(cryptedToken)
		if(err != nil) {
			context.AbortWithStatus(http.StatusBadRequest)
			log.Println(err)
			return
		}

		decbytes, err := CryptManager.M_Decrypt(bytes)
		if(err != nil) {
			context.AbortWithStatus(http.StatusBadRequest)
			log.Println(err)
			return
		}

		accessToken =  string(decbytes)
	}

	err := GitLabManager.M_TestGitLabData(gitLabURL, projectName, namespace, accessToken)
	if(err != nil) {
		context.JSON(http.StatusInternalServerError, "GitGame Error: " + err.Error())
		return
	} else {
		go func() {
			gitLabData , err := GitLabManager.M_GetGitLabData(gitLabURL, projectName, namespace, accessToken)
			if(err != nil) {
				log.Println(err)
			}

			gitGameState, err := InterpreterManager.M_Interpret(gitLabData)
			if(err != nil){
				log.Println(err)
			}

			err = OutputManager.M_SaveAsWikiPage(gitLabURL, projectName, namespace, accessToken, gitGameState, context.Request.Host)
			if(err != nil) {
				log.Println(err)
			}
		}()
		context.JSON(http.StatusOK, "Game, Started. Please wait at least 10 seconds and look into your wiki for #gitGame Result")
	}

}
