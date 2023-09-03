package give

import (
	"database/sql"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	_ "github.com/mattn/go-sqlite3"
	"strconv"
	"strings"
)

func GiveButtonInBase() []string {
	var givButton []string // функция для получения кнопок которые отправятся в бд
	database, err := sql.Open("sqlite3", "./info.db")
	if err != nil {
		fmt.Printf("Ошибка в giveButton %s \n", err)
	}
	fmt.Printf("БАЗА %T \n", database)

	rows, err := database.Query("SELECT button FROM myDb") // запрос делаем
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

	return givButton
}

func CheckingdbButton(buttonrequest string) bool { // проверяет в базе значения если значения есть тогда мы отпрвляет в ответ yes bool и запускаем функцию которая выполнит запрос и сформирует ответ красивый =)
	database, err := sql.Open("sqlite3", "./info.db")
	var boolyesno []string
	if err != nil {
		fmt.Printf("Ошибка в giveButton %s \n", err)
	}
	fmt.Printf("БАЗА %T \n", database)

	result, errorka := database.Query("SELECT button FROM myDb WHERE button=?", buttonrequest) // пока пойдет и эта часть потом обновим если надобу дет

	var but string
	for result.Next() {
		result.Scan(&but)
		boolyesno = append(boolyesno, but)
	}
	//fmt.Printf("%s ВООООТ запрос по БД прошлись вижу вот", boolyesno)
	var sumLenElemen int
	for i := 0; i < len(boolyesno); i++ { // Проверяет кол во ленов в массиве если больше 1 вернем true
		sumLenElemen += 1
	}
	if sumLenElemen >= 1 {
		return true
	}

	if errorka != nil {
		fmt.Printf("\n А вот ошибка такая  %s ", errorka)
	}

	defer database.Close()
	return false // если true выше не отработает то вернуть нилл

}

//TODO Нужно сделать запрос который будет возвращать и формировать ответ по кнопке если она есть в бд Сформировать просто ответ крутой  discription and photo
//TODO и надо делать проверку на ошибку если фоток нет чтобы прога не падала
//TODO теперь сделать проверку на фото и вообще сначало фотки распарсить и возвратить хммм

func GiveDescriptionButton(buttonsearch string) string { // пока будем получать просто описание
	database, err := sql.Open("sqlite3", "./info.db")
	defer database.Close()

	var descriptionitog []string
	if err != nil {
		fmt.Printf("Ошибка в giveButton %s \n", err)
	}
	//fmt.Printf("БАЗА %T \n", database)

	result, errorka := database.Query("SELECT description FROM myDb WHERE button=?", buttonsearch) // пока пойдет и эта часть потом обновим если надобу дет

	var decriptionval string
	for result.Next() {
		result.Scan(&decriptionval)
		descriptionitog = append(descriptionitog, decriptionval)
	}
	stringitog := descriptionitog[0]
	//
	//fmt.Printf("%s ВООООТ запрос по БД прошлись вижу вот", descriptionitog)
	//
	if errorka != nil {
		//fmt.Printf("\n Результат запроса такой %s. А вот ошибка такая  %s ", result, errorka)

	}

	return stringitog // если true выше не отработает то вернуть нилл
}

func GivePhotoButton(buttonsearch string, addCaptionInPhoto string) []interface{} { // достает все фотки из кнопки и передает их в групповое сообщение
	database, err := sql.Open("sqlite3", "./info.db")
	defer database.Close()

	var photoitog []string
	if err != nil {
		fmt.Printf("Ошибка в giveButton %s \n", err)
	}
	fmt.Printf("БАЗА %T \n", database)

	result, _ := database.Query("SELECT photo FROM myDb WHERE button=?", buttonsearch) // пока пойдет и эта часть потом обновим если надобу дет

	var photoval string
	for result.Next() {
		result.Scan(&photoval)
		photoitog = append(photoitog, photoval)
	}
	str := photoitog[0]
	var arrInterface []interface{}
	var colvoCicle int
	if len(str) > 0 && str[len(str)-1] == ',' { // сделаем проверку на то что последний символ запятая и есть ли там вообще данные
		str = str[:len(str)-1] // удаляем последнию запятую из фоток =_) тк выпадает ошибка и фотки не присылаются
		//fmt.Printf("Вижу такое описание фоток : %s", str)
		words := strings.Split(str, ",")
		for _, word := range words {
			if colvoCicle == 0 {
				zeroPhoto := tgbotapi.NewInputMediaPhoto(tgbotapi.FileID(word))
				zeroPhoto.Caption = addCaptionInPhoto
				arrInterface = append(arrInterface, zeroPhoto)
			} else {
				arrInterface = append(arrInterface, tgbotapi.NewInputMediaPhoto(tgbotapi.FileID(word)))
			}
			colvoCicle++ // считаем чтобы если первый проход тогда мы добавим описание к нулевой кнопке
		}
		//fmt.Print(arrInterface)
		return arrInterface
	} else {
		words := strings.Split(str, ",")
		for _, word := range words {
			if colvoCicle == 0 {
				zeroPhoto := tgbotapi.NewInputMediaPhoto(tgbotapi.FileID(word))
				zeroPhoto.Caption = addCaptionInPhoto
				arrInterface = append(arrInterface, zeroPhoto)
			} else {
				arrInterface = append(arrInterface, tgbotapi.NewInputMediaPhoto(tgbotapi.FileID(word)))
			}
			colvoCicle++ // считаем чтобы если первый проход тогда мы добавим описание к нулевой кнопке

		}
		fmt.Print(arrInterface)
		return arrInterface
	}
	return arrInterface

}

