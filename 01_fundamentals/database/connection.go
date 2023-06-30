package database

import "fmt"

var conn string

func init() {
	fmt.Println("Init starting")
	conn = "SQL-Server"
}

func GetConnection() string {
	return conn
}