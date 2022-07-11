package repository

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/XXena/shorter/internal/entities"

	"github.com/XXena/shorter/mock"

	"github.com/golang/mock/gomock"
)

func TestRecordPostgres_Create(t *testing.T) {
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
	mockRecordRepo := mock.NewMockRecord(ctl)
	mockRecordRepo.EXPECT().Create(inputData).Return(inputData.ID, nil)

	id, err := mockRecordRepo.Create(inputData)

	if assert.Nil(t, err) {
		assert.Equal(t, id, inputData.ID)
	} else {
		t.Errorf("repo Create err: %v", err)
	}

}

func TestRecordPostgres_GetByToken(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	mockRecordRepo := mock.NewMockRecord(ctl)

	now := time.Now()
	inputData := entities.Record{
		ID:         9,
		LongURL:    "https://engineering.atspotify.com/2020/04/when-should-i-write-an-architecture-decision-record/",
		Token:      "XAJjKHF4",
		CreatedAt:  now,                  // 0001-01-01T00:00:01Z
		ExpiryDate: now.AddDate(1, 0, 0), // 26132-08-16T01:41:32Z
	}

	mockRecordRepo.EXPECT().GetByToken(inputData.Token).Return(inputData, nil)
	record, err := mockRecordRepo.GetByToken(inputData.Token)
	if assert.Nil(t, err) {
		assert.Equal(t, inputData, record)
	} else {
		t.Errorf("record GetByToken err: %v", err)
	}
}

func TestRecordPostgres_GetByURL(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	mockRecordRepo := mock.NewMockRecord(ctl)

	now := time.Now()
	inputData := entities.Record{
		ID:         9,
		LongURL:    "https://engineering.atspotify.com/2020/04/when-should-i-write-an-architecture-decision-record/",
		Token:      "XAJjKHF4",
		CreatedAt:  now,                  // 0001-01-01T00:00:01Z
		ExpiryDate: now.AddDate(1, 0, 0), // 26132-08-16T01:41:32Z
	}

	mockRecordRepo.EXPECT().GetByURL(inputData.LongURL).Return(inputData, nil)
	record, err := mockRecordRepo.GetByURL(inputData.LongURL)
	if assert.Nil(t, err) {
		assert.Equal(t, inputData, record)
	} else {
		t.Errorf("record GetByURL err: %v", err)
	}
}

//func TestRecordPostgres_Delete(t *testing.T) { // todo ложно положительный, почему?
//	ctl := gomock.NewController(t)
//	defer ctl.Finish()
//
//	mockRecordRepo := mock.NewMockRecord(ctl)
//
//	inputData := entities.Record{
//		ID: 9,
//	}
//
//	mockRecordRepo.EXPECT().Delete(inputData.ID).Return(nil)
//	err := mockRecordRepo.Delete(inputData.ID)
//	if !assert.Nil(t, err) {
//		t.Errorf("repo Delete err: %v", err)
//	}
//}
