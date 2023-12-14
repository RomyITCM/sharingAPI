package need_approve

import (
	"context"
	"database/sql"
	"smile-service/entities"

	"go.uber.org/zap"
)

const execSPDataNeedApproveReject = `exec [sp_sales_order_rejected]$1,$2,$3`

func GetDataNeedApproveReject(ctx context.Context,
	db *sql.DB,
	log *zap.Logger,
	dataNeedApproveReject *entities.DataNeedApproveReject,
) error {
	rows, err := db.QueryContext(
		ctx,
		execSPDataNeedApproveReject,
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
