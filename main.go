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

	router := server.Router{}

	// TODO: create character and associated routes (modify attrs, spells, etc.)

	// TODO: going to need a lot of .Use() statements
	router.Use("character", routes.Character)
	server.InitServer().ListenAndServe()
}
