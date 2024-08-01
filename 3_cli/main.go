package main

import (
	"3_cli/bins"
	"3_cli/files"
	"3_cli/storage"
	"fmt"
	"math/rand/v2"
)

func promtData(promt string) string {
	var res string
	fmt.Println(promt + ": ")
	fmt.Scan(&res)

	return res
}

func main() {
	storage := storage.NewStorage(files.NewJsonDb("data.json"))
	binName := promtData("Введите название файла")
	myBin, err := bins.NewBin(rand.Int32(), binName, true)
	if err != nil {
		fmt.Println("Неверный формат названия файла")
	}
	doc, err := myBin.BinToBytes()
	if err != nil {
		fmt.Println("Не удалось преобразовать в JSON")
	}

	storage.SaveToStorage(doc)
}
