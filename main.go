package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/0x-1/GitGame/Services/GameService"
	"github.com/0x-1/GitGame/Managers/GameManager"
	"github.com/xanzy/go-gitlab"
	"log"
	"os"
)
func main() {
	log.SetOutput(os.Stdout)
	gin.SetMode(gin.DebugMode)
	engine := gin.Default()

	engine.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, "pong")
	})

	//Add Services
	GameService.M_InitGameService(engine)

	//test
	git := gitlab.NewClient(nil, "S8WXCdS2yrJwhSZ_C-oH")
	git.SetBaseURL("https://inf-git.fh-rosenheim.de/")
	project, err := GameManager.M_GetProjectByName(git,"unity")
	if(err==nil) {
		log.Println("Project: ", project.ID)
		issues, err := GameManager.M_GetProjectIssues(git, project.ID)
		if(err == nil) {
			log.Println("Issues:", len(issues))
			contributors, err := GameManager.M_GetProjectContributors(git, project.ID)
			if(err==nil) {
				log.Println("Contributors:", len(contributors))

				events, err := GameManager.M_GetProjectEvents(git, project.ID)
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
	}

	//engine.Run(":54321")
}
