package ar_overdue

import (
	"context"
	"database/sql"
	"smile-service/entities"

	"go.uber.org/zap"
)

const execSPDataCustArOverdueGetrows = `exec [sp_smile_ar_overdue_getrows] $1,$2`

func GetDataCustArOverdueGetrows(
	ctx context.Context,
	db *sql.DB,
	customer_no string,
	ship_to string,
	log *zap.Logger,
) ([]*entities.DataCustArOverdueGetrows, error) {
	rows, err := db.QueryContext(
		ctx,
		execSPDataCustArOverdueGetrows,
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

	data_customers_ar_overdue := make([]*entities.DataCustArOverdueGetrows, 0)
	for rows.Next() {
		data_customer_ar_overdue := &entities.DataCustArOverdueGetrows{}

		if err := rows.Scan(
			&data_customer_ar_overdue.DocNo,
			&data_customer_ar_overdue.CustPONo,
			&data_customer_ar_overdue.Overdue,
			&data_customer_ar_overdue.Outstanding,
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
