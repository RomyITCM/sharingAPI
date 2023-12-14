package need_approve

import (
	"context"
	"database/sql"
	"smile-service/entities"

	"go.uber.org/zap"
)

const execSPDataNeedApproveProcess = `exec [sp_smile_sales_order_approve_insert]$1,$2,$3`

func GetDataNeedApproveProcess(ctx context.Context,
	db *sql.DB,
	log *zap.Logger,
	dataNeedApprove *entities.DataNeedApprove,
) (*entities.TransNo, error) {
	rows, err := db.QueryContext(
		ctx,
		execSPDataNeedApproveProcess,
		dataNeedApprove.TransNo,
		dataNeedApprove.CreatedBy,
		dataNeedApprove.CreaterdByIp,
	)

	if err != nil {
		log.Info("Exec DB",
			zap.Any("Exec DB", err),
		)
		return nil, err
	}
	defer rows.Close()

	result := &entities.TransNo{}
	for rows.Next() {
		if err := rows.Scan(
			&result.MsgType,
			&result.MsgError,
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

	return result, nil

}
