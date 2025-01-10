package server

import (
	"github.com/kingtingthegreat/top-fetch-old/middleware"
	"github.com/kingtingthegreat/top-fetch-old/router"
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
