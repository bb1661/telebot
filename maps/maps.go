package maps

import (
	"fmt"
)

func Test() {
	fmt.Println("Done")
}

var CommandLeveling map[string]bool = map[string]bool{
	"Добавить маржу, добавить НДС":             true,
	"Добавить маржу, убрать НДС":               true,
	"Убрать маржу, добавить НДС":               true,
	"Убрать маржу, убрать НДС":                 true,
	"Добавить маржу":                           true,
	"Убрать маржу":                             true,
	"Добавить НДС":                             true,
	"Убрать НДС":                               true,
	"Настройки":                                true,
	"/calculator Добавить маржу, добавить НДС": true,
	"/calculator Добавить маржу, убрать НДС":   true,
	"/calculator Убрать маржу, добавить НДС":   true,
	"/calculator Убрать маржу, убрать НДС":     true,
	"/calculator Добавить маржу":               true,
	"/calculator Убрать маржу":                 true,
	"/calculator Добавить НДС":                 true,
	"/calculator Убрать НДС":                   true,
	"/calculator":                              true,
}

var Marks map[string]string = map[string]string{
	"/calculator":  "calc",
	"/calculator ": "calc2",
	"/calculator Добавить маржу, добавить НДС": "calc3",
	"/calculator Добавить маржу, убрать НДС":   "calc3",
	"/calculator Убрать маржу, добавить НДС":   "calc3",
	"/calculator Убрать маржу, убрать НДС":     "calc3",
	"/calculator Добавить маржу":               "calc3",
	"/calculator Убрать маржу":                 "calc3",
	"/calculator Добавить НДС":                 "calc3",
	"/calculator Убрать НДС":                   "calc3",
}
