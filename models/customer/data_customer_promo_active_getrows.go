package customer

import (
	"context"
	"database/sql"
	"smile-service/entities"

	"go.uber.org/zap"
)

const execSPDataCustomerPromoActive = `exec [sp_smile_promo_active_getrows]$1,$2`

func GetDataCustomerPromoActive(ctx context.Context, db *sql.DB, customer_no, customer_type string, log *zap.Logger) ([]*entities.DataCustomerPromoActive, error) {
	rows, err := db.QueryContext(
		ctx,
		execSPDataCustomerPromoActive,
		customer_no,
		customer_type,
	)

	if err != nil {
		log.Info("Exec DB",
			zap.Any("Exec DB", err),
		)
		return nil, err
	}
	defer rows.Close()

	data_customers_promo_active := make([]*entities.DataCustomerPromoActive, 0)
	for rows.Next() {
		data_customer_promo_active := &entities.DataCustomerPromoActive{}

		if err := rows.Scan(
			&data_customer_promo_active.PromoCode,
			&data_customer_promo_active.PromoName,
			&data_customer_promo_active.PromoDesc,
			&data_customer_promo_active.StartPromoDate,
			&data_customer_promo_active.EndPromoDate,
		); err != nil {
			log.Info("scan rows",
				zap.Any("rows", err),
			)
			return nil, err
		}

		data_customers_promo_active = append(data_customers_promo_active, data_customer_promo_active)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return data_customers_promo_active, nil

}
