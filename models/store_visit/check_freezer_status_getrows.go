package store_visit

import (
	"context"
	"database/sql"
	"smile-service/entities"

	"go.uber.org/zap"
)

const execSPCheckFreezerStatus = `exec [sp_smile_check_freezer_status_getrows]`

func GetDataCheckFreezerStatus(ctx context.Context, db *sql.DB, log *zap.Logger) ([]*entities.DataCheckFreezerStatus, error) {
	rows, err := db.QueryContext(
		ctx,
		execSPCheckFreezerStatus,
	)

	if err != nil {
		log.Info("Exec DB",
			zap.Any("Exec DB", err),
		)
		return nil, err
	}
	defer rows.Close()

	data_check_freezers := make([]*entities.DataCheckFreezerStatus, 0)

	for rows.Next() {
		data_check_freezer := &entities.DataCheckFreezerStatus{}

		if err := rows.Scan(
			&data_check_freezer.Value,
			&data_check_freezer.Text,
		); err != nil {
			log.Info("scan rows",
				zap.Any("rows", err),
			)
			return nil, err
		}

		data_check_freezers = append(data_check_freezers, data_check_freezer)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return data_check_freezers, nil
}
