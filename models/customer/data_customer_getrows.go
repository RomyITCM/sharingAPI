package customer

import (
	"context"
	"database/sql"
	"smile-service/entities"

	"go.uber.org/zap"
)

const execSPDataApprovedCustomer = `exec [sp_smile_approved_customer_getrows] $1`

func GetRowsDataCustomer(ctx context.Context, db *sql.DB,
	search string, log *zap.Logger) ([]*entities.DataApprovedCustomer, error) {

	rows, err := db.QueryContext(
		ctx,
		execSPDataApprovedCustomer,
		search,
	)

	if err != nil {
		log.Info("Exec DB",
			zap.Any("Exec DB", err),
		)
		return nil, err
	}
	defer rows.Close()

	data_customers := make([]*entities.DataApprovedCustomer, 0)
	for rows.Next() {
		data_customer := &entities.DataApprovedCustomer{}

		if err := rows.Scan(
			&data_customer.CustomerNo,
			&data_customer.Status,
			&data_customer.CustomerName,
			&data_customer.CustomerType,
			&data_customer.VatRegNo,
			&data_customer.StreetAddress,
			&data_customer.City,
			&data_customer.District,
			&data_customer.OutletCount,
			&data_customer.PicCount,
		); err != nil {
			log.Info("scan rows",
				zap.Any("rows", err),
			)
			return nil, err
		}

		data_customers = append(data_customers, data_customer)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return data_customers, nil
}
