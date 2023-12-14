package customer

import (
	"context"
	"database/sql"
	"smile-service/entities"

	"go.uber.org/zap"
)

const execSPDataCustomerAddress = `exec [sp_smile_customer_address_getrows] $1, $2`

func GetRowsCustomerAddress(ctx context.Context, db *sql.DB,
	customerNo string, search string, log *zap.Logger) ([]*entities.DataCustomerAddress, error) {

	rows, err := db.QueryContext(
		ctx,
		execSPDataCustomerAddress,
		customerNo,
		search,
	)

	if err != nil {
		log.Info("Exec DB",
			zap.Any("Exec DB", err),
		)
		return nil, err
	}
	defer rows.Close()

	customerAddresses := make([]*entities.DataCustomerAddress, 0)
	for rows.Next() {
		customerAddress := &entities.DataCustomerAddress{}

		if err := rows.Scan(
			&customerAddress.Id,
			&customerAddress.AliasName,
			&customerAddress.AliasNameFull,
			&customerAddress.StreetAddress,
			&customerAddress.ProvinceId,
			&customerAddress.ProvinceName,
			&customerAddress.CityId,
			&customerAddress.CityName,
			&customerAddress.DistrictId,
			&customerAddress.DistrictName,
			&customerAddress.Latitude,
			&customerAddress.Longitude,
			&customerAddress.StoreFreezer,
			&customerAddress.BenFreezer,
		); err != nil {
			log.Info("scan rows",
				zap.Any("rows", err),
			)
			return nil, err
		}

		customerAddresses = append(customerAddresses, customerAddress)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return customerAddresses, nil
}
