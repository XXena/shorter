package services

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/XXena/shorter/internal/entities"
	"github.com/XXena/shorter/mock"
	"github.com/golang/mock/gomock"
)

func TestRecordService_Create(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	now := time.Now()
	inputData := entities.Record{
		ID:         9,
		LongURL:    "https://engineering.atspotify.com/2020/04/when-should-i-write-an-architecture-decision-record/",
		Token:      "XAJjKHF4",
		CreatedAt:  now,                  // 0001-01-01T00:00:01Z
		ExpiryDate: now.AddDate(1, 0, 0), // 26132-08-16T01:41:32Z
	}

	mockRecordService := mock.NewMockRecordRepo(ctl)
	mockRecordService.EXPECT().Create(inputData).Return(inputData.Token, nil)

	token, err := mockRecordService.Create(inputData)

	if assert.Nil(t, err) {
		assert.Equal(t, token, inputData.Token)
	} else {
		t.Errorf("service Create err: %v", err)
	}

}

func TestRecordService_ForwardToCreate(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	now := time.Now()
	inputData := entities.Record{
		ID:         9,
		LongURL:    "https://engineering.atspotify.com/2020/04/when-should-i-write-an-architecture-decision-record/",
		Token:      "XAJjKHF4",
		CreatedAt:  now,                  // 0001-01-01T00:00:01Z
		ExpiryDate: now.AddDate(1, 0, 0), // 26132-08-16T01:41:32Z
	}

	mockRecordService := mock.NewMockRecordRepo(ctl)
	mockRecordService.EXPECT().ForwardToCreate(inputData.LongURL, inputData.ExpiryDate).Return([]byte(inputData.Token), nil)

	tokenBytes, err := mockRecordService.ForwardToCreate(inputData.LongURL, inputData.ExpiryDate)

	if assert.Nil(t, err) {
		assert.Equal(t, tokenBytes, []byte(inputData.Token))
	}
}

func TestRecordService_GetByURL(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	now := time.Now()
	inputData := entities.Record{
		ID:         9,
		LongURL:    "https://engineering.atspotify.com/2020/04/when-should-i-write-an-architecture-decision-record/",
		Token:      "XAJjKHF4",
		CreatedAt:  now,                  // 0001-01-01T00:00:01Z
		ExpiryDate: now.AddDate(1, 0, 0), // 26132-08-16T01:41:32Z
	}

	mockRecordService := mock.NewMockRecordRepo(ctl)
	mockRecordService.EXPECT().GetByURL(inputData.LongURL).Return(inputData.Token, nil)

	token, err := mockRecordService.GetByURL(inputData.LongURL)

	if assert.Nil(t, err) {
		assert.Equal(t, token, inputData.Token)
	} else {
		t.Errorf("service GetByURL err: %v", err)
	}

}

func TestRecordService_Redirect(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	now := time.Now()
	inputData := entities.Record{
		ID:         9,
		LongURL:    "https://engineering.atspotify.com/2020/04/when-should-i-write-an-architecture-decision-record/",
		Token:      "XAJjKHF4",
		CreatedAt:  now,                  // 0001-01-01T00:00:01Z
		ExpiryDate: now.AddDate(1, 0, 0), // 26132-08-16T01:41:32Z
	}

	mockRecordService := mock.NewMockRecordRepo(ctl)
	mockRecordService.EXPECT().Redirect(inputData.Token).Return(inputData.LongURL, nil)

	url, err := mockRecordService.Redirect(inputData.Token)

	if assert.Nil(t, err) {
		assert.Equal(t, url, inputData.LongURL)
	} else {
		t.Errorf("service Redirect err: %v", err)
	}
}
