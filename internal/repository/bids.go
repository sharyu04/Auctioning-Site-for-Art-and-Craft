package repository

import (
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type Bids struct {
	Id         uuid.UUID
	Artwork_id uuid.UUID
	Amount     float64
	Status     uuid.UUID
	Bidder_id  uuid.UUID
	Created_at time.Time
}

type BidStorer interface {
	CreateBid(bid Bids) (Bids, error)
	GetBidStatus(bidStatusName string) (BidStatus, error)
}

type BidStatus struct {
	Id   uuid.UUID
	Name string
}

type bidStore struct {
	DB *sqlx.DB
}

func NewBidRepo(db *sqlx.DB) BidStorer {
	return &bidStore{
		DB: db,
	}
}

func (bs *bidStore) CreateBid(bid Bids) (Bids, error) {
	bid.Id = uuid.New()
	bid.Created_at = time.Now()
	err := bs.DB.QueryRow("INSERT INTO bids(id, artwork_id, amount, status, bidder_id, created_at) VALUES($1, $2, $3, $4, $5, $6) RETURNING id",
		bid.Id, bid.Artwork_id, bid.Amount, bid.Status, bid.Bidder_id, bid.Created_at).Scan(&bid.Id)

	if err != nil {
		return Bids{}, err
	}
	return bid, nil
}

func (bs *bidStore) GetBidStatus(bidStatusName string) (BidStatus, error) {

	var bidStatus BidStatus

	rows, err := bs.DB.Query("SELECT id, name FROM bidstatus where name = $1", bidStatusName)
	if err != nil {
		return BidStatus{}, err
	}

	for rows.Next() {
		err := rows.Scan(&bidStatus.Id, &bidStatus.Name)
		if err != nil {
			return BidStatus{}, err
		}
	}

	return bidStatus, nil

}
