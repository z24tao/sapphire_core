package server

import (
	"github.com/z24tao/sapphire_core/world"
	"net/http"
)

func boardHandler(w http.ResponseWriter, r *http.Request) {
	serveData(w, world.GetDefaultBoardState())
}
