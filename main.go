package main

import (
	"log"
	"net/http"
	"os"
	"turskek-workout-app/config"
	"turskek-workout-app/routes"
)

func main() {
	//cfg := config.InitConfig()
	//port := cfg.Server.Port
	port := os.Getenv("PORT")

	if port == "" {
		cfg := config.InitConfig(true)
		port = cfg.Server.Port
	}

	// Handle routes
	http.Handle("/", routes.Handlers())

	// serve
	log.Printf("Server up on port '%s'", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))

}
