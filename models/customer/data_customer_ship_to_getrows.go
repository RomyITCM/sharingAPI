package customer

import (
	"context"
	"database/sql"
	"smile-service/entities"

	"go.uber.org/zap"
)

const execSPDataCustomerShipTo = `exec [sp_smile_customer_ship_to_getrows]$1,$2`

func GetDataCustomerShipTo(ctx context.Context, db *sql.DB, customer string, search string, log *zap.Logger) ([]*entities.DataCustomerShipTo, error) {
	rows, err := db.QueryContext(
		ctx,
		execSPDataCustomerShipTo,
		customer,
		search,
	)

	if err != nil {
		log.Info("Exec DB",
			zap.Any("Exec DB", err),
		)
		return nil, err
	}
	defer rows.Close()

	data_customers_ship_to := make([]*entities.DataCustomerShipTo, 0)
	for rows.Next() {
		data_customer_ship_to := &entities.DataCustomerShipTo{}

		if err := rows.Scan(
			&data_customer_ship_to.ShipToId,
			&data_customer_ship_to.ShipToName,
			&data_customer_ship_to.ShipToAddress,
			&data_customer_ship_to.PriceZone,
			&data_customer_ship_to.SiteId,
			&data_customer_ship_to.SiteName,
			&data_customer_ship_to.SiteAddress,
			&data_customer_ship_to.ShipToStatus,
			&data_customer_ship_to.TripZone,
		); err != nil {
			log.Info("scan rows",
				zap.Any("rows", err),
			)
			return nil, err
		}

		data_customers_ship_to = append(data_customers_ship_to, data_customer_ship_to)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return data_customers_ship_to, nil

}
