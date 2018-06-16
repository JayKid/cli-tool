package main

// Alias the object containing an alias and its command
type Alias struct {
	Alias   string   `json:"alias"`
	Command []string `json:"command"`
	Path    string   `json:"path"`
}
