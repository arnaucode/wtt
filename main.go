package main

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"time"

	"github.com/fatih/color"
)

var directoryPath string
var filePath string

func main() {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	directoryPath = usr.HomeDir + "/.wtt"
	filePath = directoryPath + "/work.json"
	readProjects()

	//get the command line parameters
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "new":
			//create project os.Args[2]
			if len(os.Args) > 2 {
				projectName := os.Args[2]
				color.Green("creating new project: " + projectName)
				err := newProject(projectName)
				if err != nil {
					break
				}
				saveWork()
			} else {
				color.Red("No project name specified")
			}
			break
		case "list":
			//list projects
			listProjects()
			break
		case "start":
			//check if already there is any project started
			if work.CurrentProjectName != "" {
				color.Red("Can not start project, already project " + work.CurrentProjectName + " running")
				break
			}
			//start project os.Args[2]
			if len(os.Args) > 2 {
				projectName := os.Args[2]
				i := getProjectIByName(projectName)
				if i < 0 {
					color.Red("Project name: " + projectName + ", no exists")
					break
				}
				var newStreak Streak
				newStreak.Start = time.Now()
				work.Projects[i].Streaks = append(work.Projects[i].Streaks, newStreak)
				work.CurrentProjectName = projectName
				saveWork()
				fmt.Println("starting to work in project " + work.Projects[i].Name)
			}
			break
		case "stop":
			if work.CurrentProjectName == "" {
				color.Red("no project started to stop")
				break
			}
			i := getProjectIByName(work.CurrentProjectName)
			if i < 0 {
				color.Red("Project name: " + work.CurrentProjectName + ", no exists")
				break
			}
			j := len(work.Projects[i].Streaks) - 1
			work.Projects[i].Streaks[j].End = time.Now()
			work.Projects[i].Streaks[j].Duration = time.Now().Sub(work.Projects[i].Streaks[j].Start)
			work.CurrentProjectName = ""
			saveWork()
			fmt.Print("Worked ")
			fmt.Print(work.Projects[i].Streaks[j].Duration)
			fmt.Println(" in the project " + work.Projects[i].Name)

			//stop project os.Args[2]
			break
		default:
			color.Red("no option selected")
			os.Exit(1)
		}
	} else {
		color.Red("no option selected")
		os.Exit(1)
	}
}
