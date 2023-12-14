package transaction

import (
	"context"
	"database/sql"
	"smile-service/entities"

	"go.uber.org/zap"
)

const execSPDataTransactionPromo = `exec [sp_smile_promo_insert]$1,$2,$3`

func InsertTransactionPromo(
	ctx context.Context,
	db *sql.DB,
	data_transaction_promo *entities.DataTransactionPromo,
	log *zap.Logger) error {
	rows, err := db.QueryContext(
		ctx,
		execSPDataTransactionPromo,
		data_transaction_promo.ShipTo,
		data_transaction_promo.SalesMan,
		data_transaction_promo.PaymentTerm,
	)

	if err != nil {
		log.Info(
			"Exec DB", zap.Any("Exec Db", err),
		)

		return err
	}
	defer rows.Close()

	return nil
}
