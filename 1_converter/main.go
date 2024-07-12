package main

import (
	"fmt"
)

func main() {
	usdToEur := 0.92
	usdToRub := 87.0
	eurToRub := usdToRub / usdToEur
	fmt.Print(eurToRub)

	initCurrency, targetCurrency, sumForConvertation := userInputData()
	fmt.Print(initCurrency, targetCurrency, sumForConvertation)
}

func userInputData() (string, string, float64) {
	var initCurrency string
	var targetCurrency string
	var sumForConvertation float64
	fmt.Print("Введите валюту, из которой хотите конвертировать (USD, RUB, EURO)")
	fmt.Scan(&initCurrency)
	fmt.Print("Введите валюту, которую хотите получить (USD, RUB, EURO)")
	fmt.Scan(&targetCurrency)
	fmt.Print("Введите сумму, которую хотите рассчитать")
	fmt.Scan(&sumForConvertation)

	return initCurrency, targetCurrency, sumForConvertation
}

func currencyConverter(initValue string, targetValue string, sum float64) {

}
