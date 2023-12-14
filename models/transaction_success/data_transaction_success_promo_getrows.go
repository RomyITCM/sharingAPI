package transaction_success

import (
	"context"
	"database/sql"
	"smile-service/entities"

	"go.uber.org/zap"
)

const execSPDataTransactionSuccesPromo = `exec [sp_smile_transaction_success_promo_getrows]$1`

func GetDataTransactionSuccesPromo(ctx context.Context, db *sql.DB, trans_no string, log *zap.Logger) ([]*entities.DataTransactionSuccessPromo, error) {
	rows, err := db.QueryContext(
		ctx,
		execSPDataTransactionSuccesPromo,
		trans_no)

	if err != nil {
		log.Info("Exec DB",
			zap.Any("Exec DB", err),
		)
		return nil, err
	}
	defer rows.Close()

	data_transaction_promos := make([]*entities.DataTransactionSuccessPromo, 0)
	for rows.Next() {
		data_transaction_promo := &entities.DataTransactionSuccessPromo{}

		if err := rows.Scan(
			&data_transaction_promo.PromoCode,
			&data_transaction_promo.PromoName,
			&data_transaction_promo.ResultDesc,
			&data_transaction_promo.AmountDisc,
		); err != nil {
			log.Info("scan rows",
				zap.Any("rows", err),
			)
			return nil, err
		}

		data_transaction_promos = append(data_transaction_promos, data_transaction_promo)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return data_transaction_promos, nil

}
