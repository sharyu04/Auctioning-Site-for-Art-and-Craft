package bid

import (
	"errors"

	"github.com/google/uuid"
	"github.com/sharyu04/Auctioning-Site-for-Art-and-Craft/internal/pkg/dto"
	"github.com/sharyu04/Auctioning-Site-for-Art-and-Craft/internal/repository"
)

type service struct {
	bidRepo repository.BidStorer
}

type Service interface {
	CreateBid(bidDetails dto.CreateBidRequest) (bid repository.Bids, err error)
}

func NewService(bidRepo repository.BidStorer) Service {
	return &service{
		bidRepo: bidRepo,
	}
}

func (bs *service) CreateBid(bidDetails dto.CreateBidRequest) (bid repository.Bids, err error) {

	highestBid, starting_price, err := bs.bidRepo.GetHighestBid(bidDetails.Artwork_id)
	if err != nil {
		return repository.Bids{}, err
	}

	if highestBid == 0 {
		err = errors.New("Starting price is %f : . Bid above the starting price")
		return repository.Bids{}, err
	}

	if highestBid > bidDetails.Amount {
		err = errors.New("Highest bid is %f : . Bid above the highest bid")
		return repository.Bids{}, err
	}

	status, err := bs.bidRepo.GetBidStatus("live")
	artworkId, _ := uuid.Parse(bidDetails.Artwork_id)
	bidder, _ := uuid.Parse(bidDetails.Bidder_id)
	bidInfo := repository.Bids{
		Artwork_id: artworkId,
		Amount:     bidDetails.Amount,
		Status:     status.Id,
		Bidder_id:  bidder,
	}

	return bs.bidRepo.CreateBid(bidInfo)
}
