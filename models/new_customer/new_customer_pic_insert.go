package new_customer

import (
	"context"
	"database/sql"
	"smile-service/entities"

	"go.uber.org/zap"
)

const (
	execSPCustomerPICInsert = `exec [sp_smile_customer_pic_insert] $1,$2,$3,$4,$5,$6,$7,$8,$9,$10`
)

func DataNewCustomerPicInsert(ctx context.Context, db *sql.DB,
	data_new_customer_pic *entities.InfoNewCustomerPic,
	log *zap.Logger) (*entities.PicId, error) {

	rows, err := db.QueryContext(ctx, execSPCustomerPICInsert,
		data_new_customer_pic.CustomerId,
		data_new_customer_pic.Honorific,
		data_new_customer_pic.PicName,
		data_new_customer_pic.Position,
		data_new_customer_pic.PhoneNo,
		data_new_customer_pic.Email,
		data_new_customer_pic.OutletId,
		data_new_customer_pic.CustomerAddressId,
		data_new_customer_pic.CreatedBy,
		data_new_customer_pic.CreatedByIp,
	)

	if err != nil {
		log.Info(
			"Exec DB", zap.Any("Exec Db", err),
		)

		return nil, err
	}
	defer rows.Close()

	pic_id := &entities.PicId{}
	for rows.Next() {
		if err := rows.Scan(
			&pic_id.PicId,
		); err != nil {
			log.Info("scan rows",
				zap.Any("rows", err),
			)
			return nil, err
		}
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return pic_id, nil
}
