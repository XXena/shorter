package services

import (
	"errors"
	"testing"

	"github.com/XXena/shorter/internal/entities"

	"github.com/XXena/shorter/test"

	"github.com/stretchr/testify/assert"

	"github.com/XXena/shorter/mock"
	"github.com/golang/mock/gomock"
)

func Test__RecordService_Create_pass(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	inputData := test.NewFakeRecord()
	mockRecordService := mock.NewMockRecordServiceInterface(ctl)
	mockRecordService.EXPECT().Create(inputData).Return(inputData.Token, nil)
	token, err := mockRecordService.Create(inputData)
	if err != nil {
		t.Fatal(err)
	}

	if assert.Nil(t, err) {
		assert.Equal(t, token, inputData.Token)
	}

}

func Test__RecordService_Create_fails_on_empty_record(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	inputData := entities.Record{}
	mockRecordService := mock.NewMockRecordServiceInterface(ctl)
	mockRecordService.EXPECT().Create(inputData).Return(inputData.Token, errors.New("error: invalid argument"))
	token, err := mockRecordService.Create(inputData)
	assert.Equal(t, token, "")
	assert.NotNil(t, err)
}

func Test__RecordService_ForwardToCreate_pass(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	inputData := test.NewFakeRecord()
	mockRecordService := mock.NewMockRecordServiceInterface(ctl)
	mockRecordService.EXPECT().ForwardToCreate(inputData.LongURL, inputData.ExpiryDate).Return([]byte(inputData.Token), nil)
	tokenBytes, err := mockRecordService.ForwardToCreate(inputData.LongURL, inputData.ExpiryDate)
	if err != nil {
		t.Fatal(err)
	}

	if assert.Nil(t, err) {
		assert.Equal(t, tokenBytes, []byte(inputData.Token))
	}
}

func Test__RecordService_ForwardToCreate_fails_if_no_url(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	inputData := entities.Record{}
	mockRecordService := mock.NewMockRecordServiceInterface(ctl)
	mockRecordService.EXPECT().ForwardToCreate(inputData.LongURL, inputData.ExpiryDate).Return([]byte(inputData.Token), errors.New("error: invalid argument"))
	tokenBytes, err := mockRecordService.ForwardToCreate(inputData.LongURL, inputData.ExpiryDate)

	assert.NotNil(t, err)
	assert.Equal(t, []byte{}, tokenBytes)
}

func Test__RecordService_GetByURL_pass(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	inputData := test.NewFakeRecord()
	mockRecordService := mock.NewMockRecordServiceInterface(ctl)
	mockRecordService.EXPECT().GetByURL(inputData.LongURL).Return(inputData.Token, nil)
	token, err := mockRecordService.GetByURL(inputData.LongURL)
	if err != nil {
		t.Fatal(err)
	}

	if assert.Nil(t, err) {
		assert.Equal(t, token, inputData.Token)
	}

}

func Test__RecordService_GetByURL_fails_if_no_url(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	inputData := entities.Record{}
	mockRecordService := mock.NewMockRecordServiceInterface(ctl)
	mockRecordService.EXPECT().GetByURL(inputData.LongURL).Return(inputData.Token, errors.New("no rows in sql set"))
	token, err := mockRecordService.GetByURL(inputData.LongURL)
	assert.NotNil(t, err)
	assert.Equal(t, "", token)
}

func Test__RecordService_Redirect_pass(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	inputData := test.NewFakeRecord()
	mockRecordService := mock.NewMockRecordServiceInterface(ctl)
	mockRecordService.EXPECT().Redirect(inputData.Token).Return(inputData.LongURL, nil)
	url, err := mockRecordService.Redirect(inputData.Token)
	if err != nil {
		t.Fatal(err)
	}

	if assert.Nil(t, err) {
		assert.Equal(t, url, inputData.LongURL)
	}
}

func Test__RecordService_Redirect_fails_if_no_url(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	inputData := entities.Record{}
	mockRecordService := mock.NewMockRecordServiceInterface(ctl)
	mockRecordService.EXPECT().Redirect(inputData.Token).Return(inputData.LongURL, errors.New("no rows in sql set"))
	url, err := mockRecordService.Redirect(inputData.Token)
	assert.Equal(t, "", url)
	assert.NotNil(t, err)
}
