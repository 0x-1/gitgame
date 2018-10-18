package GitService

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/0x-1/GitGame/Managers/GameManager"
	"github.com/0x-1/GitGame/Managers/CryptManager"
	"encoding/base64"
)

func M_InitGameService(engine *gin.Engine) {
	//engine.POST("/requestGame", m_OnRequestGame)

	engine.GET("/gitgame/start/:projectName/:accessToken", m_OnStartGame)
	engine.GET("/gitgame/update/:projectName/:cryptedToken", m_OnUpdateGame)
}

//TODO RETURN OK AFTER SANITY CHECK ONLY (token ok?, project name ok?) and do the rest in a go routine!

func m_OnStartGame(context *gin.Context) {
	//claims := jwt.ExtractClaims(context)
	//userID := claims["id"].(string)
	projectName := context.Param("projectName")
	accessToken := context.Param("accessToken")
	//var request Models.RequestGame
	//err := context.BindJSON(&request)

	if(len(projectName) <= 0 || len(accessToken) <= 0) {
		context.AbortWithStatus(http.StatusBadRequest)
	}

	err := GameManager.M_TestGame(projectName, accessToken)
	if(err != nil) {
		context.AbortWithStatus(http.StatusBadRequest)
	}

	go func() {
		err , game := GameManager.M_GetGame(projectName, accessToken)
		if(err == nil) {
			GameManager.M_SaveGame(game, accessToken)
		}
	}()

	context.JSON(http.StatusOK, "Game, Started. Please wait at least 10 seconds and look into your wiki for #gitGame Result")
}

func m_OnUpdateGame(context *gin.Context) {
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

	err = GameManager.M_TestGame(projectName, string(accessToken))
	if(err != nil) {
		context.AbortWithStatus(http.StatusBadRequest)
	}

	go func() {
		err , game := GameManager.M_GetGame(projectName, string(accessToken))
		if(err == nil) {
			GameManager.M_SaveGame(game, string(accessToken))
		}
	}()
}
