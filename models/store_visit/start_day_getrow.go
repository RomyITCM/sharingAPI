package store_visit

import (
	"context"
	"database/sql"
	"smile-service/entities"

	"go.uber.org/zap"
)

const execSPDataStartDayGetrow = `exec [sp_smile_start_day_getrow]$1`

func GetDataStartDay(ctx context.Context, db *sql.DB, id string, log *zap.Logger) ([]*entities.DataStartDayGetrow, error) {
	rows, err := db.QueryContext(
		ctx,
		execSPDataStartDayGetrow,
		id,
	)

	if err != nil {
		log.Info(
			"Exec DB",
			zap.Any("Exec DB", err),
		)
		return nil, err
	}
	defer rows.Close()

	data_start_days := make([]*entities.DataStartDayGetrow, 0)
	for rows.Next() {
		data_start_day := &entities.DataStartDayGetrow{}

		if err := rows.Scan(
			&data_start_day.Id,
			&data_start_day.Pic,
			&data_start_day.VisitDate,
			&data_start_day.Kilometer,
			&data_start_day.PicKilometer,
			&data_start_day.Vehicle,
			&data_start_day.PicVehiclePlate,
			&data_start_day.Address,
			&data_start_day.AddressLatitude,
			&data_start_day.AddressLongitude,
		); err != nil {
			log.Info("scan rows",
				zap.Any("rows", err),
			)
			return nil, err
		}

		data_start_days = append(data_start_days, data_start_day)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return data_start_days, nil
}
