package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	calc "telebot2/calcs"
	kb "telebot2/keyboards"
	rp "telebot2/replies"
	sql "telebot2/sql"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

var msg string   // Сообщение для отправки
var ptext string // previous message - предыдущая полученная команда

func main() {
	botStop := false
	calc.Test()
	sql.Test()
	rp.Test()

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
			reply := ""

			// ID чата/диалога.
			// Может быть идентификатором как чата с пользователем
			// (тогда он равен UserID) так и публичного чата/канала
			ChatID := update.Message.Chat.ID
			message := tgbotapi.NewMessage(ChatID, reply)
			//key := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Command())
			// Текст сообщения

			Text := update.Message.Text
			//key.ReplyMarkup = ""
			log.Printf("[%s] %d %s", UserName, ChatID, Text)
			answerMessage := tgbotapi.MessageConfig{}

			switch {
			case ptext == "/calculator":
				switch {
				case strings.HasPrefix(Text, "Добавить маржу, добавить НДС"):
					msg = "Введите цену"
				case strings.HasPrefix(Text, "Добавить маржу, убрать НДС"):
					msg = "Введите цену"
				case strings.HasPrefix(Text, "Убрать маржу, добавить НДС"):
					msg = "Введите цену"
				case strings.HasPrefix(Text, "Убрать маржу, убрать НДС"):
					msg = "Введите цену"
				case strings.HasPrefix(Text, "Добавить маржу"):
					msg = "Введите цену"
				case strings.HasPrefix(Text, "Убрать маржу"):
					msg = "Введите цену"
				case strings.HasPrefix(Text, "Добавить НДС"):
					msg = "Введите цену"
				case strings.HasPrefix(Text, "Убрать НДС"):
					msg = "Введите цену"
				case strings.HasPrefix(Text, "Главное меню"):
					msg = "Возврат в главное меню"
				case strings.HasPrefix(Text, "Настройки"):
					msg = "Возврат в главное меню" // Ввод в базу данных о марже/ндс клиента
				default:
					msg = "Возврат в главное меню"
				}

			case ptext == "Добавить маржу, добавить НДС":
				if price, err := strconv.ParseFloat(value(Text), 64); err == nil {
					price, _ = calc.Marjin(price, 0.3, true)
					price, _ = calc.Vat(price, 0.2, true)
					msg = fmt.Sprintf("%f", price)
				}
			case ptext == "Добавить маржу, убрать НДС":
				if price, err := strconv.ParseFloat(value(Text), 64); err == nil {
					price, _ = calc.Marjin(price, 0.3, true)
					price, _ = calc.Vat(price, 0.2, false)
					msg = fmt.Sprintf("%f", price)
				}
			case ptext == "Убрать маржу, добавить НДС":
				if price, err := strconv.ParseFloat(value(Text), 64); err == nil {
					price, _ = calc.Marjin(price, 0.3, false)
					price, _ = calc.Vat(price, 0.2, true)
					msg = fmt.Sprintf("%f", price)
				}
			case ptext == "Убрать маржу, убрать НДС":
				if price, err := strconv.ParseFloat(value(Text), 64); err == nil {
					price, _ = calc.Marjin(price, 0.3, false)
					price, _ = calc.Vat(price, 0.2, false)
					msg = fmt.Sprintf("%f", price)
				}
			case ptext == "Добавить маржу":
				if price, err := strconv.ParseFloat(value(Text), 64); err == nil {
					price, _ = calc.Marjin(price, 0.3, true)
					msg = fmt.Sprintf("%f", price)
				}
			case ptext == "Убрать маржу":
				if price, err := strconv.ParseFloat(value(Text), 64); err == nil {
					price, _ = calc.Marjin(price, 0.3, false)
					msg = fmt.Sprintf("%f", price)
				}
			case ptext == "Добавить НДС":
				if price, err := strconv.ParseFloat(value(Text), 64); err == nil {
					price, _ = calc.Vat(price, 0.2, true)
					msg = fmt.Sprintf("%f", price)
				}
			case ptext == "Убрать НДС":
				if price, err := strconv.ParseFloat(value(Text), 64); err == nil {
					price, _ = calc.Vat(price, 0.2, false)
					msg = fmt.Sprintf("%f", price)
				}

			case Text == "/help":
				msg = rp.Help()

			case Text == "/calculator":
				msg = rp.CalculatorInit()
				answerMessage.ReplyMarkup = kb.CalcKb

			case Text == "/testsql ":
				sql.Test()
			case Text == "/testrep ":
				rp.Test()
			case Text == "/testcalc ":
				calc.Test()
			case Text == "/testkeyboards ":
				kb.Test()
			case Text == "/testbuttons":
				answerMessage.ReplyMarkup = kb.NumericKeyboard
			case Text == "bb": //стоп бот
				botStop = true
				msg = "Бот остановлен, бб"
			default:
				msg = "Неизвестная команда. Для помощи /help"
			}
			//bot.Send(tgbotapi.NewMessage(ChatID, "-----"))
			ptext = Text
			reply = msg
			// Созадаем сообщение
			message = tgbotapi.NewMessage(ChatID, reply)
			message.ReplyMarkup = answerMessage.ReplyMarkup
			bot.Send(message)

			//message := tgbotapi.NewMessage(ChatID, reply)
			// и отправляем его
			//bot.Send(key)

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
