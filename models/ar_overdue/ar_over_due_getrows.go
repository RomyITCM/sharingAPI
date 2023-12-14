package ar_overdue

import (
	"context"
	"database/sql"
	"smile-service/entities"

	"go.uber.org/zap"
)

const execSPDataArOverDueGetrows = `exec [sp_smile_ar_over_due_getrows] $1,$2,$3,$4`

func GetDataArOverDue(ctx context.Context, db *sql.DB, CustomerNo string, BillTo string, BankCode string,
	Search string, log *zap.Logger,
) ([]*entities.DataArOverDueGetrows, error) {
	rows, err := db.QueryContext(
		ctx,
		execSPDataArOverDueGetrows,
		CustomerNo,
		BillTo,
		BankCode,
		Search,
	)

	if err != nil {
		log.Info("Exec DB",
			zap.Any("Exec DB", err),
		)
		return nil, err
	}
	defer rows.Close()

	data_ar_over_due_rows := make([]*entities.DataArOverDueGetrows, 0)
	for rows.Next() {
		data_ar_over_due_row := &entities.DataArOverDueGetrows{}

		if err := rows.Scan(
			&data_ar_over_due_row.TransNo,
			&data_ar_over_due_row.DocNo,
			&data_ar_over_due_row.TransDate,
			&data_ar_over_due_row.JvDate,
			&data_ar_over_due_row.CustomerName,
			&data_ar_over_due_row.BillTo,
			&data_ar_over_due_row.BillToName,
			&data_ar_over_due_row.ShipTo,
			&data_ar_over_due_row.ShipToName,
			&data_ar_over_due_row.TotalInvoice,
			&data_ar_over_due_row.Ppn,
			&data_ar_over_due_row.TotalAmount,
			&data_ar_over_due_row.PaymentAmount,
			&data_ar_over_due_row.TotalAmountOutstanding,
			&data_ar_over_due_row.DueDate,
			&data_ar_over_due_row.JatuhTempo,
			&data_ar_over_due_row.OneToFourteenDetail,
			&data_ar_over_due_row.FifteenToTwentyOneDetail,
			&data_ar_over_due_row.TwentyTwoToThirtyDetail,
			&data_ar_over_due_row.ThirtyoneToSixtyDetail,
			&data_ar_over_due_row.SixtyoneToninetyDetail,
			&data_ar_over_due_row.MoreThanNinetyDetail,
		); err != nil {
			log.Info("scan rows",
				zap.Any("rows", err),
			)
			return nil, err
		}

		data_ar_over_due_rows = append(data_ar_over_due_rows, data_ar_over_due_row)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return data_ar_over_due_rows, nil

}
