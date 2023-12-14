package need_approve

import (
	"context"
	"database/sql"
	"smile-service/entities"

	"go.uber.org/zap"
)

const execSPCustomerRequestReject = `exec [sp_smile_customer_request_reject] $1, $2, $3`

func RejectCustomerRequest(ctx context.Context,
	db *sql.DB,
	log *zap.Logger,
	dataNeedApproveReject *entities.DataNeedApproveReject,
) error {

	rows, err := db.QueryContext(
		ctx,
		execSPCustomerRequestReject,
		dataNeedApproveReject.TransNo,
		dataNeedApproveReject.Reason,
		dataNeedApproveReject.CancelledBy,
	)

	if err != nil {
		log.Info("Exec DB",
			zap.Any("Exec DB", err),
		)
		return err
	}
	defer rows.Close()

	if err = rows.Err(); err != nil {
		return err
	}

	return nil
}
