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
	statement, err := database.Prepare("INSERT INTO bani(buttun) VALUES (?)") //statement - заявление перевод
	if err != nil {
		fmt.Printf("Ошибка добавление кнопокв в AddButton в   %s \n", err)
	}
	statement.Exec(button)

	return "Успешно добавили кнопку проверяй"
}

//TODO Добавить возможность добовлять описание кнопок

func AddDicriptionButton(button string) string { //Добавить возможность добовлять описание кнопок
	database, err := sql.Open("sqlite3", "./info.db")

	if err != nil {
		fmt.Printf("Ошибка подключение в UpdateButton  %s \n", err)
	}
	defer database.Close()

	if err != nil {
		fmt.Printf("Ошибка добавление кнопокв в UpdateButton в   %s \n", err)
	}

	result, errorka := database.Exec("UPDATE bani SET buttun=? WHERE buttun=?", newButton, whichButton) // пока пойдет и эта часть потом обновим если надобу дет

	return "Успешно добавили описание к кнопке проверяй"

}

//TODO Добавить возможность добавлять фоток
