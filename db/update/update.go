package update

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	logger "tg_bot_golang/logger"
)

func AddDescriptionInButton(button string, newDescription string) string {
	database, err := sql.Open("sqlite3", "./info.db")
	if err != nil {
		errStr := "AddDescriptionInButton() Ошибка при подключение к бд: " + err.Error() // err.Error() переводим ошибку в строку
		logger.Error.Println(errStr)
	}
	defer database.Close()

	statement, err := database.Prepare("UPDATE myDb SET description=newDescription WHERE button = button VALUES (?,?) ") //statement - заявление перевод
	if err != nil {
		errStr := "AddDescriptionInButton() Ошибка Update описания для кнопки: " + err.Error() // err.Error() переводим ошибку в строку
		logger.Error.Println(errStr)
	}
	statement.Exec(newDescription, button)
	logger.Info.Println("AddDescriptionInButton() Успешно обновили описание к кнопке")
	return "Добавил описание "

}

func UpdateButton(whichButton string, newButton string) string { // будет обновлять только однку кнопку не трогая описание
	database, err := sql.Open("sqlite3", "./info.db")
	if err != nil {
		errStr := "UpdateButton() Ошибка при подключение к бд: " + err.Error() // err.Error() переводим ошибку в строку
		logger.Error.Println(errStr)
	}

	defer database.Close()

	_, err = database.Exec("UPDATE myDb SET button=? WHERE button=?", newButton, whichButton) // пока пойдет и эта часть потом обновим если надобу дет
	if err != nil {
		errStr := "UpdateButton() Ошибка Update описания для кнопки: " + err.Error() // err.Error() переводим ошибку в строку
		logger.Error.Println(errStr)
	}
	logger.Info.Println("UpdateButton() Успешно обновили название к кнопке")

	return "Обновили кнопку: " + whichButton + " На: " + newButton

}

func UpdateDescriptionButton(button string, description string) string { //Добавить возможность добовлять описание кнопок
	database, err := sql.Open("sqlite3", "./info.db")
	if err != nil {
		errStr := "UpdateDescriptionButton() Ошибка при подключение к бд: " + err.Error() // err.Error() переводим ошибку в строку
		logger.Error.Println(errStr)
	}
	defer database.Close()

	_, err = database.Exec("UPDATE myDb SET description=? WHERE button=?", description, button) // пока пойдет и эта часть потом обновим если надобу дет
	if err != nil {
		errStr := "UpdateDescriptionButton() Ошибка при Update описанияк кнопке: " + err.Error() // err.Error() переводим ошибку в строку
		logger.Error.Println(errStr)
	}
	logger.Info.Println("UpdateDescriptionButton() Успешно обновили название к кнопке")

	return "Успешно добавили описание к кнопке проверяй"
}

func UpdateAddPhotoInButton(button string, stringPhotoId string) string { // добоавляет по одной фотке к тому что есть + забирает все фотки которые были и добавляет.
	database, err := sql.Open("sqlite3", "./info.db")
	if err != nil {
		errStr := "UpdateAddPhotoInButton() Ошибка при подключение к бд: " + err.Error() // err.Error() переводим ошибку в строку
		logger.Error.Println(errStr)
	}
	defer database.Close()
	result, err := database.Query("SELECT photo FROM myDb WHERE button=?", button) // пока пойдет и эта часть потом обновим если надобу дет
	if err != nil {
		errStr := "UpdateAddPhotoInButton() Ошибка при запросе фотографии из базы данных: " + err.Error() // err.Error() переводим ошибку в строку
		logger.Error.Println(errStr)
	}
	var photoitog string
	var photoval string
	for result.Next() {
		result.Scan(&photoval)
		//fmt.Printf("В цикле бегаю смотрю данные: %s", result.Scan(photoval))
		photoitog += photoval
	}
	str := photoitog
	//fmt.Printf("Перед обновление получил данные по фоткам вот они : %s", str)
	if len(str) > 0 {
		if str[len(str)-1] == ',' {
			stringPhotoId = str + stringPhotoId + ","
			_, err = database.Exec("UPDATE myDb SET photo=? WHERE button=?", stringPhotoId, button) // пока пойдет и эта часть потом обновим если надобу дет
			if err != nil {
				errStr := "UpdateAddPhotoInButton() Ошибка Update фоток к кнопке: " + err.Error() // err.Error() переводим ошибку в строку
				logger.Error.Println(errStr)
			}
			logger.Info.Println("UpdateAddPhotoInButton() Успешно добавили фотографию к тем что были")

			return "Добавил еще одну фотографию к тем которые были"
		} else { // проверка если запятой в конце нет по какой то причине то мы возьмем да и допишем эту запятую в конец встроки
			stringPhotoId = str + "," + stringPhotoId + ","
			_, err = database.Exec("UPDATE myDb SET photo=? WHERE button=?", stringPhotoId, button) // пока пойдет и эта часть потом обновим если надобу дет
			if err != nil {
				errStr := "UpdateAddPhotoInButton() Ошибка Update фоток к кнопке: " + err.Error() // err.Error() переводим ошибку в строку
				logger.Error.Println(errStr)
			}
			logger.Info.Println("UpdateAddPhotoInButton() Успешно Добавил еще одну фотографию к тем которые были")

			return "Добавил еще одну фотографию к тем которые были"
		}
	} else {
		stringPhotoId = stringPhotoId + ","
		_, err = database.Exec("UPDATE myDb SET photo=? WHERE button=?", stringPhotoId, button) // пока пойдет и эта часть потом обновим если надобу дет
		if err != nil {
			errStr := "UpdateAddPhotoInButton() Ошибка Update фоток к кнопке: " + err.Error() // err.Error() переводим ошибку в строку
			logger.Error.Println(errStr)
		}
		logger.Info.Println("UpdateAddPhotoInButton() Успешно Добавил фото,фото раньше не было")
		return "Добавил фото,фото раньше не было"
	}

}

func UpdateDbGreetings(UpdateDbGreetings string) string { //Добавить возможность добовлять описание кнопок
	database, err := sql.Open("sqlite3", "./info.db")
	if err != nil {
		errStr := "UpdateDbGreetings() Ошибка поключение к бд: " + err.Error() // err.Error() переводим ошибку в строку
		logger.Error.Println(errStr)
	}
	defer database.Close()

	firsid := 1
	_, err = database.Exec("UPDATE DbGreetings SET Greetings=? WHERE id=?", UpdateDbGreetings, firsid) // пока пойдет и эта часть потом обновим если надобу дет
	if err != nil {
		errStr := "UpdateDbGreetings() апдейта описания приветствия в таблице DbGreetings: " + err.Error() // err.Error() переводим ошибку в строку
		logger.Error.Println(errStr)
	}
	logger.Info.Println("Успешно добавили приветствие проверить можно отправив команду /start в роли пользователя а не админа")
	return "Успешно добавили приветствие проверить можно отправив команду /start в роли пользователя а не админа"
}
