package drop

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	dbgive "tg_bot_golang/db/give"
)

func DropAllTablemyDb() {
	database, _ := sql.Open("sqlite3", "./info.db")
	defer database.Close()
	statematn, _ := database.Prepare("DELETE FROM myDb") // добавляет значения
	statematn.Exec()
}

func DropOneButton(button string) string {
	database, _ := sql.Open("sqlite3", "./info.db")
	defer database.Close()
	statematn, err := database.Prepare("DELETE FROM myDb WHERE `button` = ?") // добавляет значения
	statematn.Exec(button)
	itog := "Удалил кнопку: " + button

	if err != nil {
		return err.Error()

	}
	return itog
}

//Todo сделать удаление всех фото из одной кнопки

func DropAllPhotoButton(button string) string {
	database, _ := sql.Open("sqlite3", "./info.db")
	defer database.Close()
	stringDropPhoto := ""
	_, err := database.Exec("UPDATE myDb SET photo=? WHERE button=?", stringDropPhoto, button) // пока пойдет и эта часть потом обновим если надобу дет
	itog := "Удалил фотки все для кнопки: " + button
	if err != nil {
		return err.Error()
	}
	return itog

}

func DropOneAdmin(idAdmin string) string { // TODO здесь хорошобы сделать проверку а есть ли такой админ в базе данных ?
	if idAdmin == "1" {
		return "Нельзя удалить главного админа "
	}
	admins := dbgive.GiveDBAdministratorsChecID()

	for _, value := range admins {
		if value == idAdmin {
			database, _ := sql.Open("sqlite3", "./info.db")
			defer database.Close()

			// Создаем SQL-запрос DELETE
			stmt, err := database.Prepare("DELETE FROM DBAdministrators WHERE id = ?")
			if err != nil {
				log.Fatal(err)
			}
			defer stmt.Close()

			// Выполняем запрос DELETE с передачей идентификатора
			_, err = stmt.Exec(idAdmin)
			if err != nil {
				log.Fatal(err)
			}
			return "Успешно удалили Администратора с идентификатором " + idAdmin

		}

	}
	return "Администратора с таким id нет в базе данных"

}
