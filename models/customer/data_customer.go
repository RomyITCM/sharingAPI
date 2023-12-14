package customer

import (
	"context"
	"database/sql"
	"smile-service/entities"

	"go.uber.org/zap"
)

const execSPDataCustomer = `exec [sp_smile_customer_getrows]$1`

func GetDataCustomer(ctx context.Context, db *sql.DB, search string, log *zap.Logger) ([]*entities.DataCustomer, error) {
	rows, err := db.QueryContext(
		ctx,
		execSPDataCustomer,
		search)

	if err != nil {
		log.Info("Exec DB",
			zap.Any("Exec DB", err),
		)
		return nil, err
	}
	defer rows.Close()

	data_customers := make([]*entities.DataCustomer, 0)
	for rows.Next() {
		data_customer := &entities.DataCustomer{}

		if err := rows.Scan(
			&data_customer.CustomerNo,
			&data_customer.CustomerName,
			&data_customer.CustomerAddress,
			&data_customer.CustomerType,
			&data_customer.CustomerTypeDesc,
			&data_customer.CustomerStatus,
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
