package handlers

import (
	"3_cli/api"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

type Entry struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

func HandleFlags(apiClient *api.Api, createBin, updateBin, deleteBin, getBin, getList bool, fileFlag, nameOfBin, binID string) {
	switch {
	case createBin:
		if fileFlag == "" || nameOfBin == "" {
			log.Fatal("Пожалуйста, укажите --file и --name для создания Bin")
		}

		CreateBinHandler(apiClient, fileFlag, nameOfBin)

	case updateBin:
		if fileFlag == "" || binID == "" {
			log.Fatal("Пожалуйста, укажите --file и --id для обновления Bin")
		}

		UpdateBinHandler(apiClient, fileFlag, binID)

	case deleteBin:
		if binID == "" {
			log.Fatal("Пожалуйста, укажите --id для удаления Bin")
		}

		DeleteBinHandler(apiClient, binID)

	case getBin:
		if binID == "" {
			log.Fatal("Пожалуйста, укажите --id для получения Bin")
		}

		GetBinHandler(apiClient, binID)

	case getList:
		ListBins("local.json")

	default:
		log.Fatal("Необходимо указать один из флагов: --create, --update, --delete, --get, --list")
	}
}

func CreateBinHandler(apiClient *api.Api, filePath, binName string) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Ошибка при чтении файла: %v", err)
	}

	response, err := apiClient.PostData("https://api.jsonbin.io/v3/b", binName, json.RawMessage(data))
	if err != nil {
		log.Fatalf("Ошибка при создании Bin: %v", err)
	}

	var result map[string]interface{}
	if err := json.Unmarshal(response, &result); err != nil {
		log.Fatalf("Ошибка при разборе ответа: %v", err)
	}

	metadata, ok := result["metadata"].(map[string]interface{})
	if !ok {
		log.Fatalf("Ответ не содержит корректного 'metadata': %v", result)
	}

	binID, ok := metadata["id"].(string)
	if !ok {
		log.Fatalf("Ответ не содержит корректного 'id' в 'metadata': %v", metadata)
	}

	saveBinToFile(binName, binID)
	fmt.Printf("Bin создан с ID: %s\n", binID)
}

func UpdateBinHandler(apiClient *api.Api, filePath, binID string) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Ошибка при чтении файла: %v", err)
	}

	url := fmt.Sprintf("https://api.jsonbin.io/v3/b/%s", binID)

	_, err = apiClient.PutData(url, json.RawMessage(data))
	if err != nil {
		log.Fatalf("Ошибка при обновлении Bin: %v", err)
	}

	fmt.Printf("Bin с ID %s успешно обновлен\n", binID)
}

func DeleteBinHandler(apiClient *api.Api, binID string) {
	url := fmt.Sprintf("https://api.jsonbin.io/v3/b/%s", binID)

	_, err := apiClient.DeleteData(url)
	if err != nil {
		log.Fatalf("Ошибка при удалении Bin: %v", err)
	}

	if err := removeBinFromFile(binID); err != nil {
		log.Fatalf("Ошибка при удалении записи из файла: %v", err)
	}

	fmt.Printf("Bin с ID %s успешно удален\n", binID)
}

func GetBinHandler(apiClient *api.Api, binID string) {
	url := fmt.Sprintf("https://api.jsonbin.io/v3/b/%s", binID)

	response, err := apiClient.GetData(url)
	if err != nil {
		log.Fatalf("Ошибка при получении Bin: %v", err)
	}

	fmt.Printf("Данные Bin: %s\n", response)
}

func ListBins(filePath string) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Ошибка при чтении файла: %v", err)
	}

	fmt.Printf("Список Bin:\n%s\n", data)
}

func saveBinToFile(binName, binID string) {

	filePath := "local.json"

	if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
		log.Fatalf("Ошибка при создании директории: %v", err)
	}

	var data map[string]Entry
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		data = make(map[string]Entry)
	} else {
		file, err := os.Open(filePath)
		if err != nil {
			log.Fatalf("Ошибка при открытии файла: %v", err)
		}
		defer file.Close()

		decoder := json.NewDecoder(file)
		if err := decoder.Decode(&data); err != nil && err.Error() != "EOF" {
			log.Fatalf("Ошибка при чтении файла: %v", err)
		}
	}

	// Обновляем данные
	if data == nil {
		data = make(map[string]Entry)
	}
	data[binID] = Entry{
		Name: binName,
		ID:   binID,
	}

	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatalf("Ошибка при открытии файла для записи: %v", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(data); err != nil {
		log.Fatalf("Ошибка при записи в файл: %v", err)
	}

	fmt.Printf("Bin сохранен в файл %s\n", filePath)
}

func removeBinFromFile(binID string) error {
	filePath := "local.json"

	var data map[string]Entry
	file, err := os.Open(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&data); err != nil && err.Error() != "EOF" {
		return fmt.Errorf("ошибка при чтении файла: %w", err)
	}

	delete(data, binID)

	file, err = os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(data); err != nil {
		return fmt.Errorf("ошибка при записи в файл: %w", err)
	}

	return nil
}
