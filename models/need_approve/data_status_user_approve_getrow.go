package need_approve

import (
	"context"
	"database/sql"
	"smile-service/entities"

	"go.uber.org/zap"
)

const execSPStatusUserApprove = `exec [sp_smile_status_approved_getrow] $1,$2`

func GetStatusUserApprove(ctx context.Context, db *sql.DB, user_no string,
	customer_request_no string, log *zap.Logger) ([]*entities.DataUserApprove, error) {

	rows, err := db.QueryContext(
		ctx,
		execSPStatusUserApprove,
		user_no,
		customer_request_no,
	)

	if err != nil {
		log.Info(
			"Exec DB",
			zap.Any("Exec DB", err),
		)
		return nil, err
	}
	defer rows.Close()

	data_user_approves := make([]*entities.DataUserApprove, 0)
	for rows.Next() {
		data_user_approve := &entities.DataUserApprove{}

		if err := rows.Scan(
			&data_user_approve.UserId,
			&data_user_approve.Status,
			&data_user_approve.Description,
		); err != nil {
			log.Info("scan rows",
				zap.Any("rows", err),
			)
			return nil, err
		}

		data_user_approves = append(data_user_approves, data_user_approve)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return data_user_approves, nil
}
