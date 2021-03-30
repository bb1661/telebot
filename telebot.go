package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	calc "telebot2/calcs"
	rep "telebot2/replies"
	sql "telebot2/sql"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

func main() {
	botStop := false
	calc.Test()
	sql.Test()
	rep.Test()

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

	for !botStop {
		select {
		case update := <-upd:
			// Пользователь, который написал боту
			UserName := update.Message.From.UserName

			// ID чата/диалога.
			// Может быть идентификатором как чата с пользователем
			// (тогда он равен UserID) так и публичного чата/канала
			ChatID := update.Message.Chat.ID

			// Текст сообщения

			Text := strings.ToLower(update.Message.Text)

			log.Printf("[%s] %d %s", UserName, ChatID, Text)

			switch {
			case Text == "/help":
				Text = rep.Help()
			case Text == "/calculator":
				Text = rep.CalculatorInit()
			case strings.HasPrefix(Text, "/pmpn "):
				if price, err := strconv.ParseFloat(value(Text), 64); err == nil {
					price, _ = calc.Marjin(price, 0.3, true)
					price, _ = calc.Vat(price, 0.2, true)
					Text = fmt.Sprintf("%f", price)
				}
			case strings.HasPrefix(Text, "/pmmn "):
				if price, err := strconv.ParseFloat(value(Text), 64); err == nil {
					price, _ = calc.Marjin(price, 0.3, true)
					price, _ = calc.Vat(price, 0.2, false)
					Text = fmt.Sprintf("%f", price)
				}
			case strings.HasPrefix(Text, "/mmpn "):
				if price, err := strconv.ParseFloat(value(Text), 64); err == nil {
					price, _ = calc.Marjin(price, 0.3, false)
					price, _ = calc.Vat(price, 0.2, true)
					Text = fmt.Sprintf("%f", price)
				}
			case strings.HasPrefix(Text, "/mmmn "):
				if price, err := strconv.ParseFloat(value(Text), 64); err == nil {
					price, _ = calc.Marjin(price, 0.3, false)
					price, _ = calc.Vat(price, 0.2, false)
					Text = fmt.Sprintf("%f", price)
				}
			case Text == "/testsql ":
				sql.Test()
			case Text == "/testrep ":
				rep.Test()
			case Text == "/testcalc ":
				calc.Test()
			case Text == "bb": //стоп бот
				botStop = true
			default:
				Text = "Неизвестная команда. Для помощи /help"
			}

			// Ответим пользователю его же сообщением
			reply := Text
			// Созадаем сообщение
			msg := tgbotapi.NewMessage(ChatID, reply)
			// и отправляем его
			bot.Send(msg)

		}

	}
}

func value(text string) (val string) {
	if i := strings.IndexByte(text, ' '); i >= 0 {
		s := text[(i + 1):]
		return s
	} else {
		return "error getting value"
	}
}
