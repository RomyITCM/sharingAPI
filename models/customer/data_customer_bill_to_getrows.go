package customer

import (
	"context"
	"database/sql"
	"smile-service/entities"

	"go.uber.org/zap"
)

const execSPDataCustomerBillTo = `exec [sp_smile_customer_bill_to_getrows]$1,$2`

func GetDataCustomerBillTo(ctx context.Context, db *sql.DB, customer string, search string, log *zap.Logger) ([]*entities.DataCustomerBillTo, error) {
	rows, err := db.QueryContext(
		ctx,
		execSPDataCustomerBillTo,
		customer,
		search,
	)

	if err != nil {
		log.Info("Exec DB",
			zap.Any("Exec DB", err),
		)
		return nil, err
	}
	defer rows.Close()

	data_customers_bill_to := make([]*entities.DataCustomerBillTo, 0)
	for rows.Next() {
		data_customer_bill_to := &entities.DataCustomerBillTo{}

		if err := rows.Scan(
			&data_customer_bill_to.BillToId,
			&data_customer_bill_to.BillToName,
			&data_customer_bill_to.BillToAddress,
		); err != nil {
			log.Info("scan rows",
				zap.Any("rows", err),
			)
			return nil, err
		}

		data_customers_bill_to = append(data_customers_bill_to, data_customer_bill_to)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return data_customers_bill_to, nil

}
