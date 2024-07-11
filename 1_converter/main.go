package main

import (
	"fmt"
)

func main() {
	usdToEur := 0.92
	usdToRub := 87.0
	eurToRub := usdToRub / usdToEur
	fmt.Print(eurToRub)
}
