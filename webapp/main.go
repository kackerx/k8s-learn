package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer, "webapp")
	})

	fmt.Println("web app running [9000]")
	if err := http.ListenAndServe(":9000", nil); err != nil {
		panic(err)
	}
}
