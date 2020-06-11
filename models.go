package models

//Greeting provides entry point information
type Greeting struct {
	Version string   `json:"version"`
	Links   []string `json:"links"`
}
