package server

import "github.com/OlivierArgentieri/go_killprocess/middlewares"

func (server *Server) initRoutes() {
	// route
	server.Router.HandleFunc("/kill/{pid}", middlewares.SetMiddlewareJSON(server.KillProcess)).Methods("POST")
	server.Router.HandleFunc("/processes", middlewares.SetMiddlewareJSON(server.GetProcesses)).Methods("GET")
}
