package update

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

func AddDescriptionInButton(button string, newDescription string) string {
	database, err := sql.Open("sqlite3", "./info.db")
	//if err != nil {
	//	fmt.Printf("Ошибка подключение в AddButton  %s \n", err)
	//}
	defer database.Close()

	statement, err := database.Prepare("UPDATE bani SET description=newDescription WHERE buttun = button VALUES (?,?) ") //statement - заявление перевод
	if err != nil {
		fmt.Printf("Ошибка добавление кнопокв в AddButton в   %s \n", err)
	}
	statement.Exec(newDescription, button)

	return "Добавил описание "

}
