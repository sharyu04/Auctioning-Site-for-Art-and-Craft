package artwork

// import (
// 	"testing"

// 	"github.com/google/uuid"
// 	"github.com/sharyu04/Auctioning-Site-for-Art-and-Craft/internal/pkg/dto"
// 	"github.com/sharyu04/Auctioning-Site-for-Art-and-Craft/internal/repository"
// 	"github.com/sharyu04/Auctioning-Site-for-Art-and-Craft/internal/repository/mocks"
// 	"github.com/stretchr/testify/mock"
// )

// func TestCreateArtwork(t *testing.T) {
// 	artworkRepo := mocks.NewArtworkStorer(t)
// 	service := NewService(artworkRepo)

// 	tests := []struct {
// 		name            string
// 		input           dto.CreateArtworkRequest
// 		setup           func(artworkMock *mocks.ArtworkStorer)
// 		isErrorExpected bool
// 	}{
// 		{
// 			name: "Success for create artwork",
// 			input: dto.CreateArtworkRequest{
// 				Name:           "Bidder1's test painting",
// 				Description:    "This is bidder 1's third Painting",
// 				Image:          "https://c4.wallpaperflare.com/wallpaper/889/503/640/iron-man-painting-nano-suit-artwork-wallpaper-preview.jpg",
// 				Starting_price: 10000,
// 				Duration:       2,
// 				Category:       "Pencil_Art",
// 			},
// 			setup: func(artworkMock *mocks.ArtworkStorer) {
// 				artworkMock.On("GetCategory", mock.Anything).Return(repository.Category{}, nil).Once()
// 				artworkMock.On("CreateArtwork", mock.Anything).Return(repository.Artworks{}, nil).Once()
// 			},
// 			isErrorExpected: false,
// 		},
// 	}

// 	for _, test := range tests {
// 		t.Run(test.name, func(t *testing.T) {
// 			test.setup(artworkRepo)

// 			// test service
// 			_, err := service.CreateArtwork(test.input)

// 			if (err != nil) != test.isErrorExpected {
// 				t.Errorf("Test Failed, expected error to be %v, but got err %v", test.isErrorExpected, err != nil)
// 			}
// 		})
// 	}
// }

// func TestGetArtworks(t *testing.T) {
// 	artworkRepo := mocks.NewArtworkStorer(t)
// 	service := NewService(artworkRepo)

// 	tests := []struct {
// 		name            string
// 		start           int
// 		count           int
// 		category        string
// 		setup           func(userMock *mocks.ArtworkStorer)
// 		isErrorExpected bool
// 	}{
// 		{
// 			name:     "Success for get artworks",
// 			start:    0,
// 			count:    4,
// 			category: "",
// 			setup: func(artworkMock *mocks.ArtworkStorer) {
// 				Id, _ := uuid.Parse("010d9c50-646c-48af-8cb1-2db613102c4d")
// 				// artworkRepo.On("GetFilterArtworks", mock.Anything, mock.Anything, mock.Anything).Return([]dto.GetArtworkResponse{
// 				// 	{
// 				// 		Id:             Id,
// 				// 		Name:           "Bidder1's test painting",
// 				// 		Description:    "This is bidder 1's third Painting",
// 				// 		Image:          "https://c4.wallpaperflare.com/wallpaper/889/503/640/iron-man-painting-nano-suit-artwork-wallpaper-preview.jpg",
// 				// 		Starting_price: 10000,
// 				// 		Category:       "Pencil_Art",
// 				// 		Highest_bid:    20000,
// 				// 	},
// 				// }, nil).Once()
// 				artworkRepo.On("GetAllArtworks", mock.Anything, mock.Anything).Return([]dto.GetArtworkResponse{
// 					{
// 						Id:             Id,
// 						Name:           "Bidder1's test painting",
// 						Description:    "This is bidder 1's third Painting",
// 						Image:          "https://c4.wallpaperflare.com/wallpaper/889/503/640/iron-man-painting-nano-suit-artwork-wallpaper-preview.jpg",
// 						Starting_price: 10000,
// 						Category:       "Pencil_Art",
// 						Highest_bid:    20000,
// 					},
// 				}, nil).Once()
// 			},
// 			isErrorExpected: false,
// 		},
// 	}

// 	for _, test := range tests {
// 		t.Run(test.name, func(t *testing.T) {
// 			test.setup(artworkRepo)

// 			// test service
// 			_, err := service.GetArtworks(test.category, test.start, test.count)

// 			if (err != nil) != test.isErrorExpected {
// 				t.Errorf("Test Failed, expected error to be %v, but got err %v", test.isErrorExpected, err != nil)
// 			}
// 		})
// 	}
// }

// func TestGetArtworkByID(t *testing.T) {
// 	artworkRepo := mocks.NewArtworkStorer(t)
// 	service := NewService(artworkRepo)

// 	tests := []struct {
// 		name            string
// 		id              string
// 		setup           func(artworkMock *mocks.ArtworkStorer)
// 		isErrorExpected bool
// 	}{
// 		{
// 			name: "Success for get artwork by id",
// 			id:   "010d9c50-646c-48af-8cb1-2db613102c4d",
// 			setup: func(artworkMock *mocks.ArtworkStorer) {
// 				artworkRepo.On("GetArtworkById", mock.Anything).Return(dto.GetArtworkResponse{}, nil).Once()
// 			},
// 			isErrorExpected: false,
// 		},
// 	}

// 	for _, test := range tests {
// 		t.Run(test.name, func(t *testing.T) {
// 			test.setup(artworkRepo)

// 			// test service
// 			_, err := service.GetArtworkByID(test.id)

// 			if (err != nil) != test.isErrorExpected {
// 				t.Errorf("Test Failed, expected error to be %v, but got err %v", test.isErrorExpected, err != nil)
// 			}
// 		})
// 	}
// }

// func TestDeleteArtworkById(t *testing.T) {
// 	artworkRepo := mocks.NewArtworkStorer(t)
// 	service := NewService(artworkRepo)

// 	tests := []struct {
// 		name            string
// 		id              string
// 		owner_id        string
// 		role            string
// 		setup           func(artworkMock *mocks.ArtworkStorer)
// 		isErrorExpected bool
// 	}{
// 		{
// 			name:     "Success for get artwork by id",
// 			id:       "010d9c50-646c-48af-8cb1-2db613102c4d",
// 			owner_id: "010d9c50-646c-48af-8cb1-2db613102c4d",
// 			role:     "admin",
// 			setup: func(artworkMock *mocks.ArtworkStorer) {
// 				artworkRepo.On("DeleteArtworkById", mock.Anything, mock.Anything, mock.Anything).Return(nil).Once()
// 			},
// 			isErrorExpected: false,
// 		},
// 	}

// 	for _, test := range tests {
// 		t.Run(test.name, func(t *testing.T) {
// 			test.setup(artworkRepo)

// 			// test service
// 			err := service.DeleteArtworkById(test.id, test.owner_id, test.role)

// 			if (err != nil) != test.isErrorExpected {
// 				t.Errorf("Test Failed, expected error to be %v, but got err %v", test.isErrorExpected, err != nil)
// 			}
// 		})
// 	}
// }
