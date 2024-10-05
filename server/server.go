package server

import (
	"github.com/kingtingthegreat/top-fetch/middleware"
	"github.com/kingtingthegreat/top-fetch/router"
	"net/http"
)

func Server() *http.Server {
	router := router.Router()

	server := http.Server{
		Addr:    ":8080",
		Handler: middleware.Logger(router),
	}

	return &server
}
