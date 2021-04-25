package main

import (
	// "fmt"
	"log"
	// "strconv"
	// "strings"
	"io/ioutil"
	calc "telebot2/calcs"
	kb "telebot2/keyboards"
	maps "telebot2/maps"
	rp "telebot2/replies"
	sql "telebot2/sql"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
	yaml "gopkg.in/yaml.v3"
)

var msg string // Сообщение для отправки
type Session struct {
	currentTag     string
	currentCommand string
}

type Config struct {
	Token string `yaml:"token"`
	Db    struct {
		DBURL    string `yaml:"dburl"`
		Server   string `yaml:"server"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Database string `yaml:"database"`
	}
}

var s map[int64]*Session

func main() {
	s = make(map[int64]*Session)

	//Config.yaml import
	textfile, err := ioutil.ReadFile("config.yml")
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	config := Config{}
	err3 := yaml.Unmarshal([]byte(textfile), &config)
	if err3 != nil {
		log.Fatalf("error: %v", err)
	}

	//Тесты подключения различных модулей + неудаление того, что пригодится
	botStop := false
	calc.Test()
	sql.Test()
	rp.Test()
	maps.Test()

	err2 := sql.SqlCon(config.Db.Server,
		config.Db.User,
		config.Db.Password,
		config.Db.Port,
		config.Db.Database,
	)

	if err2 != nil {
		log.Panic(err)
	}
	// подключаемся к боту с помощью токена
	bot, err := tgbotapi.NewBotAPI(config.Token)
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

			//Поиск сессии пользователя
			_, ok := s[ChatID]
			if !ok {
				s[ChatID] = &Session{ //CurrentTag currentCommand
					"menu", "",
				}
			}

			//Сборка команды с путем
			log.Printf("[%s] %d %s", UserName, ChatID, Text)
			addKb := tgbotapi.MessageConfig{}

			if maps.CommandLeveling[s[ChatID].currentCommand] {

				s[ChatID].currentCommand = s[ChatID].currentCommand + " " + Text

			} else {
				s[ChatID].currentCommand = Text

			}

			//Вывод команды
			message = tgbotapi.NewMessage(ChatID, "| Введена команда: "+s[ChatID].currentCommand+" | Tag: "+s[ChatID].currentTag)
			bot.Send(message)

			switch {
			/*
				добавить тэги и описать их
				в первую очередь свитч должен проходить по тегам.

			*/
			case s[ChatID].currentTag == "calc/needCalc":
				msg = calc.NeedCalc(s[ChatID].currentCommand, 0.2, 0.3)
				s[ChatID].currentTag = "menu"
			case s[ChatID].currentTag == "calc/needPrice":
				msg = "Введите цену"
				s[ChatID].currentTag = "calc/needCalc"
				addKb.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
			case s[ChatID].currentTag == "login/pasteSQL":

				_, err = sql.CreateUser("NameTest", "LastNameTest", s[ChatID].currentCommand, ChatID)
				if err != nil {
					log.Println("Error creating user: ", err.Error())
				}
				msg = "user created"
				s[ChatID].currentTag = "menu"
			default:

				switch {

				case Text == "/help":
					msg = rp.Help()
				case s[ChatID].currentCommand == "/calculator":
					msg = rp.CalculatorInit()
					addKb.ReplyMarkup = kb.CalcKb
					s[ChatID].currentTag = "calc/needPrice"
				case s[ChatID].currentCommand == "/testsql":
					sql.Test()
				case s[ChatID].currentCommand == "/testrep":
					rp.Test()
				case s[ChatID].currentCommand == "/testcalc":
					calc.Test()
				case s[ChatID].currentCommand == "/testkeyboards":
					kb.Test()
				case s[ChatID].currentCommand == "/testkeyboards2":
					addKb.ReplyMarkup = kb.InlineKeyboard
					msg = "тест кб"
				case s[ChatID].currentCommand == "/testbuttons":
					addKb.ReplyMarkup = kb.NumericKeyboard
				case s[ChatID].currentCommand == "":

				case s[ChatID].currentCommand == "bb": //стоп бот
					botStop = true
					msg = "Бот остановлен, бб"
				case s[ChatID].currentCommand == "/login":
					userExists, _ := sql.CheckUser(ChatID)
					if userExists == 0 {
						msg = "Необходим логин. Введите почту."
						s[ChatID].currentTag = "login/pasteSQL"
					} else {
						msg = "Пользователь существует"
						s[ChatID].currentTag = "menu"
					}

				default:
					s[ChatID].currentTag = "menu"
					msg = "Неизвестная команда. Для помощи /help"
				}

			}

			reply = msg
			// Созадаем сообщение
			message = tgbotapi.NewMessage(ChatID, reply)
			message.ReplyMarkup = addKb.ReplyMarkup
			bot.Send(message)

		}

	}
}
