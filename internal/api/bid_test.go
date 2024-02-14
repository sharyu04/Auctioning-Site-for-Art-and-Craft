package api

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sharyu04/Auctioning-Site-for-Art-and-Craft/internal/app/bid/mocks"
	"github.com/sharyu04/Auctioning-Site-for-Art-and-Craft/internal/repository"
	"github.com/stretchr/testify/mock"
)

func TestCreateBidHandler(t *testing.T) {
	bidSvc := mocks.NewService(t)
	CreateBidHandler := createBidHandler(bidSvc)

	tests := []struct {
		name               string
		input              string
		user_id            string
		setup              func(mock *mocks.Service)
		expectedStatusCode int
	}{
		{
			name: "Success for create user",
			input: `{
						"artwork_id" : "468977257870097654290",
						"amount" : 10000,
                    }`,
			user_id: "587693875980298540",
			setup: func(mockSvc *mocks.Service) {
				mockSvc.On("CreateBid", mock.Anything).Return(repository.Bids{}, nil).Once()
			},
			expectedStatusCode: http.StatusOK,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.setup(bidSvc)

			req, err := http.NewRequest("POST", "/bid/create", bytes.NewBuffer([]byte(test.input)))
			if err != nil {
				t.Fatal(err)
				return
			}

			req.Header.Set("user_id", test.user_id)

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(CreateBidHandler)
			handler.ServeHTTP(rr, req)

			fmt.Println("Error")

			if rr.Result().StatusCode != test.expectedStatusCode {
				t.Errorf("Expected %d but got %d", test.expectedStatusCode, rr.Result().StatusCode)
			}
		})
	}
}
