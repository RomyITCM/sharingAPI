package area

import (
	"context"
	"database/sql"
	"smile-service/entities"

	"go.uber.org/zap"
)

const execSPCityGetRows = `exec [sp_smile_city_getrows] $1`

func GetRowsDataCity(ctx context.Context, db *sql.DB, provinceId string,
	log *zap.Logger) ([]*entities.DataArea, error) {

	rows, err := db.QueryContext(
		ctx,
		execSPCityGetRows,
		provinceId,
	)

	if err != nil {
		log.Info("Exec DB",
			zap.Any("Exec DB", err),
		)
		return nil, err
	}
	defer rows.Close()

	data_cities := make([]*entities.DataArea, 0)
	for rows.Next() {
		data_city := &entities.DataArea{}

		if err := rows.Scan(
			&data_city.Id,
			&data_city.Name,
		); err != nil {
			log.Info("scan rows",
				zap.Any("rows", err),
			)
			return nil, err
		}

		data_cities = append(data_cities, data_city)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return data_cities, nil
}
