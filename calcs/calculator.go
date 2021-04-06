package calculator

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func Test() {
	fmt.Println("Done")
}

func NeedCalc(command string, vat float64, marjin float64) string {
	var msg string

	var mValue float64
	var VatValue float64
	var text string
	command = strings.ReplaceAll(command, ",", ".")
	switch {

	case strings.HasPrefix(command, "/calculator Добавить маржу, добавить НДС "):
		command = strings.ReplaceAll(command, "/calculator Добавить маржу, убрать НДС ", "")
		if command, err := strconv.ParseFloat(command, 64); err == nil {

			command, mValue = Marjin(command, 0.3, true)

			command, VatValue = Vat(command, 0.2, true)
			msg = fmt.Sprintf("%f", math.Ceil(command*100)/100)
		}
	case strings.HasPrefix(command, "/calculator Добавить маржу, убрать НДС "):
		command = strings.ReplaceAll(command, "/calculator Добавить маржу, убрать НДС ", "")
		if command, err := strconv.ParseFloat(command, 64); err == nil {
			command, mValue = Marjin(command, 0.3, true)
			command, VatValue = Vat(command, 0.2, false)
			msg = fmt.Sprintf("%f", math.Ceil(command*100)/100)
		}
	case strings.HasPrefix(command, "/calculator Убрать маржу, добавить НДС "):
		command = strings.ReplaceAll(command, "/calculator Убрать маржу, добавить НДС ", "")
		if command, err := strconv.ParseFloat(command, 64); err == nil {
			command, mValue = Marjin(command, 0.3, false)
			command, VatValue = Vat(command, 0.2, true)
			msg = fmt.Sprintf("%f", math.Ceil(command*100)/100)
		}
	case strings.HasPrefix(command, "/calculator Убрать маржу, убрать НДС "):
		command = strings.ReplaceAll(command, "/calculator Убрать маржу, убрать НДС ", "")
		if command, err := strconv.ParseFloat(command, 64); err == nil {
			command, mValue = Marjin(command, 0.3, false)
			command, VatValue = Vat(command, 0.2, false)
			msg = fmt.Sprintf("%f", math.Ceil(command*100)/100)
		}
	case strings.HasPrefix(command, "/calculator Добавить маржу "):
		command = strings.ReplaceAll(command, "/calculator Добавить маржу ", "")

		if command, err := strconv.ParseFloat(command, 64); err == nil {
			command, mValue = Marjin(command, 0.3, true)
			msg = fmt.Sprintf("%f", math.Ceil(command*100)/100)
		}
	case strings.HasPrefix(command, "/calculator Убрать маржу "):
		command = strings.ReplaceAll(command, "/calculator Убрать маржу ", "")
		if command, err := strconv.ParseFloat(command, 64); err == nil {
			command, mValue = Marjin(command, 0.3, false)
			msg = fmt.Sprintf("%f", math.Ceil(command*100)/100)
		}
	case strings.HasPrefix(command, "/calculator Добавить НДС "):
		command = strings.ReplaceAll(command, "/calculator Добавить НДС ", "")
		if command, err := strconv.ParseFloat(command, 64); err == nil {
			command, VatValue = Vat(command, 0.2, true)
			msg = fmt.Sprintf("%f", math.Ceil(command*100)/100)
		}
	case strings.HasPrefix(command, "/calculator Убрать НДС "):
		command = strings.ReplaceAll(command, "/calculator Убрать НДС ", "")
		if command, err := strconv.ParseFloat(command, 64); err == nil {
			command, VatValue = Vat(command, 0.2, false)
			msg = fmt.Sprintf("%f", math.Ceil(command*100)/100)
		}

	default:
		msg = "Команда с ошибкой"
	}

	if VatValue != 0 {
		text = text + " В том числе НДС: " + fmt.Sprintf("%f", math.Ceil(VatValue*100)/100) + "\n"
	}
	if mValue != 0 {
		text = text + " В том числе маржа: " + fmt.Sprintf("%f", math.Ceil(mValue*100)/100) + "\n"
	}
	return ("Итоговая цена: " + msg + text)

}

func Vat(command float64, VatValue float64, plusMinus bool) (float64, float64) { //Добавить ндс true, убрать false. Возврат цена с/без ндс , размер ндс
	var VatAmount float64
	if plusMinus {
		VatAmount = command * VatValue
		command = command * (1 + VatValue)

	} else {
		command = command / (1 + VatValue)
		VatAmount = command * VatValue
	}

	return command, VatAmount
}

func Marjin(command float64, MarjinValue float64, plusMinus bool) (float64, float64) { //Добавить Marjin true, убрать false. Возврат цена с/без маржм , размер маржи
	var MarjinAmount float64

	if plusMinus {
		command = command / (1 - MarjinValue)
		MarjinAmount = command * MarjinValue

	} else {

		MarjinAmount = command * MarjinValue
		command = command * (1 - MarjinValue)

	}

	return command, MarjinAmount
}
