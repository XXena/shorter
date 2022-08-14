package handlers

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/XXena/shorter/config"
)

func (h *Handler) Fetch(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		h.logger.Debug(fmt.Errorf("form parameter parsing failed: %w", err))
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	url := r.Form.Get(config.UrlFormParameter)
	body, err := h.service.RecordRepo.GetByURL(url)
	if err == nil {
		_, err = w.Write([]byte(body))
		if err != nil {
			h.logger.Error(fmt.Errorf("internal error: %w", err))
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		return
	}

	// todo добавить типизацию ошибок и проверять тип ошибки
	// когда не найден URL, следует создать новую запись:
	if strings.Contains(err.Error(), "no rows in result set") {
		expiry, _ := time.Parse(time.RFC3339, r.Form.Get(config.ExpiryFormParameter))
		token, err := h.service.ForwardToCreate(url, expiry)
		if err != nil {
			h.logger.Error(fmt.Errorf("create new record failed: %w", err))
			http.Error(w, err.Error(), http.StatusInternalServerError)

			return
		}

		w.WriteHeader(http.StatusCreated)
		_, err = w.Write(token)
		if err != nil {
			h.logger.Error(fmt.Errorf("response write error: %w", err))
			http.Error(w, err.Error(), http.StatusInternalServerError)

			return
		}

		return
	}

	h.logger.Error(fmt.Errorf("GetByURL failed with unknown error: %w", err))
	http.Error(w, err.Error(), http.StatusInternalServerError)

}
