package services

import (
	"testing"

	"github.com/XXena/shorter/test"

	"github.com/stretchr/testify/assert"

	"github.com/XXena/shorter/mock"
	"github.com/golang/mock/gomock"
)

func TestRecordService_Create(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	inputData := test.NewFakeRecord()
	mockRecordService := mock.NewMockRecordRepo(ctl)
	mockRecordService.EXPECT().Create(inputData).Return(inputData.Token, nil)
	token, err := mockRecordService.Create(inputData)
	if err != nil {
		t.Fatal(err)
	}

	if assert.Nil(t, err) {
		assert.Equal(t, token, inputData.Token)
	}

}

func TestRecordService_ForwardToCreate(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	inputData := test.NewFakeRecord()
	mockRecordService := mock.NewMockRecordRepo(ctl)
	mockRecordService.EXPECT().ForwardToCreate(inputData.LongURL, inputData.ExpiryDate).Return([]byte(inputData.Token), nil)
	tokenBytes, err := mockRecordService.ForwardToCreate(inputData.LongURL, inputData.ExpiryDate)
	if err != nil {
		t.Fatal(err)
	}

	if assert.Nil(t, err) {
		assert.Equal(t, tokenBytes, []byte(inputData.Token))
	}
}

func TestRecordService_GetByURL(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	inputData := test.NewFakeRecord()
	mockRecordService := mock.NewMockRecordRepo(ctl)
	mockRecordService.EXPECT().GetByURL(inputData.LongURL).Return(inputData.Token, nil)
	token, err := mockRecordService.GetByURL(inputData.LongURL)
	if err != nil {
		t.Fatal(err)
	}

	if assert.Nil(t, err) {
		assert.Equal(t, token, inputData.Token)
	}

}

func TestRecordService_Redirect(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	inputData := test.NewFakeRecord()
	mockRecordService := mock.NewMockRecordRepo(ctl)
	mockRecordService.EXPECT().Redirect(inputData.Token).Return(inputData.LongURL, nil)
	url, err := mockRecordService.Redirect(inputData.Token)
	if err != nil {
		t.Fatal(err)
	}

	if assert.Nil(t, err) {
		assert.Equal(t, url, inputData.LongURL)
	}
}
