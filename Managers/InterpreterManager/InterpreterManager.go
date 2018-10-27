package InterpreterManager

import (
	"errors"
	"github.com/0x-1/GitGame/Models"
	"gopkg.in/alecthomas/kingpin.v2"
	"log"
	"strings"
)

var (
	app       = kingpin.New("interpreter", "gitgame config file interpreter")
	debug     = app.Command("debug", "enable debug mode")
	debugBool = debug.Arg("boolean", "enable or disable bool").Required().Bool()

	level       = app.Command("level", "level element")
	levelExp = level.Arg("exp", "experience required").Required().Int()

	questCmd       = app.Command("quest", "quest element")
	questScope          = questCmd.Arg("scope", "player").Required().String()
	questType       = questCmd.Arg("type", "issue").Required().String()
	questExp        = questCmd.Arg("exp", "exp reward").Required().Int()
	questConstraint = questCmd.Arg("constraint", "closed|opened").String()

	archievementCmd = app.Command("archievement", "archievement element")
	archievementScope = archievementCmd.Arg("scope", "project").Required().String()
	archievementType = archievementCmd.Arg("type", "milestone").Required().String()
	archievementExp = archievementCmd.Arg("exp", "exp reward").Required().Int()
	archievementConstraint = archievementCmd.Arg("constraint", "milestone:opened").String()
)

func M_Interpret(gitLabData Models.GitLabData) (Models.GitGameState, error) {

	//Split by line
	lines := strings.Split(gitLabData.ConfigFileContent, "\n")

	//Create GameState
	var state Models.GitGameState
	for _, member := range gitLabData.Members {
		state.Players = append(state.Players, Models.Player{MemberData: member})
	}

	//Line by Line
	for _, line := range lines {
		newState, err := m_InterpretLine(state, gitLabData, line)
		if err != nil {
			return state, err
		} else {
			state = newState
		}
	}

	return state, nil
}

func m_InterpretFile() (error) {

	for {
		break
	}
	return nil
}

func m_InterpretLine(currentState Models.GitGameState, gitData Models.GitLabData, cmd string) (Models.GitGameState, error) {

	//Comments
	cmdNoComment := strings.Split(cmd, "//")
	if (len(cmdNoComment[0]) <= 0) {
		return currentState, nil
	}

	//Parsing Error
	parsed, err := app.Parse(strings.Fields(cmdNoComment[0]))
	if (err != nil) {
		return currentState, err
	}

	//Parsing OK, Switch Commands
	parsedString := kingpin.MustParse(parsed, err)
	switch parsedString {

	//Level Add Command
	case level.FullCommand():
		{
			currentState.Levels = append(currentState.Levels, Models.Level{RequiredEXP: *levelExp})
		}

		//Quest Add Command
	case questCmd.FullCommand():
		{
			newState, err := m_InterpretQuestAdd(currentState, gitData, cmd)
			if(err != nil) {
				return currentState, err
			}

			currentState = newState
		}
	case archievementCmd.FullCommand():
		{
			newState, err := m_InterpretArchievementAdd(currentState, gitData, cmd)
			if(err != nil) {
				return currentState, err
			}
			currentState = newState
		}
	default:
		{
			return currentState, errors.New("unknown command in: "+cmd)
		}
	}

	return currentState, nil
}

func m_InterpretQuestAdd(currentState Models.GitGameState, gitData Models.GitLabData, cmd string)(Models.GitGameState, error){
	var newState Models.GitGameState
	newState = currentState
	var todo Models.Todo
	switch *questType {
	case "issue":
		{

				if len(*questConstraint)>= 0 {
					switch *questConstraint {
					case "opened":
						{
							if(*questScope == "player") {
								todo.Description = "Player Open Issue"
								todo.Experience = *questExp
								todo.Done = false
								for _, event := range gitData.Events {
									if (event.TargetType != "Issue") {
										continue
									}
									if (event.ActionName != "opened") {
										continue
									}

									for pindex, player := range newState.Players {
										if (event.AuthorID == player.MemberData.ID) {
											newState.Players[pindex].Experience += *questExp
										}
									}
								}
								newState.Todos = append(newState.Todos, todo)
								return newState, nil
							} else {
								return newState, errors.New("unknown scope in: " + cmd)
							}

						}
					case "closed":
						{
							if(*questScope == "player") {
								todo.Description = "Player Close Issue"
								todo.Experience = *questExp
								todo.Done = false
								for _, event := range gitData.Events {
									if (event.TargetType != "Issue") {
										continue
									}
									if (event.ActionName != "closed") {
										continue
									}
									for pindex, player := range newState.Players {
										if (event.AuthorID == player.MemberData.ID) {
											newState.Players[pindex].Experience += *questExp
										}
									}
								}
								newState.Todos = append(newState.Todos, todo)
								return newState, nil
							} else {
								return newState, errors.New("unknown scope in: " + cmd)
							}
						}
					default:
						{
							return newState, errors.New("unknown issue constaint in: " + cmd)
						}
					}
				} else {
					return newState, errors.New("issue required to have a constaint in: " + cmd)
				}
		}
	/*case "commit":
		{
			for _, event := range gitData.Events {
				if (event.ActionName == "pushed to") {
					for pindex, player := range currentState.Players {
						if (event.AuthorID == player.MemberData.ID) {
							currentState.Players[pindex].Experience += *questExp
						}
					}
				}
			}
			return currentState, nil
		}*/

	default:
		return currentState, errors.New("unknown quest type in: " + cmd)
	}
}

func m_InterpretArchievementAdd(currentState Models.GitGameState, gitData Models.GitLabData, cmd string) (Models.GitGameState, error) {
	var newState Models.GitGameState
	newState = currentState
	var todo Models.Todo
	switch (*archievementType) {
	case "milestone":
		{
			if len(*archievementConstraint) >= 0 {
				switch *archievementConstraint {
				case "opened":
					{
						if (*archievementScope == "project") {
							todo.Description = "Open Milestone for Project."
							todo.Experience = *archievementExp
							todo.Done = false

							for _, event := range gitData.Events {
								if (event.TargetType != "Milestone") {
									continue
								}
								if (event.ActionName != "opened") {
									continue
								}

								for pindex, _ := range newState.Players {
									newState.Players[pindex].Experience += *archievementExp
								}
								todo.Done = true
							}
							newState.Todos = append(newState.Todos, todo)
							return newState, nil
						} else {
							return newState, errors.New("unknown scope in: " + cmd)
						}

					}
				default:
					{
						return newState, errors.New("unknown milestone constaint in: " + cmd)
					}
				}
			} else {
				return newState, errors.New("milsestone required to have a constaint in: " + cmd)
			}

			return newState, nil
		}
	default:
		{
			return newState, errors.New("unknown archievement type in: " + cmd)
		}
	}
}

/*case "milestone":
{
for _, event := range gitData.Events {
if (event.TargetType == "Milestone" && event.ActionName == "opened") {
for pindex, player := range currentState.Players {
if (event.AuthorID == player.MemberData.ID) {
currentState.Players[pindex].Experience += *questaddexp
}
}
}
}
return currentState, nil
}*/
