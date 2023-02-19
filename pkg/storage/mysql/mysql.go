package storage

import (
	"database/sql"
	"fmt"
	"log"
	"tg-bot/pkg/telegram/types"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func mysqlConn(dbName string) *sql.DB {
	db, err := sql.Open("mysql", types.CFG.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

	return db
}

func InsertUserInfo(chatID int, username string) {
	database := mysqlConn(types.CFG.DBName)

	defer database.Close()

	time := time.Now().Unix()

	result, err := database.Prepare("INSERT IGNORE INTO bot (chat_id, username, timestamp) VALUES (?, ?, ?)")
	if err != nil {
		fmt.Println("Error while prepating sql request: ", err)
	}
	_, err = result.Exec(chatID, username, time)
	if err != nil {
		fmt.Println("error while creating sql request: ", err)
	}
}
