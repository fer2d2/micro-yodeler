# MicroYodeler

MicroYodeler is an extremely simple client-server application written in Golang
to run remote scripts between two *blind* docker containers. It could be used,
for example, to execute a database script inside your MySQL container from a PHP
container.

This project is intended for development purposes only, and it **should not be
used in production environments**.

## How it works

1. Create a directory in your target container with your scripts. Scripts must be executable and contain a *shebang* (like `#!/bin/bash` or `#!/usr/bin/env python`)
2. Run the server in your target container with `yodel server -d /your/scripts/directory -p 5058`
3. Run the client in your source container with `yodel client -p 5058 -h container_hostname -f script_name.sh`

The *yodel client* will return the script exit code, so it could be used in some automation tools like *rake* to simplify your tasks workflow.

### Command line flags
```
usage: yodel <command> [<args>]
The most commonly used yodel commands are:
 server  Starts the server on the specified scripts directory
   Usage of server:
    -d string
      	Directory to look for the executable scripts
    -p int
      	Port to listen (default 5058)
 client  Executes a script from the remote server
   Usage of client:
     -f string
       	Script to execute
     -h string
       	Hostname of the remote container (default "localhost")
     -p int
       	Port of the server in the remote container (default 5058)
```
