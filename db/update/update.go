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

	statement, err := database.Prepare("UPDATE myDb SET description=newDescription WHERE button = button VALUES (?,?) ") //statement - заявление перевод
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

	result, errorka := database.Exec("UPDATE myDb SET button=? WHERE button=?", newButton, whichButton) // пока пойдет и эта часть потом обновим если надобу дет
	if errorka == nil {
		fmt.Printf("\n Результат запроса такой %s. А вот ошибка такая  %s ", result, errorka)

	}

	return "Обновили кнопку: " + whichButton + " На: " + newButton

}

func UpdateDescriptionButton(button string, description string) string { //Добавить возможность добовлять описание кнопок
	database, err := sql.Open("sqlite3", "./info.db")
	if err != nil {
		fmt.Printf("Ошибка подключение в AddDescriptionButton  %s \n", err)
	}
	defer database.Close()
	if err != nil {
		fmt.Printf("Ошибка добавление кнопокв в AddDescriptionButton в   %s \n", err)
	}
	result, errorka := database.Exec("UPDATE myDb SET description=? WHERE button=?", description, button) // пока пойдет и эта часть потом обновим если надобу дет
	fmt.Printf("\n Результат запроса такой %s. А вот ошибка такая  %s ", result, errorka)
	return "Успешно добавили описание к кнопке проверяй"
}

func UpdateAddPhotoInButton(button string, stringPhotoId string) string { // добоавляет по одной фотке к тому что есть + забирает все фотки которые были и добавляет.
	database, err := sql.Open("sqlite3", "./info.db")
	if err != nil {
		fmt.Printf("Ошибка подключение в AddButton  %s \n", err)
	}
	defer database.Close()
	// TODO взять инфо из базы и добавить к ней еще одно фото =)
	result, _ := database.Query("SELECT photo FROM myDb WHERE button=?", button) // пока пойдет и эта часть потом обновим если надобу дет
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
			result, errorka := database.Exec("UPDATE myDb SET photo=? WHERE button=?", stringPhotoId, button) // пока пойдет и эта часть потом обновим если надобу дет
			fmt.Printf("\n Результат запроса такой при добавление кнопки %s. А вот ошибка такая  %s ", result, errorka)
			fmt.Printf("Отправим данные старые + новые : %s", stringPhotoId)
			return "Добавил еще одну фотографию к тем которые были"
		} else { // проверка если запятой в конце нет по какой то причине то мы возьмем да и допишем эту запятую в конец встроки
			stringPhotoId = str + "," + stringPhotoId + ","
			result, errorka := database.Exec("UPDATE myDb SET photo=? WHERE button=?", stringPhotoId, button) // пока пойдет и эта часть потом обновим если надобу дет
			fmt.Printf("\n Результат запроса такой при добавление кнопки %s. А вот ошибка такая  %s ", result, errorka)
			fmt.Printf("Отправим данные старые + новые : %s", stringPhotoId)
			return "Добавил еще одну фотографию к тем которые были"
		}
	} else {
		stringPhotoId = stringPhotoId + ","
		database.Exec("UPDATE myDb SET photo=? WHERE button=?", stringPhotoId, button) // пока пойдет и эта часть потом обновим если надобу дет
		//result, errorka := эта часть для строки выше нужна
		//fmt.Printf("\n Результат запроса такой при добавление кнопки %s. А вот ошибка такая  %s ", result, errorka)
		return "Добавил фото,фото раньше не было"
	}

}

func UpdateDbGreetings(UpdateDbGreetings string) string { //Добавить возможность добовлять описание кнопок
	database, err := sql.Open("sqlite3", "./info.db")
	if err != nil {
		fmt.Printf("Ошибка подключение в AddDescriptionButton  %s \n", err)
	}
	defer database.Close()
	if err != nil {
		fmt.Printf("Ошибка добавление кнопокв в AddDescriptionButton в   %s \n", err)
	}
	firsid := 1
	result, errorka := database.Exec("UPDATE DbGreetings SET Greetings=? WHERE id=?", UpdateDbGreetings, firsid) // пока пойдет и эта часть потом обновим если надобу дет
	fmt.Printf("\n Результат запроса такой %s. А вот ошибка такая  %s ", result, errorka)
	return "Успешно добавили приветствие проверить можно отправив команду /start в роли пользователя а не админа"
}
