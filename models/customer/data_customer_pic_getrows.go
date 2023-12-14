package customer

import (
	"context"
	"database/sql"
	"smile-service/entities"

	"go.uber.org/zap"
)

const execSPCustomerPicGetRows = `exec [sp_smile_approved_customer_pic_getrows] $1`

func GetRowsDataCustomerPic(ctx context.Context, db *sql.DB,
	customerNo string, log *zap.Logger) ([]*entities.DataNewCustomerPic, error) {

	rows, err := db.QueryContext(
		ctx,
		execSPCustomerPicGetRows,
		customerNo,
	)

	if err != nil {
		log.Info(
			"Exec DB",
			zap.Any("Exec DB", err),
		)
		return nil, err
	}
	defer rows.Close()

	data_customer_pics := make([]*entities.DataNewCustomerPic, 0)
	for rows.Next() {
		data_customer_pic := &entities.DataNewCustomerPic{}

		if err := rows.Scan(
			&data_customer_pic.PicId,
			&data_customer_pic.Honorific,
			&data_customer_pic.PicName,
			&data_customer_pic.Position,
			&data_customer_pic.PhoneNo,
			&data_customer_pic.Email,
			&data_customer_pic.CustomerAddressId,
			&data_customer_pic.CustomerAddressName,
		); err != nil {
			log.Info("scan rows",
				zap.Any("rows", err),
			)
			return nil, err
		}

		data_customer_pics = append(data_customer_pics, data_customer_pic)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return data_customer_pics, nil
}
