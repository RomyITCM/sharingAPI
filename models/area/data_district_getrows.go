package area

import (
	"context"
	"database/sql"
	"smile-service/entities"

	"go.uber.org/zap"
)

const execSPDistrictGetRows = `exec [sp_smile_district_getrows] $1`

func GetRowsDataDistrict(ctx context.Context, db *sql.DB, cityId string,
	log *zap.Logger) ([]*entities.DataArea, error) {

	rows, err := db.QueryContext(
		ctx,
		execSPDistrictGetRows,
		cityId,
	)

	if err != nil {
		log.Info("Exec DB",
			zap.Any("Exec DB", err),
		)
		return nil, err
	}
	defer rows.Close()

	data_districts := make([]*entities.DataArea, 0)
	for rows.Next() {
		data_district := &entities.DataArea{}

		if err := rows.Scan(
			&data_district.Id,
			&data_district.Name,
		); err != nil {
			log.Info("scan rows",
				zap.Any("rows", err),
			)
			return nil, err
		}

		data_districts = append(data_districts, data_district)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return data_districts, nil
}
