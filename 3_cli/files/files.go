package files

import (
	"encoding/json"
	"fmt"
	"os"
)

type JsonDb struct {
	filename string
}

func NewJsonDb(name string) *JsonDb {
	return &JsonDb{
		filename: name,
	}
}

func (db *JsonDb) Read() ([]byte, error) {
	var parsedData json.RawMessage
	data, err := os.ReadFile(db.filename)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &parsedData)
	if err != nil {
		return nil, err
	}
	return parsedData, nil
}

func (db *JsonDb) Write(content []byte) {
	file, err := os.Create(db.filename)

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
