package api

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sharyu04/Auctioning-Site-for-Art-and-Craft/internal/app/artwork/mocks"
	"github.com/sharyu04/Auctioning-Site-for-Art-and-Craft/internal/repository"
	"github.com/stretchr/testify/mock"
)

func TestCreateArtworkHandler(t *testing.T) {
	artworkSvc := mocks.NewService(t)
	CreateUserHandler := createArtworkHandler(artworkSvc)

	tests := []struct {
		name               string
		input              string
		setup              func(mock *mocks.Service)
		expectedStatusCode int
	}{
		{
			name: "Success for create artwork",
			input: `{
						"name" : "user1",
						"description" : "Marwadi",
						"image": "sharayu6@gmail.com",
						"password": "Root"
                    }`,
			setup: func(mockSvc *mocks.Service) {
				mockSvc.On("CreateUser", mock.Anything, mock.Anything).Return(repository.Artworks{}, nil).Once()
			},
			expectedStatusCode: http.StatusOK,
		},
		// {
		//     name:               "Fail Invalid json",
		//     input:              "",
		//     setup:              func(mockSvc *mocks.Service) {},
		//     expectedStatusCode: http.StatusBadRequest,
		// },

	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.setup(artworkSvc)

			req, err := http.NewRequest("POST", "/user/signup", bytes.NewBuffer([]byte(test.input)))
			if err != nil {
				t.Fatal(err)
				return
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(CreateUserHandler)
			handler.ServeHTTP(rr, req)

			fmt.Println("Error")

			if rr.Result().StatusCode != test.expectedStatusCode {
				t.Errorf("Expected %d but got %d", test.expectedStatusCode, rr.Result().StatusCode)
			}
		})
	}
}
