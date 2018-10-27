package main

import (
	"github.com/0x-1/GitGame/Services/GitLabService"
	"github.com/gin-gonic/gin"
	"net/http"
)



func main() {
	//data, err := CryptManager.M_Encrypt([]byte("S8WXCdS2yrJwhSZ_C-oH"))

	/*err, gitData := GameManager.M_GetGame("https://inf-git.fh-rosenheim.de/", "unity", "S8WXCdS2yrJwhSZ_C-oH")
	if(err != nil) {
		log.Println(err)
		return
	}

	var gameState Models.GitGameState
	for _, member := range gitData.Members {
		gameState.Players = append(gameState.Players, Models.Player{MemberData:member, Experience:0})
	}


	//Player Init

	//startState.Players = data.Project.




	res, err := InterpreterManager.M_Interpret(gameState,gitData, "quest add issue 100 open //first quest")
	if(err == nil) {
		log.Println("yey", res.Players[0].MemberData.Username, res.Players[0].Experience)
	} else {
		log.Println(".gitGame config error:",err)
	}*/

	//return
	//switch kingpin.MustParse(app.Parse([]string{"debug","2"})) {
	//case debug.FullCommand():
	//	log.Println(*debugBool)
	//}
	//log.SetOutput(os.Stdout)
	gin.SetMode(gin.DebugMode)
	engine := gin.Default()

	engine.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, "pong")
	})

	//Add Services
	GitLabService.M_InitGitLabService(engine)

	engine.Run(":54321")
	//test


	/*git := gitlab.NewClient(nil, "S8WXCdS2yrJwhSZ_C-oH")
	git.SetBaseURL("https://inf-git.fh-rosenheim.de/")
	project, err := GitLabManager.m_GetProjectByName(git,"unity")
	if(err==nil) {
		log.Println("Project: ", project.ID)

		_, err := GitLabManager.m_GetFile(git, project.ID, project.DefaultBranch)
		if(err != nil) {
			log.Println(err)
		}

		issues, err := GitLabManager.m_GetProjectIssues(git, project.ID)
		if(err == nil) {
			log.Println("Issues:", len(issues))
			contributors, err := GitLabManager.m_GetProjectContributors(git, project.ID)
			if(err==nil) {
				log.Println("Contributors:", len(contributors))
				events, err := GitLabManager.m_GetProjectEvents(git, project.ID)
				if(err==nil) {
					for _, event := range events {
						//log.Println(event.CreatedAt, event.AuthorUsername,event.ActionName,event.TargetType, event.TargetTitle)
						if(event.TargetType=="Issue") {
							for _, issue := range issues {
								if(issue.ID == event.TargetID) {
									for _, assignee := range issue.Assignees {
										if(assignee.ID == event.AuthorID) {
											log.Println("EXP!, assignee closed assigned issue")
										}
									}
								}
							}
						}
					}
				}
			} else {

			}
		} else {

		}
	} else {
		log.Println(err)
	}*/

	//engine.Run(":54321")
}
