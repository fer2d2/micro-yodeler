package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/fer2d2/micro-yodeler/client"
	"github.com/fer2d2/micro-yodeler/server"
)

func main() {
	var port int
	var script string
	var hostname string
	var path string

	serverCommand := flag.NewFlagSet("server", flag.ExitOnError)
	clientCommand := flag.NewFlagSet("client", flag.ExitOnError)

	serverCommand.IntVar(&port, "p", 5058, "Port to listen")
	serverCommand.StringVar(&path, "d", "", "Directory to look for the executable scripts")

	clientCommand.IntVar(&port, "p", 5058, "Port to listen")
	clientCommand.StringVar(&script, "f", "", "Script to execute")
	clientCommand.StringVar(&hostname, "h", "localhost", "Hostname of the remote server")

	if len(os.Args) == 1 {
		fmt.Println("usage: yodel <command> [<args>]")
		fmt.Println("The most commonly used yodel commands are: ")
		fmt.Println(" server  Starts the server on the specified scripts directory")
		fmt.Println(" client  Executes a script from the remote server")
		return
	}

	switch os.Args[1] {
	case "server":
		serverCommand.Parse(os.Args[2:])
		server.Run(port, path)
	case "client":
		clientCommand.Parse(os.Args[2:])
		client.Run(port, script, hostname)
	default:
		fmt.Printf("%q is not valid command.\n", os.Args[1])
		os.Exit(2)
	}
}
