package server

import (
	"log"
	"net/http"
	"time"

	"github.com/amat-andre/Server/internal/handlers"
)

type Server struct {
	Log *log.Logger
	Serv http.Server
}

func NewRout(logger *log.Logger) *Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.MainHandler)
	mux.HandleFunc("/upload", handlers.UploadHandler)

	server := Server{
		Log: logger,
		Serv: http.Server{
			Addr: ":8080",
			Handler: mux,
			ErrorLog: logger,
			ReadTimeout:  5 * time.Second,
        	WriteTimeout: 10 * time.Second,
        	IdleTimeout:  15 * time.Second,
		},
	}

	return &server
}
