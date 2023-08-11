package createdb

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

// create db только создание
func СreateTable() {
	database, _ := sql.Open("sqlite3", "./info.db")
	defer database.Close()
	statematn, _ := database.Prepare("CREATE TABLE IF NOT EXISTS bani(id INTEGER PRIMARY KEY ,buttun TEXT, description TEXT, photo TEXT)")
	statematn.Exec()
}
