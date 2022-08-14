package handlers

import (
	"math/rand"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/XXena/shorter/internal/entities"
	"github.com/XXena/shorter/mock"
	"github.com/XXena/shorter/pkg/logger"
	"github.com/XXena/shorter/test"
	"github.com/golang/mock/gomock"

	"github.com/XXena/shorter/internal/services"
)

func TestHandler_Fetch(t *testing.T) {
	type fields struct {
		//service *services.Service
		service *mock.MockRecordRepo
		logger  logger.Interface
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}

	// prepare data:
	now := time.Now()
	rand.Seed(time.Now().UnixNano())
	str := test.RandSeq(10)
	token, err := services.GenerateShortLink(str, nil)

	if err != nil {
		t.Fatalf("generate short link error: %v", err)
	}

	inputData := []entities.Record{
		// case new record
		{
			ID:         rand.Int(),
			LongURL:    "https://" + str + "/",
			Token:      token,
			CreatedAt:  now,
			ExpiryDate: now.AddDate(1, 0, 0),
		},
		// case existing record
		{
			ID:         9,
			LongURL:    "https://engineering.atspotify.com/2020/04/when-should-i-write-an-architecture-decision-record/",
			Token:      "XAJjKHF4",
			CreatedAt:  now,                  // 0001-01-01T00:00:01Z
			ExpiryDate: now.AddDate(1, 0, 0), // 26132-08-16T01:41:32Z
		},
	}

	// mock the record service: // todo как лучше замокать сервис? может быть иначе?
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	mockRecordService := mock.NewMockRecordRepo(ctl)
	mockRecordService.EXPECT().GetByURL(inputData[1].LongURL).Return(inputData[1].Token, nil)
	mockRecordService.EXPECT().Create(inputData[0]).Return(inputData[0].Token, nil)

	// prepare requests for each case:
	requestsSet, err := test.MakePOSTRequestSet([]test.RequestParams{
		// case new record
		{
			"url",
			inputData[0].LongURL,
		},

		// case existing record
		{
			"url",
			inputData[1].LongURL,
		},
	})

	if err != nil {
		t.Errorf("Handler Fetch err: %v", err)
	}

	// create response recorder:
	respRecoder := httptest.NewRecorder()

	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"create new short url",
			fields{
				service: mockRecordService,
				logger:  nil,
			},
			args{
				w: respRecoder,
				r: requestsSet[0],
			}},
		{"get existing short url",
			fields{
				service: mockRecordService,
				logger:  nil,
			},
			args{
				w: respRecoder,
				r: requestsSet[1],
			}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//h := NewHandler()
			// todo как ему (хэндлеру) подсунуть мок сервиса? или надо сгенерить свои моки для хэндлера (или описать вручную)?
			//h := &Handler{
			//	service: tt.fields.service,
			//	logger:  tt.fields.logger,
			//}
			//h.Fetch(tt.args.w, tt.args.r)
		})
	}
}
