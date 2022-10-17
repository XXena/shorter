package handlers

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/XXena/shorter/mock"
	"github.com/XXena/shorter/test"
	"github.com/golang/mock/gomock"
)

func Test__Redirect_fails_if_no_url(t *testing.T) {
	ctrl := gomock.NewController(t)
	service := mock.NewMockRecordServiceInterface(ctrl)
	service.EXPECT().Redirect(gomock.Any()).Return("", errors.New("error: invalid argument"))
	_, err := service.Redirect("")
	if err != nil {
		t.Log(err)
	}
}

func Test__Redirect_pass_on_token(t *testing.T) {
	inputData := test.NewFakeRecord()
	ctrl := gomock.NewController(t)
	service := mock.NewMockRecordServiceInterface(ctrl)
	service.EXPECT().Redirect(inputData.Token).Return(inputData.LongURL, nil)
	_, err := service.Redirect(inputData.Token)
	if err != nil {
		t.Fatal(err)
	}
}

func Test__Redirect_OK_on_token(t *testing.T) {
	inputData := test.NewFakeRecord()
	req := httptest.NewRequest(http.MethodGet, "/"+inputData.Token, nil)
	w := httptest.NewRecorder()

	ctrl := gomock.NewController(t)
	mockHandler := mock.NewMockHandlerInterface(ctrl)
	mockHandler.EXPECT().Redirect(w, req).Return()

	mockHandler.Redirect(w, req)
	if status := w.Code; status != http.StatusOK {
		t.Errorf("redirect handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}
