package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func configKeyIsProper(configKey string) bool {
	properConfigKeys := [3]string{"port", "static_directory", "log_file_path"}

	for _, properConfigKey := range properConfigKeys {
		if configKey == properConfigKey {
			return true
		}
	}
	return false
}

func setConfigs(configs map[string]string, configFilePath string) {
	configs["port"] = "8888"
	configs["static_directory"] = "./static"
	// NOTE will default to stdout in log.go
	configs["log_file_path"] = ""

	if configFilePath == "" {
		return
	}

	configFileContentsByteArray, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		fmt.Printf("[CONF FATAL] %s\n", err)
		os.Exit(1)
	}

	configFileLines := strings.Split(string(configFileContentsByteArray), "\n")
	// config should look like
	// port:8080
	// log_file_path:/var/log/simpleStatic.log
	// etc...
	for configFileLineNumber, configFileLine := range configFileLines {
		if configFileLine == "\n" || configFileLine == "" || configFileLine[0] == '#' {
			continue
		}

		configKeyValue := strings.Split(configFileLine, ":")

		if len(configKeyValue) != 2 {
			fmt.Printf("[CONF WARN] Improper config given in %s on line %d.\n %s\n", configFilePath, configFileLineNumber, configFileLine)
			continue
		}

		key := configKeyValue[0]
		value := configKeyValue[1]

		if !configKeyIsProper(key) {
			fmt.Printf("[CONF WARN] Improper key \"%s\" on line %d in %s. Should only be port, static_directory, or log_file_path\n", key, configFileLineNumber, configFilePath)
			continue
		}
		configs[key] = value
	}
}
