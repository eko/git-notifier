package main

import (
	"fmt"
	"os"
	"os/exec"
    "io/ioutil"
    "strings"
)

// A commit structure
type Commit struct {
	Link       string
	Author     string
	Message    string
}

// Initializes current repository
func (r *Repository) Initialize(mainDirectory string) {
	r.Directory = fmt.Sprintf("%s/%s", mainDirectory, r.Name)
}

// Fetches current repository commit SHA1
func (r *Repository) FetchCurrentSha1() string {
    var sha1 []byte

    _, e := os.Stat(r.Directory)

    if e != nil {
        fmt.Printf("Cloning repository %s in %s\n", r.Name, r.Directory)

        command   := fmt.Sprintf("git clone --depth=1 -n -b %s %s %s && cd %s && git log --pretty=oneline -1 | cut -d ' ' -f 1 | tr -d '\n'", r.Branch, r.Git, r.Directory, r.Directory)
        output, e := exec.Command("sh", "-c", command).Output()
        check_error(e)

        fmt.Printf("Repository %s cloned to SHA1: %s\n", r.Name, output)

        // Write last SHA1
        sha1 = []byte(output)
        e = ioutil.WriteFile(fmt.Sprintf("%s/SHA1", r.Directory), sha1, 0644)
        check_error(e)
    } else {
        sha1, e = ioutil.ReadFile(fmt.Sprintf("%s/SHA1", r.Directory))
        check_error(e)
    }

    return string(sha1)
}

// Fetches and retrieve current repository new commit SHA1
func (r *Repository) FetchLastSha1() string {
    var sha1 []byte

    command   := fmt.Sprintf("cd %s && git pull -q origin %s && git log --pretty=oneline -1 | cut -d ' ' -f 1 | tr -d '\n'", r.Directory, r.Branch)
    output, e := exec.Command("sh", "-c", command).Output()
    check_error(e)

    fmt.Printf("Repository %s updated to SHA1: %s\n", r.Name, output)

    sha1 = []byte(output)
    //e = ioutil.WriteFile(fmt.Sprintf("%s/SHA1", r.Directory), sha1, 0644)
    //check_error(e)

    return string(sha1)
}

// Returns an array of commits corresponding to diff between two sha1
func (r *Repository) GetDiff(fromSha1 string, toSha1 string) []Commit {
    commits := make([]Commit, 0)

    command   := fmt.Sprintf("cd %s && git log %s...%s --pretty=oneline --format='%%H#separator#%%an#separator#%%s'", r.Directory, fromSha1, toSha1)
    output, e := exec.Command("sh", "-c", command).Output()
    check_error(e)

    results := strings.Split(string(output), "\n")

    for _, result := range results {
    	if len(result) > 0 {
	        parts := strings.Split(result, "#separator#")
	        commits = append(commits, Commit{fmt.Sprintf("%s%s", r.CommitUrl, parts[0]), parts[1], parts[2]})
        }
    }

    return commits
}