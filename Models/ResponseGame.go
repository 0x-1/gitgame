package Models

import "github.com/xanzy/go-gitlab"

type ResponseGame struct {
	users []gitlab.User
	project gitlab.Project
}