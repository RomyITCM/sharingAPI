package smd

import (
	"context"
	"database/sql"
	"smile-service/entities"

	"go.uber.org/zap"
)

const execSPDataSMDFreezer = `exec [sp_smile_smd_freezer_getrows]$1,$2,$3,$4`

func GetDataSMDFreezerList(ctx context.Context, db *sql.DB,
	customerNo string,
	billTo string,
	shipTo string,
	createdBy string,
	log *zap.Logger) ([]*entities.DataSMDFreezerList, error) {
	rows, err := db.QueryContext(
		ctx,
		execSPDataSMDFreezer,
		customerNo, billTo, shipTo, createdBy,
	)

	if err != nil {
		log.Info("Exec DB",
			zap.Any("Exec DB", err),
		)
		return nil, err
	}
	defer rows.Close()

	data_smd_frerezer_lists := make([]*entities.DataSMDFreezerList, 0)

	for rows.Next() {
		data_smd_freezer_list := &entities.DataSMDFreezerList{}

		if err := rows.Scan(
			&data_smd_freezer_list.Freezer,
			&data_smd_freezer_list.Location,
			&data_smd_freezer_list.Power,
			&data_smd_freezer_list.Condition,
			&data_smd_freezer_list.OutsideCondition,
			&data_smd_freezer_list.TransNo,
			&data_smd_freezer_list.SerialNo,
		); err != nil {
			log.Info("scan rows",
				zap.Any("rows", err),
			)
			return nil, err
		}

		data_smd_frerezer_lists = append(data_smd_frerezer_lists, data_smd_freezer_list)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return data_smd_frerezer_lists, nil
}
