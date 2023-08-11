package createdb

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

// create db только создание  // *sql.DB
func СreateTable() {
	database, err := sql.Open("sqlite3", "./info.db")
	if err != nil {
		fmt.Printf("Ошибка в create bd %s \n", err)
	}
	defer database.Close()
	statematn, err := database.Prepare("CREATE TABLE IF NOT EXISTS bani(id INTEGER PRIMARY KEY ,buttun TEXT, description TEXT, photo TEXT)")
	if err != nil {
		fmt.Printf("Ошибка при создание базы запросом    %s \n", err)
	}
	statematn.Exec()
}
