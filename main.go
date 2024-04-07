package main

import (
	"net"
	"net/http"
)

func main() {
	l, err := net.Listen("tcp4", ":9090")

	if err != nil {
		panic(err)
	}

	http.HandleFunc("/*", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world"))
	})

	http.Serve(l, nil)
}
