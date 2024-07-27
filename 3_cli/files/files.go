package files

import (
	"encoding/json"
	"fmt"
	"os"
)

func ReadFile(name string) ([]byte, error) {
	var parsedData json.RawMessage
	data, err := os.ReadFile(name)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &parsedData)
	if err != nil {
		return nil, err
	}
	return parsedData, nil
}

func WriteFile(content []byte, fileName string) {
	file, err := os.Create(fileName)

	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	_, err = file.Write(content)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Запись успешна")
}
