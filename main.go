package main

// TODO Сделать команду которая очищает все состояния то есть очищает переменныые stagekomand и nameButtonTheUp
import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	_ "github.com/mattn/go-sqlite3"
	"log"
	btcreate "tg_bot_golang/createbutton"
	dbadd "tg_bot_golang/db/adddate"
	dbcreate "tg_bot_golang/db/createdb"
	dbdrop "tg_bot_golang/db/drop"
	dbgive "tg_bot_golang/db/give"
	dbupdate "tg_bot_golang/db/update"
)

func Info(data []string) {
	fmt.Printf("%s получаю из получение кнопок", data)
}
func main() {
	dbcreate.СreateTable()
	//будем отправлять смс типо да хорошо добавьте кнопку =)

	// тест здесь будут функции и команды которые приходят иметация стайтд машин ? Ну попробуем
	var stagekomand string          // будет пустая команда если заполнена будем её затирать если используем
	var stagenameButtonTheUp string // типа тут мы обновим кнопку на это навание

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
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Такой команды нет") // для отправки копии смс

		//TODO сделать проверку вначале на есть ли такая кнопка в меню как бы регируем на запросы пользователей

		check_button := dbgive.CheckingdbButton(update.Message.Text)
		if check_button == true {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "TRUE") // для отправки копии смс
			if _, err := bot.Send(msg); err != nil {
				fmt.Println(err)
			}
		}

		switch stagekomand { //для проверки пришедших команд смсок точнее ну будем дальше смотреть
		case "Добавить:":
			stagekomand = ""
			itog := dbadd.AddButton(update.Message.Text) //что передадим в добавить?
			msg.Text = itog                              // передаю инфу в текст
		case "Обновить кнопку:":
			if stagenameButtonTheUp == "" { // если название еще не ввели то мы запрашиваем кнопку А вот если ввели ниже
				stagenameButtonTheUp = update.Message.Text // пишем какую кнопку будем обновлять
				msg.Text = "Осталось написать новое название"
			} else { // а вот если ввели то мы формируем простой запрос из пердыдущей и новой кнопки а так же очищаем наше "Состояние" Чтобы прога работала заново
				answer := dbupdate.UpdateButton(stagenameButtonTheUp, update.Message.Text)
				msg.Text = answer
				//а теперь очистим наши состояния
				stagekomand, stagenameButtonTheUp = "", "" // просто сделали их пустыми =)
			}

		}

		//if strings.HasPrefix(update.Message.Text, "Удалить кнопку:") { // удаляет кнопку
		//	vartext := update.Message.Text[28:]
		//	itog := dbdrop.DropOneButton(vartext)
		//	msg.Text = itog // передаю инфу в текст
		//
		//	fmt.Printf("я зашел в удалить кнопку текст такой:%s \n", vartext)
		//}

		//if strings.HasPrefix(update.Message.Text, "Удалить всё") {
		//	msg.Text = "Удалили данные"
		//	dbdrop.DropAllTableBani()
		//
		//}

		switch update.Message.Text { //для проверки пришедших команд смсок точнее ну будем дальше смотреть
		case "open":
			fmt.Println("сейчас выдам кнопки")
			buttonBase := tgbotapi.NewReplyKeyboard(btcreate.CreateButton(dbgive.GiveButtonInBase()))
			msg.ReplyMarkup = buttonBase
		case "close":
			msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
		case "Добавить:":
			stagekomand = "Добавить:"
			vartext := "Напишите название кнопки"
			msg.Text = vartext
		case "Удалить кнопку:":
			vartext := update.Message.Text[28:]
			itog := dbdrop.DropOneButton(vartext)
			msg.Text = itog // передаю инфу в текст
			fmt.Printf("я зашел в удалить кнопку текст такой:%s \n", vartext)
		case "Удалить всё":
			msg.Text = "Удалили данные"
			dbdrop.DropAllTableBani()
		case "Обновить кнопку:":
			// обновляет значения кнопки
			//vartext := update.Message.Text[28:]
			//itog := dbdrop.DropOneButton(vartext)
			stagekomand = "Обновить кнопку:"
			msg.Text = "Какую кнопку обновим? " // передаю инфу в текст

		}

		if _, err := bot.Send(msg); err != nil {
			fmt.Println(err)
		}
	}
}
