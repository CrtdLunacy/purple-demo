package main

import (
	"fmt"
)

func getCurrenciesList(currencies []string, initCurrency string) []string {
	availableCurrencies := make([]string, 0, len(currencies)-1)
	for _, currency := range currencies {
		if currency != initCurrency {
			availableCurrencies = append(availableCurrencies, currency)
		}
	}

	return availableCurrencies
}

func getExchangeRate(fromCurrency, toCurrency string) float64 {
	rates := map[string]map[string]float64{
		"USD": {
			"EUR": 0.85,
			"RUB": 74.0,
		},
		"EUR": {
			"USD": 1 / 0.85,
			"RUB": 87.0,
		},
		"RUB": {
			"USD": 1 / 74.0,
			"EUR": 1 / 87.0,
		},
	}
	return rates[fromCurrency][toCurrency]
}

func exchangeCurrency(initValue string, targetValue string, sum float64) float64 {
	exchangeRate := getExchangeRate(initValue, targetValue)
	targetSum := sum * exchangeRate
	fmt.Printf("Конвертация: %.2f %s = %.2f %s\n", sum, initValue, targetSum, targetValue)

	return targetSum
}

func currencyConverter() {
	var initCurrency string
	var targetCurrency string
	var sumForConvertation float64
	var isRepeatEchange string
	currenciesList := []string{"USD", "RUB", "EUR"}

	for {
		for {
			availableCurrencies := getCurrenciesList(currenciesList, initCurrency)
			fmt.Printf("Введите валюту, из которой хотите конвертировать %s: ", availableCurrencies)
			fmt.Scan(&initCurrency)
			if initCurrency == "USD" || initCurrency == "RUB" || initCurrency == "EUR" {
				break
			}
			initCurrency = ""
			fmt.Println("Неверная валюта. Пожалуйста, попробуйте снова.")
		}

	mainLoop:
		for {
			fmt.Println("Введите сумму, которую хотите рассчитать: ")
			_, err := fmt.Scan(&sumForConvertation)

			if err != nil {
				fmt.Println("Неверный ввод. Пожалуйста, введите числовое значение.")
				var discard string
				fmt.Scanln(&discard)
				continue
			}

			switch {
			case sumForConvertation < 0:
				fmt.Println("Сумма должна быть больше 0")
			case sumForConvertation > 1000000000:
				fmt.Println("Максимально доступная сумма для рассчетов = 1000000000, укажите меньшую сумму")
			default:
				break mainLoop
			}
		}

		for {
			availableCurrencies := getCurrenciesList(currenciesList, initCurrency)
			fmt.Printf("Введите валюту, в которую хотите конвертировать %s: ", availableCurrencies)
			fmt.Scan(&targetCurrency)
			if targetCurrency == "USD" || targetCurrency == "RUB" || targetCurrency == "EUR" && targetCurrency != initCurrency {
				break
			}
			fmt.Println("Неверная валюта. Пожалуйста, попробуйте снова.")
		}

		exchangeCurrency(initCurrency, targetCurrency, sumForConvertation)

		fmt.Println("Хотите повторить рассчеты (y/n)?: ")
		fmt.Scan(&isRepeatEchange)

		if isRepeatEchange == "N" || isRepeatEchange == "n" {
			break
		}
		initCurrency = ""
	}
}

func main() {
	currencyConverter()
}
