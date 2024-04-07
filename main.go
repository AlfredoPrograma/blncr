package main

import (
	"fmt"
	"net/http"
)

func main() {
	config, err := readConfig()

	if err != nil {
		logFatal(err.Error())
	}

	if err != nil {
		logFatal(err.Error())
	}

	// Handling request and forwarding to servers
	http.HandleFunc("/*", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world"))
	})

	logInfo(fmt.Sprintf("Load balancer running at port %v", config.Port))
	http.ListenAndServe(fmt.Sprintf(":%v", config.Port), nil)
}
