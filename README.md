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

Some details about this configuration:

* `frequency` corresponds to frequency your repositories will be checked for new commits (in seconds),
* `repositories` is a list of your repositories with some information about it: name, branch and git url:
    * `name` is a repository name,
    * `logo` is an (optional) logo you can specify to appear on your notification,
    * `branch` is the branch you want to be checked for new commits,
    * `git` is the Git repository URL of your project,
    * `commit_url` is the HTTP URL which will be used to view the commits detail once you will click on the notification.

# Run

```bash
$ go build
$ ./git-notifier
```

I suggest you to add it to the User program at startup in order to avoid launching it manually.
