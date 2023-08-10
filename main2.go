//package main
//
//import (
//	"log"
//
//	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
//)
//
//func main() {
//	bot, err := tgbotapi.NewBotAPI("5975063523:AAFagQJfXf3z-zgA0JjHPusoGhjjXIYOyEI")
//	if err != nil {
//		log.Panic(err)
//	}
//
//	log.Printf("Authorized on account %s", bot.Self.UserName)
//
//	u := tgbotapi.NewUpdate(0)
//	u.Timeout = 60
//
//	updates := bot.GetUpdatesChan(u)
//
//	for update := range updates {
//		if update.Message != nil { // If we got a message
//			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
//
//			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
//			msg.ReplyToMessageID = update.Message.MessageID
//
//			bot.Send(msg)
//		}
//	}
//}

package main

import (
	"database/sql"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"strconv"
)

// create db

func createTable() {
	database, _ := sql.Open("sqlite3", "./info.db")
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
		fmt.Printf("%d: %s ,%s \n", id, but, phot)
	}

}

//	type testbut struct {
//		name
//		namebutton string
//	}
//
//	func NewReplyKeyboard(rows ...[]KeyboardButton) ReplyKeyboardMarkup {
//		var keyboard [][]KeyboardButton
//
//		keyboard = append(keyboard, rows...)
//
//		return ReplyKeyboardMarkup{
//			ResizeKeyboard: true,
//			Keyboard:       keyboard,
//		}
//	}

func main() {
	//fmt.Printf("%T :переменная button\n", button)
	fmt.Printf("%T :переменная тест2\n", test2)

	fmt.Printf("%T :переменная numeric02\n", numeric02)
	fmt.Printf("%T :переменная numeric01\n", numeric01)
	fmt.Printf("%T :переменная numericKeyboard\n", numericKeyboard)
	//createTable() // вызвали и создали базу
	bot, err := tgbotapi.NewBotAPI("")
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
			msg.ReplyMarkup = test2
		case "close":
			msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
		}

		if _, err := bot.Send(msg); err != nil {
			fmt.Println(err)
		}
	}
}

func createButton() []tgbotapi.KeyboardButton {
	var itog []tgbotapi.KeyboardButton
	//a := tgbotapi.NewKeyboardButton("44")
	//b := tgbotapi.NewKeyboardButton("55")
	//c := tgbotapi.NewKeyboardButton("66")
	//itog = append(itog, a)
	//itog = append(itog, b)
	//itog = append(itog, c)

	for i := 0; i < 5; i++ {
		itog = append(itog, tgbotapi.NewKeyboardButton(strconv.Itoa(i)))
	}

	return itog
}

//
//var button = createButton()
//
//var value1 = []tgbotapi.KeyboardButton{}

//var button = []tgbotapi.KeyboardButton{tgbotapi.NewKeyboardButton("11"), tgbotapi.NewKeyboardButton("22")}

var test2 = tgbotapi.NewReplyKeyboard(createButton())

var numeric02 = tgbotapi.NewKeyboardButton("1")
var numeric01 = tgbotapi.NewKeyboardButtonRow(
	numeric02,
)
var numericKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("1"),
		tgbotapi.NewKeyboardButton("2"),
		tgbotapi.NewKeyboardButton("3"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("4"),
		tgbotapi.NewKeyboardButton("5"),
		tgbotapi.NewKeyboardButton("6"),
	),
)
