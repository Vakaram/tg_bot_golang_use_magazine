package createdb

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

// create db только создание  // *sql.DB
func СreatemyDb() {
	database, err := sql.Open("sqlite3", "./info.db")
	if err != nil {
		fmt.Printf("Ошибка в create bd %s \n", err)
	}
	defer database.Close()
	statematn, err := database.Prepare("CREATE TABLE IF NOT EXISTS myDb(id INTEGER PRIMARY KEY ,button TEXT, description TEXT, photo TEXT)")
	if err != nil {
		fmt.Printf("Ошибка при создание базы запросом    %s \n", err)
	}
	statematn.Exec()
}

func СreateDbGreetings() {
	database, err := sql.Open("sqlite3", "./info.db")
	if err != nil {
		fmt.Printf("Ошибка в create bd %s \n", err)
	}
	defer database.Close()
	statematn, err := database.Prepare("CREATE TABLE IF NOT EXISTS DbGreetings(id INTEGER PRIMARY KEY ,button TEXT, Greetings TEXT)")

	if err != nil {
		fmt.Printf("Ошибка при создание базы запросом    %s \n", err)
	}
	statematn.Exec()
	// тут реализуем добавление заглушки в id номер 1 и проверку и добавления если её там нет и сразу будем читать от туда инфу при старьте бота
	var descriptionitog []string
	firsdescrirption := 1                                                                           // ищем всегда по 1 id                                                                             // берем инфу из первого id и его же будем обновлять всегддддааа=  )
	result, err := database.Query("SELECT Greetings FROM DbGreetings WHERE id=?", firsdescrirption) // пока пойдет и эта часть потом обновим если надобу дет
	result.Columns()
	var but string
	for result.Next() {
		result.Scan(&but)
		descriptionitog = append(descriptionitog, but)
	}
	//fmt.Printf("%s ВООООТ запрос по БД прошлись вижу вот", boolyesno)
	if len(descriptionitog) < 1 { // когда данных нет запишем новые
		// при создание закидывает заглушку и всегда её печатает
		// если человек создал приветсвтие то тогда данные будут браться так же новые из функции другой
		statement, err := database.Prepare("INSERT INTO DbGreetings(Greetings) VALUES (?)") //statement - заявление перевод
		if err != nil {
			fmt.Printf("Ошибка добавление кнопокв в AddButton в   %s \n", err)
		}
		zagluska := "Приветствую вас в нашем боте вот такие кнопки у нас есть. Нажав кнопку вы получите развернутое описание"
		statement.Exec(zagluska)
	}
}

func СreateDBAdministrators(chiefadminBot string) { // сюда передаем id константу чтобы её записывать как главного пользователя если он есть то не записывать// будет создавать базу админов по ней будем ходить читать и добавлять админов
	database, err := sql.Open("sqlite3", "./info.db")
	if err != nil {
		fmt.Printf("Ошибка в create bd %s \n", err)
	}
	defer database.Close()
	statematn, err := database.Prepare("CREATE TABLE IF NOT EXISTS DBAdministrators(id INTEGER PRIMARY KEY ,idAdmin TEXT, nameAdmin TEXT)")

	if err != nil {
		fmt.Printf("Ошибка при создание базы запросом    %s \n", err)
	}
	statematn.Exec()
	// тут реализуем добавление заглушки в id номер 1 и проверку и добавления если её там нет и сразу будем читать от туда инфу при старьте бота
	var descriptionitog []string
	firstAdmin := 1                                                                              // ищем всегда по 1 id                                                                             // берем инфу из первого id и его же будем обновлять всегддддааа=  )
	result, err := database.Query("SELECT idAdmin FROM DBAdministrators WHERE id=?", firstAdmin) // пока пойдет и эта часть потом обновим если надобу дет
	result.Columns()
	var but string
	for result.Next() {
		result.Scan(&but)
		descriptionitog = append(descriptionitog, but)
	}
	//fmt.Printf("%s ВООООТ запрос по БД прошлись вижу вот", boolyesno)
	if len(descriptionitog) < 1 { // когда данных нет запишем новые
		// при создание закидывает заглушку и всегда её печатает
		// если человек создал приветсвтие то тогда данные будут браться так же новые из функции другой
		statement, err := database.Prepare("INSERT INTO DBAdministrators(idAdmin ,nameAdmin) VALUES (?,?)") //statement - заявление перевод

		if err != nil {
			fmt.Printf("Ошибка добавление кнопокв в AddButton в   %s \n", err)
		}
		nameChefAdmin := "Это вы"
		statement.Exec(chiefadminBot, nameChefAdmin)
	}
}
