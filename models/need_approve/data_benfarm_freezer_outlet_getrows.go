package need_approve

import (
	"context"
	"database/sql"
	"smile-service/entities"

	"go.uber.org/zap"
)

const execSPBenfarmFreezerOutlet = `exec [sp_smile_customer_address_benfarm_getrows] $1`

func GetRowsBenfarmFreezerOutlet(ctx context.Context, db *sql.DB, outlet_id string,
	log *zap.Logger) ([]*entities.DataBenfarmFreezer, error) {

	rows, err := db.QueryContext(
		ctx,
		execSPBenfarmFreezerOutlet,
		outlet_id,
	)

	if err != nil {
		log.Info(
			"Exec DB",
			zap.Any("Exec DB", err),
		)
		return nil, err
	}
	defer rows.Close()

	freezers := make([]*entities.DataBenfarmFreezer, 0)
	for rows.Next() {
		freezer := &entities.DataBenfarmFreezer{}

		if err := rows.Scan(
			&freezer.FreezerId,
			&freezer.CustomerId,
			&freezer.CustomerName,
			&freezer.OutletId,
			&freezer.OutletName,
			&freezer.FreezerOrigin,
			&freezer.FreezerType,
			&freezer.RequestedAmount,
		); err != nil {
			log.Info("scan rows",
				zap.Any("rows", err),
			)
			return nil, err
		}

		freezers = append(freezers, freezer)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return freezers, nil
}
