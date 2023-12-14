package site

import (
	"context"
	"database/sql"
	"smile-service/entities"

	"go.uber.org/zap"
)

const execSPDataSite = `exec [sp_smile_site_getrows]$1`

func GetDataSite(ctx context.Context, db *sql.DB, search string, log *zap.Logger) ([]*entities.DataSite, error) {
	rows, err := db.QueryContext(
		ctx,
		execSPDataSite,
		search)

	if err != nil {
		log.Info("Exec DB",
			zap.Any("Exec DB", err),
		)
		return nil, err
	}
	defer rows.Close()

	data_sites := make([]*entities.DataSite, 0)
	for rows.Next() {
		data_site := &entities.DataSite{}

		if err := rows.Scan(
			&data_site.SiteId,
			&data_site.SiteName,
			&data_site.SiteAddress,
		); err != nil {
			log.Info("scan rows",
				zap.Any("rows", err),
			)
			return nil, err
		}

		data_sites = append(data_sites, data_site)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return data_sites, nil

}
