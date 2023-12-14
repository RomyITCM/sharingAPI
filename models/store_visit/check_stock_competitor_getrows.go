package store_visit

import (
	"context"
	"database/sql"
	"smile-service/entities"

	"go.uber.org/zap"
)

const execSPCheckStockCompetitorList = `exec [sp_smile_check_stock_competitor_product_getrows]$1,$2,$3`

func GetDataCheckStockCompetitorList(ctx context.Context, db *sql.DB, shipTo string, skuArticle string, search string, log *zap.Logger) ([]*entities.DataCheckStockCompetitorGetrows, error) {
	rows, err := db.QueryContext(
		ctx,
		execSPCheckStockCompetitorList,
		shipTo,
		skuArticle,
		search,
	)

	if err != nil {
		log.Info("Exec DB",
			zap.Any("Exec DB", err),
		)
		return nil, err
	}
	defer rows.Close()

	data_check_stocks := make([]*entities.DataCheckStockCompetitorGetrows, 0)

	for rows.Next() {
		data_check_stock := &entities.DataCheckStockCompetitorGetrows{}

		if err := rows.Scan(
			&data_check_stock.ArticleNumber,
			&data_check_stock.UrlImage,
			&data_check_stock.ArticleDescription,
			&data_check_stock.SalesPrice,
			&data_check_stock.Stock,
			&data_check_stock.Checked,
		); err != nil {
			log.Info("scan rows",
				zap.Any("rows", err),
			)
			return nil, err
		}

		data_check_stocks = append(data_check_stocks, data_check_stock)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return data_check_stocks, nil
}
