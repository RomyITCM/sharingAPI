package area

import (
	"context"
	"database/sql"
	"smile-service/entities"

	"go.uber.org/zap"
)

const execSPSubdistrictGetrows = `exec [sp_smile_villages_getrows] $1`

func GetRowsDataSubdistrict(ctx context.Context, db *sql.DB, districtId string,
	log *zap.Logger) ([]*entities.DataArea, error) {

	rows, err := db.QueryContext(
		ctx,
		execSPSubdistrictGetrows,
		districtId,
	)

	if err != nil {
		log.Info("Exec DB",
			zap.Any("Exec DB", err),
		)
		return nil, err
	}
	defer rows.Close()

	data_subdistricts := make([]*entities.DataArea, 0)
	for rows.Next() {
		data_subdistrict := &entities.DataArea{}

		if err := rows.Scan(
			&data_subdistrict.Id,
			&data_subdistrict.Name,
		); err != nil {
			log.Info("scan rows",
				zap.Any("rows", err),
			)
			return nil, err
		}

		data_subdistricts = append(data_subdistricts, data_subdistrict)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return data_subdistricts, nil
}
