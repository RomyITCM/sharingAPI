package area

import (
	"context"
	"database/sql"
	"smile-service/entities"

	"go.uber.org/zap"
)

const execSPRegionGetrows = `exec [sp_smile_region_getrows] $1`

func GetRowsDataRegion(ctx context.Context, db *sql.DB, search string,
	log *zap.Logger) ([]*entities.DataMasterRegion, error) {

	rows, err := db.QueryContext(
		ctx,
		execSPRegionGetrows,
		search,
	)

	if err != nil {
		log.Info("Exec DB",
			zap.Any("Exec DB", err),
		)
		return nil, err
	}
	defer rows.Close()

	data_regions := make([]*entities.DataMasterRegion, 0)
	for rows.Next() {
		data_region := &entities.DataMasterRegion{}

		if err := rows.Scan(
			&data_region.RegionId,
			&data_region.RegionName,
		); err != nil {
			log.Info("scan rows",
				zap.Any("rows", err),
			)
			return nil, err
		}

		data_regions = append(data_regions, data_region)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return data_regions, nil
}
