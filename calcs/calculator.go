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
	switch {

	case strings.HasPrefix("/calculator Добавить маржу, добавить НДС ", command):
		command = strings.Replace(command, "/calculator Добавить маржу, убрать НДС ", "", -1)
		if command, err := strconv.ParseFloat(command, 64); err == nil {
			command, _ = Marjin(command, 0.3, true)
			command, _ = Vat(command, 0.2, true)
			msg = fmt.Sprintf("%f", command)
		}
	case strings.HasPrefix("/calculator Добавить маржу, убрать НДС ", command):
		command = strings.Replace(command, "/calculator Добавить маржу, убрать НДС ", "", -1)
		if command, err := strconv.ParseFloat(command, 64); err == nil {
			command, _ = Marjin(command, 0.3, true)
			command, _ = Vat(command, 0.2, false)
			msg = fmt.Sprintf("%f", command)
		}
	case strings.HasPrefix("/calculator Убрать маржу, добавить НДС ", command):
		command = strings.Replace(command, "/calculator Убрать маржу, добавить НДС ", "", -1)
		if command, err := strconv.ParseFloat(command, 64); err == nil {
			command, _ = Marjin(command, 0.3, false)
			command, _ = Vat(command, 0.2, true)
			msg = fmt.Sprintf("%f", command)
		}
	case strings.HasPrefix("/calculator Убрать маржу, убрать НДС ", command):
		command = strings.Replace(command, "/calculator Убрать маржу, убрать НДС ", "", -1)
		if command, err := strconv.ParseFloat(command, 64); err == nil {
			command, _ = Marjin(command, 0.3, false)
			command, _ = Vat(command, 0.2, false)
			msg = fmt.Sprintf("%f", command)
		}
	case strings.HasPrefix("/calculator Добавить маржу ", command):
		command = strings.Replace(command, "/calculator Добавить маржу ", "", -1)

		if command, err := strconv.ParseFloat(command, 64); err == nil {
			command, _ = Marjin(command, 0.3, true)
			msg = fmt.Sprintf("%f", command)
		}
	case strings.HasPrefix("/calculator Убрать маржу ", command):
		command = strings.Replace(command, "/calculator Убрать маржу ", "", -1)
		if command, err := strconv.ParseFloat(command, 64); err == nil {
			command, _ = Marjin(command, 0.3, false)
			msg = fmt.Sprintf("%f", command)
		}
	case strings.HasPrefix("/calculator Добавить НДС ", command):
		command = strings.Replace(command, "/calculator Добавить НДС ", "", -1)
		if command, err := strconv.ParseFloat(command, 64); err == nil {
			command, _ = Vat(command, 0.2, true)
			msg = fmt.Sprintf("%f", command)
		}
	case strings.HasPrefix("/calculator Убрать НДС ", command):
		command = strings.Replace(command, "/calculator Убрать НДС ", "", -1)
		if command, err := strconv.ParseFloat(command, 64); err == nil {
			command, _ = Vat(command, 0.2, false)
			msg = fmt.Sprintf("%f", command)
		}
	}
	return msg

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

	return math.Ceil(command*100) / 100, math.Ceil(VatAmount*100) / 100
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

	return math.Ceil(command*100) / 100, math.Ceil(MarjinAmount*100) / 100
}
