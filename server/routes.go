package server

import "github.com/OlivierArgentieri/go_killprocess/middlewares"

func (server *Server) initRoutes() {
	// route
	server.Router.HandleFunc("/kill/{pid}", middlewares.SetMiddlewareJSON(server.KillProcess)).Methods("POST")

}
