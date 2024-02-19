package api

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sharyu04/Auctioning-Site-for-Art-and-Craft/internal/app/artwork/mocks"
	"github.com/sharyu04/Auctioning-Site-for-Art-and-Craft/internal/pkg/dto"
	"github.com/sharyu04/Auctioning-Site-for-Art-and-Craft/internal/repository"
	"github.com/stretchr/testify/mock"
)

func TestCreateArtworkHandler(t *testing.T) {
	artworkSvc := mocks.NewService(t)
	CreateArtworkHandler := createArtworkHandler(artworkSvc)

	tests := []struct {
		name               string
		input              string
		user_id            string
		setup              func(mock *mocks.Service)
		expectedStatusCode int
	}{
		{
			name: "Success for create artwork",
			input: `{
						"name":"Bidder1's test painting",
						"description": "This is bidder 1's third Painting",
						"image":"https://c4.wallpaperflare.com/wallpaper/889/503/640/iron-man-painting-nano-suit-artwork-wallpaper-preview.jpg",
						"starting_price": 10000,
						"duration":2,
						"category":"Pencil_Art"
                    }`,
			user_id: "cb16658c-5ea5-4d0f-b8f4-de4b71eb0518",
			setup: func(mockSvc *mocks.Service) {
				mockSvc.On("CreateArtwork", mock.Anything).Return(repository.Artworks{}, nil).Once()
			},
			expectedStatusCode: http.StatusOK,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.setup(artworkSvc)

			req, err := http.NewRequest("POST", "/artwork/create", bytes.NewBuffer([]byte(test.input)))
			if err != nil {
				t.Fatal(err)
				return
			}

			req.Header.Set("user_id", test.user_id)

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(CreateArtworkHandler)
			handler.ServeHTTP(rr, req)

			fmt.Println("Error")

			if rr.Result().StatusCode != test.expectedStatusCode {
				t.Errorf("Expected %d but got %d", test.expectedStatusCode, rr.Result().StatusCode)
			}
		})
	}
}

func TestGetArtworksHandler(t *testing.T) {
	artworkSvc := mocks.NewService(t)
	GetArtworksHandler := GetArtworksHandler(artworkSvc)

	tests := []struct {
		name               string
		start              int
		count              int
		category           string
		setup              func(mock *mocks.Service)
		expectedStatusCode int
	}{
		{
			name:  "Success for get artwork",
			start: 0,
			count: 4,
			setup: func(mockSvc *mocks.Service) {
				mockSvc.On("GetArtworks", mock.Anything, mock.Anything, mock.Anything).Return(nil, nil).Once()
			},
			expectedStatusCode: http.StatusOK,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.setup(artworkSvc)

			req, err := http.NewRequest("GET", fmt.Sprintf("/artworks?start=%d&count=%d&category=%s", test.start, test.count, test.category), bytes.NewBuffer([]byte("")))
			if err != nil {
				t.Fatal(err)
				return
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(GetArtworksHandler)
			handler.ServeHTTP(rr, req)

			fmt.Println("Error")

			if rr.Result().StatusCode != test.expectedStatusCode {
				t.Errorf("Expected %d but got %d", test.expectedStatusCode, rr.Result().StatusCode)
			}
		})
	}
}

func TestDeleteArtworkHandler(t *testing.T) {
	artworkSvc := mocks.NewService(t)
	DeleteArtworkHandler := DeleteArtworkHandler(artworkSvc)

	tests := []struct {
		name               string
		id                 string
		user_id            string
		role               string
		setup              func(mock *mocks.Service)
		expectedStatusCode int
	}{
		{
			name:    "Success for get artwork",
			id:      "16c8a24-4103-4148-881f-bf70391fb897",
			user_id: "16c8a24-4103-4148-881f-bf70391fb897",
			role:    "admin",
			setup: func(mockSvc *mocks.Service) {
				mockSvc.On("DeleteArtworkById", mock.Anything, mock.Anything, mock.Anything).Return(nil).Once()
			},
			expectedStatusCode: http.StatusOK,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.setup(artworkSvc)

			req, err := http.NewRequest("DELETE", fmt.Sprintf("/artwork/%s", test.id), bytes.NewBuffer([]byte("")))
			if err != nil {
				t.Fatal(err)
				return
			}

			req.Header.Set("user_id", test.user_id)
			req.Header.Set("role", test.role)

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(DeleteArtworkHandler)
			handler.ServeHTTP(rr, req)

			fmt.Println("Error")

			if rr.Result().StatusCode != test.expectedStatusCode {
				t.Errorf("Expected %d but got %d", test.expectedStatusCode, rr.Result().StatusCode)
			}
		})
	}
}

func TestGetArtworkByIdHandler(t *testing.T) {
	artworkSvc := mocks.NewService(t)
	GetArtworkByIdHandler := GetArtworkByIdHandler(artworkSvc)

	tests := []struct {
		name               string
		id                 string
		setup              func(mock *mocks.Service)
		expectedStatusCode int
	}{
		{
			name: "Success for get artwork",
			id:   "16c8a24-4103-4148-881f-bf70391fb897",
			setup: func(mockSvc *mocks.Service) {
				mockSvc.On("GetArtworkByID", mock.Anything).Return(dto.GetArtworkResponse{}, nil).Once()
			},
			expectedStatusCode: http.StatusOK,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.setup(artworkSvc)

			req, err := http.NewRequest("GET", fmt.Sprintf("/artwork/%s", test.id), bytes.NewBuffer([]byte("")))
			if err != nil {
				t.Fatal(err)
				return
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(GetArtworkByIdHandler)
			handler.ServeHTTP(rr, req)

			fmt.Println("Error")

			if rr.Result().StatusCode != test.expectedStatusCode {
				t.Errorf("Expected %d but got %d", test.expectedStatusCode, rr.Result().StatusCode)
			}
		})
	}
}
