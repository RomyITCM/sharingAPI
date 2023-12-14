package new_customer

import (
	"context"
	"database/sql"
	"smile-service/entities"

	"go.uber.org/zap"
)

const execSPPicResponsibility = `exec [sp_smile_pic_responsibility_getrows] $1`

func GetRowsPicResponsibility(ctx context.Context, db *sql.DB, customer_id string,
	log *zap.Logger) ([]*entities.DataPicResponsibility, error) {

	rows, err := db.QueryContext(
		ctx,
		execSPPicResponsibility,
		customer_id,
	)

	if err != nil {
		log.Info(
			"Exec DB",
			zap.Any("Exec DB", err),
		)
		return nil, err
	}
	defer rows.Close()

	responsibilities := make([]*entities.DataPicResponsibility, 0)
	for rows.Next() {
		responsibility := &entities.DataPicResponsibility{}

		if err := rows.Scan(
			&responsibility.Id,
			&responsibility.Name,
		); err != nil {
			log.Info("scan rows",
				zap.Any("rows", err),
			)
			return nil, err
		}

		responsibilities = append(responsibilities, responsibility)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return responsibilities, nil
}
