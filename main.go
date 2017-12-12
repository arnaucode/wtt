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
		case "new", "n":
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
		case "list", "ls":
			if len(os.Args) > 2 {
				param := os.Args[2]
				if param == "-a" {
					listProjectsDetails()
				}
			} else {
				//list projects
				listProjects()
			}
			break
		case "start", "s":
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
			} else {
				color.Red("No project name to start selected")
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
			color.Green("Worked " + work.Projects[i].Streaks[j].Duration.String() + " in the project " + work.Projects[i].Name)
			break
		case "rm":
			if len(os.Args) > 2 {
				projectName := os.Args[2]
				if work.CurrentProjectName == projectName {
					work.CurrentProjectName = ""
				}
				deleteProject(projectName)
				saveWork()
				color.Yellow("Project " + projectName + " deleted")
			} else {
				color.Red("no project name specified")
			}

			break
		case "current", "c":
			if work.CurrentProjectName != "" {
				fmt.Print("Current working project: ")
				color.Blue(work.CurrentProjectName)
			} else {
				fmt.Println("No current working project.")
			}
		case "help", "h":
			fmt.Println("./wtt new {projectname}")
			fmt.Println("./wtt ls")
			fmt.Println("./wtt ls -a")
			fmt.Println("./wtt start {projectname}")
			fmt.Println("./wtt stop")
			fmt.Println("./wtt rm")
			fmt.Println("./wtt current")
			fmt.Println("./wtt help")
		default:
			color.Red("option not exists")
			os.Exit(1)
		}
	} else {
		color.Red("no option selected")
		fmt.Println("Can run 'help' for commands information")
		fmt.Println("./wtt help")
		os.Exit(1)
	}
}
