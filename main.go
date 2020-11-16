package main

//https://medium.com/@hugo.bjarred/rest-api-with-golang-mux-mysql-c5915347fa5b
import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func main() {
	db, err = sql.Open("mysql", "<user>:<password>@tcp(127.0.0.1:3306)/<dbname>")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
}
