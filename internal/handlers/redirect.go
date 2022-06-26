package handlers

import (
	"fmt"
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

		h.logger.Error(fmt.Errorf("redirect failed: %w", err))
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte(fmt.Sprintf("internal error: %s", err.Error())))
		if err != nil {
			h.logger.Error(fmt.Errorf("internal error: %w", err))
			return
		}
		return
	}

	http.Redirect(w, r, longURL, http.StatusFound)
	return

}
