package customer

import (
	"context"
	"database/sql"
	"smile-service/entities"

	"go.uber.org/zap"
)

const execSPDataCustomerAddressBillToByCustomer = `exec [sp_smile_customer_address_bill_to_by_customer_getrows] $1`

func GetRowsCustomerAddressBillToByCustomer(ctx context.Context, db *sql.DB,
	customerNo string, log *zap.Logger) ([]*entities.CustomerAddressBillToByCustomer, error) {

	rows, err := db.QueryContext(
		ctx,
		execSPDataCustomerAddressBillToByCustomer,
		customerNo,
	)

	if err != nil {
		log.Info("Exec DB",
			zap.Any("Exec DB", err),
		)
		return nil, err
	}
	defer rows.Close()

	billTos := make([]*entities.CustomerAddressBillToByCustomer, 0)
	for rows.Next() {
		billTo := &entities.CustomerAddressBillToByCustomer{}

		if err := rows.Scan(
			&billTo.Id,
			&billTo.BillingName,
		); err != nil {
			log.Info("scan rows",
				zap.Any("rows", err),
			)
			return nil, err
		}

		billTos = append(billTos, billTo)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return billTos, nil
}
