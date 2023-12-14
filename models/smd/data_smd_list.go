package smd

import (
	"context"
	"database/sql"
	"smile-service/entities"

	"go.uber.org/zap"
)

const execSPDataSMD = `exec [sp_smile_smd_getrows]$1`

func GetDataSMD(ctx context.Context, db *sql.DB, search string, log *zap.Logger) ([]*entities.DataSMDList, error) {
	rows, err := db.QueryContext(
		ctx,
		execSPDataSMD,
		search,
	)

	if err != nil {
		log.Info("Exec DB",
			zap.Any("Exec DB", err),
		)
		return nil, err
	}
	defer rows.Close()

	data_smd_lists := make([]*entities.DataSMDList, 0)

	for rows.Next() {
		data_smd_list := &entities.DataSMDList{}

		if err := rows.Scan(
			&data_smd_list.ShipToName,
			&data_smd_list.ShipToNo,
			&data_smd_list.CustomerName,
			&data_smd_list.CustomerNo,
			&data_smd_list.BillToName,
			&data_smd_list.BillToNo,
			&data_smd_list.ShipToAddress,
			&data_smd_list.LastReport,
		); err != nil {
			log.Info("scan rows",
				zap.Any("rows", err),
			)
			return nil, err
		}

		data_smd_lists = append(data_smd_lists, data_smd_list)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return data_smd_lists, nil
}
