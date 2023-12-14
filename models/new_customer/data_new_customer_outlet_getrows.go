package new_customer

import (
	"context"
	"database/sql"
	"smile-service/entities"

	"go.uber.org/zap"
)

const execSPCustomerOutletById = `exec [sp_smile_new_customer_outlet_getrows] $1`

func GetDataNewCustomerOutlet(ctx context.Context, db *sql.DB, customer_id string, log *zap.Logger) ([]*entities.DataNewCustomerOutlet, error) {
	rows, err := db.QueryContext(
		ctx,
		execSPCustomerOutletById,
		customer_id,
	)

	if err != nil {
		log.Info(
			"Exec DB",
			zap.Any("Exec DB", err),
		)
		return nil, err
	}
	defer rows.Close()

	data_new_customer_outlets := make([]*entities.DataNewCustomerOutlet, 0)
	for rows.Next() {
		data_new_customer_outlet := &entities.DataNewCustomerOutlet{}

		if err := rows.Scan(
			&data_new_customer_outlet.OutletId,
			&data_new_customer_outlet.OutletName,
			&data_new_customer_outlet.Address,
			&data_new_customer_outlet.ProvinceId,
			&data_new_customer_outlet.CityId,
			&data_new_customer_outlet.DistrictId,
			&data_new_customer_outlet.Latitude,
			&data_new_customer_outlet.Longitude,
			&data_new_customer_outlet.StoreFreezer,
			&data_new_customer_outlet.BenFreezer,
		); err != nil {
			log.Info("scan rows",
				zap.Any("rows", err),
			)
			return nil, err
		}

		data_new_customer_outlets = append(data_new_customer_outlets, data_new_customer_outlet)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return data_new_customer_outlets, nil
}
