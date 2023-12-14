package store_visit

import (
	"context"
	"database/sql"
	"smile-service/entities"

	"go.uber.org/zap"
)

const execSPCheckFreezerList = `exec [sp_smile_check_freezer_getrows]$1,$2`

func GetDataCheckFreezerList(ctx context.Context, db *sql.DB, ship_to string, search string, log *zap.Logger) ([]*entities.DataCheckFreezerGetrows, error) {
	rows, err := db.QueryContext(
		ctx,
		execSPCheckFreezerList,
		ship_to,
		search,
	)

	if err != nil {
		log.Info("Exec DB",
			zap.Any("Exec DB", err),
		)
		return nil, err
	}
	defer rows.Close()

	data_check_freezers := make([]*entities.DataCheckFreezerGetrows, 0)

	for rows.Next() {
		data_check_freezer := &entities.DataCheckFreezerGetrows{}

		if err := rows.Scan(
			&data_check_freezer.TransNo,
			&data_check_freezer.SerialNo,
			&data_check_freezer.CustomerNo,
			&data_check_freezer.BillTo,
			&data_check_freezer.ShipTo,
			&data_check_freezer.ArticleNo,
			&data_check_freezer.ArticleDescription,
			&data_check_freezer.FreezerAvailable,
			&data_check_freezer.FreezerFrom,
			&data_check_freezer.FreezerUse,
			&data_check_freezer.FreezerLocation,
			&data_check_freezer.FreezerChecked,
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
