package area

import (
	"context"
	"database/sql"
	"smile-service/entities"

	"go.uber.org/zap"
)

const execSPProvinceGetRows = `exec [sp_smile_province_getrows]`

func GetRowsDataProvince(ctx context.Context, db *sql.DB, log *zap.Logger) ([]*entities.DataArea, error) {

	rows, err := db.QueryContext(
		ctx,
		execSPProvinceGetRows,
	)

	if err != nil {
		log.Info("Exec DB",
			zap.Any("Exec DB", err),
		)
		return nil, err
	}
	defer rows.Close()

	data_provinces := make([]*entities.DataArea, 0)
	for rows.Next() {
		data_province := &entities.DataArea{}

		if err := rows.Scan(
			&data_province.Id,
			&data_province.Name,
		); err != nil {
			log.Info("scan rows",
				zap.Any("rows", err),
			)
			return nil, err
		}

		data_provinces = append(data_provinces, data_province)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return data_provinces, nil
}
