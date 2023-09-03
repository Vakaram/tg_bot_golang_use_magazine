package adddate

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

func AddButton(button string) string {
	database, err := sql.Open("sqlite3", "./info.db")
	if err != nil {
		fmt.Printf("Ошибка подключение в AddButton  %s \n", err)
	}

	defer database.Close()
	statement, err := database.Prepare("INSERT INTO myDb(button) VALUES (?)") //statement - заявление перевод
	if err != nil {
		fmt.Printf("Ошибка добавление кнопокв в AddButton в   %s \n", err)
	}
	statement.Exec(button)

	return "Успешно добавили кнопку проверяй"
}

func AddAdmin(idAdmin, nameAdmin string) string {
	database, err := sql.Open("sqlite3", "./info.db")
	if err != nil {
		fmt.Printf("Ошибка подключение в AddButton  %s \n", err)
	}
	defer database.Close()

	statement, err := database.Prepare("INSERT INTO DBAdministrators (idAdmin, nameAdmin) VALUES (?,?)") //statement - заявление перевод
	if err != nil {
		fmt.Printf("Ошибка добавление админа в базу данных   %s \n", err)
	}
	statement.Exec(idAdmin, nameAdmin)

	return "Успешно добавили админа"
}
