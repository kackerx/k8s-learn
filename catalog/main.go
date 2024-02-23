package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer, "hello catalog v2")
	})
	http.HandleFunc("/items", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")
		data := []string{"apple", "yell", "orid", "v2", "v2", "v2"}
		dataJson, err := json.Marshal(data)
		if err != nil {
			return
		}

		writer.Write(dataJson)
	})

	fmt.Println("catalog running in [3000]")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		panic(err)
	}
}
