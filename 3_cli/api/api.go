package api

import "3_cli/config"

type Api struct {
	Key string
}

func NewApi() *Api {
	return &Api{
		Key: config.ReadFromEnv("KEY"),
	}
}
