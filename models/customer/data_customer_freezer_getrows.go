package customer

import (
	"context"
	"database/sql"
	"smile-service/entities"

	"go.uber.org/zap"
)

const execSPDataCustomerFreezer = `exec [sp_smile_customer_address_freezers_getrows] $1, $2`

func GetRowsCustomerFreezer(ctx context.Context, db *sql.DB, customerAddressId string,
	freezerId string, log *zap.Logger) ([]*entities.DataNewCustomerFreezer, error) {

	rows, err := db.QueryContext(
		ctx,
		execSPDataCustomerFreezer,
		customerAddressId,
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

	data_customer_freezers := make([]*entities.DataNewCustomerFreezer, 0)
	for rows.Next() {
		data_customer_freezer := &entities.DataNewCustomerFreezer{}

		if err := rows.Scan(
			&data_customer_freezer.FreezerId,
			&data_customer_freezer.OutletId,
			&data_customer_freezer.CustomerId,
			&data_customer_freezer.CustomerAddressId,
			&data_customer_freezer.FreezerOrigin,
			&data_customer_freezer.FreezerType,
			&data_customer_freezer.RequestedAmount,
		); err != nil {
			log.Info("scan rows",
				zap.Any("rows", err),
			)
			return nil, err
		}

		data_customer_freezers = append(data_customer_freezers, data_customer_freezer)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return data_customer_freezers, nil
}
