package adddate

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	logger "tg_bot_golang/logger"
)

func AddButton(button string) string {
	database, err := sql.Open("sqlite3", "./info.db")
	if err != nil {
		//fmt.Printf("Ошибка подключение в AddButton  %s \n", err)
		errStr := "AddButton() Ошибка при открытии базы данных:" + err.Error() // err.Error() переводим ошибку в строку
		logger.Error.Println(errStr)
	}

	defer database.Close()
	statement, err := database.Prepare("INSERT INTO myDb(button) VALUES (?)") //statement - заявление перевод
	if err != nil {
		//fmt.Printf("Ошибка добавление кнопокв в AddButton в   %s \n", err)
		errStr := "AddButton() Ошибка при добавление данных: " + err.Error() // err.Error() переводим ошибку в строку
		logger.Error.Println(errStr)
	}
	statement.Exec(button)
	errStr := "AddButton() Успешно добавили кнопку проверяй: " + button // err.Error() переводим ошибку в строку
	logger.Info.Println(errStr)
	return "Успешно добавили кнопку проверяй"

}

func AddAdmin(idAdmin, nameAdmin string) string {
	database, err := sql.Open("sqlite3", "./info.db")
	if err != nil {
		fmt.Printf("Ошибка подключение в AddButton  %s \n", err)
		errStr := "AddAdmin() Ошибка открытия базы данных: " + err.Error() // err.Error() переводим ошибку в строку
		logger.Error.Println(errStr)
	}
	defer database.Close()

	statement, err := database.Prepare("INSERT INTO DBAdministrators (idAdmin, nameAdmin) VALUES (?,?)") //statement - заявление перевод
	if err != nil {
		fmt.Printf("Ошибка добавление админа в базу данных   %s \n", err)
		errStr := "AddAdmin() Ошибка добавления данных в базу данных DBAdministrators : " + err.Error() // err.Error() переводим ошибку в строку
		logger.Error.Println(errStr)
	}
	statement.Exec(idAdmin, nameAdmin)
	errStr := "AddAdmin() Успешно добавили кнопку проверяй id админа: " + idAdmin + "Имя админа :" + nameAdmin // err.Error() переводим ошибку в строку
	logger.Info.Println(errStr)

	return "Успешно добавили админа"
}
