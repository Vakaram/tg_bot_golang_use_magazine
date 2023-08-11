package give

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

func GiveButtonInBase() []string {
	var givButton []string // функция для получения кнопок которые отправятся в бд
	database, err := sql.Open("sqlite3", "./info.db")
	if err != nil {
		fmt.Printf("Ошибка в giveButton %s \n", err)
	}
	fmt.Printf("БАЗА %T \n", database)

	rows, err := database.Query("SELECT buttun FROM bani") // запрос делаем
	if err != nil {
		fmt.Printf("Ошибка при получение кнопок из базы give button   %s \n", err)
	}
	defer database.Close()

	var but string
	for rows.Next() {
		rows.Scan(&but)
		givButton = append(givButton, but)
		//fmt.Printf("%d: %s ,%s \n", id, but, phot)
	}

	fmt.Printf("Такой массив кнопок получаю сейчас %s \n", givButton)
	return givButton
}
