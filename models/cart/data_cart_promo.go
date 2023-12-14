package cart

import (
	"context"
	"database/sql"
	"smile-service/entities"

	"go.uber.org/zap"
)

const execSPDataCartPromo = `exec [sp_smile_promo_getrow]$1,$2`
const execSPDataCartPromoEdit = `exec [sp_smile_promo_edit_getrow]$1`

func GetDataCartPromo(ctx context.Context, db *sql.DB, shipTo, salesMan string, log *zap.Logger) ([]*entities.CartPromo, error) {
	rows, err := db.QueryContext(
		ctx,
		execSPDataCartPromo,
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

	data_cart_promos := make([]*entities.CartPromo, 0)
	for rows.Next() {
		data_cart_promo := &entities.CartPromo{}

		if err := rows.Scan(
			&data_cart_promo.PromoCode,
			&data_cart_promo.PromoName,
			&data_cart_promo.ResultDesc,
			&data_cart_promo.AmountDisc,
			&data_cart_promo.IsPrimary,
		); err != nil {
			log.Info("scan rows",
				zap.Any("rows", err),
			)
			return nil, err
		}

		data_cart_promos = append(data_cart_promos, data_cart_promo)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return data_cart_promos, nil

}

func GetDataCartPromoEdit(ctx context.Context, db *sql.DB, transNo string, log *zap.Logger) ([]*entities.CartPromo, error) {
	rows, err := db.QueryContext(
		ctx,
		execSPDataCartPromoEdit,
		transNo,
	)

	if err != nil {
		log.Info("Exec DB",
			zap.Any("Exec DB", err),
		)
		return nil, err
	}
	defer rows.Close()

	data_cart_promos := make([]*entities.CartPromo, 0)
	for rows.Next() {
		data_cart_promo := &entities.CartPromo{}

		if err := rows.Scan(
			&data_cart_promo.PromoCode,
			&data_cart_promo.PromoName,
			&data_cart_promo.ResultDesc,
			&data_cart_promo.AmountDisc,
			&data_cart_promo.IsPrimary,
		); err != nil {
			log.Info("scan rows",
				zap.Any("rows", err),
			)
			return nil, err
		}

		data_cart_promos = append(data_cart_promos, data_cart_promo)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return data_cart_promos, nil

}
