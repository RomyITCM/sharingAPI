package need_approve

import (
	"context"
	"database/sql"
	"smile-service/entities"

	"go.uber.org/zap"
)

const execSPNeedApproveCustomer = `exec [sp_smile_need_approve_customer_getrows] $1`

func GetNeedApproveCustomer(ctx context.Context, db *sql.DB,
	search string, log *zap.Logger) ([]*entities.DataNeedApproveCustomerRequest, error) {

	rows, err := db.QueryContext(
		ctx,
		execSPNeedApproveCustomer,
		search,
	)

	if err != nil {
		log.Info("Exec DB",
			zap.Any("Exec DB", err),
		)
		return nil, err
	}
	defer rows.Close()

	data_need_approve_customers := make([]*entities.DataNeedApproveCustomerRequest, 0)
	for rows.Next() {
		data_need_approve_customer := &entities.DataNeedApproveCustomerRequest{}

		if err := rows.Scan(
			&data_need_approve_customer.CustomerRequestNo,
			&data_need_approve_customer.Status,
			&data_need_approve_customer.CustomerName,
			&data_need_approve_customer.CustomerType,
			&data_need_approve_customer.VatRegNo,
			&data_need_approve_customer.StreetAddress,
			&data_need_approve_customer.CityId,
			&data_need_approve_customer.DistrictId,
			&data_need_approve_customer.OutletCount,
			&data_need_approve_customer.PicCount,
		); err != nil {
			log.Info("scan rows",
				zap.Any("rows", err),
			)
			return nil, err
		}

		data_need_approve_customers = append(data_need_approve_customers, data_need_approve_customer)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return data_need_approve_customers, nil
}
