package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func find(slice []string, targetItem string) bool {
	for _, value := range slice {
		if value == targetItem {
			return true
		}
	}
	return false
}

func average(numbers []float64) float64 {
	sum := 0.0
	for _, num := range numbers {
		sum += num
	}
	return sum / float64(len(numbers))
}

func sum(numbers []float64) float64 {
	sum := 0.0
	for _, num := range numbers {
		sum += num
	}
	return sum
}

func median(numbers []float64) float64 {
	sort.Float64s(numbers)
	length := len(numbers)
	if length == 0 {
		return 0
	}
	if length%2 == 1 {
		return numbers[length/2]
	}
	return (numbers[length/2-1] + numbers[length/2]) / 2.0
}

func choiseOperation() string {
	operationsList := []string{"AVG", "SUM", "MED"}
	var selectedOperation string

	for {
		fmt.Printf("Выберите операцию, которую необходимо выполнить %s: ", operationsList)
		fmt.Scan(&selectedOperation)
		if find(operationsList, selectedOperation) {
			break
		}
		selectedOperation = ""
		fmt.Println("Такая операция недоступна. Пожалуйста, попробуйте снова.")
	}

	return selectedOperation
}

func getOperationDataList() []float64 {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Укажите через запятую любое количество чисел, с которыми хотите выполнить операцию: ")
		operationDataList, _ := reader.ReadString('\n')
		trimmedData := strings.TrimSpace(operationDataList)

		dataArray := strings.Split(trimmedData, ",")
		numbers := make([]float64, 0, len(dataArray))

		for _, str := range dataArray {
			trimmedStr := strings.TrimSpace(str)
			num, err := strconv.ParseFloat(trimmedStr, 64)
			if err != nil {
				fmt.Println("Ошибка парсинга числа:", err)
				continue
			}
			numbers = append(numbers, num)
		}

		if len(numbers) > 0 {
			return numbers
		}

		fmt.Println("Такая операция недоступна. Пожалуйста, попробуйте снова.")
	}
}

func getResultBySelectedOperation(operation string, operationData []float64) float64 {
	operationsMap := map[string]func([]float64) float64{
		"AVG": average,
		"SUM": sum,
		"MED": median,
	}

	if operationFunc, exists := operationsMap[operation]; exists {
		return operationFunc(operationData)
	}
	return 0
}

func calculation() {
	var isOn string

	for {
		selectedOperation := choiseOperation()

		operationDataList := getOperationDataList()

		result := getResultBySelectedOperation(selectedOperation, operationDataList)

		fmt.Println("Результат = ", result)

		fmt.Println("Продолжим расчеты? (y/n): ")
		fmt.Scan(&isOn)

		if isOn != "y" || isOn == "Y" {
			break
		}
	}
}

func main() {
	calculation()
}
