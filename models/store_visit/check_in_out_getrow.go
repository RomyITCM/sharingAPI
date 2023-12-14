package store_visit

import (
	"context"
	"database/sql"
	"smile-service/entities"

	"go.uber.org/zap"
)

const execSPCheckInOutGetrow = `exec [sp_smile_check_in_out_getrow]$1,$2`

func GetDataCheckInOut(ctx context.Context, db *sql.DB, id string, ptype string, log *zap.Logger) ([]*entities.DataCheckInOutGetrow, error) {
	rows, err := db.QueryContext(
		ctx,
		execSPCheckInOutGetrow,
		id,
		ptype,
	)

	if err != nil {
		log.Info(
			"Exec DB",
			zap.Any("Exec DB", err),
		)
		return nil, err
	}
	defer rows.Close()

	data_check_in_outs := make([]*entities.DataCheckInOutGetrow, 0)
	for rows.Next() {
		data_check_in_out := &entities.DataCheckInOutGetrow{}

		if err := rows.Scan(
			&data_check_in_out.ID,
			&data_check_in_out.VisitID,
			&data_check_in_out.CheckInTime,
			&data_check_in_out.CheckInAddress,
			&data_check_in_out.CheckInLatitude,
			&data_check_in_out.CheckInLongitude,
			&data_check_in_out.CheckOutTime,
			&data_check_in_out.CheckOutAddress,
			&data_check_in_out.CheckOutLatitude,
			&data_check_in_out.CheckOutLongitude,
			&data_check_in_out.PicCheckInOut,
		); err != nil {
			log.Info("scan rows",
				zap.Any("rows", err),
			)
			return nil, err
		}

		data_check_in_outs = append(data_check_in_outs, data_check_in_out)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return data_check_in_outs, nil
}
