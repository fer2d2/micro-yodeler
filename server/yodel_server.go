package server

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
	"strconv"
)

func Run(port int, path string) {
	scriptRunner := NewScriptRunner(path)
	rpc.Register(scriptRunner)
	rpc.HandleHTTP()

	server, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	if err != nil {
		log.Fatal("listen error: ", err)
	}

	http.Serve(server, nil)
}
