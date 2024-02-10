package repository

import (
	"errors"
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
	GetHighestBid(artWorkId string) (float64, float64, error)
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

	var ownerId uuid.UUID

	rows, err := bs.DB.Query("SELECT owner_id FROM artworks where id = $1", bid.Artwork_id)
	for rows.Next() {
		err := rows.Scan(&ownerId)
		if err != nil {
			return Bids{}, err
		}
	}

	if ownerId == bid.Bidder_id {
		err = errors.New("You cannot bid on your own artwork listing")
		return Bids{}, err
	}

	rows, err = bs.DB.Query("SELECT * FROM bids where artwork_id = $1 and bidder_id = $2", bid.Artwork_id, bid.Bidder_id)
	if err != nil {
		return Bids{}, err
	}

	i := 0
	for rows.Next() {
		i++
	}

	if i != 0 {
		err = errors.New("Bid from user already exists!")
		return Bids{}, err
	}

	bid.Id = uuid.New()
	bid.Created_at = time.Now()
	err = bs.DB.QueryRow("INSERT INTO bids(id, artwork_id, amount, status, bidder_id, created_at) VALUES($1, $2, $3, $4, $5, $6) RETURNING id",
		bid.Id, bid.Artwork_id, bid.Amount, bid.Status, bid.Bidder_id, bid.Created_at).Scan(&bid.Id)

	if err != nil {
		return Bids{}, err
	}

	err = bs.DB.QueryRow("UPDATE artworks SET highest_bid = $1 where id = $2",
		bid.Id, bid.Artwork_id).Scan()
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

func (as *bidStore) GetHighestBid(artWorkId string) (float64, float64, error) {

	rows, err := as.DB.Query("Select highest_bid, starting_price from artworks where id = $1", artWorkId)
	if err != nil {
		return 0, 0, err
	}

	var highestBidId uuid.UUID
	var starting_price float64
	for rows.Next() {
		err = rows.Scan(&highestBidId, &starting_price)
		if err != nil {
			return 0, 0, err
		}
	}

	if highestBidId == uuid.Nil {
		return 0, starting_price, nil
	}

	rows, err = as.DB.Query("Select amount from bids where id = $1", highestBidId)
	if err != nil {
		return 0, 0, err
	}
	var highestBid float64
	for rows.Next() {
		err = rows.Scan(&highestBid)
		if err != nil {
			return 0, 0, err
		}
	}

	return highestBid, starting_price, nil
}
