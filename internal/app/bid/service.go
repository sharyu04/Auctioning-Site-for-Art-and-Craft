package bid

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/sharyu04/Auctioning-Site-for-Art-and-Craft/internal/pkg/apperrors"
	"github.com/sharyu04/Auctioning-Site-for-Art-and-Craft/internal/pkg/dto"
	"github.com/sharyu04/Auctioning-Site-for-Art-and-Craft/internal/repository"
)

type service struct {
	bidRepo repository.BidStorer
}

type Service interface {
	CreateBid(bidDetails dto.CreateBidRequest) (bid repository.Bids, err error)
	UpdateBid(updateRequest dto.UpdateBidRequest, bidder_id string) (repository.Bids, error)
}

func NewService(bidRepo repository.BidStorer) Service {
	return &service{
		bidRepo: bidRepo,
	}
}

func (bs *service) CreateBid(bidDetails dto.CreateBidRequest) (bid repository.Bids, err error) {

	if bidDetails.Artwork_id == "" {
		return repository.Bids{}, apperrors.BadRequest{ErrorMsg: "ArtworkId Missing"}
	}

	highestBid, starting_price, err := bs.bidRepo.GetHighestBid(bidDetails.Artwork_id)
	if err != nil {
		return repository.Bids{}, err
	}

	if highestBid == 0 {
		if starting_price > bidDetails.Amount {
			errMsg := fmt.Sprintf("Bid must be equal to or above the starting price (%.2f)", starting_price)
			err := errors.New(errMsg)
			return repository.Bids{}, err
		}
	}

	if highestBid >= bidDetails.Amount {
		errMsg := fmt.Sprintf("Bid must be above the Highest bid (%.2f)", highestBid)
		return repository.Bids{}, apperrors.BadRequest{ErrorMsg: errMsg}
	}

	status, err := bs.bidRepo.GetBidStatus("live")
	if err != nil {
		return repository.Bids{}, err
	}
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

func (bs *service) UpdateBid(updateRequest dto.UpdateBidRequest, bidder_id string) (repository.Bids, error) {
	highestBid, _, err := bs.bidRepo.GetHighestBid(updateRequest.ArtworkId)
	if err != nil {
		return repository.Bids{}, err
	}

	if highestBid >= updateRequest.Amount {
		errMsg := fmt.Sprintf("Bid must be above the Highest bid (%.2f)", highestBid)
		return repository.Bids{}, apperrors.BadRequest{ErrorMsg: errMsg}
	}

	return bs.bidRepo.UpdateBid(updateRequest, bidder_id)

}
