package main

import "github.com/OlivierArgentieri/go_killprocess/server"

func main() {

	ap := server.Server{}
	ap.Run(":8081")
}
