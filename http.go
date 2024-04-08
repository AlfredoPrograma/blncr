package main

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
)

func handleConnection(conn net.Conn, balancer *Balancer) {
	connReader := bufio.NewReader(conn)
	req, err := http.ReadRequest(connReader)

	if err != nil {
		logWarn(err.Error())
	}

	logDebug(fmt.Sprintf("Request: %#v", req))
	logDebug(fmt.Sprintf("Server: %v", balancer.getServer()))
}
