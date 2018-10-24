package GitLabService

import (
	"github.com/gin-gonic/gin"
	"github.com/xanzy/go-gitlab"
	"log"
	"net/http"
	"github.com/0x-1/GitGame/Managers/GameManager"
	"github.com/0x-1/GitGame/Managers/CryptManager"
	"encoding/base64"
)

func M_InitGameService(engine *gin.Engine) {
	//engine.POST("/requestGame", m_OnRequestGame)

	engine.GET("/gitgame/start/:projectName/:accessToken", m_OnStartGame)
	engine.GET("/gitgame/update/:projectName/:cryptedToken", m_OnUpdateGame)
	engine.GET("/gitgame/test", m_OnTest)
}

func m_OnTest(context *gin.Context) {
	err, gitData := GameManager.M_GetGame("https://inf-git.fh-rosenheim.de/", "unity", "S8WXCdS2yrJwhSZ_C-oH")
	if(err != nil) {
		log.Println(err)
		return
	}

	var list []gitlab.ContributionEvent
	for _,element := range gitData.Events {
		if(element.TargetType=="Issue" && element.ActionName=="opened") {
			list = append(list, element)
		}

	}
	context.JSON(http.StatusOK, list)
}

//TODO RETURN OK AFTER SANITY CHECK ONLY (token ok?, project name ok?) and do the rest in a go routine!

func m_OnStartGame(context *gin.Context) {

	gitLabURL := context.DefaultQuery("url", "https://gitlab.com")
	projectName := context.Param("projectName")
	accessToken := context.Param("accessToken")

	if(len(projectName) <= 0 || len(accessToken) <= 0) {
		context.AbortWithStatus(http.StatusBadRequest)
	}

	err := GameManager.M_TestGame(gitLabURL, projectName, accessToken)
	if(err != nil) {
		context.AbortWithStatus(http.StatusBadRequest)
	}

	go func() {
		err , game := GameManager.M_GetGame(gitLabURL, projectName, accessToken)
		if(err == nil) {
			log.Println("yey")
			GameManager.M_SaveGame(game, accessToken)
		} else {
			log.Println(err)
		}
	}()

	context.JSON(http.StatusOK, "Game, Started. Please wait at least 10 seconds and look into your wiki for #gitGame Result")
}

func m_OnUpdateGame(context *gin.Context) {
	gitLabURL := context.DefaultQuery("url", "https://gitlab.com")
	projectName := context.Param("projectName")
	cryptedToken := context.Param("cryptedToken")

	if(len(projectName) <= 0 || len(cryptedToken) <= 0) {
		context.AbortWithStatus(http.StatusBadRequest)
	}

	bytes, err := base64.StdEncoding.DecodeString(cryptedToken)
	if(err != nil) {
		context.AbortWithStatus(http.StatusBadRequest)
	}

	accessToken, err := CryptManager.M_Decrypt(bytes)
	if(err != nil) {
		context.AbortWithStatus(http.StatusBadRequest)
	}

	err = GameManager.M_TestGame(gitLabURL, projectName, string(accessToken))
	if(err != nil) {
		context.AbortWithStatus(http.StatusBadRequest)
	}

	go func() {
		err , game := GameManager.M_GetGame(gitLabURL, projectName, string(accessToken))
		if(err == nil) {
			GameManager.M_SaveGame(game, string(accessToken))
		}
	}()
}