package store_visit

import (
	"context"
	"database/sql"
	"smile-service/entities"

	"go.uber.org/zap"
)

const execSPDataStoreVisit = `exec [sp_smile_store_visit_getrows]$1,$2`

func GetDataStoreVisit(ctx context.Context, db *sql.DB, pic string, visitDate string, log *zap.Logger) ([]*entities.DataStoreVisit, error) {
	rows, err := db.QueryContext(
		ctx,
		execSPDataStoreVisit,
		pic,
		visitDate,
	)

	if err != nil {
		log.Info(
			"Exec DB",
			zap.Any("Exec DB", err),
		)
		return nil, err
	}
	defer rows.Close()

	data_store_visits := make([]*entities.DataStoreVisit, 0)
	for rows.Next() {
		data_store_visit := &entities.DataStoreVisit{}

		if err := rows.Scan(
			&data_store_visit.Id,
			&data_store_visit.DetailID,
			&data_store_visit.StartDayAddress,
			&data_store_visit.StartDayTime,
			&data_store_visit.EndDayAddress,
			&data_store_visit.EndDayTime,
			&data_store_visit.CustomerNo,
			&data_store_visit.CustomerName,
			&data_store_visit.CustomerAddress,
			&data_store_visit.CustomerType,
			&data_store_visit.CustomerStatus,
			&data_store_visit.CustomerTypeDesc,
			&data_store_visit.ShipTo,
			&data_store_visit.StoreAdddress,
			&data_store_visit.StoreName,
			&data_store_visit.ShipToName,
			&data_store_visit.CheckInTime,
			&data_store_visit.CheckInAddress,
			&data_store_visit.CheckOutTime,
			&data_store_visit.CheckOutAddress,
			&data_store_visit.ShipToID,
			&data_store_visit.PriceZone,
			&data_store_visit.SiteId,
			&data_store_visit.SiteName,
			&data_store_visit.SiteAddress,
			&data_store_visit.TripZone,
			&data_store_visit.ShipToStatus,
			&data_store_visit.Latitude,
			&data_store_visit.Longitude,
			&data_store_visit.TotalEfectiveCall,
			&data_store_visit.TotalSales,
			&data_store_visit.TotalAvailableStore,
			&data_store_visit.TotalUnAvailableStore,
		); err != nil {
			log.Info("scan rows",
				zap.Any("rows", err),
			)
			return nil, err
		}

		data_store_visits = append(data_store_visits, data_store_visit)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return data_store_visits, nil
}
