package cart

import (
	"context"
	"database/sql"
	"smile-service/entities"

	"go.uber.org/zap"
)

const execSPDataCartBill = `exec [sp_smile_cart_bill_getrows]$1,$2`
const execSPDataCartBillEdit = `exec [sp_smile_cart_bill_edit_getrows]$1`

func GetDataCartBill(ctx context.Context, db *sql.DB, shipTo, salesMan string, log *zap.Logger) ([]*entities.CartBill, error) {
	rows, err := db.QueryContext(
		ctx,
		execSPDataCartBill,
		shipTo,
		salesMan,
	)

	if err != nil {
		log.Info("Exec DB",
			zap.Any("Exec DB", err),
		)
		return nil, err
	}
	defer rows.Close()

	data_cart_bills := make([]*entities.CartBill, 0)
	for rows.Next() {
		data_cart_bill := &entities.CartBill{}

		if err := rows.Scan(
			&data_cart_bill.TotalAmount,
			&data_cart_bill.Disc,
			&data_cart_bill.Tax,
			&data_cart_bill.TotalBill,
		); err != nil {
			log.Info("scan rows",
				zap.Any("rows", err),
			)
			return nil, err
		}

		data_cart_bills = append(data_cart_bills, data_cart_bill)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return data_cart_bills, nil

}

func GetDataCartBillEdit(ctx context.Context, db *sql.DB, transNo string, log *zap.Logger) ([]*entities.CartBill, error) {
	rows, err := db.QueryContext(
		ctx,
		execSPDataCartBillEdit,
		transNo,
	)

	if err != nil {
		log.Info("Exec DB",
			zap.Any("Exec DB", err),
		)
		return nil, err
	}
	defer rows.Close()

	data_cart_bills := make([]*entities.CartBill, 0)
	for rows.Next() {
		data_cart_bill := &entities.CartBill{}

		if err := rows.Scan(
			&data_cart_bill.TotalAmount,
			&data_cart_bill.Disc,
			&data_cart_bill.Tax,
			&data_cart_bill.TotalBill,
		); err != nil {
			log.Info("scan rows",
				zap.Any("rows", err),
			)
			return nil, err
		}

		data_cart_bills = append(data_cart_bills, data_cart_bill)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return data_cart_bills, nil

}
