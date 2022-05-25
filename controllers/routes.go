package controllers

import "github.com/OlivierArgentieri/go_killprocess/middlewares"

func (server *Server) initRoutes() {
	// route

	// process api
	server.Router.HandleFunc("/kill/{pid}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareCORS(server.KillProcess))).Methods("POST")
	server.Router.HandleFunc("/processes", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareCORS(server.GetProcesses))).Methods("GET")

	// service api
	server.Router.HandleFunc("/restartservice", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareCORS(server.RestartServices))).Methods("POST")
	server.Router.HandleFunc("/services", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareCORS(server.GetServices))).Methods("GET")
}
