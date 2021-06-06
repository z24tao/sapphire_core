package server

import (
	"fmt"
	"log"
	"net/http"
)

func init() {
	http.HandleFunc("/board", boardHandler)
	go startServer()
}

func startServer() {
	err := http.ListenAndServe(":8080", nil)
	log.Fatal(err)
}

func serveData(w http.ResponseWriter, data interface{}) {
	dataStr := fmt.Sprintf("%v", data)
	_, err := w.Write([]byte(dataStr))
	handleErr(err)
}

func handleErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
