package main

import (
	"context"
	"log"

	"github.com/TMaTTeO99/UrlHealthMonitor/API/controllers"
	"github.com/TMaTTeO99/UrlHealthMonitor/Repository/connection"
	"github.com/TMaTTeO99/UrlHealthMonitor/config"
)

func main() {

	// Load configuration variable
	config, err := config.LoadConfig()

	if err != nil {
		log.Fatal("Environment variable error: ", err)
	}

	// Start DB connection
	conn, err := connection.Connect(config)

	// Check DB connection
	if err != nil {
		log.Fatal("Db connection error: ", err)
	}

	// Create the urls table
	connection.CreateUrlTable(conn)

	// Start web server
	controllers.StartWebServer(config, conn)

	// Close DB connection at the program's end
	defer conn.Close(context.Background())
}
