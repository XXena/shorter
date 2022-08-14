package repository

import (
	"testing"
	"time"

	"github.com/pkg/errors"

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
	if err != nil {
		t.Fatal(err)
	}

	if assert.Nil(t, err) {
		assert.Equal(t, id, inputData.ID)
	}

}

func TestRecordPostgres_CreateFail(t *testing.T) { // todo negative tests
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	inputData := entities.Record{}
	mockRecordRepo := mock.NewMockRecord(ctl)
	mockRecordRepo.EXPECT().Create(inputData).Return(inputData.ID, nil)

	_, err := mockRecordRepo.Create(inputData)
	mockRecordRepo.EXPECT().Create(inputData).Return(0, errors.New("unique constraint violation"))

	_, err = mockRecordRepo.Create(inputData)

	assert.NotNil(t, err)

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
	if err != nil {
		t.Fatal(err)
	}

	if assert.Nil(t, err) {
		assert.Equal(t, inputData, record)
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
	if err != nil {
		t.Fatal(err)
	}

	if assert.Nil(t, err) {
		assert.Equal(t, inputData, record)
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
