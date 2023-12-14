package price_zone

import (
	"context"
	"database/sql"
	"smile-service/entities"

	"go.uber.org/zap"
)

const execSPPriceZone = `exec [sp_smile_price_zone_getrows] $1`

func GetRowsPriceZone(ctx context.Context, db *sql.DB, search string,
	log *zap.Logger) ([]*entities.DataPriceZone, error) {

	rows, err := db.QueryContext(
		ctx,
		execSPPriceZone,
		search,
	)

	if err != nil {
		log.Info("Exec DB",
			zap.Any("Exec DB", err),
		)
		return nil, err
	}
	defer rows.Close()

	priceZones := make([]*entities.DataPriceZone, 0)
	for rows.Next() {
		priceZone := &entities.DataPriceZone{}

		if err := rows.Scan(
			&priceZone.ZoneId,
			&priceZone.ZoneName,
		); err != nil {
			log.Info("scan rows",
				zap.Any("rows", err),
			)
			return nil, err
		}

		priceZones = append(priceZones, priceZone)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return priceZones, nil
}
