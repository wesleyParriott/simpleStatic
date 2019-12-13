# simpleStatic 

A simple and easy to use Static Web Server

### Version 0.1.0

## Installation
_NOTE: this is not tested on windows!_

```
go get simpleStatic
cd ~/go/src/github.com/wesleyParriott/simpleStatic
go build && go install
cd -
```

# Usage
```
simpleStatic /home/wesley/.config/simpleStatic.conf
```

# Config 
The configuration format is just key:value seperated by newlines. Comments are denoted by _#_.

## Example of cofiguration with the configuration defaults
```
# the port that the simpleStatic web server opens and runs on
port:8888
# the static server where the html files would be located
static_directory:./static
# the log file path is where the simpleStatic webserver is going to log to
# NOTE: if the log file path is empty it will default to stdin
log_file_path:
```
