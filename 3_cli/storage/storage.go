package storage

import (
	"3_cli/encrypter"
	"encoding/json"
	"fmt"
	"time"
)

type Db interface {
	Read() ([]byte, error)
	Write([]byte)
}

type Storage struct {
	UpdatedAt time.Time
	db        Db
	enc       encrypter.Encrypter
}

func NewStorage(db Db, enc encrypter.Encrypter) *Storage {
	file, err := db.Read()
	if err != nil {
		return &Storage{
			UpdatedAt: time.Now(),
			db:        db,
			enc:       enc,
		}
	}

	data := enc.Decrypt(file)
	var storage Storage
	err = json.Unmarshal(data, &storage)
	fmt.Printf("data is %s", storage.db)
	if err != nil {
		fmt.Println("Не удалось обработать файл базы данных")
		return &Storage{
			UpdatedAt: time.Now(),
			db:        db,
			enc:       enc,
		}
	}

	return &Storage{
		db:  db,
		enc: enc,
	}
}

func (storage *Storage) SaveToStorage(file []byte) {
	encData := storage.enc.Encrypt(file)
	storage.db.Write(encData)
}

func (storage *Storage) ReadFromStorage() {
	storage.db.Read()
}
