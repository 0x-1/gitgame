package OutputManager

import (
	"encoding/base64"
	"github.com/0x-1/GitGame/Managers/CryptManager"
	"github.com/0x-1/GitGame/Managers/GitLabManager"
	"github.com/0x-1/GitGame/Models"
	"github.com/xanzy/go-gitlab"
	"time"
)

func M_SaveAsWikiPage(gitLabURL string, projectName string, accessToken string, state Models.GitGameState, gitGameHost string, gitLabHost string)(error) {
	git := gitlab.NewClient(nil, accessToken)
	git.SetBaseURL(gitLabURL)

	project, err := GitLabManager.M_GetProjectByName(git,projectName)

	var content string


	content += "#GitGame Result created at: "+time.Now().String()+"<br/>"

	cryptedTokenBytes, err := CryptManager.M_Encrypt([]byte(accessToken))
	cryptedToken := base64.URLEncoding.EncodeToString(cryptedTokenBytes)



	content += "#Update this file by opening this link once: [link]("+gitGameHost+"/gitgame/update/"+projectName+"/"+cryptedToken+"?url="+gitLabHost+")<br/>"
	content += "#Players:<br/>"
	for _,player := range state.Players {
		content += player.MemberData.Name+"<br/>"
	}

	err = GitLabManager.M_CreateEditWikiPage(git, project.ID, content)
	if (err != nil) {
		return err
	}

	return nil
}
