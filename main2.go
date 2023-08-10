package main

import (
	"database/sql"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func createButton(arrayButton []string) []tgbotapi.KeyboardButton { // получает массив значений и выдает кнопки меню для тг
	var itog []tgbotapi.KeyboardButton
	//a := tgbotapi.NewKeyboardButton("44") /// я тут разложил все а потом цикл понял как писать =)
	//b := tgbotapi.NewKeyboardButton("55")
	//c := tgbotapi.NewKeyboardButton("66")
	//itog = append(itog, a)
	//itog = append(itog, b)
	//itog = append(itog, c)
	////
	//lenarray := len(arrayButton)
	//for i := 0; i < lenarray; i++ {
	//	itog = append(itog, tgbotapi.NewKeyboardButton(strconv.Itoa(arrayButton[i])))
	//}

	for _, v := range arrayButton {
		itog = append(itog, tgbotapi.NewKeyboardButton(v))
	}

	return itog
}

// create db

func createTable() {
	database, _ := sql.Open("sqlite3", "./info.db")
	defer database.Close()

	statematn, _ := database.Prepare("CREATE TABLE IF NOT EXISTS bani(id INTEGER PRIMARY KEY ,buttun TEXT, description TEXT, photo TEXT)")
	statematn.Exec()
	statematn, _ = database.Prepare("INSERT INTO bani (buttun, description, photo) VALUES (?,?,?)") // добавляет значения
	statematn.Exec("First", "Second", "lol")
	rows, _ := database.Query("SELECT id,buttun,photo FROM bani") // запрос делаем

	var id int
	var but string
	var phot string
	for rows.Next() {
		rows.Scan(&id, &but, &phot)
		fmt.Printf("%d: %s x,%s это из бд\n ", id, but, phot)
	}
}

func giveButtonInBase() []string {
	var givButton []string // функция для получения кнопок которые отправятся в бд

	database, _ := sql.Open("sqlite3", "./info.db")
	rows, _ := database.Query("SELECT buttun FROM bani") // запрос делаем
	defer database.Close()

	var but string
	for rows.Next() {
		rows.Scan(&but)
		givButton = append(givButton, but)
		//fmt.Printf("%d: %s ,%s \n", id, but, phot)
	}
	return givButton
}

func main() {

	buttonBase := tgbotapi.NewReplyKeyboard(createButton(giveButtonInBase()))
	//fmt.Printf("%T :переменная button\n", button)
	//fmt.Printf("%T :переменная тест2\n", test2)
	//
	//fmt.Printf("%T :переменная numeric02\n", numeric02)
	//fmt.Printf("%T :переменная numeric01\n", numeric01)
	//fmt.Printf("%T :переменная numericKeyboard\n", numericKeyboard)
	//var rows *sql.Rows
	createTable() // вызвали и создали базу
	////здесь получаем из крейт тейбл наши данные и их распаршиваем в нашу функцию а потом передаем инфо в кнопки =)
	//var id int
	//for rows.Next() {
	//	rows.Scan(&id)
	//	fmt.Printf("%d:\n", id)
	//}

	bot, err := tgbotapi.NewBotAPI("5975063523:AAFagQJfXf3z-zgA0JjHPusoGhjjXIYOyEI")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore non-Message updates
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)

		switch update.Message.Text {
		case "open":
			msg.ReplyMarkup = buttonBase
		case "close":
			msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
		}

		if _, err := bot.Send(msg); err != nil {
			fmt.Println(err)
		}
	}
}

//
//var button = createButton()
//
//var value1 = []tgbotapi.KeyboardButton{}

//var button = []tgbotapi.KeyboardButton{tgbotapi.NewKeyboardButton("11"), tgbotapi.NewKeyboardButton("22")}
