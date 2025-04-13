package helpers

import (
	"context"
	"database/sql"
	"log"

	"github.com/mmhtoo/go-logger-api/config"
)

func WithTx(
	ctx context.Context,
	database *config.Database,
	callback func(trxn *sql.Tx) (any,error),
) (any,error) {
	tx, err := database.Connection.BeginTx(ctx,nil)
	if err != nil {
		return nil, err
	}
	defer func(){
		if r := recover(); r != nil {
			log.Printf("Error while executing transaction: %s", r)
			tx.Rollback()
		} else if err != nil {
			_ = tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	return callback(tx)
}