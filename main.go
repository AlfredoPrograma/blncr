package main

import (
	"fmt"
	"net"
)

func main() {
	config, err := readConfig()

	if err != nil {
		logFatal(err.Error())
	}

	if err != nil {
		logFatal(err.Error())
	}

	// Start listening requests via TCP
	listener, err := net.Listen("tcp", fmt.Sprintf(":%v", config.Port))

	if err != nil {
		logFatal(err.Error())
	}

	for {
		conn, err := listener.Accept()

		if err != nil {
			logWarn(fmt.Sprintf("Cannot handle incoming connection %s", conn.RemoteAddr()))
		}

		logInfo(fmt.Sprintf("Handling connection %s", conn.RemoteAddr()))
	}
}
