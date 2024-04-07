package main

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
)

func handleConnection(conn net.Conn) {
	connReader := bufio.NewReader(conn)
	req, err := http.ReadRequest(connReader)

	if err != nil {
		logWarn(err.Error())
	}

	logDebug(fmt.Sprintf("%#v", req))
}
