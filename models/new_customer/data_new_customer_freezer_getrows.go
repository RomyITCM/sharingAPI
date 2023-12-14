package new_customer

import (
	"context"
	"database/sql"
	"smile-service/entities"

	"go.uber.org/zap"
)

const execSPDataNewCustomerFreezer = `exec [sp_smile_outlet_freezer_getrows] $1, $2`

func GetDataNewCustomerFreezer(ctx context.Context, db *sql.DB, outletId string, freezerId string, log *zap.Logger) ([]*entities.DataNewCustomerFreezer, error) {
	rows, err := db.QueryContext(
		ctx,
		execSPDataNewCustomerFreezer,
		outletId,
		freezerId,
	)

	if err != nil {
		log.Info(
			"Exec DB",
			zap.Any("Exec DB", err),
		)
		return nil, err
	}
	defer rows.Close()

	data_new_customer_freezers := make([]*entities.DataNewCustomerFreezer, 0)
	for rows.Next() {
		data_new_customer_freezer := &entities.DataNewCustomerFreezer{}

		if err := rows.Scan(
			&data_new_customer_freezer.FreezerId,
			&data_new_customer_freezer.OutletId,
			&data_new_customer_freezer.CustomerId,
			&data_new_customer_freezer.CustomerAddressId,
			&data_new_customer_freezer.FreezerOrigin,
			&data_new_customer_freezer.FreezerType,
			// &data_new_customer_freezer.FreezerLocation,
			// &data_new_customer_freezer.FreezerUse,
			&data_new_customer_freezer.RequestedAmount,
		); err != nil {
			log.Info("scan rows",
				zap.Any("rows", err),
			)
			return nil, err
		}

		data_new_customer_freezers = append(data_new_customer_freezers, data_new_customer_freezer)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return data_new_customer_freezers, nil
}
