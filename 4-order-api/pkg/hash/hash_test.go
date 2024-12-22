package hash

import (
	"strings"
	"testing"
)

func TestHashString(t *testing.T) {
	// Тестовая строка для хэширования
	input := "password123"

	// Вызов функции
	result, err := HashString(input)
	if err != nil {
		t.Fatalf("HashString returned an error: %v", err)
	}

	// Проверка формата результата
	parts := strings.Split(result, ":")
	if len(parts) != 2 {
		t.Fatalf("Result format is incorrect, expected two parts separated by ':', got: %s", result)
	}

	salt := parts[0]
	hash := parts[1]

	// Проверяем, что соль не пустая
	if len(salt) == 0 {
		t.Fatalf("Salt is empty")
	}

	// Проверяем, что хэш не пустой
	if len(hash) == 0 {
		t.Fatalf("Hash is empty")
	}
}
