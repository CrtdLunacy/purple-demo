package storage

import "3_cli/files"

func SaveToStorage(file []byte) {
	files.WriteFile(file, "data.json")
}

func ReadFromStorage() {
	files.ReadFile("data.json")
}
