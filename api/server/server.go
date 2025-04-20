package server

import (
	"app/config"
	"app/routes"
	"fmt"
	"log"
)

func Run() {
	router := routes.Router()
	address := fmt.Sprintf("%s:%s", config.Server.Host, config.Server.Port)

	router.Run(address)
	log.Printf("Server running on %s", address)
}
