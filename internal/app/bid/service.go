package bid

import (
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
	status, err := bs.bidRepo.GetBidStatus("live")
	artworkOwner, _ := uuid.Parse(bidDetails.Artwork_id)
	bidder, _ := uuid.Parse(bidDetails.Bidder_id)
	bidInfo := repository.Bids{
		Artwork_id: artworkOwner,
		Amount:     bidDetails.Amount,
		Status:     status.Id,
		Bidder_id:  bidder,
	}

	return bs.bidRepo.CreateBid(bidInfo)
}
