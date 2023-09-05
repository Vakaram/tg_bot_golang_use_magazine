package drop

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	dbgive "tg_bot_golang/db/give"
	logger "tg_bot_golang/logger"
)

func DropAllTablemyDb() {
	database, err := sql.Open("sqlite3", "./info.db")
	if err != nil {
		errStr := "DropAllTablemyDb() Ошибка открытия базы данных: " + err.Error() // err.Error() переводим ошибку в строку
		logger.Error.Println(errStr)
	}
	defer database.Close()
	statematn, err := database.Prepare("DELETE FROM myDb") // добавляет значения
	if err != nil {
		errStr := "DropAllTablemyDb() Ошибка при запросе удаления базы данных: " + err.Error() // err.Error() переводим ошибку в строку
		logger.Error.Println(errStr)
	}
	logger.Info.Println("DropAllTablemyDb() Удаление таблицы myDB")
	statematn.Exec()

}

func DropOneButton(button string) string { // удалить однку кнопку
	database, err := sql.Open("sqlite3", "./info.db")
	if err != nil {
		errStr := "DropOneButton() Ошибка открытия базы данных: " + err.Error() // err.Error() переводим ошибку в строку
		logger.Error.Println(errStr)
	}
	defer database.Close()
	statematn, err := database.Prepare("DELETE FROM myDb WHERE `button` = ?") // добавляет значения
	if err != nil {
		errStr := "DropOneButton() Ошибка при запросе удаления одной кнопки: " + err.Error() // err.Error() переводим ошибку в строку
		logger.Error.Println(errStr)
	}
	statematn.Exec(button)
	itog := "Удалил кнопку: " + button
	logger.Info.Println("DropOneButton() Успешно удалили кнопку")
	return itog
}

func DropAllPhotoButton(button string) string { // удалить все фотки
	database, err := sql.Open("sqlite3", "./info.db")
	if err != nil {
		errStr := "DropAllPhotoButton() Ошибка при открытие базы : " + err.Error() // err.Error() переводим ошибку в строку
		logger.Error.Println(errStr)
	}
	defer database.Close()
	stringDropPhoto := ""
	_, err = database.Exec("UPDATE myDb SET photo=? WHERE button=?", stringDropPhoto, button) // пока пойдет и эта часть потом обновим если надобу дет
	itog := "Удалил фотки все для кнопки: " + button
	if err != nil {
		errStr := "DropAllPhotoButton() Ошибка при запросе удаления всех фоток: " + err.Error() // err.Error() переводим ошибку в строку
		logger.Error.Println(errStr)
	}
	logger.Info.Println("DropAllPhotoButton() Успешно удалили все кнопки ")

	return itog

}

func DropOneAdmin(idAdmin string) string { //
	if idAdmin == "1" {
		return "Нельзя удалить главного админа "
	}
	admins := dbgive.GiveDBAdministratorsChecID()

	for _, value := range admins {
		if value == idAdmin {
			database, err := sql.Open("sqlite3", "./info.db")
			if err != nil {
				errStr := "DropOneAdmin() Ошибка при открытие базы : " + err.Error() // err.Error() переводим ошибку в строку
				logger.Error.Println(errStr)
			}
			defer database.Close()

			// Создаем SQL-запрос DELETE
			stmt, err := database.Prepare("DELETE FROM DBAdministrators WHERE id = ?")
			if err != nil {
				errStr := "DropOneAdmin() Ошибка при удаление админа пришел id для удаления  : " + idAdmin + err.Error() // err.Error() переводим ошибку в строку
				logger.Error.Println(errStr)
			}
			defer stmt.Close()

			// Выполняем запрос DELETE с передачей идентификатора
			_, err = stmt.Exec(idAdmin)
			if err != nil {
				errStr := "DropOneAdmin() Ошибка при удаление админа пришел id для удаления zzz  : " + idAdmin + err.Error() // err.Error() переводим ошибку в строку
				logger.Error.Println(errStr)
			}
			logger.Info.Println("DropOneAdmin() Успешно удалили администратора с ID= " + idAdmin)
			return "Успешно удалили Администратора с идентификатором " + idAdmin
		}
	}
	logger.Info.Println("DropOneAdmin() Администратора с таким id нет в базе данных ID= " + idAdmin)
	return "Администратора с таким id нет в базе данных"
}
