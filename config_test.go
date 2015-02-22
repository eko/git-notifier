// A Git notifier for OS X
//
// Author: Vincent Composieux <vincent.composieux@gmail.com>

package main

import (
    "testing"
)

// Tests getting the configuration values
func TestGetConfiguration(t *testing.T) {
    config := GetConfiguration()

    if (config.Frequency != 5) {
        t.Fail()
    }

    if (len(config.Repositories) != 1) {
        t.Fail()
    }

    repository := config.Repositories[0]

    if (repository.Name != "Git-Notifier") {
        t.Fail()
    }

    if (repository.Logo != "git.png") {
        t.Fail()
    }

    if (repository.Branch != "master") {
        t.Fail()
    }

    if (repository.Git != "git@github.com:eko/git-notifier.git") {
        t.Fail()
    }

    if (repository.CommitUrl != "https://github.com/eko/git-notifier/commit/") {
        t.Fail()
    }
}
