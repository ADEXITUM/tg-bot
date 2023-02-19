package types

import "github.com/go-sql-driver/mysql"

var CFG = mysql.Config{
	User:   "root",
	Passwd: "root",
	Net:    "tcp",
	Addr:   "192.168.56.1:3306",
	DBName: "bot",
}
