package main

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"os"

	"golang.org/x/crypto/acme/autocert"
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

func loggingMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		infof("%s from %s at %s", r.Method, r.Host, r.URL)
		next.ServeHTTP(w, r)
	})
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
	// if setLogFile is given a empty string it defaults to stdout
	setLogFile(logFilePath)
	infof("simpleStatic is Serving from %s on :%s", staticDirectory, port)

	httpsManager := &autocert.Manager{
		Cache:      autocert.DirCache("/home/wesley/go/src/simpleStatic/testcache"),
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist("localhost.com"),
	}

	// NOTE:
	// we could make the error logger the logger in the log file
	// there's a global variable named logger
	// shouldn't bite me in the ass
	server := &http.Server{
		// Addr: ":https",
		Addr: ":443",
		// TLSConfig: httpsManager.TLSConfig(),
		TLSConfig: &tls.Config{
			ServerName:     "localhost",
			GetCertificate: httpsManager.GetCertificate,
		},
		Handler:  loggingMiddleWare(http.FileServer(http.Dir(staticDirectory))),
		ErrorLog: logger,
	}

	err := server.ListenAndServeTLS("", "")

	// err := http.ListenAndServe(":"+port, loggingMiddleWare(http.FileServer(http.Dir(staticDirectory))))

	if err != nil {
		fatal(err.Error())
	}
}
