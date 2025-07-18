package main

import (
	"os"

	routes "example.com/gottrpg/v2/routes"
	server "example.com/gottrpg/v2/server"
)

func main() {
	// Check for command line argument to run dice roller
	if len(os.Args) > 1 && os.Args[1] == "dice" {
		rollDiceInteractive()
		return
	}

	if len(os.Args) > 1 && os.Args[1] == "serve" {
		router := server.Router{}
		router.Use("character", routes.Character)
		server.InitServer().ListenAndServe()
	}
}
