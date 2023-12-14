package product

import (
	"context"
	"database/sql"
	"smile-service/entities"

	"go.uber.org/zap"
)

const execSPDataProduct = `exec [sp_smile_data_product_getrows]$1,$2,$3`
const execSPDataProductEdit = `exec [sp_smile_data_product_edit_getrows]$1`

func GetDataProduct(ctx context.Context, db *sql.DB, site_id, zone_id, sales_man string, log *zap.Logger) ([]*entities.DataProduct, error) {
	rows, err := db.QueryContext(
		ctx,
		execSPDataProduct,
		site_id,
		zone_id,
		sales_man,
	)

	if err != nil {
		log.Info("Exec DB",
			zap.Any("Exec DB", err),
		)
		return nil, err
	}
	defer rows.Close()

	data_products := make([]*entities.DataProduct, 0)
	for rows.Next() {
		data_product := &entities.DataProduct{}

		if err := rows.Scan(
			&data_product.ArticleNumber,
			&data_product.UrlImage,
			&data_product.ArticleDescription,
			&data_product.SalesPrice,
			&data_product.Stock,
			&data_product.QtyBook,
		); err != nil {
			log.Info("scan rows",
				zap.Any("rows", err),
			)
			return nil, err
		}

		data_products = append(data_products, data_product)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return data_products, nil

}

func GetDataProductEdit(ctx context.Context, db *sql.DB, trans_no string, log *zap.Logger) ([]*entities.DataProduct, error) {
	rows, err := db.QueryContext(
		ctx,
		execSPDataProductEdit,
		trans_no,
	)

	if err != nil {
		log.Info("Exec DB",
			zap.Any("Exec DB", err),
		)
		return nil, err
	}
	defer rows.Close()

	data_products := make([]*entities.DataProduct, 0)
	for rows.Next() {
		data_product := &entities.DataProduct{}

		if err := rows.Scan(
			&data_product.ArticleNumber,
			&data_product.UrlImage,
			&data_product.ArticleDescription,
			&data_product.SalesPrice,
			&data_product.Stock,
			&data_product.QtyBook,
		); err != nil {
			log.Info("scan rows",
				zap.Any("rows", err),
			)
			return nil, err
		}

		data_products = append(data_products, data_product)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return data_products, nil

}
