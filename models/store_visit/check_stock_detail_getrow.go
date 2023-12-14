package store_visit

import (
	"context"
	"database/sql"
	"smile-service/entities"

	"go.uber.org/zap"
)

const execSPCheckStockDetailGetrow = `exec [sp_smile_check_stock_product_stock_getrow]$1,$2`

const execSPCheckStockDetailImagesGetrow = `exec [sp_smile_check_stock_product_images_getrow]$1,$2,$3,$4`

func GetDataCheckStockDetailGetrow(ctx context.Context, db *sql.DB, trans_no string, article_no string, log *zap.Logger) ([]*entities.DataCheckStockDetailGetrow, error) {
	rows, err := db.QueryContext(
		ctx,
		execSPCheckStockDetailGetrow,
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

	data_check_stocks := make([]*entities.DataCheckStockDetailGetrow, 0)
	// data_check_stocks := make([]interface{}, 0)

	for rows.Next() {
		data_check_stock := &entities.DataCheckStockDetailGetrow{}
		/////////////////////////////////////////////////////////////////////////////////////

		// yg lama dikomen dulu
		if err := rows.Scan(
			&data_check_stock.ID,
			&data_check_stock.TransNo,
			&data_check_stock.ArticleNumber,
			&data_check_stock.ExpiredDate,
			&data_check_stock.QtyStock,
			&data_check_stock.Condition,
			&data_check_stock.NoteStock,
			&data_check_stock.SalesReturnNo,
			&data_check_stock.NoteSalesReturn,
		); err != nil {
			log.Info("scan rows",
				zap.Any("rows", err),
			)
			return nil, err
		}

		data_check_stocks = append(data_check_stocks, data_check_stock)
		////////////////////////////////////////////////////////////////////////////////////////////////////////////////
		rows2, err := db.QueryContext(
			ctx,
			execSPCheckStockDetailImagesGetrow,
			trans_no,
			article_no,
			data_check_stock.Condition,
			data_check_stock.ID,
		)

		if err != nil {
			log.Info("Exec DB",
				zap.Any("Exec DB", err),
			)
			return nil, err
		}
		defer rows2.Close()

		data_check_stock_images := make([]*entities.DataCheckStockDetailImageGetrow, 0)

		for rows2.Next() {
			data_check_stock_image := &entities.DataCheckStockDetailImageGetrow{}

			if err := rows2.Scan(
				&data_check_stock_image.ImageStock,
			); err != nil {
				log.Info("scan rows",
					zap.Any("rows", err),
				)
				return nil, err
			}

			data_check_stock_images = append(data_check_stock_images, data_check_stock_image)
			// data_check_stock.ImageStock = *data_check_stock_images[0]

		}
		if err = rows2.Err(); err != nil {
			return nil, err
		}
		/////////////////////////////////////

		data_check_stock.ImageStock = data_check_stock_images

	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return data_check_stocks, nil
}
