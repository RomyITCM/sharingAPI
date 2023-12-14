package transaction_success

import (
	"context"
	"database/sql"
	"smile-service/entities"

	"go.uber.org/zap"
)

const execSPDataTransactionSucces = `exec [sp_smile_transaction_success_getrow]$1`

func GetDataTransactionSuccesHeader(ctx context.Context, db *sql.DB, trans_no string, log *zap.Logger) ([]*entities.DataTransactionSuccessHeader, error) {

	rows, err := db.QueryContext(
		ctx,
		execSPDataTransactionSucces,
		trans_no)

	if err != nil {
		log.Info("Exec DB",
			zap.Any("Exec DB", err),
		)
		return nil, err
	}
	defer rows.Close()

	data_transactions := make([]*entities.DataTransactionSuccessHeader, 0)
	for rows.Next() {
		data_transaction := &entities.DataTransactionSuccessHeader{}

		if err := rows.Scan(
			&data_transaction.TransNo,
			&data_transaction.TransDate,
			&data_transaction.CustPoNo,
			&data_transaction.CustPoDate,
			&data_transaction.PoExpDate,
			&data_transaction.ReqDate,
			&data_transaction.PaymentTerm,
			&data_transaction.Amount,
			&data_transaction.Disc,
			&data_transaction.Vat,
			&data_transaction.TotalAmount,
			&data_transaction.BillTo,
			&data_transaction.BillToAddress,
			&data_transaction.ShipTo,
			&data_transaction.ShipToAddress,
			&data_transaction.TotalCarton,
			&data_transaction.Status,
			&data_transaction.Attachment,
			&data_transaction.CustType,
		); err != nil {
			log.Info("scan rows",
				zap.Any("rows", err),
			)
			return nil, err
		}

		data_transactions = append(data_transactions, data_transaction)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return data_transactions, nil

}
