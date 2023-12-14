package transaction

import (
	"context"
	"database/sql"
	"smile-service/entities"

	"go.uber.org/zap"
)

const execSPDataTransaction = `exec [sp_smile_order_transaction_insert]$1,$2,$3,$4`

func InsertTransaction(
	ctx context.Context,
	db *sql.DB,
	data_transaction *entities.DataTransaction,
	log *zap.Logger) (*entities.QtyCart, error) {
	rows, err := db.QueryContext(
		ctx,
		execSPDataTransaction,
		data_transaction.ShipTo,
		data_transaction.ArticleNo,
		data_transaction.Qty,
		data_transaction.SalesMan,
	)

	if err != nil {
		log.Info(
			"Exec DB", zap.Any("Exec Db", err),
		)

		return nil, err
	}
	defer rows.Close()

	qty_cart := &entities.QtyCart{}
	for rows.Next() {
		if err := rows.Scan(
			&qty_cart.QtyCart,
		); err != nil {
			log.Info("scan rows",
				zap.Any("rows", err),
			)
			return nil, err
		}
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return qty_cart, nil
}
