package store_visit

import (
	"context"
	"database/sql"
	"smile-service/entities"

	"go.uber.org/zap"
)

const execSPDataStoreList = `exec [sp_smile_store_list_getrows]$1,$2`

func GetDataStoreList(ctx context.Context, db *sql.DB, search string, deptCode string, log *zap.Logger) ([]*entities.DataStoreList, error) {
	rows, err := db.QueryContext(
		ctx,
		execSPDataStoreList,
		search,
		deptCode,
	)

	if err != nil {
		log.Info("Exec DB",
			zap.Any("Exec DB", err),
		)
		return nil, err
	}
	defer rows.Close()

	data_store_lists := make([]*entities.DataStoreList, 0)

	for rows.Next() {
		data_store_list := &entities.DataStoreList{}

		if err := rows.Scan(
			&data_store_list.ShipToName,
			&data_store_list.ShipToNo,
			&data_store_list.CustomerName,
			&data_store_list.CustomerNo,
			&data_store_list.CustomerAddress,
			&data_store_list.CustomerType,
			&data_store_list.CustomerStatus,
			&data_store_list.CustomerTypeDesc,
			&data_store_list.BillToName,
			&data_store_list.BillToNo,
			&data_store_list.ShipToAddress,
			&data_store_list.Latitude,
			&data_store_list.Longitude,
			&data_store_list.LastReport,
			&data_store_list.Status,
			&data_store_list.PriceZone,
			&data_store_list.SiteId,
			&data_store_list.SiteName,
			&data_store_list.SiteAddress,
			&data_store_list.TripZone,
			&data_store_list.ShipToStatus,
			&data_store_list.LastPic,
			&data_store_list.LastCheckIn,
		); err != nil {
			log.Info("scan rows",
				zap.Any("rows", err),
			)
			return nil, err
		}

		data_store_lists = append(data_store_lists, data_store_list)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return data_store_lists, nil
}
