package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func prettyPrintAlias(alias Alias) {
	fmt.Println("trigger:")
	fmt.Printf("./tool %s\n", alias.alias)
	fmt.Println("")
	fmt.Println("command:")
	fmt.Println(alias.command)
	fmt.Println("")
	if len(alias.path) > 0 {
		fmt.Println("executed in this path:")
		fmt.Println(alias.path)
		fmt.Println("")
	}
}

func main() {
	// This would be read from a JSON
	aliases := []Alias{
		{alias: "b", command: []string{"go", "build", "-o", "main"}, path: os.Getenv("PWD")},
		{alias: "e", command: []string{"echo", "\"foo\""}},
	}
	list := flag.Bool("l", false, "List the aliases available")
	alias := flag.String("r", "", "Run alias")
	flag.Parse()

	if *list {
		fmt.Printf("Aliases:\n\n")
		for _, alias := range aliases {
			prettyPrintAlias(alias)
		}
	} else if len(*alias) > 0 {
		for _, aliasFromConfiguration := range aliases {
			if aliasFromConfiguration.alias == *alias {
				command := exec.Command(aliasFromConfiguration.command[0], aliasFromConfiguration.command[1:]...)

				if len(aliasFromConfiguration.path) > 0 {
					command.Dir = aliasFromConfiguration.path
				}

				output, err := command.CombinedOutput()
				log.Printf("Running command and waiting for it to finish...")
				if err != nil {
					os.Stderr.WriteString(err.Error())
				}
				fmt.Println(string(output))
			}
		}
	} else {
		flag.PrintDefaults()
	}

	os.Exit(1)
}
