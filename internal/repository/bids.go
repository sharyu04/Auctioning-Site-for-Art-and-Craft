package repository

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/sharyu04/Auctioning-Site-for-Art-and-Craft/internal/pkg/apperrors"
	"github.com/sharyu04/Auctioning-Site-for-Art-and-Craft/internal/pkg/dto"
)

type Bids struct {
	Id         uuid.UUID `db:"id"`
	Artwork_id uuid.UUID `db:"artwork_id"`
	Amount     float64   `db:"amount"`
	Status     uuid.UUID `db:"status"`
	Bidder_id  uuid.UUID `db:"bidder_id"`
	Created_at time.Time `db:"created_at"`
	//updated at
}

type BidStorer interface {
	CreateBid(bid Bids) (Bids, error)
	GetBidStatus(bidStatusName string) (BidStatus, error)
	GetHighestBid(artWorkId string) (float64, float64, error)
	UpdateBid(bid dto.UpdateBidRequest, bidder_id string) (Bids, error)
	DeleteBid(user_id uuid.UUID, role string, bid_id uuid.UUID) error
}

type BidStatus struct {
	Id   uuid.UUID `db:"id"`
	Name string    `db:"name"`
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
		// err = errors.New("You cannot bid on your own artwork listing")
		return Bids{}, apperrors.BadRequest{ErrorMsg: "Cannot bid on your own artwork"}
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
		// err = errors.New("Bid from user already exists!")
		return Bids{}, apperrors.BadRequest{ErrorMsg: "Duplicate bid"}
	}

	bid.Id = uuid.New()
	bid.Created_at = time.Now()
	err = bs.DB.QueryRow("INSERT INTO bids(id, artwork_id, amount, status, bidder_id, created_at) VALUES($1, $2, $3, $4, $5, $6) RETURNING id",
		bid.Id, bid.Artwork_id, bid.Amount, bid.Status, bid.Bidder_id, bid.Created_at).Scan(&bid.Id)

	if err != nil {
		return Bids{}, err
	}

	var updateQueryId uuid.UUID
	err = bs.DB.QueryRow("UPDATE artworks SET highest_bid = $1 where id = $2 RETURNING id", bid.Id, bid.Artwork_id).Scan(&updateQueryId)
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

func (bs *bidStore) GetHighestBid(artWorkId string) (float64, float64, error) {

	artWorkIdUuid, _ := uuid.Parse(artWorkId)
	rows, err := bs.DB.Query("Select highest_bid, starting_price from artworks where id = $1", artWorkIdUuid)
	if err != nil {
		return 0, 0, err
	}

	var highestBidId uuid.UUID
	var starting_price float64
	i := 0
	for rows.Next() {
		err = rows.Scan(&highestBidId, &starting_price)
		if err != nil {
			return 0, 0, err
		}
		i++
	}

	if i == 0 {
		errMsg := fmt.Sprintf("Artworks not found with id: %v", artWorkIdUuid)
		return 0, 0, apperrors.BadRequest{ErrorMsg: errMsg}
	}

	if highestBidId == uuid.Nil {
		return 0, starting_price, nil
	}

	rows, err = bs.DB.Query("Select amount from bids where id = $1", highestBidId)
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

func (bs *bidStore) UpdateBid(bid dto.UpdateBidRequest, bidder_id string) (Bids, error) {

	var bidId uuid.UUID

	err := bs.DB.QueryRow("UPDATE bids SET amount = $1 where artwork_id = $2 and bidder_id = $3 RETURNING id",
		bid.Amount, bid.ArtworkId, bidder_id).Scan(&bidId)

	if err != nil {
		// err = errors.New("No bid exist on the artwork from the logged in user!")
		return Bids{}, apperrors.BadRequest{ErrorMsg: "Bid not found"}
	}

	// if bidId == uuid.Nil {
	// 	err = errors.New("No bid exist on the artwork from the logged in user!")
	// 	return Bids{}, err
	// }

	var updateQueryId uuid.UUID
	err = bs.DB.QueryRow("UPDATE artworks SET highest_bid = $1 where id = $2 RETURNING id", bidId, bid.ArtworkId).Scan(&updateQueryId)
	if err != nil {
		return Bids{}, err
	}

	rows, err := bs.DB.Query("SELECT * FROM bids where id = $1", bidId)
	if err != nil {
		return Bids{}, err
	}

	var res Bids
	for rows.Next() {
		err = rows.Scan(&res.Id, &res.Artwork_id, &res.Amount, &res.Status, &res.Bidder_id, &res.Created_at)
		if err != nil {
			return Bids{}, err
		}
	}

	return res, nil
}

func (bs *bidStore) GetBidsByArtworkId(artwork_id uuid.UUID) ([]Bids, error) {

	rows, err := bs.DB.Query("SELECT * FROM bids where artwork_id = $1", artwork_id)
	if err != nil {
		return []Bids{}, err
	}

	var bidsList []Bids
	for rows.Next() {
		var res Bids
		err = rows.Scan(&res.Id, &res.Artwork_id, &res.Amount, &res.Status, &res.Bidder_id, &res.Created_at)
		if err != nil {
			return []Bids{}, err
		}
		bidsList = append(bidsList, res)
	}

	return bidsList, nil
}

func (bs *bidStore) DeleteBid(user_id uuid.UUID, role string, bid_id uuid.UUID) error {
	// get artwork id and check for authorization
	var artwork_id uuid.UUID
	if role == "user" {
		row, err := bs.DB.Query("Select artwork_id from bids where bidder_id = $1 and id = $2", user_id, bid_id)
		if err != nil {
			return err
		}

		i := 0
		for row.Next() {
			if err := row.Scan(&artwork_id); err != nil {
				return err
			}
			if artwork_id == uuid.Nil {
				return apperrors.UnAuthorizedAccess{ErrorMsg: "Unauthorized Delete"}
			}
			i++
		}
		if i == 0 {
			return apperrors.UnAuthorizedAccess{ErrorMsg: "Unauthorized Delete"}
		}

	} else {
		row, err := bs.DB.Query("Select artwork_id from bids where id = $1", bid_id)

		if err != nil {
			return err
		}

		i := 0
		for row.Next() {
			if err := row.Scan(&artwork_id); err != nil {
				return err
			}
			if artwork_id == uuid.Nil {
				return apperrors.UnAuthorizedAccess{ErrorMsg: "Unauthorized Delete"}
			}
			i++
		}

		if i == 0 {
			return apperrors.BadRequest{ErrorMsg: "No such bid exists"}
		}

	}

	//set new highest bid in the artwork
	var newHighestBid uuid.UUID
	row, err := bs.DB.Query("Select id, MAX(amount) from bids group by id having artwork_id = $1 and id != $2", artwork_id, bid_id)
	if err != nil {
		return err
	}
	for row.Next() {
		var amount float64
		if err := row.Scan(&newHighestBid, &amount); err != nil {
			return err
		}
	}

	if newHighestBid == uuid.Nil {
		_, err := bs.DB.Query("Update artworks set highest_bid = null where id = $1", artwork_id)
		if err != nil {
			return err
		}
	} else {
		_, err := bs.DB.Query("Update artworks set highest_bid = $1 where id = $2", newHighestBid, artwork_id)
		if err != nil {
			return err
		}
	}

	//delete bid
	_, err = bs.DB.Query("Delete from bids where id = $1", bid_id)
	if err != nil {
		return err
	}

	return nil

}
