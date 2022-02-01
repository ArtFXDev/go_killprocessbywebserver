package main

import "github.com/OlivierArgentieri/go_killprocess/server"

func main() {
	app := server.Server{}
	app.Run(":5119")
}
