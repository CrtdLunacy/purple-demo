package api

import (
	"3_cli/config"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Api struct {
	Key string
}

func NewApi() *Api {
	return &Api{
		Key: config.ReadFromEnv("KEY"),
	}
}

func (api *Api) doRequest(method, url, name string, data interface{}) ([]byte, error) {
	client := &http.Client{}

	var req *http.Request
	var err error

	if data != nil {
		jsonData, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}
		req, err = http.NewRequest(method, url, bytes.NewBuffer(jsonData))
		if err != nil {
			return nil, err
		}
	} else {
		req, err = http.NewRequest(method, url, nil)
	}

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Master-Key", api.Key)
	if method == "POST" {
		req.Header.Set("X-Bin-Name", name)
	}

	fmt.Println("Request Headers:")
	for name, values := range req.Header {
		for _, value := range values {
			fmt.Printf("%s: %s\n", name, value)
		}
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (api *Api) GetData(url string) ([]byte, error) {
	return api.doRequest("GET", url, "", nil)
}

func (api *Api) DeleteData(url string) ([]byte, error) {
	return api.doRequest("DELETE", url, "", nil)
}

func (api *Api) PutData(url string, data interface{}) ([]byte, error) {
	return api.doRequest("PUT", url, "", data)
}

func (api *Api) PostData(url, name string, data interface{}) ([]byte, error) {
	return api.doRequest("POST", url, name, data)
}
