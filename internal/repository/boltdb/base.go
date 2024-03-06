package repository

import (
	"context"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/sharyu04/Auctioning-Site-for-Art-and-Craft/internal/repository"
)

type BaseRepository struct {
	DB *sqlx.DB
}

type BaseTransaction struct {
	tx *sqlx.Tx
}

func (repo *BaseRepository) BeginTx(ctx context.Context) (repository.Transaction, error) {
	txObj, err := repo.DB.BeginTxx(ctx, nil)
	if err != nil {
		log.Printf("Error occoured while initiating database transaction: %v", err.Error())
		return nil, err
	}

	return &BaseTransaction{
		tx: txObj,
	}, nil
}

func (repo *BaseRepository) HandleTransaction(ctx context.Context, tx repository.Transaction, incomingErr error) (err error){
    if incomingErr!=nil{
        err = tx.Rollback();
        if err!=nil{
            log.Printf("Error occured while rollback database transaction: %v", err.Error())
            return
        }
        return
    }

    err = tx.Commit()
    if err!=nil{
        log.Printf("Error occured while commit database transaction: %v", err.Error())
        return
    }
    return
}

func (repo *BaseTransaction) Commit() error {
    return repo.tx.Commit()
}
func (repo *BaseTransaction) Rollback() error{
    return repo.tx.Rollback()
}
