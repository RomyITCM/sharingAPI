package cart

import (
	"context"
	"database/sql"
	"smile-service/entities"

	"go.uber.org/zap"
)

const execSPDataCartPaymentTerm = `exec [sp_smile_payment_term_getrows]`

func GetDataCartPaymentTerm(ctx context.Context, db *sql.DB, log *zap.Logger) ([]*entities.CartPaymentTerm, error) {
	rows, err := db.QueryContext(
		ctx,
		execSPDataCartPaymentTerm,
	)

	if err != nil {
		log.Info("Exec DB",
			zap.Any("Exec DB", err),
		)
		return nil, err
	}
	defer rows.Close()

	data_cart_payment_terms := make([]*entities.CartPaymentTerm, 0)
	for rows.Next() {
		data_cart_payment_term := &entities.CartPaymentTerm{}

		if err := rows.Scan(
			&data_cart_payment_term.ID,
			&data_cart_payment_term.Title,
			&data_cart_payment_term.Description,
		); err != nil {
			log.Info("scan rows",
				zap.Any("rows", err),
			)
			return nil, err
		}

		data_cart_payment_terms = append(data_cart_payment_terms, data_cart_payment_term)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return data_cart_payment_terms, nil

}
