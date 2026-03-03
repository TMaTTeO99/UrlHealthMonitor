package connection

/*
	Package to manage the db standar operations like:
	1. Start Conncetion
	2. End Connection

*/

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

func Connect() (*pgx.Conn, error) {

	prefix := os.Getenv("PREFIX")
	user := os.Getenv("USER_NAME")
	pass := os.Getenv("USER_PASSWD")
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	dbName := os.Getenv("DB_NAME")

	addr := fmt.Sprintf("%s://%s:%s@%s:%s/%s", prefix, user, pass, host, port, dbName)
	return pgx.Connect(context.Background(), addr)
}

func CreateUrlTable(conn *pgx.Conn) pgx.Row {

	row, err := conn.Query(context.Background(), "CREATE TABLE urls"+
		"(id Integer PRIMARY KEY, url VARCHAR(255))")

	if err != nil {
		log.Fatal("ERROR: ", err)
	}

	return row
}
