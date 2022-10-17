package handlers

import (
	"fmt"
	"net/http"
	"strings"
)

func (h *handler) Redirect(w http.ResponseWriter, r *http.Request) {
	shortURL := r.URL.RequestURI()
	longURL, err := h.service.Redirect(shortURL)

	if err == nil {
		http.Redirect(w, r, longURL, http.StatusFound)

		return
	}

	if strings.Contains(err.Error(), "no rows in result set") {
		w.WriteHeader(http.StatusNotFound)
		_, err = w.Write([]byte(fmt.Sprintf("Short link %s not found, create new one", shortURL)))

		return
	}

	h.logger.Error(fmt.Errorf("redirect failed: %w", err))
	http.Error(w, err.Error(), http.StatusInternalServerError)

	return

}
