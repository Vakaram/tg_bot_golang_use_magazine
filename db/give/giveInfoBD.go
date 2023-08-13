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

func CheckingdbButton(buttonrequest string) bool { // проверяет в базе значения если значения есть тогда мы отпрвляет в ответ yes bool и запускаем функцию которая выполнит запрос и сформирует ответ красивый =)
	database, err := sql.Open("sqlite3", "./info.db")
	var boolyesno []string
	if err != nil {
		fmt.Printf("Ошибка в giveButton %s \n", err)
	}
	fmt.Printf("БАЗА %T \n", database)

	result, errorka := database.Query("SELECT buttun FROM bani WHERE buttun=?", buttonrequest) // пока пойдет и эта часть потом обновим если надобу дет

	var but string
	for result.Next() {
		result.Scan(&but)
		boolyesno = append(boolyesno, but)
	}
	fmt.Printf("%s ВООООТ запрос по БД прошлись вижу вот", boolyesno)
	var sumLenElemen int
	for i := 0; i < len(boolyesno); i++ { // Проверяет кол во ленов в массиве если больше 1 вернем true
		sumLenElemen += 1
	}
	if sumLenElemen >= 1 {
		return true
	}

	fmt.Printf("\n Результат запроса такой %s. А вот ошибка такая  %s ", result, errorka)

	defer database.Close()
	return false // если true выше не отработает то вернуть нилл

}
