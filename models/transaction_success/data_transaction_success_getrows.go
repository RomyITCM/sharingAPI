package transaction_success

import (
	"context"
	"database/sql"
	"smile-service/entities"

	"go.uber.org/zap"
)

const execSPDataTransactionSuccesDetail = `exec [sp_smile_transaction_success_getrows]$1`

func GetDataTransactionSuccesDetail(ctx context.Context, db *sql.DB, trans_no string, log *zap.Logger) ([]*entities.DataTransactionSuccessDetail, error) {
	rows, err := db.QueryContext(
		ctx,
		execSPDataTransactionSuccesDetail,
		trans_no)

	if err != nil {
		log.Info("Exec DB",
			zap.Any("Exec DB", err),
		)
		return nil, err
	}
	defer rows.Close()

	data_transaction_details := make([]*entities.DataTransactionSuccessDetail, 0)
	for rows.Next() {
		data_transaction_detail := &entities.DataTransactionSuccessDetail{}

		if err := rows.Scan(
			&data_transaction_detail.ArticleDesc,
			&data_transaction_detail.Qty,
			&data_transaction_detail.Uom,
			&data_transaction_detail.Price,
			&data_transaction_detail.TotalPrice,
		); err != nil {
			log.Info("scan rows",
				zap.Any("rows", err),
			)
			return nil, err
		}

		data_transaction_details = append(data_transaction_details, data_transaction_detail)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return data_transaction_details, nil

}
