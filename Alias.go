package main

import "fmt"

// Alias the object containing an alias and its command
type Alias struct {
	Alias   string   `json:"alias"`
	Command []string `json:"command"`
	Path    string   `json:"path"`
}

// PrettyPrint is the method that encapsulates how to display the alias in a human readable way
func (alias Alias) PrettyPrint() {
	fmt.Println("trigger:")
	fmt.Printf("./tool %s\n", alias.Alias)
	fmt.Println("")
	fmt.Println("command:")
	fmt.Println(alias.Command)
	fmt.Println("")
	if len(alias.Path) > 0 {
		fmt.Println("executed in this path:")
		fmt.Println(alias.Path)
		fmt.Println("")
	}
}
