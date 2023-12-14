package new_customer

import (
	"context"
	"database/sql"
	"smile-service/entities"

	"go.uber.org/zap"
)

const execSPCustomerPicById = `exec [sp_smile_customer_pic_getrows] $1`

func GetDataNewCustomerPic(ctx context.Context, db *sql.DB, customer_id string, log *zap.Logger) ([]*entities.DataNewCustomerPic, error) {
	rows, err := db.QueryContext(
		ctx,
		execSPCustomerPicById,
		customer_id,
	)

	if err != nil {
		log.Info(
			"Exec DB",
			zap.Any("Exec DB", err),
		)
		return nil, err
	}
	defer rows.Close()

	data_new_customer_pics := make([]*entities.DataNewCustomerPic, 0)
	for rows.Next() {
		data_new_customer_pic := &entities.DataNewCustomerPic{}

		if err := rows.Scan(
			&data_new_customer_pic.PicId,
			&data_new_customer_pic.Honorific,
			&data_new_customer_pic.PicName,
			&data_new_customer_pic.Position,
			&data_new_customer_pic.PhoneNo,
			&data_new_customer_pic.Email,
			&data_new_customer_pic.OutletId,
			&data_new_customer_pic.OutletName,
			&data_new_customer_pic.CustomerAddressId,
			&data_new_customer_pic.CustomerAddressName,
		); err != nil {
			log.Info("scan rows",
				zap.Any("rows", err),
			)
			return nil, err
		}

		data_new_customer_pics = append(data_new_customer_pics, data_new_customer_pic)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return data_new_customer_pics, nil
}
