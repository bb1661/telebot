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
