package store_visit

import (
	"context"
	"database/sql"
	"smile-service/entities"

	"go.uber.org/zap"
)

const execSPCheckFreezerGetrow = `exec [sp_smile_check_freezer_getrow]$1,$2`

func GetDataCheckFreezerGetrow(ctx context.Context, db *sql.DB, trans_no string, serial_no string, log *zap.Logger) ([]*entities.DataCheckFreezerGetrow, error) {
	rows, err := db.QueryContext(
		ctx,
		execSPCheckFreezerGetrow,
		trans_no,
		serial_no,
	)

	if err != nil {
		log.Info("Exec DB",
			zap.Any("Exec DB", err),
		)
		return nil, err
	}
	defer rows.Close()

	data_check_freezers := make([]*entities.DataCheckFreezerGetrow, 0)

	for rows.Next() {
		data_check_freezer := &entities.DataCheckFreezerGetrow{}

		if err := rows.Scan(
			&data_check_freezer.TransNo,
			&data_check_freezer.Mode,
			&data_check_freezer.ArticleDescription,
			&data_check_freezer.SerialNo,
			&data_check_freezer.ImagesType,
			&data_check_freezer.ImagesName,
			&data_check_freezer.FreezerAvailable,
			&data_check_freezer.AvailableNote,
			&data_check_freezer.Brand,
			&data_check_freezer.NoteBrand,
			&data_check_freezer.Type,
			&data_check_freezer.Capacity,
			&data_check_freezer.Status,
			&data_check_freezer.NoteStatus,
			&data_check_freezer.Location,
			&data_check_freezer.NoteLocation,
			&data_check_freezer.IceThickness,
			&data_check_freezer.Temperature,
			&data_check_freezer.FreezerUse,
			&data_check_freezer.NoteFreezerUse,
			&data_check_freezer.CreatedBy,
			&data_check_freezer.CreatedDate,
			&data_check_freezer.CreatedByIP,
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
