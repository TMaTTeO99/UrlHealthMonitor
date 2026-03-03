package main

import (
	"fmt"
	"log"

	"github.com/TMaTTeO99/UrlHealthMonitor/Repository/connection"
)

func main() {

	// Start DB connection
	conn, err := connection.Connect()

	if err != nil {
		log.Fatal("Db connection error: ", err)
	}

	row := connection.CreateUrlTable(conn)
	fmt.Print(row)
}
