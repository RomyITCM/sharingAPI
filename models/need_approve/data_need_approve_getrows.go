package need_approve

import (
	"context"
	"database/sql"
	"smile-service/entities"

	"go.uber.org/zap"
)

const execSPDataNeedApprove = `exec [sp_smile_need_approve_getrows]$1,$2`

func GetDataNeedApprove(ctx context.Context, db *sql.DB, user_id, search string, log *zap.Logger) ([]*entities.DataNeedApproveList, error) {
	rows, err := db.QueryContext(
		ctx,
		execSPDataNeedApprove,
		user_id,
		search,
	)

	if err != nil {
		log.Info("Exec DB",
			zap.Any("Exec DB", err),
		)
		return nil, err
	}
	defer rows.Close()

	data_need_approves := make([]*entities.DataNeedApproveList, 0)
	for rows.Next() {
		data_need_approve := &entities.DataNeedApproveList{}

		if err := rows.Scan(
			&data_need_approve.TransNo,
			&data_need_approve.TransDate,
			&data_need_approve.ShipTo,
			&data_need_approve.ReqDelvDate,
			&data_need_approve.TotalPrice,
			&data_need_approve.ArticleDesc,
			&data_need_approve.Qty,
			&data_need_approve.Note,
			&data_need_approve.UrlImage,
			&data_need_approve.UserName,
		); err != nil {
			log.Info("scan rows",
				zap.Any("rows", err),
			)
			return nil, err
		}

		data_need_approves = append(data_need_approves, data_need_approve)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return data_need_approves, nil

}
