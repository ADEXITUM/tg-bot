package storage

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func mysqlConn(dbName string) *sql.DB {
	cfg := mysql.Config{
		User:   "root",
		Passwd: "root",
		Net:    "tcp",
		Addr:   "192.168.56.1:3306",
		DBName: "bot",
	}
	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
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

//TODO: 1) Complete mysqlConn()
//		2) Force deleting of previous database entry for each userID
func InsertUserInfo(chatID int, username string) {
	database := mysqlConn("bot")

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
