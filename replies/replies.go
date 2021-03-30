package replies

import (
	"fmt"
)

func Test() {
	fmt.Println("Done")
}

func Help() (msg string) {
	return `Доступные команды: 
	/login
	/help
	/calculator`
}

func CalculatorInit() (msg string) {
	return `Можете посчитать цену:
	/pmpn цена - добавить маржу, добавить НДС 
	/pmmn цена - добавить маржу, убрать НДС
	/mmpn цена - убрать маржу, добавить НДС 
	/mmmn цена - убрать маржу, убрать НДС`
}
