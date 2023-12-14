package need_approve

import (
	"context"
	"database/sql"
	"smile-service/entities"

	"go.uber.org/zap"
)

const (
	execSPApproveCustomerRequest = `exec [sp_smile_customer_request_approval_insert] $1, $2, $3, $4`
)

func ApproveCustomerRequest(
	ctx context.Context, db *sql.DB,
	data_new_customer *entities.DataNeedApprove,
	log *zap.Logger) (*entities.CustomerNo, error) {

	rows, err := db.QueryContext(
		ctx, execSPApproveCustomerRequest,
		data_new_customer.TransNo,
		data_new_customer.NextStatus,
		data_new_customer.CreatedBy,
		data_new_customer.CreaterdByIp,
	)

	if err != nil {
		log.Info(
			"Exec DB", zap.Any("Exec Db", err),
		)

		return nil, err
	}
	defer rows.Close()

	customer_no := &entities.CustomerNo{}
	for rows.Next() {
		if err := rows.Scan(
			&customer_no.CustomerNo,
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

	return customer_no, nil
}
