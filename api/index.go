package handler

import (
	"github.com/kingtingthegreat/top-fetch-old/server"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	server := server.Server()
	server.Handler.ServeHTTP(w, r)
}
