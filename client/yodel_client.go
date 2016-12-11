package client

import (
	"fmt"
	"log"
	"net/rpc"
	"os"
	"strconv"

	"github.com/fer2d2/micro-yodeler/command"
)

func Run(port int, script string, hostname string) {
	client, err := rpc.DialHTTP("tcp", hostname+":"+strconv.Itoa(port))
	if err != nil {
		log.Fatal("dialing: ", err)
	}

	var reply command.Result
	err = client.Call("ScriptRunner.Execute", script, &reply)
	if err != nil {
		log.Fatal("Error: ", err)
	}

	fmt.Printf("StdOut=%v\nStdErr=%v\n\n", reply.StdOut, reply.StdErr)

	os.Exit(reply.ExitCode)
}
