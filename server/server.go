package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/OlivierArgentieri/go_killprocess/middlewares"
	"github.com/gorilla/mux"
)

type Server struct {
	Router *mux.Router
}

func (server *Server) initRoutes() {
	// route
	server.Router.HandleFunc("/kill/{pid}", middlewares.SetMiddlewareJSON(server.KillProcess)).Methods("POST")

}

func (server *Server) init() {
	server.Router = mux.NewRouter()
}

func (server *Server) Run(addr string) {
	server.init()
	server.initRoutes()

	fmt.Println("Listening at", addr)
	log.Fatal(http.ListenAndServe(addr, server.Router))
}
