package store_visit

import (
	"context"
	"database/sql"
	"smile-service/entities"

	"go.uber.org/zap"
)

const execSPCheckStockList = `exec [sp_smile_check_stock_product_getrows]$1,$2,$3`

func GetDataCheckStockList(ctx context.Context, db *sql.DB, ship_to string, search string, created_by string, log *zap.Logger) ([]*entities.DataCheckStockGetrows, error) {
	rows, err := db.QueryContext(
		ctx,
		execSPCheckStockList,
		ship_to,
		search,
		created_by,
	)

	if err != nil {
		log.Info("Exec DB",
			zap.Any("Exec DB", err),
		)
		return nil, err
	}
	defer rows.Close()

	data_check_stocks := make([]*entities.DataCheckStockGetrows, 0)

	for rows.Next() {
		data_check_stock := &entities.DataCheckStockGetrows{}

		if err := rows.Scan(
			&data_check_stock.TransNo,
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
