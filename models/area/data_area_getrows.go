package area

import (
	"context"
	"database/sql"
	"smile-service/entities"

	"go.uber.org/zap"
)

const execSPAreaGetRows = `exec [sp_smile_area_getrows] $1,$2`

func GetRowsArea(ctx context.Context, db *sql.DB, regionId string,
	search string, log *zap.Logger) ([]*entities.DataMasterArea, error) {

	rows, err := db.QueryContext(
		ctx,
		execSPAreaGetRows,
		regionId,
		search,
	)

	if err != nil {
		log.Info("Exec DB",
			zap.Any("Exec DB", err),
		)
		return nil, err
	}
	defer rows.Close()

	data_areas := make([]*entities.DataMasterArea, 0)
	for rows.Next() {
		data_area := &entities.DataMasterArea{}

		if err := rows.Scan(
			&data_area.AreaId,
			&data_area.AreaName,
		); err != nil {
			log.Info("scan rows",
				zap.Any("rows", err),
			)
			return nil, err
		}

		data_areas = append(data_areas, data_area)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return data_areas, nil
}
