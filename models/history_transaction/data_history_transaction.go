package history_transaction

import (
	"context"
	"database/sql"
	"smile-service/entities"

	"go.uber.org/zap"
)

const execSPDataHistoryTransaction = `exec [sp_smile_history_transaction_getrows]$1,$2,$3,$4`

func GetDataHistoryTransaction(ctx context.Context, db *sql.DB, page string, time_stamp string, search string, user_id string, log *zap.Logger) ([]*entities.DataHistoryTransaction, error) {
	rows, err := db.QueryContext(
		ctx,
		execSPDataHistoryTransaction,
		page,
		time_stamp,
		search,
		user_id,
	)

	if err != nil {
		log.Info("Exec DB",
			zap.Any("Exec DB", err),
		)
		return nil, err
	}
	defer rows.Close()

	data_history_transactions := make([]*entities.DataHistoryTransaction, 0)
	for rows.Next() {
		data_history_transaction := &entities.DataHistoryTransaction{}

		if err := rows.Scan(
			&data_history_transaction.TransNo,
			&data_history_transaction.TransDate,
			&data_history_transaction.Status,
			&data_history_transaction.ShipTo,
			&data_history_transaction.PODate,
			&data_history_transaction.TotalPrice,
			&data_history_transaction.ArticleDesc,
			&data_history_transaction.Qty,
			&data_history_transaction.Note,
			&data_history_transaction.UrlImage,
		); err != nil {
			log.Info("scan rows",
				zap.Any("rows", err),
			)
			return nil, err
		}

		data_history_transactions = append(data_history_transactions, data_history_transaction)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return data_history_transactions, nil

}
