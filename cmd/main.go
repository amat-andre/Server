package main

import (
	"log"
	"net/http"

	s "github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	
	logger := log.New(log.Writer(), "", log.LstdFlags|log.Lshortfile) 
	server := s.NewRout(logger)

	err := http.ListenAndServe(server.Serv.Addr, server.Serv.Handler)
	if err != nil {
		logger.Fatal(err)
	}
}

