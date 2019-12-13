package main

import (
	"fmt"
	"net/http"
	"os"
)

const version = "0.1.0"

func printUsage() {
	usage := `Usage: simpleStatic [FILE | -h | -v]\n
simpleStatic is a simple and easy to configure http static server.

give simpleStatic a path to a configuration file like so:
	simpleStatic /home/wesley/.config/simpleStatic.conf

a configuration file is lines of key/value pairs delimited by a ':'
the defaults are like this:
	port:8888
	static_directory:./static
	log_file_path:
	# NOTE: if the log file path is empty it will default to stdin
	# Also, that there are no spaces between the key and the value
		`
	fmt.Println(usage)
}

func main() {
	configFilePath := ""

	if len(os.Args) >= 2 {
		if os.Args[1] == "--help" || os.Args[1] == "-h" {
			printUsage()
			os.Exit(0)
		} else if os.Args[1] == "--version" || os.Args[1] == "-v" {
			fmt.Println(version)
			os.Exit(0)
		}
		configFilePath = os.Args[1]
	}

	configs := make(map[string]string)
	setConfigs(configs, configFilePath)
	port := configs["port"]
	staticDirectory := configs["static_directory"]
	logFilePath := configs["log_file_path"]

	// NOTE:
	// if setLogFile is given a empty string it defaults to sdout

	setLogFile(logFilePath)
	infof("simpleStatic is Serving from %s on :%s", staticDirectory, port)
	err := http.ListenAndServe(":"+port, http.FileServer(http.Dir(staticDirectory)))
	if err != nil {
		fatal(err.Error())
	}
}
