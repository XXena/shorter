package handlers

import (
	"testing"

	"github.com/XXena/shorter/test"
	"github.com/pkg/errors"

	mock "github.com/XXena/shorter/mock"
	"github.com/golang/mock/gomock"
)

func Test__Redirect_fails_if_no_url(t *testing.T) {
	ctrl := gomock.NewController(t)
	service := mock.NewMockRecordRepo(ctrl)
	service.EXPECT().Redirect(gomock.Any()).Return("", errors.New("error: invalid argument"))
	_, err := service.Redirect("")
	if err != nil {
		t.Log(err)
	}
}

func Test__Redirect_pass_on_token(t *testing.T) {
	inputData := test.NewFakeRecord()
	ctrl := gomock.NewController(t)
	service := mock.NewMockRecordRepo(ctrl)
	service.EXPECT().Redirect(inputData.Token).Return(inputData.LongURL, nil)
	_, err := service.Redirect(inputData.Token)
	if err != nil {
		t.Fatal(err)
	}
}
