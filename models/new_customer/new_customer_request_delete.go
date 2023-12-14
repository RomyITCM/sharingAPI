package new_customer

import (
	"context"
	"database/sql"
	"smile-service/entities"

	"go.uber.org/zap"
)

const (
	execSPDataCustomerDelete = `exec [sp_smile_new_customer_request_delete] $1`
)

func DeleteCustomerRequest(
	ctx context.Context, db *sql.DB,
	data_new_customer *entities.CustomerRequestNo,
	log *zap.Logger) error {

	rows, err := db.QueryContext(
		ctx, execSPDataCustomerDelete,
		data_new_customer.CustomerRequestNo,
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
