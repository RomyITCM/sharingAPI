package need_approve

import (
	"context"
	"database/sql"
	"smile-service/entities"

	"go.uber.org/zap"
)

const (
	execSPApproveCustomerAddress = `exec [sp_smile_approve_customer_address_insert] $1,$2,$3,$4`
)

func ApproveCustomerAddress(
	ctx context.Context, db *sql.DB,
	data_customer_address *entities.DataNeedApprove,
	log *zap.Logger) error {

	rows, err := db.QueryContext(
		ctx, execSPApproveCustomerAddress,
		data_customer_address.TransNo,
		data_customer_address.NextStatus,
		data_customer_address.CreatedBy,
		data_customer_address.CreaterdByIp,
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
