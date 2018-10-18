package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/0x-1/GitGame/Services/GitService"
	"github.com/0x-1/GitGame/Managers/GitLabManager"
	"github.com/xanzy/go-gitlab"
	"log"
	"os"
	"gopkg.in/alecthomas/kingpin.v2"
)


var (
	app = kingpin.New("interpreter", "gitgame config file interpreter")
	debug = app.Command("debug", "enable debug mode")
	debugBool = debug.Arg("boolean", "enable or disable bool").Required().Bool()
)
func main() {
	//data, err := CryptManager.M_Encrypt([]byte("S8WXCdS2yrJwhSZ_C-oH"))




	return
	switch kingpin.MustParse(app.Parse([]string{"debug","2"})) {
	case debug.FullCommand():
		log.Println(*debugBool)
	}

	//os.Exit(0)
	return
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
	project, err := GitLabManager.M_GetProjectByName(git,"unity")
	if(err==nil) {
		log.Println("Project: ", project.ID)

		_, err := GitLabManager.M_GetFile(git, project.ID, project.DefaultBranch)
		if(err != nil) {
			log.Println(err)
		}

		issues, err := GitLabManager.M_GetProjectIssues(git, project.ID)
		if(err == nil) {
			log.Println("Issues:", len(issues))
			contributors, err := GitLabManager.M_GetProjectContributors(git, project.ID)
			if(err==nil) {
				log.Println("Contributors:", len(contributors))
				events, err := GitLabManager.M_GetProjectEvents(git, project.ID)
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
