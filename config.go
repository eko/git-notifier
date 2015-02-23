// A Git notifier for OS X
//
// Author: Vincent Composieux <vincent.composieux@gmail.com>

package main

import (
    "encoding/json"
    "io/ioutil"
    "time"
)

const FILENAME  = "config.json"

// Configuration JSON main structure
type Configuration struct {
    Frequency    time.Duration `json:"frequency"`
    Repositories []Repository `json:"repositories"`
}

// Repository JSON structure
type Repository struct {
    Name      string `json:"name"`
    Logo      string `json:"logo"`
    Branch    string `json:"branch"`
    Git       string `json:"git"`
    CommitUrl string `json:"commit_url"`
    Directory string
}

// Returns configuration from JSON file
func GetConfiguration() *Configuration {
    file, e := ioutil.ReadFile(FILENAME)
    check_error(e)
 
    config := &Configuration{}
    e = json.Unmarshal(file, &config)
    check_error(e)

    return config
}
