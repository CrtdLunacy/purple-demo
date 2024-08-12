package config

import "os"

func ReadFromEnv(key string) string {
	envDataByKey := os.Getenv(key)
	if key == "" {
		panic("Ключ не может быть пустым значением")
	}
	return envDataByKey
}
