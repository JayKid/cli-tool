package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func getArguments() (userConfigurationPath *string, showList *bool, userSuppliedAlias *string) {
	userConfigurationPath = flag.String("c", "", "Supply configuration path")
	showList = flag.Bool("l", false, "List the aliases available")
	userSuppliedAlias = flag.String("r", "", "Run alias")
	flag.Parse()
	return
}

func determineConfigurationPath(userConfigurationPath *string) (configurationPath string) {
	toolExecutable, err := os.Executable()
	if err != nil {
		panic(err)
	}
	toolPath := filepath.Dir(toolExecutable)

	configurationPath = toolPath + "/config/aliases.json"

	if len(*userConfigurationPath) > 0 {
		configurationPath = *userConfigurationPath
	}
	return
}

func parseAliasesFromConfiguration(configurationPath string) (aliases []Alias) {
	rawJSON, err := ioutil.ReadFile(configurationPath)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	aliasesBody := []byte(rawJSON)
	aliases = make([]Alias, 0)
	json.Unmarshal(aliasesBody, &aliases)
	return
}

func printAliases(aliases []Alias) {
	fmt.Printf("Aliases:\n\n")
	for _, alias := range aliases {
		alias.PrettyPrint()
	}
}

func runAlias(aliases []Alias, userSuppliedAlias *string) {
	for _, aliasFromConfiguration := range aliases {
		if aliasFromConfiguration.Alias == *userSuppliedAlias {
			command := exec.Command(aliasFromConfiguration.Command[0], aliasFromConfiguration.Command[1:]...)

			if len(aliasFromConfiguration.Path) > 0 {
				if aliasFromConfiguration.Path == "PWD" {
					command.Dir = os.Getenv("PWD")
				} else {
					command.Dir = aliasFromConfiguration.Path

				}
			}

			output, err := command.CombinedOutput()
			log.Printf("Running command and waiting for it to finish...")
			if err != nil {
				os.Stderr.WriteString(err.Error())
			}
			fmt.Println(string(output))
		}
	}
}

func main() {

	userConfigurationPath, showList, userSuppliedAlias := getArguments()
	configurationPath := determineConfigurationPath(userConfigurationPath)
	aliases := parseAliasesFromConfiguration(configurationPath)

	if *showList {
		printAliases(aliases)
	} else if len(*userSuppliedAlias) > 0 {
		runAlias(aliases, userSuppliedAlias)
	} else {
		flag.PrintDefaults()
	}

	os.Exit(1)
}
