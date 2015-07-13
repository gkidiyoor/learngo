# echoserver

Simple HTTP server that echoes back the request body

## Usage:

* `./echoserver`
Starts the server on `localhost:9000`

* `./echoserver -port="8080"`
Starts the server on `localhost:8080`
_NOTE: `port` needs to be specified as a string in the arguments_

* `./echoserver --help`
Displays the usage information

## Testing:

Use `curl` to test GET and POST requests as follows:

* `curl localhost:9000` should return a welcome message

* `curl -X POST localhost:9000` should return a blank line

* `curl --data "key=value" localhost:9000` should return "key=value" as response

