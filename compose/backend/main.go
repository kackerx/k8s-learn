package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer, "webapp")
	})

	http.HandleFunc("/user", func(writer http.ResponseWriter, request *http.Request) {
		data := map[string]string{"name": "user"}
		dataByte, err := json.Marshal(&data)
		if err != nil {
			fmt.Fprintln(writer, err)
			return
		}

		writer.Write(dataByte)
	})

	fmt.Println("web app running [9000]")
	if err := http.ListenAndServe(":9000", nil); err != nil {
		panic(err)
	}
}
