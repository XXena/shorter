package handlers

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/XXena/shorter/config"

	"github.com/XXena/shorter/internal/entities"
)

func (h *Handler) Fetch(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		h.logger.Error(fmt.Errorf("form parameter parsing failed: %w", err))
		return
	}

	url := r.Form.Get(config.FormParameter)
	body, err := h.service.Record.GetByURL(url)
	if err != nil {
		// todo добавить типизацию ошибок и проверять тип ошибки
		// если ошибка - не найден URL, то следует создать новую запись:
		if strings.Contains(err.Error(), "no rows in result set") {
			now := time.Now()
			Record := entities.Record{
				LongURL:   url,
				CreatedAt: now,
				// срок действия - год после создания:
				ExpiryDate: now.AddDate(1, 0, 0), // todo нужен настраиваемый срок действия
			}
			body, err := h.service.Record.Create(Record) // todo вернуть только хэш или полный адрес (http://localhost:8080/UK4Rr3X6)?
			w.WriteHeader(http.StatusCreated)
			_, err = w.Write([]byte(body))
			if err != nil {
				h.logger.Error(fmt.Errorf("create new record failed: %w", err))

				w.WriteHeader(http.StatusInternalServerError)
				_, err := w.Write([]byte(fmt.Sprintf("internal error: %s", err.Error())))
				if err != nil {
					h.logger.Error(fmt.Errorf("internal error: %w", err))
					return
				}
				return
			}
			return
		}
		_, err = w.Write([]byte(body))
		if err != nil {
			h.logger.Error(fmt.Errorf("GetByURL failed: %w", err))
			w.WriteHeader(http.StatusInternalServerError)
			_, err := w.Write([]byte(fmt.Sprintf("internal error: %s", err.Error())))
			if err != nil {
				h.logger.Error(fmt.Errorf("internal error: %w", err))
				return
			}
			return
		}
	}

	_, err = w.Write([]byte(body))
	if err != nil {
		h.logger.Error(fmt.Errorf("internal error: %w", err))
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte(fmt.Sprintf("internal error: %s", err.Error())))
		if err != nil {
			h.logger.Error(fmt.Errorf("internal error: %w", err))
			return
		}
		return
	}

}
