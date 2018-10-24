package Models

import "github.com/xanzy/go-gitlab"

type GitGameState struct {
	Levels []Level
	Players []Player
}

type Player struct {
	MemberData gitlab.ProjectMember
	Experience int
}

type Level struct {
	RequiredEXP int
}