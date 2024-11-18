package middleware

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.SetFormatter(&log.JSONFormatter{})

		if err := r.ParseForm(); err != nil {
			log.Error(err)
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		log.WithFields(log.Fields{
			"Host":    r.Host,
			"Method":  r.Method,
			"Path":    r.URL.Path,
			"Payload": r.PostForm,
		}).Info("Logger received")

		next.ServeHTTP(w, r)
	})
}
