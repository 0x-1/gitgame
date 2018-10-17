package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/0x-1/GitGame/Services/GitService"
	"github.com/0x-1/GitGame/Managers/DataManager"
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
	GitService.M_InitGameService(engine)

	engine.Run(":54321")
	//test
	git := gitlab.NewClient(nil, "S8WXCdS2yrJwhSZ_C-oH")
	git.SetBaseURL("https://inf-git.fh-rosenheim.de/")
	project, err := DataManager.M_GetProjectByName(git,"unity")
	if(err==nil) {
		log.Println("Project: ", project.ID)

		_, err := DataManager.M_GetFile(git, project.ID, project.DefaultBranch)
		if(err != nil) {
			log.Println(err)
		}

		issues, err := DataManager.M_GetProjectIssues(git, project.ID)
		if(err == nil) {
			log.Println("Issues:", len(issues))
			contributors, err := DataManager.M_GetProjectContributors(git, project.ID)
			if(err==nil) {
				log.Println("Contributors:", len(contributors))
				events, err := DataManager.M_GetProjectEvents(git, project.ID)
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
