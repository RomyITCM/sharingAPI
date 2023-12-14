package need_approve

import (
	"context"
	"database/sql"
	"smile-service/entities"

	"go.uber.org/zap"
)

const execSPNeedApproveOutlet = `exec [sp_smile_need_approve_outlet_getrows] $1`

func GetNeedApproveOutlet(ctx context.Context, db *sql.DB,
	search string, log *zap.Logger) ([]*entities.DataNeedApproveOutlet, error) {

	rows, err := db.QueryContext(
		ctx,
		execSPNeedApproveOutlet,
		search,
	)

	if err != nil {
		log.Info("Exec DB",
			zap.Any("Exec DB", err),
		)
		return nil, err
	}
	defer rows.Close()

	data_need_approve_outlets := make([]*entities.DataNeedApproveOutlet, 0)

	for rows.Next() {
		data_need_approve_outlet := &entities.DataNeedApproveOutlet{}

		if err := rows.Scan(
			&data_need_approve_outlet.OutletId,
			&data_need_approve_outlet.OutletName,
			&data_need_approve_outlet.Address,
			&data_need_approve_outlet.ProvinceId,
			&data_need_approve_outlet.CityId,
			&data_need_approve_outlet.DistrictId,
			&data_need_approve_outlet.Latitude,
			&data_need_approve_outlet.Longitude,
			&data_need_approve_outlet.StoreFreezer,
			&data_need_approve_outlet.BenFreezer,
		); err != nil {
			log.Info("scan rows",
				zap.Any("rows", err),
			)
			return nil, err
		}

		data_need_approve_outlets = append(data_need_approve_outlets, data_need_approve_outlet)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return data_need_approve_outlets, nil
}
