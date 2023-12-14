package smd

import (
	"context"
	"database/sql"
	"smile-service/entities"

	"go.uber.org/zap"
)

const execSPDataSMDFreezerGetRow = `exec [sp_smile_smd_freezer_getrow]$1,$2`

func GetDataSMDFreezer(ctx context.Context, db *sql.DB,
	transNo string,
	serialNo string,
	log *zap.Logger) ([]*entities.DataSMDFreezer, error) {
	rows, err := db.QueryContext(
		ctx,
		execSPDataSMDFreezerGetRow,
		transNo, serialNo,
	)

	if err != nil {
		log.Info("Exec DB",
			zap.Any("Exec DB", err),
		)
		return nil, err
	}
	defer rows.Close()

	data_smd_frerezers := make([]*entities.DataSMDFreezer, 0)

	for rows.Next() {
		data_smd_freezer := &entities.DataSMDFreezer{}

		if err := rows.Scan(
			&data_smd_freezer.FreezerAvailable,
			&data_smd_freezer.SerialNo,
			&data_smd_freezer.Merk,
			&data_smd_freezer.NoteMerk,
			&data_smd_freezer.Capacity,
			&data_smd_freezer.Location,
			&data_smd_freezer.Access,
			&data_smd_freezer.NoteAccess,
			&data_smd_freezer.Power,
			&data_smd_freezer.NotePower,
			&data_smd_freezer.Condition,
			&data_smd_freezer.NoteCondition,
			&data_smd_freezer.OutsideCondition,
			&data_smd_freezer.NoteOutsideCondition,
			&data_smd_freezer.Suhu,
		); err != nil {
			log.Info("scan rows",
				zap.Any("rows", err),
			)
			return nil, err
		}

		data_smd_frerezers = append(data_smd_frerezers, data_smd_freezer)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return data_smd_frerezers, nil
}
