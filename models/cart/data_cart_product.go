package cart

import (
	"context"
	"database/sql"
	"smile-service/entities"

	"go.uber.org/zap"
)

const execSPDataCartProduct = `exec [sp_smile_cart_product_getrows]$1,$2`
const execSPDataCartProductEdit = `exec [sp_smile_cart_product_edit_getrows]$1`

func GetDataCartProduct(ctx context.Context, db *sql.DB, shipTo, salesMan string, log *zap.Logger) ([]*entities.CartProduct, error) {
	rows, err := db.QueryContext(
		ctx,
		execSPDataCartProduct,
		shipTo,
		salesMan,
	)

	if err != nil {
		log.Info("Exec DB",
			zap.Any("Exec DB", err),
		)
		return nil, err
	}
	defer rows.Close()

	data_cart_products := make([]*entities.CartProduct, 0)
	for rows.Next() {
		data_cart_product := &entities.CartProduct{}

		if err := rows.Scan(
			&data_cart_product.ArticleDesc,
			&data_cart_product.Qty,
			&data_cart_product.Uom,
			&data_cart_product.Price,
			&data_cart_product.Total,
			&data_cart_product.Image,
		); err != nil {
			log.Info("scan rows",
				zap.Any("rows", err),
			)
			return nil, err
		}

		data_cart_products = append(data_cart_products, data_cart_product)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return data_cart_products, nil

}

func GetDataCartProductEdit(ctx context.Context, db *sql.DB, transNo string, log *zap.Logger) ([]*entities.CartProduct, error) {
	rows, err := db.QueryContext(
		ctx,
		execSPDataCartProductEdit,
		transNo,
	)

	if err != nil {
		log.Info("Exec DB",
			zap.Any("Exec DB", err),
		)
		return nil, err
	}
	defer rows.Close()

	data_cart_products := make([]*entities.CartProduct, 0)
	for rows.Next() {
		data_cart_product := &entities.CartProduct{}

		if err := rows.Scan(
			&data_cart_product.ArticleDesc,
			&data_cart_product.Qty,
			&data_cart_product.Uom,
			&data_cart_product.Price,
			&data_cart_product.Total,
			&data_cart_product.Image,
		); err != nil {
			log.Info("scan rows",
				zap.Any("rows", err),
			)
			return nil, err
		}

		data_cart_products = append(data_cart_products, data_cart_product)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return data_cart_products, nil

}
