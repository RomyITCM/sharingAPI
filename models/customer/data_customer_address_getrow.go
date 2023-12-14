package customer

import (
	"context"
	"database/sql"
	"smile-service/entities"

	"go.uber.org/zap"
)

const execSPDataCustomerAddressDetail = `exec [sp_smile_customer_address_getrow] $1`

func GetRowCustomerAddress(ctx context.Context, db *sql.DB,
	customerAddressId string, log *zap.Logger) ([]*entities.CustomerAddressDetail, error) {

	rows, err := db.QueryContext(
		ctx,
		execSPDataCustomerAddressDetail,
		customerAddressId,
	)

	if err != nil {
		log.Info("Exec DB",
			zap.Any("Exec DB", err),
		)
		return nil, err
	}
	defer rows.Close()

	customerAddresses := make([]*entities.CustomerAddressDetail, 0)
	for rows.Next() {
		customerAddress := &entities.CustomerAddressDetail{}

		if err := rows.Scan(
			&customerAddress.Id,
			&customerAddress.AliasName,
			&customerAddress.ALiasNameFull,
			&customerAddress.CustomerNo,
			&customerAddress.StoreArea,
			&customerAddress.StoreImg,
			&customerAddress.StreetAddress,
			&customerAddress.ProvinceId,
			&customerAddress.Province,
			&customerAddress.CityId,
			&customerAddress.City,
			&customerAddress.DistrictId,
			&customerAddress.District,
			&customerAddress.SubdistrictId,
			&customerAddress.Subdistrict,
			&customerAddress.ZipCode,
			&customerAddress.Latitude,
			&customerAddress.Longitude,
			&customerAddress.PhoneNo,
			&customerAddress.MobilePhoneNo,
			&customerAddress.Email,
			&customerAddress.Status,
			&customerAddress.Salesman,
			&customerAddress.RegionId,
			&customerAddress.AreaId,
			&customerAddress.ZoneId,
			&customerAddress.BillTo,
			&customerAddress.ShipSchedule,
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
