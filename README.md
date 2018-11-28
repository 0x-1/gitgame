# GitGame
GitGame is a web service written in Go (GoLang). It's a work-in-progess gamification Solution that features a simplistic but extendable way to gamify your repository.

`.gitgame file + http call = wiki page`


## Content
* [Installation](#Installation)
* [Configuration](#Configuration)
* [Usage](#Usage)
* [Result](#Result)

## Installation
To install Gin package, you need to install Go and set your Go workspace first.

1. Download & Install

`go get -u github.com/0x-1/gitgame`

2. Run

`go run %gopath%\src\github.com\0x-1\gitgame\main.go`

## Configuration
1. Create .gitgame File in Repository

2. Modify .gitgame with new game rules


Currently supported game rules:

* quest **[scope]** _issue|commit_ **[points]** **[constraint]**  
Quests are repeatable and rewards **[points]**  
Currently supports either _issue or commit_  
issue currently supports _opened_ as **[constraint]**  
commit currently supports _created_ as **[constraint]**

* achievement **[scope]** milestone|pipeline **[points]** **[constraint]**  
Achievements reward **[points]** only once and can't be repeated  
Currently supports either _milestone or pipeline_  
milestone currently supports _opened_ as **[constraint]**  
pipeline currently supports _created_ as **[constraint]**

* level [required points]  
This creates a new level for all players with [required points] as the absolut boundary

Example .gitgame Configuration File:
```
//Konfiguration  
level 25  
level 50  
level 75  
level 100  
level 125  
level 150  
  
quest player commit 1 created  
quest player issue 2 opened  
achievement project pipeline 60 created
```

## Usage
After installing, running and creating the .gitgame file you can call the service with a http call

`http://[service-host]/gitgame/[namespace]/[projekt]?token=[gitlab-token]&url=[gitlab-url]`

(work in progress)

## Result

(work in progress)
