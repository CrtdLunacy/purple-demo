package storage

import (
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
}

func NewStorage(db Db) *Storage {
	file, err := db.Read()
	if err != nil {
		return &Storage{
			UpdatedAt: time.Now(),
			db:        db,
		}
	}

	var storage Storage
	err = json.Unmarshal(file, &storage)
	if err != nil {
		fmt.Println("Не удалось обработать файл базы данных")
		return &Storage{
			UpdatedAt: time.Now(),
			db:        db,
		}
	}

	return &Storage{
		db: db,
	}
}

func (storage *Storage) SaveToStorage(file []byte) {
	storage.db.Write(file)
}

func (storage *Storage) ReadFromStorage() {
	storage.db.Read()
}
