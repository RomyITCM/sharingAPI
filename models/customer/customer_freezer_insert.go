package customer

import (
	"context"
	"database/sql"
	"smile-service/entities"

	"go.uber.org/zap"
)

const (
	execSPCustomerFreezerInsert = `exec [sp_smile_customer_freezer_insert]
		$1, $2, $3, $4, $5, $6, $7`
)

func DataCustomerFreezerInsert(ctx context.Context, db *sql.DB,
	data_customer_freezer *entities.InfoCustomerFreezer,
	log *zap.Logger) (*entities.FreezerId, error) {

	rows, err := db.QueryContext(ctx, execSPCustomerFreezerInsert,
		data_customer_freezer.CustomerId,
		data_customer_freezer.CustomerAddressId,
		data_customer_freezer.FreezerOrigin,
		data_customer_freezer.FreezerType,
		data_customer_freezer.RequestedAmount,
		data_customer_freezer.CreatedBy,
		data_customer_freezer.CreatedByIp,
	)

	if err != nil {
		log.Info(
			"Exec DB", zap.Any("Exec Db", err),
		)

		return nil, err
	}
	defer rows.Close()

	freezer_id := &entities.FreezerId{}
	for rows.Next() {
		if err := rows.Scan(
			&freezer_id.FreezerId,
		); err != nil {
			log.Info("scan rows",
				zap.Any("rows", err),
			)
			return nil, err
		}
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return freezer_id, nil
}
