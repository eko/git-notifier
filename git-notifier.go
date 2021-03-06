// A Git notifier for OS X
//
// Author: Vincent Composieux <vincent.composieux@gmail.com>

package main

import (
    "github.com/deckarep/gosx-notifier"
    "fmt"
    "os"
    "time"
)

const DIRECTORY = ".repositories"

// Run Git OSX Notifier
func main() {
    config := GetConfiguration()

    fmt.Printf("Git-notifier is loaded, please wait %d seconds for the first tick.\n", config.Frequency)

    ticker := time.NewTicker(time.Second * config.Frequency)

    func () {
        for _ = range ticker.C {
            fmt.Printf("Checking repositories for new commits...\n")

            CheckRepositories(config.Repositories)
        }
    }()
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

// Check for repositories to have new commits
func CheckRepositories(repositories []Repository) {
    mainDirectory := InitializeMainDirectory()

    for _, repository := range repositories {
        repository.Initialize(mainDirectory)

        currentSha1 := repository.FetchCurrentSha1()
        lastSha1    := repository.FetchLastSha1()

        if (currentSha1 != lastSha1) {
            commits := repository.GetDiff(currentSha1, lastSha1)

            for _, commit := range commits {
                SendNotification(repository, commit)
            }
        }
    }
}

// Send a notification
func SendNotification(repository Repository, commit Commit) {
    notification := gosxnotifier.NewNotification(commit.Message)

    notification.Group        = fmt.Sprintf("com.unique.gitnotifier.%s", commit.Sha1)
    notification.Title        = repository.Name
    notification.Subtitle     = commit.Author
    notification.Link         = commit.Link
    notification.ContentImage = "git.png"

    if len(repository.Logo) > 1 {
        notification.ContentImage = repository.Logo
    }

    e := notification.Push()
    check_error(e)
}

// Generic error management
func check_error(err error) {
    if err != nil { panic(err); os.Exit(1) }
}
