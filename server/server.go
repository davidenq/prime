package server

import (
	"log"
	"net/http"
)

//Server .
type Server struct {
	Port string
}

//NewHTTP .
func (s *Server) NewHTTP() {
	http.HandleFunc("/check", CheckIfIsPrime)
	http.HandleFunc("/genate", GenerateLargePrime)

	log.Fatal(http.ListenAndServe(s.Port, nil))
}
