package dto

type CreateBidRequest struct {
	Artwork_id string
	Amount     float64
	Bidder_id  string
}
