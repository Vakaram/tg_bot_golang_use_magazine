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

func UpdateButton(whichButton string, newButton string) string { // будет обновлять только однку кнопку не трогая описание
	database, err := sql.Open("sqlite3", "./info.db")

	//TODO а так же сделать проверку на кнопку есть ли такая кнопка в базе и если что возвращать ошибку что нет такой кнопки
	if err != nil {
		fmt.Printf("Ошибка подключение в UpdateButton  %s \n", err)
	}
	defer database.Close()

	if err != nil {
		fmt.Printf("Ошибка добавление кнопокв в UpdateButton в   %s \n", err)
	}

	result, errorka := database.Exec("UPDATE bani SET buttun=? WHERE buttun=?", newButton, whichButton) // пока пойдет и эта часть потом обновим если надобу дет
	fmt.Printf("\n Результат запроса такой %s. А вот ошибка такая  %s ", result, errorka)
	//statement, err := database.Prepare("UPDATE bani SET buttun=? WHERE buttun=?", newButton, whichButton) //statement - заявление перевод

	//_, err = tx.NamedExec( // скинули пример запроса сказали топчик но не пойму как работает
	//	`UPDATE  bani SET buttun=:newButton WHERE buttun=:whichButton`,
	//	map[string]interface{}{
	//		"newButton":   newButton,
	//		"whichButton": whichButton,
	//	},
	//)
	//
	//statement, err := database.Prepare("UPDATE bani SET buttun=newButton WHERE buttun=whichButton  VALUES (?, ?)") //statement - заявление перевод
	//if err != nil {
	//	fmt.Printf("Ошибка добавление кнопокв в UpdateButton в   %s \n", err)
	//}
	//
	//result, errorka := statement.Exec(newButton, whichButton)

	//fmt.Printf("\n Результат запроса такой %s. А вот ошибка такая  %s ", result, errorka)

	//statement, err := database.Prepare("UPDATE bani SET buttun='sad' WHERE buttun='Баня4'") //statement - заявление перевод
	//if err != nil {
	//	fmt.Printf("Ошибка добавление кнопокв в UpdateButton в   %s \n", err)
	//}
	//
	//result, errorka := statement.Exec()
	//fmt.Printf("\n Результат запроса такой %s. А вот ошибка такая  %s ", result, errorka)

	return "Обновили кнопку: " + whichButton + " На: " + newButton

}
