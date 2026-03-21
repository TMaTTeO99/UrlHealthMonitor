package connection

/*
	Package to manage the db standar operations like:
	1. Start Conncetion
	2. Create Urls table
	3. Check if the Urls table exists
	4. Add url inside DB
*/

import (
	"context"
	"fmt"
	"log"

	"github.com/TMaTTeO99/UrlHealthMonitor/API/models"
	"github.com/TMaTTeO99/UrlHealthMonitor/config"
	"github.com/jackc/pgx/v5"
)

const EXISTS_URLS_TABLE_QUERY = "SELECT EXISTS (SELECT FROM information_schema.tables WHERE table_schema = 'public' AND table_name = $1)"
const CREATE_URLS_TABLE_QUERY = "CREATE TABLE urls (id Integer PRIMARY KEY GENERATED ALWAYS AS IDENTITY, url VARCHAR(255), user_id INTEGER)"
const ADD_URL_QUERY = "INSERT INTO urls (url, user_id) VALUES ($1, $2)"
const GET_URLS_BY_USER_ID = "SELECT * FROM urls WHERE urls.user_id = $1"

func Connect(config *config.ConfigData) (*pgx.Conn, error) {

	addr := fmt.Sprintf("%s://%s:%s@%s:%s/%s", config.PREFIX, config.DB_USER_NAME, config.USER_PASSWD, config.DB_HOST, config.DB_PORT, config.DB_NAME)
	return pgx.Connect(context.Background(), addr)
}

func CreateUrlTable(conn *pgx.Conn) {

	if UrlsTableExists(conn) {
		return
	}

	row, err := conn.Query(context.Background(), CREATE_URLS_TABLE_QUERY)

	if err != nil {
		log.Fatal("ERROR: ", err)
	}

	defer row.Close()
}

func UrlsTableExists(conn *pgx.Conn) bool {

	var exists bool
	conn.QueryRow(context.Background(), EXISTS_URLS_TABLE_QUERY, "urls").Scan(&exists)

	return exists
}

func InsertUrl(conn *pgx.Conn, url string) error {

	_, err := conn.Exec(context.Background(), ADD_URL_QUERY, url, 11111)
	return err

}

func GetAllUlr(conn *pgx.Conn, id int) ([]models.UrlDataDTO, error) {

	row, err := conn.Query(context.Background(), GET_URLS_BY_USER_ID, id)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer row.Close()

	var urls []models.UrlDataDTO
	for row.Next() {

		var url string
		var userId, id int

		if err := row.Scan(&id, &url, &userId); err != nil {
			urls = append(urls, models.UrlDataDTO{
				Id:          id,
				UserId:      userId,
				Title:       "",
				Url:         url,
				Description: "",
				Image:       "",
			})
		}
	}

	return urls, nil

}
