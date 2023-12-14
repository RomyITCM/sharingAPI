package new_customer

import (
	"context"
	"database/sql"
	"smile-service/entities"
)

const (
	execSPFreezerUdpate = `exec [sp_smile_outlet_freezer_update]
		$1, $2, $3, $4, $5, $6`
)

func UpdateDataOutletFreezer(ctx context.Context, db *sql.DB,
	FreezerId string,
	FreezerOrigin string,
	FreezerType string,
	RequestedAmount string,
	CreatedBy string,
	CreatedByIp string,
) error {

	rows, err := db.QueryContext(ctx, execSPFreezerUdpate,
		FreezerId,
		FreezerOrigin,
		FreezerType,
		RequestedAmount,
		CreatedBy,
		CreatedByIp,
	)

	if err != nil {
		return err
	}
	defer rows.Close()

	return nil
}

func DataOutletFreezerUpdate(ctx context.Context, db *sql.DB,
	data_outlet_freezer *entities.InfoNewCustomerFreezerUpdate) error {

	err := UpdateDataOutletFreezer(ctx, db,
		data_outlet_freezer.FreezerId,
		data_outlet_freezer.FreezerOrigin,
		data_outlet_freezer.FreezerType,
		data_outlet_freezer.RequestedAmount,
		data_outlet_freezer.CreatedBy,
		data_outlet_freezer.CreatedByIp,
	)

	if err != nil {
		return err
	}

	return nil
}
