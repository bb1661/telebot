package calculator

import (
	"fmt"
	"math"
)

func TTT() {
	fmt.Println("Done")
}

func main() {
	var price float64
	var plusMinusVat bool = true
	var plusMinusMarjin bool = true
	var calcVariables map[string]float64 = map[string]float64{
		"Vat":    0.2,
		"Marjin": 0.3,
	}

	//fmt.Println(Vat(price, plusMinusVat))
	price = 100
	priceAfterVat, _ := Vat(price, calcVariables["Vat"], plusMinusVat)
	_, VatAmount := Vat(price, calcVariables["Vat"], plusMinusVat)

	priceAfterMarjin, MarjinAmount := Marjin(price, calcVariables["Marjin"], plusMinusMarjin)

	if plusMinusVat {
		fmt.Println("Цена с НДС: ", priceAfterVat)
		fmt.Println("Размер НДС: ", VatAmount)
	} else {
		fmt.Println("Цена без НДС: ", priceAfterVat)
		fmt.Println("Размер НДС: ", VatAmount)
	}

	if plusMinusMarjin {
		fmt.Println("Цена с маржой: ", priceAfterMarjin)
		fmt.Println("Размер маржи: ", MarjinAmount)
	} else {
		fmt.Println("Цена без маржи: ", priceAfterMarjin)
		fmt.Println("Размер маржи: ", MarjinAmount)
	}

	//fmt.Println("Vat")
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
