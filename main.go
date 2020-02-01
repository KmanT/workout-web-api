package main

import (
	"log"
	"net/http"
	"os"

	// change the dot to the name of the project
	// if it does not run
	"./config"
	"./routes"
)

func main() {
	// Heroku has the env variable "PORT" set
	// while chances are you do not. Keep it
	// this way.
	port := os.Getenv("PORT")

	if port == "" {
		// It will create an instance of the config
		// Then set the port correctly
		cfg := config.InitConfig(true)
		port = cfg.Server.Port
	}

	// Handle routes
	http.Handle("/", routes.Handlers())

	// serve
	log.Printf("Server up on port '%s'", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))

}
