package store_visit

import (
	"context"
	"database/sql"
	"smile-service/entities"

	"go.uber.org/zap"
)

const execSPCheckStockGetrow = `exec [sp_smile_check_stock_product_getrow]$1,$2`

func GetDataCheckStockGetrow(ctx context.Context, db *sql.DB, trans_no string, article_no string, log *zap.Logger) ([]*entities.DataCheckStockGetrow, error) {
	rows, err := db.QueryContext(
		ctx,
		execSPCheckStockGetrow,
		trans_no,
		article_no,
	)

	if err != nil {
		log.Info("Exec DB",
			zap.Any("Exec DB", err),
		)
		return nil, err
	}
	defer rows.Close()

	data_check_stocks := make([]*entities.DataCheckStockGetrow, 0)

	for rows.Next() {
		data_check_stock := &entities.DataCheckStockGetrow{}

		if err := rows.Scan(
			&data_check_stock.TransNo,
			&data_check_stock.ArticleNo,
			&data_check_stock.Type,
			&data_check_stock.StockAvailable,
			&data_check_stock.Price,
			&data_check_stock.PromoAvailable,
			&data_check_stock.MediaPromoAvailable,
			&data_check_stock.NotePromo,
			&data_check_stock.StartDatePromo,
			&data_check_stock.EndDatePromo,
			&data_check_stock.ImageType,
			&data_check_stock.ImagesName,
			&data_check_stock.GoodStock,
			&data_check_stock.BadStock,
			&data_check_stock.ExpiredStock,
			&data_check_stock.TotalStock,
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
