package InterpreterManager

import (
	"errors"
	"github.com/0x-1/GitGame/Models"
	"gopkg.in/alecthomas/kingpin.v2"
	"strings"
)

var (
	app = kingpin.New("interpreter", "gitgame config file interpreter")
	debug = app.Command("debug", "enable debug mode")
	debugBool = debug.Arg("boolean", "enable or disable bool").Required().Bool()

	level = app.Command("level", "level element")
	leveladd = level.Command("add", "add level")
	leveladdexp = leveladd.Arg("exp", "experience required").Required().Int()

	quest = app.Command("quest", "quest element")
	questadd = quest.Command("add", "add quest")
	questaddtype = questadd.Arg("type", "issue|commit|milestone|pipeline").Required().String()
	questaddexp = questadd.Arg("exp", "quest exp reward").Required().Int()
	questaddconstraint = questadd.Arg("constraint", "action like close|open|closeassigned|commentassigned|...").Required().String()
)

func M_Interpret(gitLabData Models.GitLabData)(Models.GitGameState, error) {

	//Split by line
	lines := strings.Split(gitLabData.ConfigFileContent,"\n")

	//Create GameState
	var state Models.GitGameState
	for _, member := range gitLabData.Members {
		state.Players = append(state.Players, Models.Player{MemberData:member})
	}

	//Line by Line
	for _, line := range lines {
		state, err := m_InterpretLine(state, Models.GitLabData{}, line)
		if(err != nil) {
			return state, err
		}
	}

	return state, nil
}

func m_InterpretFile()(error){

	for {
		break
	}
	return nil
}

func m_InterpretLine(currentState Models.GitGameState, gitData Models.GitLabData,cmd string) (Models.GitGameState, error) {

	//Comments
	cmdNoComment := strings.Split(cmd, "//")
	if(len(cmdNoComment[0]) <= 0) {
		return currentState, nil
	}

	//Parsing Error
	parsed, err := app.Parse(strings.Fields(cmdNoComment[0]))
	if(err != nil) {
		return currentState, err
	}

	//Parsing OK, Switch Commands
	parsedString := kingpin.MustParse(parsed, err)
	switch parsedString {

	//Level Add Command
	case leveladd.FullCommand():
		{
			currentState.Levels = append(currentState.Levels, Models.Level{RequiredEXP:*leveladdexp})
		}

	//Quest Add Command
	case questadd.FullCommand():
		{
			switch *questaddtype {
			case "issue":
				{
					for _,event := range gitData.Events {
						if(event.TargetType=="Issue" && event.ActionName=="opened") {
							for pindex,player := range currentState.Players {
								if(event.AuthorID==player.MemberData.ID) {
									currentState.Players[pindex].Experience += *questaddexp
								}
							}
						}
					}
					return currentState, nil
				}
			case "commit":
				{
					for _,event := range gitData.Events {
						if(event.ActionName=="pushed to") {
							for pindex,player := range currentState.Players {
								if(event.AuthorID==player.MemberData.ID) {
									currentState.Players[pindex].Experience += *questaddexp
								}
							}
						}
					}
					return currentState, nil
				}
			case "milestone":
				{
					for _,event := range gitData.Events {
						if(event.TargetType=="Milestone" && event.ActionName=="opened") {
							for pindex,player := range currentState.Players {
								if(event.AuthorID==player.MemberData.ID) {
									currentState.Players[pindex].Experience += *questaddexp
								}
							}
						}
					}
					return currentState, nil
				}
			case "pipeline":
				{
					for _,event := range gitData.Events {
						if(event.TargetType=="Issue" && event.ActionName=="opened") {
							for pindex,player := range currentState.Players {
								if(event.AuthorID==player.MemberData.ID) {
									currentState.Players[pindex].Experience += *questaddexp
								}
							}
						}
					}
					return currentState, nil
				}
			default:
				return currentState, errors.New("unknown quest type in command: "+cmd)
			}


			//currentState.Levels = append(currentState.Levels, Models.Level{RequiredEXP:*leveladdexp})
		}
	}

	return currentState,nil
}