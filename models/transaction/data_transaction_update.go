package transaction

import (
	"context"
	"database/sql"
	"smile-service/entities"

	"go.uber.org/zap"
)

const execSPDataTransactionUpdate = `exec [sp_smile_sales_order_detail_update]$1,$2,$3,$4`
const execSPDataTransactionDelete = `exec [sp_smile_sales_order_detail_delete]$1`
const execSPDataTransactionUpdateHeader = `exec [sp_smile_sales_order_update]$1`

func UpdateQtyTransaction(
	ctx context.Context,
	db *sql.DB,
	data_transaction *entities.UpdateQtyCart,
	log *zap.Logger) error {

	//delete detail SO
	_, err := db.QueryContext(
		ctx,
		execSPDataTransactionDelete,
		data_transaction.TransNo,
	)

	if err != nil {
		log.Info(
			"Exec DB", zap.Any("Exec Db", err),
		)

		return err
	}

	line_no := 10

	//insert data detail SO
	for _, v := range data_transaction.Detail {
		rows, err := db.Query(
			execSPDataTransactionUpdate,
			data_transaction.TransNo,
			v.ArticleNo,
			v.Qty,
			line_no,
		)

		if err != nil {
			log.Info(
				"Exec DB", zap.Any("Exec Db", err),
			)

			return err
		}
		defer rows.Close()

		line_no = line_no + 10

	}

	//update header SO
	_, err = db.QueryContext(
		ctx,
		execSPDataTransactionUpdateHeader,
		data_transaction.TransNo,
	)

	if err != nil {
		log.Info(
			"Exec DB", zap.Any("Exec Db", err),
		)

		return err
	}

	return nil
}
