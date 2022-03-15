package server

import "github.com/OlivierArgentieri/go_killprocess/middlewares"

func (server *Server) initRoutes() {
	// route
	server.Router.HandleFunc("/kill/{pid}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareCORS(server.KillProcess))).Methods("POST")
	server.Router.HandleFunc("/processes", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareCORS(server.GetProcesses))).Methods("GET")
}
