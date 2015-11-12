# sshexec

Program to SSH into remote servers and execute commands

## Usage:

* `go run sshexec.go`
Starts the program. It looks for cfg.yml file in the same directory for configuration options.

* `go run sshexec.go --cfg config.yml`
Starts the program. It looks for config.yml file in the same directory for configuration options.

## Configuration Options:

The following configuration options can be specified in the yaml file:

* `authtype` should be one of `password`, `publickey` or `sshagent`
 * `password` is for Password based authentication. It requires password field be specified in the configuration as well.
 * `publickey` is for SSH Key Pair based authentication. It requires the path to private key be specified in the configuration as well.
 * `sshagent` is for SSH Agent based authentication. This requires that the key be used with `ssh-add` before invoking this program.
* `username` is used for SSH authentication if specified, else current user is used.
* `password` is required only for Password based authentication. See above.
* `privatekeyfile` is required only for SSH Key Pair based authentication. See above.
* `hostsfile` specifies the path to text file containing list of servers on which the command(s) need to be executed. The server host names must be specified one per line.
* `port` specifies the port to connect to. This port is appended to each of the server host name in the `hostsfile`, provided the server name does not already contain a port.
* `envfile` specifies any environment variables that need to be set before executing each of the command(s). These should be specified in the format `key=value` with one key value pair per line.
* `commandsfile` specifies the command(s) that need to be executed on each of the servers. Command(s) should be specified one per line.

## Sample Configuration File:

```
authtype     : sshagent
hostsfile    : servers.txt
port         : 22
commandsfile : commands.txt
```

```
authtype     : password
username     : <username>
password     : <password>
hostsfile    : servers.txt
port         : 22
commandsfile : commands.txt
```

```
authtype       : publickey
privatekeyfile : ~/.ssh/id_rsa
hostsfile      : servers.txt
port           : 22
commandsfile   : commands.txt
```
