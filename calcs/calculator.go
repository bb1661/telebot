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

func NeedCalc(price string, command string, vat float64, marjin float64) string {
	var msg string
	switch {

	case strings.HasPrefix("/calculator Добавить маржу, добавить НДС ", command):
		price = strings.Replace(price, "/calculator Добавить маржу, убрать НДС ", "", -1)
		if price, err := strconv.ParseFloat(price, 64); err == nil {
			price, _ = Marjin(price, 0.3, true)
			price, _ = Vat(price, 0.2, true)
			msg = fmt.Sprintf("%f", price)
		}
	case strings.HasPrefix("/calculator Добавить маржу, убрать НДС ", command):
		price = strings.Replace(price, "/calculator Добавить маржу, убрать НДС ", "", -1)
		if price, err := strconv.ParseFloat(price, 64); err == nil {
			price, _ = Marjin(price, 0.3, true)
			price, _ = Vat(price, 0.2, false)
			msg = fmt.Sprintf("%f", price)
		}
	case strings.HasPrefix("/calculator Убрать маржу, добавить НДС ", command):
		price = strings.Replace(price, "/calculator Убрать маржу, добавить НДС ", "", -1)
		if price, err := strconv.ParseFloat(price, 64); err == nil {
			price, _ = Marjin(price, 0.3, false)
			price, _ = Vat(price, 0.2, true)
			msg = fmt.Sprintf("%f", price)
		}
	case strings.HasPrefix("/calculator Убрать маржу, убрать НДС ", command):
		price = strings.Replace(price, "/calculator Убрать маржу, убрать НДС ", "", -1)
		if price, err := strconv.ParseFloat(price, 64); err == nil {
			price, _ = Marjin(price, 0.3, false)
			price, _ = Vat(price, 0.2, false)
			msg = fmt.Sprintf("%f", price)
		}
	case strings.HasPrefix("/calculator Добавить маржу ", command):
		price = strings.Replace(price, "/calculator Добавить маржу ", "", -1)

		if price, err := strconv.ParseFloat(price, 64); err == nil {
			price, _ = Marjin(price, 0.3, true)
			msg = fmt.Sprintf("%f", price)
		}
	case strings.HasPrefix("/calculator Убрать маржу ", command):
		price = strings.Replace(price, "/calculator Убрать маржу ", "", -1)
		if price, err := strconv.ParseFloat(price, 64); err == nil {
			price, _ = Marjin(price, 0.3, false)
			msg = fmt.Sprintf("%f", price)
		}
	case strings.HasPrefix("/calculator Добавить НДС ", command):
		price = strings.Replace(price, "/calculator Добавить НДС ", "", -1)
		if price, err := strconv.ParseFloat(price, 64); err == nil {
			price, _ = Vat(price, 0.2, true)
			msg = fmt.Sprintf("%f", price)
		}
	case strings.HasPrefix("/calculator Убрать НДС ", command):
		price = strings.Replace(price, "/calculator Убрать НДС ", "", -1)
		if price, err := strconv.ParseFloat(price, 64); err == nil {
			price, _ = Vat(price, 0.2, false)
			msg = fmt.Sprintf("%f", price)
		}
	}
	return msg

}

func Vat(price float64, VatValue float64, plusMinus bool) (float64, float64) { //Добавить ндс true, убрать false. Возврат цена с/без ндс , размер ндс
	var VatAmount float64
	if plusMinus {
		VatAmount = price * VatValue
		price = price * (1 + VatValue)

	} else {
		price = price / (1 + VatValue)
		VatAmount = price * VatValue
	}

	return math.Ceil(price*100) / 100, math.Ceil(VatAmount*100) / 100
}

func Marjin(price float64, MarjinValue float64, plusMinus bool) (float64, float64) { //Добавить Marjin true, убрать false. Возврат цена с/без маржм , размер маржи
	var MarjinAmount float64

	if plusMinus {
		price = price / (1 - MarjinValue)
		MarjinAmount = price * MarjinValue

	} else {

		MarjinAmount = price * MarjinValue
		price = price * (1 - MarjinValue)

	}

	return math.Ceil(price*100) / 100, math.Ceil(MarjinAmount*100) / 100
}
