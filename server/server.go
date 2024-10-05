package server

import (
	"net/http"
	"github.com/kingtingthegreat/top-fetch/router"
	"github.com/kingtingthegreat/top-fetch/middleware"
)

func Server() http.Server {
	router := router.Router()

	server := http.Server{
		Addr:    ":8080",
		Handler: middleware.Logger(router),
	}

	return server
}

