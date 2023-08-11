package drop

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func DropAllTableBani() {
	database, _ := sql.Open("sqlite3", "./info.db")
	defer database.Close()
	statematn, _ := database.Prepare("DELETE FROM bani") // добавляет значения
	statematn.Exec()
}
