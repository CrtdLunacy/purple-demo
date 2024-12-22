package verify

import (
	"encoding/json"
	"fmt"
	"go/validation-api/configs"
	"go/validation-api/pkg/hash"
	myjson "go/validation-api/pkg/json"
	"go/validation-api/pkg/request"
	"go/validation-api/pkg/response"
	"io"
	"log"
	"net/http"
	"net/smtp"
	"os"

	"github.com/jordan-wright/email"
)

type VerifyHandlerDeps struct {
	*configs.Config
}

type VerifyHandler struct {
	*configs.Config
}

type JSONData struct {
	Email string `json:"email"`
	Hash  string `json:"hash"`
}

func NewVerifyHandler(router *http.ServeMux, deps VerifyHandlerDeps) {
	handler := &VerifyHandler{
		Config: deps.Config,
	}
	router.HandleFunc("POST /send", handler.Send())
	router.HandleFunc("GET /verify/{hash}", handler.Verify())
}

func (verifyHandler *VerifyHandler) Send() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		address := "http://localhost:8081/verify"
		body, err := request.HandleBody[SendRequest](w, req)
		if err != nil {
			response.ResponseJSON(w, err.Error(), http.StatusBadRequest)
			return
		}

		hashedEmail := hash.HashString(body.Email)
		data := JSONData{
			Hash:  hashedEmail,
			Email: body.Email,
		}

		err = myjson.JSONWritter(data)
		if err != nil {
			log.Fatalf("Ошибка: %v", err)
		}

		e := email.NewEmail()
		e.From = fmt.Sprintf("%s  <%s>", "Tesoviy Testovich", verifyHandler.Config.EmailSettings.SMTP_EMAIL)
		e.To = []string{body.Email}
		e.Subject = "testovoe pismo"
		e.HTML = []byte(fmt.Sprintf("<h1>Для подтверждения регистрации перейдите по ссылке %s/%s</h1>", address, hashedEmail))

		err = e.Send(fmt.Sprintf("%s:%s", verifyHandler.Config.EmailSettings.SMTP_HOST, verifyHandler.Config.EmailSettings.SMTP_PORT),
			smtp.PlainAuth(
				"",
				verifyHandler.Config.EmailSettings.SMTP_EMAIL,
				verifyHandler.Config.EmailSettings.SMTP_PASSWORD,
				verifyHandler.Config.EmailSettings.SMTP_HOST))

		if err != nil {
			log.Fatalf("Ошибка при отправке письма: %v", err)
		}

		log.Println("Письмо успешно отправлено")
	}
}

func (verifyHandler *VerifyHandler) Verify() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		hash := req.PathValue("hash")
		file, err := os.Open("data.json")
		if err != nil {
			response.ResponseJSON(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer file.Close()

		fileData, err := io.ReadAll(file)
		if err != nil {
			response.ResponseJSON(w, err.Error(), http.StatusInternalServerError)
			return
		}

		data := JSONData{}
		if err = json.Unmarshal(fileData, &data); err != nil {
			response.ResponseJSON(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if data.Hash != hash {
			w.Write([]byte("false"))
			if err = os.Remove("data.json"); err != nil {
				response.ResponseJSON(w, err.Error(), http.StatusInternalServerError)
				return
			}
		} else {
			w.Write([]byte("true"))
		}
	}
}
