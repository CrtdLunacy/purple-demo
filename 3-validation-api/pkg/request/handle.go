package request

import (
	"go/validation-api/pkg/response"
	"net/http"
)

func HandleBody[T any](w http.ResponseWriter, r *http.Request) (*T, error) {
	body, err := Decode[T](r.Body)

	if err != nil {
		response.ResponseJSON(w, err.Error(), 400)
		return nil, err
	}

	if err = IsValid(body); err != nil {
		response.ResponseJSON(w, err.Error(), 400)
		return nil, err
	}

	return &body, nil
}