func GiveDescriptionDBGreetings() (string, bool) {
	database, err := sql.Open("sqlite3", "./info.db")
	defer database.Close()

	var descriptionitog []string
	if err != nil {
		fmt.Printf("Ошибка в giveButton %s \n", err)
	}

	firsdescrirption := 1                                                                           // ищем всегда по 1 id                                                                             // берем инфу из первого id и его же будем обновлять всегддддааа=  )
	result, err := database.Query("SELECT Greetings FROM DbGreetings WHERE id=?", firsdescrirption) // пока пойдет и эта часть потом обновим если надобу дет
	result.Columns()
	var but string
	for result.Next() {
		result.Scan(&but)
		descriptionitog = append(descriptionitog, but)
	}
	//fmt.Printf("%s ВООООТ запрос по БД прошлись вижу вот", boolyesno)
	var sumLenElemen int
	for i := 0; i < len(descriptionitog); i++ { // Проверяет кол во ленов в массиве если больше 1 вернем true
		sumLenElemen += 1
	}
	if sumLenElemen >= 1 {
		str := descriptionitog[0]
		return str, true
	} else {
		str := "Нет данных о приветственном смс"
		return str, false
	}

	//var decriptionval string
	//for result.Next() {
	//	result.Scan(&decriptionval)
	//	descriptionitog = append(descriptionitog, decriptionval)
	//}
	//
	//stringitog := descriptionitog[0]
	//if errorka != nil {
	//	fmt.Printf("\n Результат запроса такой %s. А вот ошибка такая  %s ", result, errorka)
	//}
	////fmt.Printf(descriptionitog[0])
	////fmt.Printf(stringitog)
}

func GiveDBAdministratorsIDAdmin() []string {
	var givadmin []string // функция для получения кнопок которые отправятся в бд
	database, err := sql.Open("sqlite3", "./info.db")
	if err != nil {
		fmt.Printf("Ошибка в giveButton %s \n", err)
	}
	fmt.Printf("БАЗА %T \n", database)

	rows, err := database.Query("SELECT idAdmin FROM DBAdministrators") // запрос делаем
	if err != nil {
		fmt.Printf("Ошибка при получение кнопок из базы give button   %s \n", err)
	}
	defer database.Close()

	var idAdmin string
	for rows.Next() {
		rows.Scan(&idAdmin)
		givadmin = append(givadmin, idAdmin)
		//fmt.Printf("%d: %s ,%s \n", id, but, phot)
	}
	return givadmin
}

func GiveDBAdministratorsPrimaryKeyIDAdmiName() string {
	var givadminalldiscript []string // функция для получения кнопок которые отправятся в бд
	database, err := sql.Open("sqlite3", "./info.db")
	if err != nil {
		fmt.Printf("Ошибка в giveButton %s \n", err)
	}
	fmt.Printf("БАЗА %T \n", database)

	rows, err := database.Query("SELECT id,idAdmin, nameAdmin FROM DBAdministrators") // запрос делаем
	if err != nil {
		fmt.Printf("Ошибка при получение кнопок из базы give button   %s \n", err)
	}
	defer database.Close()

	for rows.Next() {
		var id int
		var idAdmin string
		var nameAdmin string

		rows.Scan(&id, &idAdmin, &nameAdmin)
		separateadmin := "№=" + strconv.Itoa(id) + " ИД=" + idAdmin + " ИМЯ=" + nameAdmin
		givadminalldiscript = append(givadminalldiscript, separateadmin)
		//fmt.Printf("%d: %s ,%s \n", id, but, phot)
	}
	allAdminStr := ""
	for _, value := range givadminalldiscript {
		allAdminStr += value + "\n"
	}
	return allAdminStr
}

func GiveDBAdministratorsChecID() []string {
	var givadminalldiscript []string // функция для получения всех существующих id
	database, err := sql.Open("sqlite3", "./info.db")
	if err != nil {
		fmt.Printf("Ошибка в giveButton %s \n", err)
	}
	fmt.Printf("БАЗА %T \n", database)

	rows, err := database.Query("SELECT id FROM DBAdministrators") // запрос делаем // получаем толкьо id
	if err != nil {
		fmt.Printf("Ошибка при получение кнопок из базы give button   %s \n", err)
	}
	defer database.Close()

	for rows.Next() {
		var id int

		rows.Scan(&id)
		givadminalldiscript = append(givadminalldiscript, strconv.Itoa(id))
		//fmt.Printf("%d: %s ,%s \n", id, but, phot)
	}

	return givadminalldiscript
}
