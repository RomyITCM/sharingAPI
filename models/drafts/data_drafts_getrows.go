package transaction_hold

import (
	"context"
	"database/sql"
	"smile-service/entities"

	"go.uber.org/zap"
)

const execSPDataTransactionHold = `exec [sp_smile_drafts_getrows]$1,$2`

func GetDataTransactionHold(ctx context.Context, db *sql.DB, salesMan, search string, log *zap.Logger) ([]*entities.DataDrafts, error) {
	rows, err := db.QueryContext(
		ctx,
		execSPDataTransactionHold,
		salesMan,
		search,
	)

	if err != nil {
		log.Info("Exec DB",
			zap.Any("Exec DB", err),
		)
		return nil, err
	}
	defer rows.Close()

	data_drafts := make([]*entities.DataDrafts, 0)
	for rows.Next() {
		data_draft := &entities.DataDrafts{}

		if err := rows.Scan(
			&data_draft.ShipTo,
			&data_draft.TotalPrice,
			&data_draft.ArticleDesc,
			&data_draft.Qty,
			&data_draft.Note,
			&data_draft.UrlImage,
			&data_draft.CustomerNo,
			&data_draft.CustomerName,
			&data_draft.CustomerAddress,
			&data_draft.CustomerType,
			&data_draft.CustomerStatus,
			&data_draft.CustomerTypeDesc,
			&data_draft.ShipToId,
			&data_draft.ShipToAddress,
			&data_draft.PriceZone,
			&data_draft.SiteId,
			&data_draft.SiteName,
			&data_draft.SiteAddress,
			&data_draft.ShipToStatus,
			&data_draft.TripZone,
		); err != nil {
			log.Info("scan rows",
				zap.Any("rows", err),
			)
			return nil, err
		}

		data_drafts = append(data_drafts, data_draft)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return data_drafts, nil

}
