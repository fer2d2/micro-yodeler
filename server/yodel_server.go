package server

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
	"strconv"

	"github.com/fer2d2/micro-yodeler/cli"
)

func Run(config *cli.Config) {
	scriptRunner := NewScriptRunner(config.Path)
	rpc.Register(scriptRunner)
	rpc.HandleHTTP()

	server, err := net.Listen("tcp", ":"+strconv.Itoa(config.Port))
	if err != nil {
		log.Fatal("listen error: ", err)
	}

	http.Serve(server, nil)
}
