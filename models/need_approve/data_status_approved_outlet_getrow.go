package need_approve

import (
	"context"
	"database/sql"
	"smile-service/entities"

	"go.uber.org/zap"
)

const execSPStatusApprovedOutlet = `exec [sp_smile_status_approved_outlet_getrow] $1, $2`

func GetDataOutletStatusApproved(ctx context.Context, db *sql.DB, user_no string,
	outlet_id string, log *zap.Logger) ([]*entities.DataApproveStatusUser, error) {

	rows, err := db.QueryContext(
		ctx,
		execSPStatusApprovedOutlet,
		user_no,
		outlet_id,
	)

	if err != nil {
		log.Info(
			"Exec DB",
			zap.Any("Exec DB", err),
		)
		return nil, err
	}
	defer rows.Close()

	status_approvals := make([]*entities.DataApproveStatusUser, 0)
	for rows.Next() {
		status_approval := &entities.DataApproveStatusUser{}

		if err := rows.Scan(
			&status_approval.UserId,
			&status_approval.Status,
			&status_approval.Description,
		); err != nil {
			log.Info("scan rows",
				zap.Any("rows", err),
			)
			return nil, err
		}

		status_approvals = append(status_approvals, status_approval)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return status_approvals, nil
}
