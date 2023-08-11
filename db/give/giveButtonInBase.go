package give

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func GiveButtonInBase() []string {
	var givButton []string // функция для получения кнопок которые отправятся в бд

	database, _ := sql.Open("sqlite3", "./info.db")
	rows, _ := database.Query("SELECT buttun FROM bani") // запрос делаем
	defer database.Close()

	var but string
	for rows.Next() {
		rows.Scan(&but)
		givButton = append(givButton, but)
		//fmt.Printf("%d: %s ,%s \n", id, but, phot)
	}
	return givButton
}
