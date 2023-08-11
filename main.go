package main

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"strings"
	btcreate "tg_bot_golang/createbutton"
	dbadd "tg_bot_golang/db/adddate"
	dbcreate "tg_bot_golang/db/createdb"
	dbdrop "tg_bot_golang/db/drop"
	dbgive "tg_bot_golang/db/give"
)

func Info(data []string) {
	fmt.Printf("%s получаю из получение кнопок", data)
}
func main() {
	dbcreate.СreateTable()
	//будем отправлять смс типо да хорошо добавьте кнопку =)

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
		if strings.HasPrefix(update.Message.Text, "Добавить:") { //добавляет кнопку
			vartext := update.Message.Text[17:] //обрезаем
			itog := dbadd.AddButton(vartext)    //что передадим в добавить?
			msg.Text = itog                     // передаю инфу в текст
			fmt.Println("я зашел в добавить")
		}

		if strings.HasPrefix(update.Message.Text, "Удалить кнопку:") { // удаляет кнопку
			vartext := update.Message.Text[28:]
			itog := dbdrop.DropOneButton(vartext)
			msg.Text = itog // передаю инфу в текст

			fmt.Printf("я зашел в удалить кнопку текст такой:%s \n", vartext)
		}

		if strings.HasPrefix(update.Message.Text, "Удалить всё") {
			msg.Text = "Удалили данные"
			dbdrop.DropAllTableBani()

		}

		switch update.Message.Text {
		case "open":
			fmt.Println("сейчас выдам кнопки")
			buttonBase := tgbotapi.NewReplyKeyboard(btcreate.CreateButton(dbgive.GiveButtonInBase()))

			msg.ReplyMarkup = buttonBase
		case "close":
			msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
		}

		if _, err := bot.Send(msg); err != nil {
			fmt.Println(err)
		}
	}
}
