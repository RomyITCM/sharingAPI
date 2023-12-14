package ar_overdue

import (
	"context"
	"database/sql"
	"smile-service/entities"

	"go.uber.org/zap"
)

const execSPDataCustArOverdueGetrow = `exec [sp_smile_ar_overdue_getrow] $1,$2`

func GetDataCustArOverdueGetrow(
	ctx context.Context,
	db *sql.DB,
	customer_no string,
	ship_to string,
	log *zap.Logger,
) ([]*entities.DataCustArOverdueGetrow, error) {
	rows, err := db.QueryContext(
		ctx,
		execSPDataCustArOverdueGetrow,
		customer_no,
		ship_to,
	)

	if err != nil {
		log.Info("Exec DB",
			zap.Any("Exec DB", err),
		)
		return nil, err
	}
	defer rows.Close()

	data_customers_ar_overdue := make([]*entities.DataCustArOverdueGetrow, 0)
	for rows.Next() {
		data_customer_ar_overdue := &entities.DataCustArOverdueGetrow{}

		if err := rows.Scan(
			&data_customer_ar_overdue.Customer,
			&data_customer_ar_overdue.CustomerName,
			&data_customer_ar_overdue.BillTo,
			&data_customer_ar_overdue.ShipTo,
			&data_customer_ar_overdue.TotalAROverdue,
		); err != nil {
			log.Info("scan rows",
				zap.Any("rows", err),
			)
			return nil, err
		}

		data_customers_ar_overdue = append(data_customers_ar_overdue, data_customer_ar_overdue)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return data_customers_ar_overdue, nil

}
