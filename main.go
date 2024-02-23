package main

import (
	"fmt"
	"net/http"
)

type Server struct {
}

func (s Server) ServeHTTP(writer http.ResponseWriter, request *http.Request) {

}

func main() {
	http.HandleFunc("/ping", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "hello %s", "pong")
	})

	http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "hello %s", "world")
	})

	http.HandleFunc("/hehe", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "hello %s", request.URL.Query().Get("key"))
	})

	if err := http.ListenAndServe(":9000", nil); err != nil {
		panic(err)
	}
}
