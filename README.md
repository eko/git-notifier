Git Notifier for OS X
=====================

[![GoDoc](http://godoc.org/github.com/eko/git-notifier?status.png)](http://godoc.org/github.com/eko/git-notifier)

A Golang Git notifier for Mac OS X.
This application will notify you when people are committing on your defined Git repositories and allows you to click on to visualize the commit.

![Notification example](../master/notification-example.jpg?raw=true)

# Installation

```bash
$ go get -u github.com/deckarep/gosx-notifier
$ git clone git@github.com:eko/git-notifier.git
```

# Configuration

You have to write a `JSON` configuration file with the following options:

```json
{
    "frequency": 60,
    "repositories": [
        {
            "name": "Git-Notifier",
            "logo": "/path/to/my-project-logo.png",
            "branch": "master",
            "git": "git@github.com:eko/git-notifier.git",
            "commit_url": "https://github.com/eko/git-notifier/commit/"
        }
    ]
}
```

* `frequency` corresponds to frequency your repositories will be checked for new commits (in seconds),
* `repositories` is a list of your repositories with some information about it: name, branch and git url.

# Run

```bash
$ go build
$ go run *.go
```
