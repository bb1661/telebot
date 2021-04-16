package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	calc "telebot2/calcs"
	kb "telebot2/keyboards"
	maps "telebot2/maps"
	rp "telebot2/replies"
	sql "telebot2/sql"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

var msg string   // Сообщение для отправки
var ptext string // previous message - предыдущая полученная команда

func main() {
	//Тесты подключения различных модулей + неудаление того, что пригодится
	botStop := false
	calc.Test()
	sql.Test()
	rp.Test()
	maps.Test()
	tag := "menu"
	var command string

	err := sql.SqlCon()
	if err != nil {
		log.Panic(err)
	}
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

			// Текст сообщения
			Text := update.Message.Text

			/*
				   paramStr := ""
					f := map[string]fn{
						"/login":      foo_login,
						"/help":       foo_help,
						"/calculator": foo_calculator,
						"/calculator/addMarjin": calc.Marjin(),
					}

					f[Text]()

			*/

			//Сборка команды с путем
			log.Printf("[%s] %d %s", UserName, ChatID, Text)
			addKb := tgbotapi.MessageConfig{}

			if maps.CommandLeveling[command] {
				command = command + " " + Text
			} else {
				command = Text
			}

			//Вывод команды
			message = tgbotapi.NewMessage(ChatID, "| Введена команда: "+command+" | Tag: "+tag)
			bot.Send(message)

			switch {
			/*
				добавить тэги и описать их
				в первую очередь свитч должен проходить по маркам.

			*/
			case tag == "calc/needCalc":
				msg = calc.NeedCalc(command, 0.2, 0.3)
				tag = "menu"
			case tag == "calc/needPrice":
				msg = "Введите цену"
				tag = "calc/needCalc"
				addKb.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
			case tag == "login/pasteSQL":

				_, err = sql.CreateUser("NameTest", "LastNameTest", command, ChatID)
				if err != nil {
					log.Println("Error creating user: ", err.Error())
				}
				msg = "user created"
				tag = "menu"
			default:

				switch {

				case Text == "/help":
					msg = rp.Help()
				case command == "/calculator":
					msg = rp.CalculatorInit()
					addKb.ReplyMarkup = kb.CalcKb
					tag = "calc/needPrice"
				case command == "/testsql":
					sql.Test()
				case command == "/testrep":
					rp.Test()
				case command == "/testcalc":
					calc.Test()
				case command == "/testkeyboards":
					kb.Test()
				case command == "/testkeyboards2":
					addKb.ReplyMarkup = kb.InlineKeyboard
					msg = "тест кб"
				case command == "/testbuttons":
					addKb.ReplyMarkup = kb.NumericKeyboard
				case command == "":

				case command == "bb": //стоп бот
					botStop = true
					msg = "Бот остановлен, бб"
				case command == "/login":
					userExists, _ := sql.CheckUser(ChatID)
					if userExists == 0 {
						msg = "Необходим логин. Введите почту."
						tag = "login/pasteSQL"
					} else {
						msg = "Пользователь существует"
						tag = "menu"
					}

				default:
					tag = "menu"
					msg = "Неизвестная команда. Для помощи /help"
				}

			}

			ptext = Text
			reply = msg
			// Созадаем сообщение
			message = tgbotapi.NewMessage(ChatID, reply)
			message.ReplyMarkup = addKb.ReplyMarkup
			bot.Send(message)

		}

	}
}

func old_main() {
	botStop := false
	calc.Test()
	sql.Test()
	rp.Test()
	maps.Test()
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
			addKb := tgbotapi.MessageConfig{}

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
				message = tgbotapi.NewMessage(ChatID, "here")
				bot.Send(message)
				if price, err := strconv.ParseFloat(Text, 64); err == nil {
					price, _ = calc.Marjin(price, 0.3, true)
					price, _ = calc.Vat(price, 0.2, true)
					msg = fmt.Sprintf("%f", price)

				}
			case ptext == "Добавить маржу, убрать НДС":
				if price, err := strconv.ParseFloat(Text, 64); err == nil {
					price, _ = calc.Marjin(price, 0.3, true)
					price, _ = calc.Vat(price, 0.2, false)
					msg = fmt.Sprintf("%f", price)
				}
			case ptext == "Убрать маржу, добавить НДС":
				if price, err := strconv.ParseFloat(Text, 64); err == nil {
					price, _ = calc.Marjin(price, 0.3, false)
					price, _ = calc.Vat(price, 0.2, true)
					msg = fmt.Sprintf("%f", price)
				}
			case ptext == "Убрать маржу, убрать НДС":
				if price, err := strconv.ParseFloat(Text, 64); err == nil {
					price, _ = calc.Marjin(price, 0.3, false)
					price, _ = calc.Vat(price, 0.2, false)
					msg = fmt.Sprintf("%f", price)
				}
			case ptext == "Добавить маржу":
				if price, err := strconv.ParseFloat(Text, 64); err == nil {
					price, _ = calc.Marjin(price, 0.3, true)
					msg = fmt.Sprintf("%f", price)
				}
			case ptext == "Убрать маржу":
				if price, err := strconv.ParseFloat(Text, 64); err == nil {
					price, _ = calc.Marjin(price, 0.3, false)
					msg = fmt.Sprintf("%f", price)
				}
			case ptext == "Добавить НДС":
				if price, err := strconv.ParseFloat(Text, 64); err == nil {
					price, _ = calc.Vat(price, 0.2, true)
					msg = fmt.Sprintf("%f", price)
				}
			case ptext == "Убрать НДС":
				if price, err := strconv.ParseFloat(Text, 64); err == nil {
					price, _ = calc.Vat(price, 0.2, false)
					msg = fmt.Sprintf("%f", price)
				}
			case Text == "/help":
				msg = rp.Help()

			case Text == "/calculator":
				msg = rp.CalculatorInit()
				addKb.ReplyMarkup = kb.CalcKb

			case Text == "/testsql ":
				sql.Test()
			case Text == "/testrep ":
				rp.Test()
			case Text == "/testcalc ":
				calc.Test()
			case Text == "/testkeyboards ":
				kb.Test()
			case Text == "/testbuttons":
				addKb.ReplyMarkup = kb.NumericKeyboard
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
			message.ReplyMarkup = addKb.ReplyMarkup
			bot.Send(message)

			//message := tgbotapi.NewMessage(ChatID, reply)
			// и отправляем его
			//bot.Send(key)

		}

	}
}

func main_test_inlinekb() {
	bot, err := tgbotapi.NewBotAPI("MyAwesomeBotToken")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, _ := bot.GetUpdatesChan(u)

	fmt.Print(".")
	for update := range updates {
		if update.CallbackQuery != nil {
			fmt.Print(update)

			bot.AnswerCallbackQuery(tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data))

			bot.Send(tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Data))
		}
		if update.Message != nil {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)

			switch update.Message.Text {
			case "open":
				msg.ReplyMarkup = kb.InlineKeyboard

			}

			bot.Send(msg)
		}
	}
}
