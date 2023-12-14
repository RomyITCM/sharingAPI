package cart

import (
	"context"
	"database/sql"
	"smile-service/entities"

	"go.uber.org/zap"
)

const execSPDataCartDelivery = `exec [sp_smile_cart_delivery_getrows]$1`
const execSPDataCartDeliveryEdit = `exec [sp_smile_cart_delivery_edit_getrows]$1`

func GetDataCartDelivery(ctx context.Context, db *sql.DB, shipTo string, log *zap.Logger) ([]*entities.CartDelivery, error) {
	rows, err := db.QueryContext(
		ctx,
		execSPDataCartDelivery,
		shipTo,
	)

	if err != nil {
		log.Info("Exec DB",
			zap.Any("Exec DB", err),
		)
		return nil, err
	}
	defer rows.Close()

	data_cart_deliveries := make([]*entities.CartDelivery, 0)
	for rows.Next() {
		data_cart_delivery := &entities.CartDelivery{}

		if err := rows.Scan(
			&data_cart_delivery.VendorNo,
			&data_cart_delivery.VendorName,
			&data_cart_delivery.ShipTo,
			&data_cart_delivery.ShipToAddress,
			&data_cart_delivery.PaymentTermId,
			&data_cart_delivery.PaymentTerm,
		); err != nil {
			log.Info("scan rows",
				zap.Any("rows", err),
			)
			return nil, err
		}

		data_cart_deliveries = append(data_cart_deliveries, data_cart_delivery)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return data_cart_deliveries, nil

}

func GetDataCartDeliveryEdit(ctx context.Context, db *sql.DB, transNo string, log *zap.Logger) ([]*entities.CartDeliveryEdit, error) {
	rows, err := db.QueryContext(
		ctx,
		execSPDataCartDeliveryEdit,
		transNo,
	)

	if err != nil {
		log.Info("Exec DB",
			zap.Any("Exec DB", err),
		)
		return nil, err
	}
	defer rows.Close()

	data_cart_deliveries := make([]*entities.CartDeliveryEdit, 0)
	for rows.Next() {
		data_cart_delivery := &entities.CartDeliveryEdit{}

		if err := rows.Scan(
			&data_cart_delivery.CustPoNo,
			&data_cart_delivery.PoDate,
			&data_cart_delivery.ShipTo,
			&data_cart_delivery.ShipToAddress,
			&data_cart_delivery.PaymentTermId,
			&data_cart_delivery.PaymentTerm,
			&data_cart_delivery.BillTo,
			&data_cart_delivery.BillToAddress,
			&data_cart_delivery.Attachment,
			&data_cart_delivery.DeliveryFrom,
			&data_cart_delivery.CreatedDate,
			&data_cart_delivery.DelvDate,
			&data_cart_delivery.PoExpDate,
			&data_cart_delivery.CustType,
		); err != nil {
			log.Info("scan rows",
				zap.Any("rows", err),
			)
			return nil, err
		}

		data_cart_deliveries = append(data_cart_deliveries, data_cart_delivery)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return data_cart_deliveries, nil

}
