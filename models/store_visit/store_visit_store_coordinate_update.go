package store_visit

import (
	"context"
	"database/sql"
	"smile-service/entities"
)

const (
	execSPUpdateStoreCoordinate = `exec [sp_smile_store_coordinate_update]$1,$2,$3,$4,$5`
)

func DataStoreCoordinateUpdate(ctx context.Context, db *sql.DB,
	CustomerNo string,
	ShipTo string,
	Latitude string,
	Longitude string,
	CreatedBy string) error {
	rows, err := db.QueryContext(ctx, execSPUpdateStoreCoordinate,
		CustomerNo,
		ShipTo,
		Latitude,
		Longitude,
		CreatedBy,
	)
	if err != nil {
		return err
	}
	defer rows.Close()

	return nil
}

func UpdateStoreCoordinate(ctx context.Context, db *sql.DB, dataUpdateCoordinate *entities.DataStoreVisitUpdateCoordinate) error {

	err := DataStoreCoordinateUpdate(ctx, db,
		dataUpdateCoordinate.CustomerNo,
		dataUpdateCoordinate.ShipTo,
		dataUpdateCoordinate.Latitude,
		dataUpdateCoordinate.Longitude,
		dataUpdateCoordinate.CreatedBy)

	if err != nil {
		return err
	}

	return nil
}
