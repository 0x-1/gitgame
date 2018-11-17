package OutputManager

import (
	"encoding/base64"
	"github.com/0x-1/GitGame/Managers/CryptManager"
	"github.com/0x-1/GitGame/Managers/GitLabManager"
	"github.com/0x-1/GitGame/Models"
	"github.com/xanzy/go-gitlab"
	"strconv"
	"time"
)

func M_SaveAsWikiPage(gitLabURL string, projectName string, nameSpace string, accessToken string, state Models.GitGameState, gitGameHost string)(error) {
	git := gitlab.NewClient(nil, accessToken)
	git.SetBaseURL(gitLabURL)

	project, err := GitLabManager.M_GetProjectByName(git,projectName, nameSpace)

	var content string


	content += "#GitGame Result created at: "+time.Now().Format(time.RFC1123)+"\n\n"

	cryptedTokenBytes, err := CryptManager.M_Encrypt([]byte(accessToken))
	cryptedToken := base64.URLEncoding.EncodeToString(cryptedTokenBytes)



	content += "#Update this file by opening this link once: [link]("+gitGameHost+"/gitgame/"+nameSpace+"/"+projectName+"?ctoken="+cryptedToken+"&url="+gitLabURL+")\n\n"
	content += "Spieler | Level(Max="+strconv.Itoa(len(state.Levels))+") | Fortschritt im Level | Punkte\n"
	content += "--- | --- | --- | ---\n"

	for _,player := range state.Players {
		level, nextLevelExp, levelPercentComplete := m_GetPlayerLevel(state.Levels, player.Experience)
		content += player.MemberData.Username+ " | "
		content += strconv.Itoa(level) + " | "
		content += m_GenerateProgressBar(levelPercentComplete,20)+" "+strconv.Itoa(levelPercentComplete)+"% | "
		content += strconv.Itoa(player.Experience)+"/"+strconv.Itoa(nextLevelExp)
		content += "\n"
	}

	content += "\n\n"
	content += "TODO:\n\n"
	content += "Beschreibung | Punkte | Status\n"
	content += "--- | --- | ---\n"

	for _, todo := range state.Todos {
		var status string
		if(todo.Done) {
			status = "Erledigt"
		} else {
			status = "Offen"
		}
		content += todo.Description + " | " + strconv.Itoa(todo.Experience) + " | " +status
		content += "\n"
	}

	err = GitLabManager.M_CreateEditWikiPage(git, project.ID, content)
	if (err != nil) {
		return err
	}

	return nil
}

func m_GenerateProgressBar(percentProgress int, maxSteps int) (string) {
	x := percentProgress*maxSteps/100
	rest := maxSteps-x

	var str string
	str += "["
	for i := 1; i <= x; i += 1 {
		str += "#"
	}
	for i := 1; i <= rest; i += 1 {
		str += ".."
	}
	str += "]"
	return str
}

func m_GetPlayerLevel(levels []Models.Level, currentExp int) (playerLevel int, nextLevelExp int ,currentLevelPercentComplete int) {
	if(len(levels)<= 0) {
		return 1, currentExp,100
	}


	var lastLevelReq int
	for _, level := range levels {
		if currentExp >= level.RequiredEXP {
			lastLevelReq = level.RequiredEXP
			playerLevel += 1
		} else {
			lastLevelReq = level.RequiredEXP
			break
		}
	}
	currentLevelPercentComplete = 100*currentExp/lastLevelReq
	if(currentLevelPercentComplete > 100) {
		currentLevelPercentComplete = 100
	}


	return playerLevel, lastLevelReq, currentLevelPercentComplete
}
