package main

import (
	"log"

	calculator "telebot2/calcs"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

func main() {
	calculator.TTT()
	// подключаемся к боту с помощью токена
	bot, err := tgbotapi.NewBotAPI("1756611769:AAEHSoOCzhHmsU--r3fDPsCVYMvi5DKaDec")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	// инициализируем канал, куда будут прилетать обновления от API
	var ucfg tgbotapi.UpdateConfig = tgbotapi.NewUpdate(0)
	ucfg.Timeout = 60
	upd, _ := bot.GetUpdatesChan(ucfg)
	// читаем обновления из канала
	botStop := false
	for botStop == false {
		select {
		case update := <-upd:
			// Пользователь, который написал боту
			UserName := update.Message.From.UserName

			// ID чата/диалога.
			// Может быть идентификатором как чата с пользователем
			// (тогда он равен UserID) так и публичного чата/канала
			ChatID := update.Message.Chat.ID

			// Текст сообщения
			Text := update.Message.Text

			if Text == "bb" { //стоп бот
				Text = "22"
				botStop = true
			}

			// if Text == "testcalc" { //стоп бот
			// 	priceAfterMarjin, marjinAmount := calculate.marjin(price, calcVariables["marjin"], plusMinusMarjin)
			// 	botStop = true
			// }

			log.Printf("[%s] %d %s", UserName, ChatID, Text)

			// Ответим пользователю его же сообщением
			reply := Text
			// Созадаем сообщение
			msg := tgbotapi.NewMessage(ChatID, reply)
			// и отправляем его
			bot.Send(msg)
		}

	}
}
