package keyboards

import (
	"fmt"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

var NumericKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Тестовый"),
		tgbotapi.NewKeyboardButton("Кейборд"),
	),
)

func Test() {
	fmt.Println("Done")
}

var CalcKb = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Добавить маржу"),
		tgbotapi.NewKeyboardButton("Убрать маржу"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Добавить НДС"),
		tgbotapi.NewKeyboardButton("Убрать НДС"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Добавить маржу, добавить НДС"),
		tgbotapi.NewKeyboardButton("Добавить маржу, убрать НДС"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Убрать маржу, добавить НДС"),
		tgbotapi.NewKeyboardButton("Убрать маржу, убрать НДС"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Главное меню"),
		tgbotapi.NewKeyboardButton("Настройки"),
	),
)

var InlineKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonURL("1.com", "http://1.com"),
		tgbotapi.NewInlineKeyboardButtonSwitch("2sw", "open 2"),
		tgbotapi.NewInlineKeyboardButtonData("3", "3"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("4", "4"),
		tgbotapi.NewInlineKeyboardButtonData("5", "5"),
		tgbotapi.NewInlineKeyboardButtonData("6", "6"),
	),
)
