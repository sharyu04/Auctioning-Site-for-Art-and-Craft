package dto

type CreateBidRequest struct {
	Artwork_id string  `json:"artwork_id"`
	Amount     float64 `json:"amount"`
	Bidder_id  string
}

type UpdateBidRequest struct {
	ArtworkId string  `json:"artwork_id"`
	Amount    float64 `json:"amount"`
}

type DeleteBidRequest struct {
	BidId string `json:"bid_id"`
}

type DeleteBidResponse struct {
	Message string `json:"message"`
}
