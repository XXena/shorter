package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/XXena/shorter/internal/entities"
)

func (h *Handler) Fetch(w http.ResponseWriter, r *http.Request) {
	url := r.URL.RequestURI()
	body, err := h.service.Record.GetByURL(url)
	if err != nil {
		// todo добавить типизацию ошибок и проверять тип ошибки
		// если ошибка - не найден URL, то следует создать новую запись:
		if strings.Contains(err.Error(), "no rows in result set") {
			now := time.Now()
			Record := entities.Record{
				LongURL:   url,
				CreatedAt: now,
				// год после создания:
				ExpiryDate: now.AddDate(1, 0, 0), // todo кастомный срок действия
			}
			body, err := h.service.Record.Create(Record)
			_, err = w.Write([]byte(body))
			if err != nil {
				log.Printf("Create new record failed: %s", err)
				w.WriteHeader(http.StatusInternalServerError)
				_, err := w.Write([]byte(fmt.Sprintf("internal error: %s", err.Error())))
				if err != nil {
					log.Println(err)
					return
				}
				return
			}
		}
		_, err = w.Write([]byte(body))
		if err != nil {
			log.Printf("GetByURL failed: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
			_, err := w.Write([]byte(fmt.Sprintf("internal error: %s", err.Error())))
			if err != nil {
				log.Println(err)
				return
			}
			return
		}
	}

	_, err = w.Write([]byte(body))
	if err != nil {
		log.Printf("GetByURL failed: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte(fmt.Sprintf("internal error: %s", err.Error())))
		if err != nil {
			log.Println(err)
			return
		}
		return
	}
}
