package new_customer

import (
	"context"
	"database/sql"
	"smile-service/entities"

	"go.uber.org/zap"
)

const execSPDataNewCustomerRequest = `exec [sp_smile_new_customer_request_getrows]$1`

func GetDataNewCustomerRequest(ctx context.Context, db *sql.DB, search string, log *zap.Logger) ([]*entities.DataNewCustomerRequest, error) {
	rows, err := db.QueryContext(
		ctx,
		execSPDataNewCustomerRequest,
		search,
	)

	if err != nil {
		log.Info(
			"Exec DB",
			zap.Any("Exec DB", err),
		)
		return nil, err
	}
	defer rows.Close()

	data_new_customer_requests := make([]*entities.DataNewCustomerRequest, 0)
	for rows.Next() {
		data_new_customer_request := &entities.DataNewCustomerRequest{}

		if err := rows.Scan(
			&data_new_customer_request.CustomerRequestNo,
			&data_new_customer_request.Status,
			&data_new_customer_request.CustomerName,
			&data_new_customer_request.CustomerType,
			&data_new_customer_request.VatRegNo,
			&data_new_customer_request.StreetAddress,
			&data_new_customer_request.CityId,
			&data_new_customer_request.DistrictId,
			&data_new_customer_request.OutletCount,
			&data_new_customer_request.PicCount,
		); err != nil {
			log.Info("scan rows",
				zap.Any("rows", err),
			)
			return nil, err
		}

		data_new_customer_requests = append(data_new_customer_requests, data_new_customer_request)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return data_new_customer_requests, nil
}
