package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func (h *Handler) Redirect(w http.ResponseWriter, r *http.Request) {
	shortURL := r.URL.RequestURI()
	longURL, err := h.service.Redirect(shortURL)

	if err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			w.WriteHeader(http.StatusNotFound)
			_, err = w.Write([]byte(fmt.Sprintf("Short link %s not found, create new one", shortURL)))
			return
		}

		log.Printf("Redirect failed: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte(fmt.Sprintf("internal error: %s", err.Error())))
		if err != nil {
			log.Println(err)
			return
		}
		return
	}

	http.Redirect(w, r, longURL, 302) // todo status code
	return

}
