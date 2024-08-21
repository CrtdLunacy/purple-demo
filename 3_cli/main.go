package main

import (
	"3_cli/api"
	"3_cli/handlers"
	"flag"
	"fmt"

	"github.com/joho/godotenv"
)

func promtData(promt string) string {
	var res string
	fmt.Println(promt + ": ")
	fmt.Scan(&res)

	return res
}

func main() {
	createBin := flag.Bool("create", false, "Создание Bin")
	updateBin := flag.Bool("update", false, "Обновление Bin")
	deleteBin := flag.Bool("delete", false, "Удаление Bin")
	getBin := flag.Bool("get", false, "Получение Bin")
	getList := flag.Bool("list", false, "Получение списка Bin")

	fileFlag := flag.String("file", "", "Название файла")
	nameOfBin := flag.String("name", "", "Название Bin")
	binID := flag.String("id", "", "ID Bin")

	flag.Parse()

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Файл окружения не загружен")
	}
	apiClient := api.NewApi()

	handlers.HandleFlags(apiClient, *createBin, *updateBin, *deleteBin, *getBin, *getList, *fileFlag, *nameOfBin, *binID)
	// storage := storage.NewStorage(files.NewJsonDb("data.vault"), *encrypter.NewEncrypter())
	// binName := promtData("Введите название файла")
	// myBin, err := bins.NewBin(rand.Int32(), binName, true)
	// if err != nil {
	// 	fmt.Println("Неверный формат названия файла")
	// }
	// doc, err := myBin.BinToBytes()
	// if err != nil {
	// 	fmt.Println("Не удалось преобразовать в JSON")
	// }

	// storage.SaveToStorage(doc)
}
