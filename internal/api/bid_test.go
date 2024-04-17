package api

// import (
// 	"bytes"
// 	"fmt"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/sharyu04/Auctioning-Site-for-Art-and-Craft/internal/app/bid/mocks"
// 	"github.com/sharyu04/Auctioning-Site-for-Art-and-Craft/internal/repository"
// 	"github.com/stretchr/testify/mock"
// )

// func TestCreateBidHandler(t *testing.T) {
// 	bidsSvc := mocks.NewService(t)
// 	CreateBidHandler := createBidHandler(bidsSvc)

// 	tests := []struct {
// 		name               string
// 		input              string
// 		user_id            string
// 		setup              func(mock *mocks.Service)
// 		expectedStatusCode int
// 	}{
// 		{
// 			name: "Success for create bid",
// 			input: `{
// 						"artwork_id" : "016c8a24-4103-4148-881f-bf70391fb897",
// 						"amount" : 40000
//                     }`,
// 			user_id: "cb16658c-5ea5-4d0f-b8f4-de4b71eb0518",
// 			setup: func(mockSvc *mocks.Service) {
// 				mockSvc.On("CreateBid", mock.Anything).Return(repository.Bids{}, nil).Once()
// 			},
// 			expectedStatusCode: http.StatusOK,
// 		},
// 	}

// 	for _, test := range tests {
// 		t.Run(test.name, func(t *testing.T) {
// 			test.setup(bidsSvc)

// 			req, err := http.NewRequest("POST", "/bid/create", bytes.NewBuffer([]byte(test.input)))
// 			if err != nil {
// 				t.Fatal(err)
// 				return
// 			}

// 			req.Header.Set("user_id", test.user_id)

// 			rr := httptest.NewRecorder()
// 			handler := http.HandlerFunc(CreateBidHandler)
// 			handler.ServeHTTP(rr, req)

// 			fmt.Println("Error")

// 			if rr.Result().StatusCode != test.expectedStatusCode {
// 				t.Errorf("Expected %d but got %d", test.expectedStatusCode, rr.Result().StatusCode)
// 			}
// 		})
// 	}
// }

// func TestUpdateBidHandler(t *testing.T) {
// 	bidsSvc := mocks.NewService(t)
// 	UpdateBidHandler := updateBidHandler(bidsSvc)

// 	tests := []struct {
// 		name               string
// 		input              string
// 		user_id            string
// 		setup              func(mock *mocks.Service)
// 		expectedStatusCode int
// 	}{
// 		{
// 			name: "Success for update bid",
// 			input: `{
// 						"artwork_id" : "016c8a24-4103-4148-881f-bf70391fb897",
// 						"amount" : 40000
//                     }`,
// 			user_id: "cb16658c-5ea5-4d0f-b8f4-de4b71eb0518",
// 			setup: func(mockSvc *mocks.Service) {
// 				mockSvc.On("UpdateBid", mock.Anything, mock.Anything).Return(repository.Bids{}, nil).Once()
// 			},
// 			expectedStatusCode: http.StatusOK,
// 		},
// 	}

// 	for _, test := range tests {
// 		t.Run(test.name, func(t *testing.T) {
// 			test.setup(bidsSvc)

// 			req, err := http.NewRequest("PUT", "/bid/update", bytes.NewBuffer([]byte(test.input)))
// 			if err != nil {
// 				t.Fatal(err)
// 				return
// 			}

// 			req.Header.Set("user_id", test.user_id)

// 			rr := httptest.NewRecorder()
// 			handler := http.HandlerFunc(UpdateBidHandler)
// 			handler.ServeHTTP(rr, req)

// 			fmt.Println("Error")

// 			if rr.Result().StatusCode != test.expectedStatusCode {
// 				t.Errorf("Expected %d but got %d", test.expectedStatusCode, rr.Result().StatusCode)
// 			}
// 		})
// 	}
// }
