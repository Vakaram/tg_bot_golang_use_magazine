package main

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	_ "github.com/mattn/go-sqlite3"
	"strconv"
	btcreate "tg_bot_golang/createbutton"
	dbadd "tg_bot_golang/db/adddate"
	dbcreate "tg_bot_golang/db/createdb"
	dbdrop "tg_bot_golang/db/drop"
	dbgive "tg_bot_golang/db/give"
	dbupdate "tg_bot_golang/db/update"
	"tg_bot_golang/logger"
)

func Info(data []string) {
	fmt.Printf("%s получаю из получение кнопок", data)
}

func main() {
	//logger.CrateLogger()
	const chiefadminBot string = "1484570227" // эта запись нужна для телеграм бота чтобы так добавился главный администратор который моэет добавлять других администраторов
	bot, err := tgbotapi.NewBotAPI("5975063523:AAFagQJfXf3z-zgA0JjHPusoGhjjXIYOyEI")
	if err != nil {
		fmt.Printf(err.Error())
	}
	dbcreate.СreatemyDB() // создаем нашу базу на новом месте =)
	dbcreate.СreateDbGreetings()
	dbcreate.СreateDBAdministrators(chiefadminBot)
	//logger.LoggerInfo("Новый запуск тесссст ")

	//Тут создаем набор команд он часто вызывается будет как константа
	var comands = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Меню"),
			tgbotapi.NewKeyboardButton("Команды"),
			tgbotapi.NewKeyboardButton("Удалить всю базу"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Добавить кнопку"),
			tgbotapi.NewKeyboardButton("Добавить описание"),
			tgbotapi.NewKeyboardButton("Удалить кнопку"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Обновить название"),
			tgbotapi.NewKeyboardButton("Добавить фото к кнопке"),
			tgbotapi.NewKeyboardButton("Удалить все фото для кнопки"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Стоп"),
			tgbotapi.NewKeyboardButton("Добавить приветствие"), // это так же работает как обновление
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Добавить администратора"), // это так же работает как обновление
			tgbotapi.NewKeyboardButton("Удалить администратора"),  // это так же работает как обновление
		),
	)

	var yesNo = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Да"),
			tgbotapi.NewKeyboardButton("Нет"),
		),
	)

	// тест здесь будут функции и команды которые приходят иметация стайтд машин ? Ну попробуем
	var stagekomand string          // будет пустая команда если заполнена будем её затирать если используем // для машин стостояния
	var stagenameButtonTheUp string // типа тут мы обновим кнопку на это навание // для id

	bot.Debug = true
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 20 // сколько будем ждать запрос ?
	updates := bot.GetUpdatesChan(u)

	for update := range updates { // тут начинмется работа бота
		if update.Message == nil { // ignore non-Message updates
			continue
		}

		logger.Info.Println("Получили запрос такой " + update.Message.Text + " От пользователя с id = " + strconv.Itoa(int(update.Message.From.ID)))

		//тут получаем всех админов и делаем сравнение с ними
		admins := dbgive.GiveDBAdministratorsIDAdmin()
		fmt.Printf("Вот что получаю из базы админов : %s\n", admins[0])
		userID := update.Message.From.ID
		reallyAdmin := false
		for _, value := range admins {
			if value == strconv.Itoa(int(userID)) {
				reallyAdmin = true
			}
		}
		fmt.Printf("Реально ли это админ ?:  %t \n", reallyAdmin)

		if reallyAdmin == true {
			logger.Info.Println("Получили запрос такой " + update.Message.Text + " Да это админ = " + strconv.Itoa(int(update.Message.From.ID)))

			buttonBase := tgbotapi.NewReplyKeyboard(btcreate.CreateButton(dbgive.GiveButtonInBase())) // формирует меню из кнопок базы которые увидит клиент //Важно что каждый раз нужно проверять в каждом сообщение

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "") // для отправки копии смс

			switch stagekomand { //для проверки пришедших команд смсок точнее ну будем дальше смотреть
			case "Добавить кнопку":
				stagekomand = ""
				itog := dbadd.AddButton(update.Message.Text) //что передадим в добавить?
				msg.Text = itog
				msg.ReplyMarkup = comands
			case "Удалить кнопку":
				stagekomand = "" // удаляем состояние
				//пердаем название кнопки для удаления
				drop := dbdrop.DropOneButton(update.Message.Text)
				msg.Text = drop
				msg.ReplyMarkup = comands
			case "Добавить приветствие":
				stagekomand = "" // удаляем состояние
				//пердаем новое приветствие в бд
				updateData := dbupdate.UpdateDbGreetings(update.Message.Text)
				if updateData != "" {
					msg.Text = updateData
					msg.ReplyMarkup = comands
				} else {
					msg.Text = "данные не удалось обновить" // пока просто как пример но надо будет делать обновления дла баз данных как то смотреть а добавились ли данные
					msg.ReplyMarkup = comands

				}
			case "Удалить все фото для кнопки":
				stagekomand = ""
				// тут должна быть функцию удаления
				dropPhotoAll := dbdrop.DropAllPhotoButton(update.Message.Text)
				msg.Text = dropPhotoAll
				msg.ReplyMarkup = comands
			case "Удалить всю базу":
				if update.Message.Text == "Да" { // если название еще не ввели то мы запрашиваем кнопку А вот если ввели ниже
					dbdrop.DropAllTablemyDb()
					msg.Text = "Удалили базу"
					msg.ReplyMarkup = comands
				} else {
					stagekomand = ""
					msg.Text = "Отменили удаление базы"
					msg.ReplyMarkup = comands
				}
			case "Добавить описание":
				if stagenameButtonTheUp == "" { // если название еще не ввели то мы запрашиваем кнопку А вот если ввели ниже
					stagenameButtonTheUp = update.Message.Text // пишем какую кнопку будем обновлять
					msg.Text = dbgive.GiveDescriptionButton(stagenameButtonTheUp) + "\n" + "Пришлите описание кнопки"

				} else { // а вот если ввели то мы формируем простой запрос из пердыдущей и новой кнопки а так же очищаем наше "Состояние" Чтобы прога работала заново
					answer := dbupdate.UpdateDescriptionButton(stagenameButtonTheUp, update.Message.Text)
					msg.ReplyMarkup = comands
					msg.Text = answer
					//а теперь очистим наши состояния
					stagekomand, stagenameButtonTheUp = "", "" // просто сделали их пустыми =)
				}
			case "Добавить администратора":
				if stagenameButtonTheUp == "" { // если название еще не ввели то мы запрашиваем кнопку А вот если ввели ниже
					stagenameButtonTheUp = update.Message.Text // сюда мы записали id нового админа дальше будем писать
					msg.Text = "Теперь пришлите имя вашего администратора для вашего удобства"
				} else { // а вот если ввели то мы формируем простой запрос из пердыдущей и новой кнопки а так же очищаем наше "Состояние" Чтобы прога работала заново
					answer := dbadd.AddAdmin(stagenameButtonTheUp, update.Message.Text)
					msg.ReplyMarkup = comands
					msg.Text = answer
					//а теперь очистим наши состояния
					stagekomand, stagenameButtonTheUp = "", "" // просто сделали их пустыми =) чтобы очистить машинное состояние
				}
			case "Удалить администратора":
				whydrpop := dbdrop.DropOneAdmin(update.Message.Text)
				msg.Text = whydrpop
			case "Добавить фото к кнопке":
				if stagenameButtonTheUp == "" { // если название еще не ввели то мы запрашиваем кнопку А вот если ввели ниже
					stagenameButtonTheUp = update.Message.Text // пишем какую кнопку будем обновлять eё передадим в функцию
					msg.Text = "Пришлите фото для кнопки, добавлять можно только по одному фото"
				} else {
					msg.ReplyMarkup = comands
					sendfunk := dbupdate.UpdateAddPhotoInButton(stagenameButtonTheUp, update.Message.Photo[3].FileID) // передает файл ид лучшего качества
					msg.Text = sendfunk
					stagekomand, stagenameButtonTheUp = "", "" // просто сделали их пустыми =)
				}
			case "Обновить название для кнопки":
				if stagenameButtonTheUp == "" { // если название еще не ввели то мы запрашиваем кнопку А вот если ввели ниже
					stagenameButtonTheUp = update.Message.Text // пишем какую кнопку будем обновлять
					msg.Text = "Осталось написать новое название кнопки"
				} else { // а вот если ввели то мы формируем простой запрос из пердыдущей и новой кнопки а так же очищаем наше "Состояние" Чтобы прога работала заново
					answer := dbupdate.UpdateButton(stagenameButtonTheUp, update.Message.Text)
					msg.Text = answer
					//а теперь очистим наши состояния
					stagekomand, stagenameButtonTheUp = "", "" // просто сделали их пустыми =)
					msg.ReplyMarkup = comands
				}
			}

			if stagekomand == "" {
				check_button := dbgive.CheckingdbButton(update.Message.Text) // эта часть отвечает за поиск информации за кнопку если есть вернем тру
				if check_button == true {
					resultDeascription := dbgive.GiveDescriptionButton(update.Message.Text)
					if resultDeascription != "" {
						arr := dbgive.GivePhotoButton(update.Message.Text, resultDeascription)
						if arr == nil {
							msg.Text = resultDeascription
							return
						}
						bot.Send(tgbotapi.NewMediaGroup(update.Message.Chat.ID, arr))
					} else {
						resultDeascription = "Нет описания и нет фото для этой кнопки"
						msg.Text = resultDeascription
					}

				}
			}

			switch update.Message.Text { // если есть совпадение команды мы делаем стейт машин =)

			case "Меню":
				isEmpty := len(buttonBase.Keyboard) == 0 // выдает ошибку считай но выше хммм
				var textInMenu string
				if isEmpty {
					textInMenu = "Пока здесь нет меню, оно скоро появится" // если пусто то вот это отрпавим
				} else {
					textInMenu = "Показываю меню"
					msg.ReplyMarkup = buttonBase
				}
				msg.Text = textInMenu

				//msg.ReplyMarkup = buttonBase
				stagekomand, stagenameButtonTheUp = "", "" // просто сделали их пустыми =)
			case "Добавить кнопку":
				stagekomand = "Добавить кнопку"
				vartext := "Напишите название кнопки"
				msg.Text = vartext
			case "Стоп":
				stagekomand = ""
				stagenameButtonTheUp = "" // очистили машинное состояние
				vartext := "Отменили и очистили ввод данных"
				msg.Text = vartext
			case "Добавить администратора":
				stagekomand = "Добавить администратора"
				vartext := "Пришлите id администратор"
				msg.Text = vartext
			case "Удалить администратора":
				stagekomand = "Удалить администратора"
				vartext := dbgive.GiveDBAdministratorsPrimaryKeyIDAdmiName() + "Пришлите № администратор для удаления"
				msg.Text = vartext
			case "Удалить кнопку":
				stagekomand = "Удалить кнопку"
				msg.Text = "Какую кнопку удалить?" // отправлем смс
				msg.ReplyMarkup = buttonBase
			case "Удалить всю базу":
				msg.Text = "Вы уверены?"
				stagekomand = "Удалить всю базу"
				msg.ReplyMarkup = yesNo
			case "Обновить название для кнопки":
				// обновляет значения кнопки
				stagekomand = "Обновить название для кнопки"
				msg.ReplyMarkup = buttonBase
				msg.Text = "Какую кнопку обновим?" // передаю инфу в текст
			case "Добавить описание":
				stagekomand = "Добавить описание"
				msg.Text = "К какой кнопке добавим описание?"
				msg.ReplyMarkup = buttonBase
			case "Добавить фото к кнопке":
				stagekomand = "Добавить фото к кнопке"
				msg.Text = "К какой кнопке добавим фото?"
				msg.ReplyMarkup = buttonBase
			case "Удалить все фото для кнопки":
				stagekomand = "Удалить все фото для кнопки"
				msg.Text = "У какой кнопки удалим все фотки?"
				msg.ReplyMarkup = buttonBase
			case "Добавить приветствие":
				stagekomand = "Добавить приветствие"
				sendDBGreetings, _ := dbgive.GiveDescriptionDBGreetings()
				msg.Text = "Пришлите новое приветствие, высылаю вам шаблон старого:\n" + sendDBGreetings

			case "К": // будет выводить список команд доступных чтобы не писать кнопки каждый раз вручную
				{
					msg.ReplyMarkup = comands
					msg.Text = "Вот доступные команды:"
				}
			}

			if _, err := bot.Send(msg); err != nil {
				fmt.Println(err)
			}

		} else { //
			logger.Info.Println("Получили запрос такой " + update.Message.Text + "Нет это не админ = " + strconv.Itoa(int(update.Message.From.ID)))
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Такой кнопки мы не нашли, ниже я вам отправил список доступных Кнопок (услуг товаров) ") // для отправки копии смс
			if stagekomand == "" {
				buttonBase := tgbotapi.NewReplyKeyboard(btcreate.CreateButton(dbgive.GiveButtonInBase())) // формирует меню из кнопок базы которые увидит клиент //Важно что каждый раз нужно проверять в каждом сообщение
				// все команды должны быьть внутри тк может произойти ошибка записи бд
				if update.Message.Text == "/start" {
					// тут реализовано если описания в базе нет то мы просто пришлем заглушку
					sendDBGreetings, _ := dbgive.GiveDescriptionDBGreetings()
					msg.Text = sendDBGreetings
					msg.ReplyMarkup = buttonBase
				} else {
					check_button := dbgive.CheckingdbButton(update.Message.Text) // эта часть отвечает за поиск информации за кнопку если есть вернем тру
					if check_button == true {
						//
						resultDeascription := dbgive.GiveDescriptionButton(update.Message.Text)
						if resultDeascription != "" {
							arr := dbgive.GivePhotoButton(update.Message.Text, resultDeascription)
							if arr == nil {
								msg.Text = resultDeascription
								return
							}
							bot.Send(tgbotapi.NewMediaGroup(update.Message.Chat.ID, arr))
						} else {
							resultDeascription = "Нет описания и нет фото для этой кнопки"
						}
					} else {
						buttonBase := tgbotapi.NewReplyKeyboard(btcreate.CreateButton(dbgive.GiveButtonInBase())) // формирует меню из кнопок базы которые увидит клиент //Важно что каждый раз нужно проверять в каждом сообщение
						msg.ReplyMarkup = buttonBase
					}
				}

			} else {
				msg.Text = "Простите в данный момент добавляют информацию в Базу данных\n"
			}

			if _, err := bot.Send(msg); err != nil {
				fmt.Println(err)
			}
		}

	}
}
