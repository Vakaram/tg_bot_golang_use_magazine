package main

//
////TODO готово сделать так чтобы бот работал постоянно и ждал смс
////TODO отображать кнопки inline
//import (
//
//	"bytes"
//	"encoding/json"
//	"fmt"
//	"io/ioutil"
//	"log"
//	"net/http"
//	"strconv"
//	"strings"
//
//)
//
//type GetMeT struct { // T добавляет для того чтобы показать что это type
//
//		Ok     bool         `json:"ok"`
//		Result GetMeResultT `json:"result"`
//	}
//
//	type GetMeResultT struct {
//		Id        int    `json:"id"`
//		IsBot     bool   `json:"is_bot"`
//		FirstName string `json:"first_name"`
//		UserName  string `json:"username"`
//	}
//
//	type GetUpdatesT struct {
//		Ok     bool                `json:"ok"`
//		Result []GetUpdatesResultT `json:"result"`
//	}
//
//	type GetUpdatesResultT struct {
//		UpdateID int                `json:"update_id"`
//		Message  GetUpdatesMessageT `json:"message,omitempty"`
//	}
//
//	type GetUpdatesMessageT struct {
//		MessageID int `json:"message_id"`
//		From      struct {
//			ID           int    `json:"id"`
//			IsBot        bool   `json:"is_bot"`
//			FirstName    string `json:"first_name"`
//			LastName     string `json:"last_name"`
//			UserName     string `json:"user_name"`
//			LanguageCode string `json:"language_code"`
//		} `json:"from"`
//		Chat struct {
//			ID        int    `json:"id"`
//			FirstName string `json:"first_name"`
//			LastName  string `json:"last_name"`
//			Username  string `json:"username"`
//			Type      string `json:"type"`
//		} `json:"chat"`
//		Data int    `json:"data"`
//		Text string `json:"text"`
//	}
//
//// кнопки
//
//	type ReplyKeyboardMarkup struct {
//		Keyboard [][]KeyboardButton `json:"keyboard"`
//	}
//
//	type KeyboardButton struct {
//		Text map[string]string `json:"text"`
//	}
//
//const telegramBaseUrl = "https://api.telegram.org/bot"
//const telegramToken = ""
//
//const methodGetMe = "getMe"
//const methodGetUpdates = "getUpdates"
//const methodSendMessage = "sendMessage"
//
//	func main() {
//		offset := 0
//		for {
//			body := getBodyByUrl(getUrlByMethod(methodGetUpdates, offset)) // получаем срез байтов п
//			get := GetUpdatesT{}                                           // проинициализировали структуру те её объявлили теперь в ние будем заполнть парсить наши данные
//			err := json.Unmarshal(body, &get)                              // пережавать нужно указатель а не просто переменную  // заполняем стркутуру из jsona // передаем ей срез байтов потом струткутру в которую хотим разобрать
//			if err != nil {
//				fmt.Println(err.Error())
//				return
//			}
//			fmt.Println(get.Result) // и так я получаю тут id ботаи другие данные из тела
//
//			//for i := 0; i < 0; i++ {
//			//	a := get.Result
//			//	fmt.Println(a)
//			//}
//
//			for _, update := range get.Result {
//				offset = update.UpdateID + 1
//				MakeRequest(strconv.Itoa(update.Message.Chat.ID), "tessst")
//
//				//// делаем эхо бота
//				//url := getUrlSendMessage(strconv.Itoa(update.Message.Chat.ID), update.Message.Text)
//				//url = string(getBodyByUrl(url))
//				//fmt.Println("Сформировал url:" + string(url)) // это эхоботовская часть
//
//				//**********************
//				//Тут попробуем сформировать кнопки чтобы отображать их в боте
//
//				//if strings.ToLower(update.Message.Text) == "Кнопки" { // пропускаем сообщение старт
//				//	// using Telegram.Bot.Types.ReplyMarkups;
//				//
//				//	ReplyKeyboardMarkup replyKeyboardMarkup = new(new[]
//				//	{
//				//		new KeyboardButton[] { "Help me", "Call me ☎️" },
//				//	})
//				//	{
//				//	ResizeKeyboard = true
//				//	};
//				//
//				//	Message sentMessage = await botClient.SendTextMessageAsync(
//				//		chatId: chatId,
//				//		text: "Choose a response",
//				//		replyMarkup: replyKeyboardMarkup,
//				//		cancellationToken: cancellationToken);
//				//
//				//	continue
//				//}
//
//				if strings.ToLower(update.Message.Text) == "go" { // пропускаем сообщение старт
//					sms := "go go go"
//
//					MakeRequest(strconv.Itoa(update.Message.Chat.ID), sms)
//					//
//					//var testbutton ReplyKeyboardMarkup
//					//m := make(map[string]string)
//					//m["button"] = "test"
//					//testbutton.Keyboard = [][]KeyboardButton{{{Text: m}}}
//
//					//url := getUrlSendMessage(strconv.Itoa(update.Message.Chat.ID), sms)
//					//fmt.Println(url)
//					//url = string(getBodyByUrl(url))
//					continue
//				}
//				if strings.Contains(update.Message.Text, "load") { // ищить корни части слов в словах
//					sms := "llllloooooad "
//					url := getUrlSendMessage(strconv.Itoa(update.Message.Chat.ID), sms)
//					fmt.Println(url)
//					url = string(getBodyByUrl(url))
//					continue
//
//				}
//			}
//		}
//	}
//
//// / пробуем отправить тело запроса тест
//func MakeRequest(chatId string, text string) {
//
//		//message := map[string]string{
//		//	"hello": "world",
//		//}
//
//		var message ReplyKeyboardMarkup
//		m := make(map[string]string)
//		m["button"] = "test"
//		message.Keyboard = [][]KeyboardButton{{{Text: m}}}
//
//		bytesRepresentation, err := json.Marshal(message)
//		if err != nil {
//			log.Fatalln(err)
//		}
//
//		resp, err := http.Post("https://api.telegram.org/bot"+telegramToken+"sendMessage"+"?chat_id="+chatId+"&text="+text, "application/json", bytes.NewBuffer(bytesRepresentation))
//		fmt.Println(resp)
//		if err != nil {
//			log.Fatalln(err)
//		}
//
//		var result map[string]interface{}
//
//		json.NewDecoder(resp.Body).Decode(&result)
//
//		log.Println(result)
//		log.Println(result["data"])
//	}
//
//// func getMurcup(chat_id string, keyb ReplyKeyboardMarkup) string { // для отправки клавиатуры
////
////	return telegramBaseUrl + telegramToken + "/" + "sendMessage" + "?chat_id=" + chat_id + "&reply_markup=" + keyb
////
//// }
//func getUrlSendMessage(chat_id string, text string) string { // принимает два обязательных параметра и формирует url это для отправки смс
//
//	return telegramBaseUrl + telegramToken + "/" + "sendMessage" + "?chat_id=" + chat_id + "&text=" + text
//
//}
//
//func getUrlByMethod(methodName string, offset_id int) string { // тут добавлю один параметр и его склею в offcet = https://api.telegram.org/bot5975063523:AAFagQJfXf3z-zgA0JjHPusoGhjjXIYOyEI/getUpdates?offset=830217713
//
//		return telegramBaseUrl + telegramToken + "/" + methodName + "?offset=" + strconv.Itoa(offset_id)
//	}
//
//	func getBodyByUrl(url string) []byte {
//		fmt.Println(url)               // вся функция возвращает тело боди запроса
//		response, err := http.Get(url) // подставляем url и типа отправляем запрос на полученные данные можно отправялть запро дальше
//		fmt.Println(response)
//		if err != nil {
//			fmt.Println(err.Error())
//		}
//		defer response.Body.Close() // мы в респон записали запрос наш и его закываем по сле выполнение программы
//		body, err := ioutil.ReadAll(response.Body)
//		fmt.Println(body)
//		if err != nil {
//			fmt.Println(err.Error())
//		}
//		return body
//	}
