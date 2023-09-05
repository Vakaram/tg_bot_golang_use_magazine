package createbutton

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	logger "tg_bot_golang/logger"
)

func CreateButton(arrayButton []string) []tgbotapi.KeyboardButton { // получает массив значений и выдает кнопки меню для тг
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
	errStr := "CreateButton() сработал его вызвали все ок:" // err.Error() переводим ошибку в строку
	logger.Info.Println(errStr)
	return itog
}
