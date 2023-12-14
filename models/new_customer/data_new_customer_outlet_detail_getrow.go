package new_customer

import (
	"context"
	"database/sql"
	"smile-service/entities"

	"go.uber.org/zap"
)

const execSPCustomerOutletDetail = `exec [sp_smile_customer_outlet_detail_getrow] $1`

func GetDataNewCustomerOutletDetail(ctx context.Context, db *sql.DB, outlet_id string, log *zap.Logger) ([]*entities.DataNewCustomerOutletDetail, error) {
	rows, err := db.QueryContext(
		ctx,
		execSPCustomerOutletDetail,
		outlet_id,
	)

	if err != nil {
		log.Info(
			"Exec DB",
			zap.Any("Exec DB", err),
		)
		return nil, err
	}
	defer rows.Close()

	data_new_customer_outlet_details := make([]*entities.DataNewCustomerOutletDetail, 0)
	for rows.Next() {
		data_new_customer_outlet_detail := &entities.DataNewCustomerOutletDetail{}

		if err := rows.Scan(
			&data_new_customer_outlet_detail.OutletId,
			&data_new_customer_outlet_detail.OutletCode,
			&data_new_customer_outlet_detail.OutletName,
			&data_new_customer_outlet_detail.CustomerId,
			&data_new_customer_outlet_detail.StoreArea,
			&data_new_customer_outlet_detail.ShipSchedule,
			&data_new_customer_outlet_detail.StoreImg,
			&data_new_customer_outlet_detail.Address,
			&data_new_customer_outlet_detail.ProvinceId,
			&data_new_customer_outlet_detail.ProvinceName,
			&data_new_customer_outlet_detail.CityId,
			&data_new_customer_outlet_detail.CityName,
			&data_new_customer_outlet_detail.DistrictId,
			&data_new_customer_outlet_detail.DistrictName,
			&data_new_customer_outlet_detail.SubdistrictId,
			&data_new_customer_outlet_detail.SubdistrictName,
			&data_new_customer_outlet_detail.Zipcode,
			&data_new_customer_outlet_detail.Latitude,
			&data_new_customer_outlet_detail.Longitude,
			&data_new_customer_outlet_detail.PhoneNo,
			&data_new_customer_outlet_detail.CellphoneNo,
			&data_new_customer_outlet_detail.Email,
			&data_new_customer_outlet_detail.Salesman,
			&data_new_customer_outlet_detail.RegionId,
			&data_new_customer_outlet_detail.RegionName,
			&data_new_customer_outlet_detail.AreaId,
			&data_new_customer_outlet_detail.AreaName,
			&data_new_customer_outlet_detail.ZoneId,
			&data_new_customer_outlet_detail.ZoneName,
			&data_new_customer_outlet_detail.Status,
			&data_new_customer_outlet_detail.ApprovalStatus,
		); err != nil {
			log.Info("scan rows",
				zap.Any("rows", err),
			)
			return nil, err
		}

		data_new_customer_outlet_details = append(data_new_customer_outlet_details, data_new_customer_outlet_detail)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return data_new_customer_outlet_details, nil
}
