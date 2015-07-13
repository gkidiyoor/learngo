# fileserver

Serves files from the specified directory on the specified local port

## Usage:

* `./fileserver`
Serves files in the `.` directory on `localhost:9000`

* `./fileserver -port="8080" -dir="/home"`
Serves files in `/home` directory on `localhost:8080`
_NOTE: Both `port` and `dir` need to be specified as a strings in the arguments_

* `./fileserver --help`
Displays the usage information

