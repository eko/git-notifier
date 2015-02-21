package main

import (
    "github.com/deckarep/gosx-notifier"
	"fmt"
	"os"
)

const DIRECTORY = ".repositories"

// Run Git OSX Notifier
func main() {
    fmt.Print("Starting git-notifier...\n")

    config        := GetConfiguration()
    mainDirectory := InitializeMainDirectory()

    for _, repository := range config.Repositories {
        repository.Initialize(mainDirectory)

        // Check if Git repository has already been cloned
        currentSha1 := repository.FetchCurrentSha1()
        lastSha1    := repository.FetchLastSha1()

        if (currentSha1 != lastSha1) {
            commits := repository.GetDiff(currentSha1, lastSha1)
            
            for _, commit := range commits {
                SendNotification(repository.Name, commit.Message, commit.Author, commit.Link, repository.Logo)
            }
        }
    }
}

// Initializes main/temp repositories directory
func InitializeMainDirectory() string {
    currentDirectory, e := os.Getwd()
    mainDirectory := fmt.Sprintf("%s/%s", currentDirectory, DIRECTORY)

    _, e = os.Stat(mainDirectory)

    if e != nil {
        e := os.Mkdir(mainDirectory, 0777)
        check_error(e)

        fmt.Printf("Directory created: %s\n", mainDirectory)
    }

    return mainDirectory
}

// Send a notification
func SendNotification(title string, subtitle string, message string, link string, logo string) {
    notification := gosxnotifier.NewNotification(message)

    notification.Title        = title
    notification.Subtitle     = subtitle
    notification.Link         = link
    notification.ContentImage = "git.png"

    if len(logo) > 1 {
        notification.ContentImage = logo
    }

    e := notification.Push()
    check_error(e)
}

// Generic error management
func check_error(err error) {
    if err != nil { panic(err); os.Exit(1) }
}
