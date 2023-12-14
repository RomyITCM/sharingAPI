package ar_overdue

import (
	"context"
	"database/sql"
	"smile-service/entities"

	"go.uber.org/zap"
)

const execSPDataCustArOverDueGetrows = `exec [sp_smile_customer_ar_over_due_getrows] $1`

func GetDataCustArOverDue(
	ctx context.Context,
	db *sql.DB,
	search string,
	log *zap.Logger,
) ([]*entities.DataCustArOverDueGetrows, error) {
	rows, err := db.QueryContext(
		ctx,
		execSPDataCustArOverDueGetrows,
		search,
	)

	if err != nil {
		log.Info("Exec DB",
			zap.Any("Exec DB", err),
		)
		return nil, err
	}
	defer rows.Close()

	data_customers_ar_over_due := make([]*entities.DataCustArOverDueGetrows, 0)
	for rows.Next() {
		data_customer_ar_over_due := &entities.DataCustArOverDueGetrows{}

		if err := rows.Scan(
			&data_customer_ar_over_due.CustomerNo,
			&data_customer_ar_over_due.CustomerName,
			&data_customer_ar_over_due.BillTo,
			&data_customer_ar_over_due.BillToName,
			&data_customer_ar_over_due.Dpp,
			&data_customer_ar_over_due.Ppn,
			&data_customer_ar_over_due.TotalAmount,
			&data_customer_ar_over_due.PaymentAmount,
			&data_customer_ar_over_due.TotalAmountOutstanding,
			&data_customer_ar_over_due.OneToFourteen,
			&data_customer_ar_over_due.FifteenToTwentyOne,
			&data_customer_ar_over_due.TwentyTwoToThirty,
			&data_customer_ar_over_due.ThirtyoneToSixty,
			&data_customer_ar_over_due.SixtyoneToninety,
			&data_customer_ar_over_due.MoreThanNinety,
		); err != nil {
			log.Info("scan rows",
				zap.Any("rows", err),
			)
			return nil, err
		}

		data_customers_ar_over_due = append(data_customers_ar_over_due, data_customer_ar_over_due)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return data_customers_ar_over_due, nil

}
