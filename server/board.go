package server

import "net/http"
import "../world"

func boardHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getBoardHandler(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func getBoardHandler(w http.ResponseWriter, r *http.Request) {
	data := world.GetDefaultBoardState()
	serveData(w, data)
}
