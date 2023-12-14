package history_transaction

import (
	"context"
	"database/sql"
	"smile-service/entities"

	"go.uber.org/zap"
)

const execSPDataHistoryTransactionDetail = `exec [sp_smile_history_transaction_getrow]$1`

func GetDataHistoryTransactionDetail(ctx context.Context, db *sql.DB, transNo string, log *zap.Logger) ([]*entities.DataHistoryTransactionDelivery, error) {
	rows, err := db.QueryContext(
		ctx,
		execSPDataHistoryTransactionDetail,
		transNo,
	)

	if err != nil {
		log.Info("Exec DB",
			zap.Any("Exec DB", err),
		)
		return nil, err
	}
	defer rows.Close()

	data_history_transactions_delivery := make([]*entities.DataHistoryTransactionDelivery, 0)
	for rows.Next() {
		data_history_transaction_delivery := &entities.DataHistoryTransactionDelivery{}

		if err := rows.Scan(
			&data_history_transaction_delivery.TransDate,
			&data_history_transaction_delivery.Salesman,
			&data_history_transaction_delivery.Status,
			&data_history_transaction_delivery.RequestRecvDate,
			&data_history_transaction_delivery.ReceiveDate,
			&data_history_transaction_delivery.ShipTo,
			&data_history_transaction_delivery.ShipToAddress,
			&data_history_transaction_delivery.CustomerPoNo,
			&data_history_transaction_delivery.AttacmentPO,
			&data_history_transaction_delivery.PaymentTerm,
			&data_history_transaction_delivery.PaymentDue,
			&data_history_transaction_delivery.PoExpDate,
			&data_history_transaction_delivery.BillTo,
			&data_history_transaction_delivery.BillToAddress,
			&data_history_transaction_delivery.Amount,
			&data_history_transaction_delivery.Disc,
			&data_history_transaction_delivery.Tax,
			&data_history_transaction_delivery.TotalPrice,
			&data_history_transaction_delivery.Reason,
			&data_history_transaction_delivery.Attachment,
		); err != nil {
			log.Info("scan rows",
				zap.Any("rows", err),
			)
			return nil, err
		}

		data_history_transactions_delivery = append(data_history_transactions_delivery, data_history_transaction_delivery)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return data_history_transactions_delivery, nil

}
