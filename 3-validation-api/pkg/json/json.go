package json

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
)

func JSONWritter[T any](data T) error {
	dir := "./"
	filename := "data.json"
	filePath := filepath.Join(dir, filename)

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = file.Write(jsonData)
	if err != nil {
		return err
	}

	log.Println("Данные успешно записаны в файл:", filePath)
	return nil
}
