package new_customer

import (
	"context"
	"database/sql"
	"smile-service/entities"

	"go.uber.org/zap"
)

const (
	execSPFreezerInsert = `exec [sp_smile_outlet_freezer_insert] 
		$1, $2, $3, $4, $5, $6, $7`
)

func DataNewFreezerInsert(ctx context.Context, db *sql.DB,
	data_new_freezer *entities.InfoNewCustomerFreezer,
	log *zap.Logger) (*entities.FreezerId, error) {

	rows, err := db.QueryContext(ctx, execSPFreezerInsert,
		data_new_freezer.CustomerId,
		data_new_freezer.OutletId,
		data_new_freezer.FreezerOrigin,
		data_new_freezer.FreezerType,
		// data_new_freezer.FreezerLocation,
		// data_new_freezer.FreezerUse,
		data_new_freezer.RequestedAmount,
		data_new_freezer.CreatedBy,
		data_new_freezer.CreatedByIp,
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
