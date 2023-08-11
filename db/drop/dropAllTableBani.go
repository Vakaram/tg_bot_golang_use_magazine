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

func DropOneButton(button string) string {
	database, _ := sql.Open("sqlite3", "./info.db")
	defer database.Close()
	statematn, err := database.Prepare("DELETE FROM bani WHERE `buttun` = ?") // добавляет значения
	statematn.Exec(button)
	itog := "Удалил кнопку: " + button

	if err != nil {
		return err.Error()

	}
	return itog
}
