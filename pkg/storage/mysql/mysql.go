package mysql

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func mysqlConn(dbName string) *sql.DB {
	return nil
}

//TODO: 1) Complete mysqlConn()
//		2) Force deleting of previous database entry for each userID
func InsertUserInfo(chatID int, username string) {
	database := mysqlConn("bot")

	defer database.Close()

	time := time.Now().Unix()
	result, err := database.Prepare("INSERT INTO user_info (chat_id, username, timestamp) VALUES (?, ?, ?)")
	if err != nil {
		fmt.Println("Error while prepating sql request: ", err)
	}
	_, err := result.Exec(chatID, username, times)
	if err != nil {
		fmt.Println("error while creating sql request: ", err)
	}
}
