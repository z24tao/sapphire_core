package server

import (
	"fmt"
	"log"
	"net/http"
)

func init() {
	http.HandleFunc("/board", boardHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
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
