package customer

import (
	"context"
	"database/sql"
	"smile-service/entities"

	"go.uber.org/zap"
)

const (
	execSPCustomerAssetInsert = `[sp_customer_asset_insert] $1, $2, $3, $4, $5, $6, $7`
)

func DataCustomerAssetInsert(ctx context.Context, db *sql.DB,
	data_customer_asset *entities.InfoCustomerAsset,
	log *zap.Logger) error {

	rows, err := db.QueryContext(
		ctx, execSPCustomerAssetInsert,
		data_customer_asset.ShipTo,
		data_customer_asset.FreezerName,
		data_customer_asset.Location,
		data_customer_asset.Usage,
		data_customer_asset.CreateBy,
		data_customer_asset.CreateDate,
		data_customer_asset.CreateIpAddress,
	)

	if err != nil {
		log.Info(
			"Exec DB", zap.Any("Exec Db", err),
		)

		return err
	}
	defer rows.Close()

	return nil
}
