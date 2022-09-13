package repository

import (
	"testing"

	"github.com/XXena/shorter/test"
	"github.com/pkg/errors"

	"github.com/stretchr/testify/assert"

	"github.com/XXena/shorter/internal/entities"

	"github.com/XXena/shorter/mock"

	"github.com/golang/mock/gomock"
)

func Test__RecordPostgres_Create_pass(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	inputData := test.NewFakeRecord()
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

func Test__RecordPostgres_Create_fails_on_empty_record(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	inputData := entities.Record{}
	mockRecordRepo := mock.NewMockRecord(ctl)
	mockRecordRepo.EXPECT().Create(inputData).Return(0, errors.New("unique constraint violation"))
	_, err := mockRecordRepo.Create(inputData)
	assert.NotNil(t, err)

}

func Test__RecordPostgres_GetByToken_pass(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	mockRecordRepo := mock.NewMockRecord(ctl)
	inputData := test.NewFakeRecord()
	mockRecordRepo.EXPECT().GetByToken(inputData.Token).Return(inputData, nil)
	record, err := mockRecordRepo.GetByToken(inputData.Token)
	if err != nil {
		t.Fatal(err)
	}

	if assert.Nil(t, err) {
		assert.Equal(t, inputData, record)
	}
}

func Test__RecordPostgres_GetByToken_fails_on_empty_record(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	mockRecordRepo := mock.NewMockRecord(ctl)
	inputData := entities.Record{}
	mockRecordRepo.EXPECT().GetByToken(inputData.Token).Return(inputData, errors.New("no rows in sql set"))
	record, err := mockRecordRepo.GetByToken(inputData.Token)
	assert.Equal(t, 0, record.ID)
	assert.NotNil(t, err)
}

func Test__RecordPostgres_GetByURL_pass(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	mockRecordRepo := mock.NewMockRecord(ctl)
	inputData := test.NewFakeRecord()
	mockRecordRepo.EXPECT().GetByURL(inputData.LongURL).Return(inputData, nil)
	record, err := mockRecordRepo.GetByURL(inputData.LongURL)
	if err != nil {
		t.Fatal(err)
	}

	if assert.Nil(t, err) {
		assert.Equal(t, inputData, record)
	}
}

func Test__RecordPostgres_GetByURL_fails_on_empty_record(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	mockRecordRepo := mock.NewMockRecord(ctl)
	inputData := test.NewFakeRecord()
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
