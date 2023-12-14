package history_transaction

import (
	"context"
	"database/sql"
	"smile-service/entities"

	"go.uber.org/zap"
)

const execSPDataHistoryTransactionProduct = `exec [sp_smile_history_transaction_product_getrows]$1`

func GetDataHistoryTransactionProduct(ctx context.Context, db *sql.DB, transNo string, log *zap.Logger) ([]*entities.DataHistoryTransactionProducts, error) {
	rows, err := db.QueryContext(
		ctx,
		execSPDataHistoryTransactionProduct,
		transNo,
	)

	if err != nil {
		log.Info("Exec DB",
			zap.Any("Exec DB", err),
		)
		return nil, err
	}
	defer rows.Close()

	data_history_transactions_product := make([]*entities.DataHistoryTransactionProducts, 0)
	for rows.Next() {
		data_history_transaction_product := &entities.DataHistoryTransactionProducts{}

		if err := rows.Scan(
			&data_history_transaction_product.ArticleDesc,
			&data_history_transaction_product.Qty,
			&data_history_transaction_product.Uom,
			&data_history_transaction_product.UnitPrice,
			&data_history_transaction_product.TotalPrice,
			&data_history_transaction_product.Image,
			&data_history_transaction_product.WhId,
			&data_history_transaction_product.SiteName,
		); err != nil {
			log.Info("scan rows",
				zap.Any("rows", err),
			)
			return nil, err
		}

		data_history_transactions_product = append(data_history_transactions_product, data_history_transaction_product)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return data_history_transactions_product, nil

}
