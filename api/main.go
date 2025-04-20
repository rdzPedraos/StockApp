package main

import (
	"app/cmd"
	"app/database/connection"
	"app/server"
)

func main() {
	connection.ConnectDatabase()

	// only if any command not was executed, run the server
	if cmd.Execute() {
		return
	}

	server.Run()
}
