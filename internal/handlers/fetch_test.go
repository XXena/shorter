package handlers

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/XXena/shorter/mock"
	"github.com/golang/mock/gomock"

	"github.com/XXena/shorter/test"
)

func Test__Fetch_OK_on_url(t *testing.T) {
	inputData := test.NewFakeRecord()
	body := strings.NewReader("url=" + inputData.LongURL)
	req := httptest.NewRequest(http.MethodPost, "/send", body)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()

	ctrl := gomock.NewController(t)
	mockHandler := mock.NewMockHandlerInterface(ctrl)
	mockHandler.EXPECT().Fetch(w, req).Return()

	mockHandler.Fetch(w, req)

	res := w.Result()
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			t.Fatal("unable to close response body")
		}
	}(res.Body)

	if status := w.Code; status != http.StatusOK {
		t.Errorf("fetch handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}
