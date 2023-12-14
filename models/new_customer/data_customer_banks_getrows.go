package new_customer

import (
	"context"
	"database/sql"
	"smile-service/entities"

	"go.uber.org/zap"
)

const execSPCustomerBanks = `exec [sp_smile_banks_getrows] $1`

func GetRowsCustomerBanks(ctx context.Context, db *sql.DB, search string,
	log *zap.Logger) ([]*entities.DataCustomerBanks, error) {

	rows, err := db.QueryContext(
		ctx,
		execSPCustomerBanks,
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

	data_customer_banks := make([]*entities.DataCustomerBanks, 0)
	for rows.Next() {
		data_customer_bank := &entities.DataCustomerBanks{}

		if err := rows.Scan(
			&data_customer_bank.Code,
			&data_customer_bank.BankName,
		); err != nil {
			log.Info("scan rows",
				zap.Any("rows", err),
			)
			return nil, err
		}

		data_customer_banks = append(data_customer_banks, data_customer_bank)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return data_customer_banks, nil
}
