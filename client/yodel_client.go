package client

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/rpc"
	"os"
	"strconv"

	"github.com/fer2d2/micro-yodeler/cli"
	"github.com/fer2d2/micro-yodeler/command"
	"github.com/fer2d2/micro-yodeler/security"
)

func Run(config *cli.Config) {
	client, err := createClient(config)
	defer client.Close()

	if err != nil {
		log.Fatal("dialing: ", err)
	}

	var reply command.Result
	err = client.Call("ScriptRunner.Execute", config.Script, &reply)
	if err != nil {
		log.Fatal("Error: ", err)
	}

	fmt.Printf("StdOut=%v\nStdErr=%v\n\n", reply.StdOut, reply.StdErr)

	os.Exit(reply.ExitCode)
}

func createClient(config *cli.Config) (*rpc.Client, error) {
	var err error
	var client *rpc.Client

	if config.Tls {
		var connection *tls.Conn
		sslConfig := security.NewSecureConfig(config.PrivateKey, config.PublicKey, false)

		connection, err = tls.Dial("tcp", config.Hostname+":"+strconv.Itoa(config.Port), sslConfig)
		client = rpc.NewClient(connection)
	} else {
		client, err = rpc.Dial("tcp", config.Hostname+":"+strconv.Itoa(config.Port))
	}

	return client, err
}
