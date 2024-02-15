package bid

import (
	"testing"

	"github.com/google/uuid"
	"github.com/sharyu04/Auctioning-Site-for-Art-and-Craft/internal/pkg/dto"
	"github.com/sharyu04/Auctioning-Site-for-Art-and-Craft/internal/repository"
	"github.com/sharyu04/Auctioning-Site-for-Art-and-Craft/internal/repository/mocks"
	"github.com/stretchr/testify/mock"
)

func TestCreateBid(t *testing.T) {
	bidRepo := mocks.NewBidStorer(t)
	service := NewService(bidRepo)

	tests := []struct {
		name            string
		input           dto.CreateBidRequest
		setup           func(bidMock *mocks.BidStorer)
		isErrorExpected bool
	}{
		{
			name: "Success for create bid",
			input: dto.CreateBidRequest{
				Artwork_id: "016c8a24-4103-4148-881f-bf70391fb897",
				Amount:     40000.0,
			},
			setup: func(bidMock *mocks.BidStorer) {
				Id, _ := uuid.Parse("010d9c50-646c-48af-8cb1-2db613102c4d")
				bidMock.On("GetHighestBid", mock.Anything).Return(20000.0, 10000.0, nil).Once()
				bidMock.On("GetBidStatus", mock.Anything).Return(repository.BidStatus{
					Id:   Id,
					Name: "user",
				}, nil).Once()
				bidMock.On("CreateBid", mock.Anything).Return(repository.Bids{}, nil).Once()
			},
			isErrorExpected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.setup(bidRepo)

			// test service
			_, err := service.CreateBid(test.input)

			if (err != nil) != test.isErrorExpected {
				t.Errorf("Test Failed, expected error to be %v, but got err %v", test.isErrorExpected, err != nil)
			}
		})
	}
}

func TestUpdateBid(t *testing.T) {
	bidRepo := mocks.NewBidStorer(t)
	service := NewService(bidRepo)

	tests := []struct {
		name            string
		input           dto.UpdateBidRequest
		bidder_id       string
		setup           func(bidMock *mocks.BidStorer)
		isErrorExpected bool
	}{
		{
			name: "Success for create bid",
			input: dto.UpdateBidRequest{
				ArtworkId: "016c8a24-4103-4148-881f-bf70391fb897",
				Amount:    40000.0,
			},
			bidder_id: "016c8a24-4103-4148-881f-bf70391fb897",
			setup: func(bidMock *mocks.BidStorer) {
				bidMock.On("GetHighestBid", mock.Anything).Return(20000.0, 10000.0, nil).Once()
				bidMock.On("UpdateBid", mock.Anything, mock.Anything).Return(repository.Bids{}, nil).Once()
			},
			isErrorExpected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.setup(bidRepo)

			// test service
			_, err := service.UpdateBid(test.input, test.bidder_id)

			if (err != nil) != test.isErrorExpected {
				t.Errorf("Test Failed, expected error to be %v, but got err %v", test.isErrorExpected, err != nil)
			}
		})
	}
}
