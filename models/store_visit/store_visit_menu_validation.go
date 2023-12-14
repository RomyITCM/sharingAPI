package store_visit

import (
	"context"
	"database/sql"
	"smile-service/entities"

	"go.uber.org/zap"
)

const execSPStoreVisitMenuValidation = `exec [sp_smile_store_visit_menu_validation]$1,$2`

func GetDataStoreVisitMenuValidation(ctx context.Context, db *sql.DB, ship_to string, created_by string, log *zap.Logger) ([]*entities.DataStoreVisitMenuValidation, error) {
	rows, err := db.QueryContext(
		ctx,
		execSPStoreVisitMenuValidation,
		ship_to,
		created_by,
	)

	if err != nil {
		log.Info("Exec DB",
			zap.Any("Exec DB", err),
		)
		return nil, err
	}
	defer rows.Close()

	data_check_menus := make([]*entities.DataStoreVisitMenuValidation, 0)

	for rows.Next() {
		data_check_menu := &entities.DataStoreVisitMenuValidation{}

		if err := rows.Scan(
			&data_check_menu.IsFreezerCheckedAll,
			&data_check_menu.IsStockCheckedAll,
			&data_check_menu.Latitude,
			&data_check_menu.Longitude,
		); err != nil {
			log.Info("scan rows",
				zap.Any("rows", err),
			)
			return nil, err
		}

		data_check_menus = append(data_check_menus, data_check_menu)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return data_check_menus, nil
}
