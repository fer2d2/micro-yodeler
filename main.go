package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/fer2d2/micro-yodeler/cli"
	"github.com/fer2d2/micro-yodeler/client"
	"github.com/fer2d2/micro-yodeler/server"
)

func main() {
	config := new(cli.Config)

	serverCommand := flag.NewFlagSet("server", flag.ExitOnError)
	clientCommand := flag.NewFlagSet("client", flag.ExitOnError)

	if notEnoughArguments() {
		showUsageMessage()
		return
	}

	switch os.Args[1] {
	case "server":
		cli.DeclareServerFlags(config, serverCommand)
		serverCommand.Parse(os.Args[2:])
		server.Run(config)
	case "client":
		cli.DeclareClientFlags(config, clientCommand)
		clientCommand.Parse(os.Args[2:])
		client.Run(config)
	default:
		fmt.Printf("%q is not valid command.\n", os.Args[1])
		os.Exit(2)
	}

}

func notEnoughArguments() bool {
	return (len(os.Args) == 1)
}

func showUsageMessage() {
	fmt.Println("usage: yodel <command> [<args>]")
	fmt.Println("The most commonly used yodel commands are: ")
	fmt.Println(" server  Starts the server on the specified scripts directory")
	fmt.Println(" client  Executes a script from the remote server")
}
