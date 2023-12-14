package new_customer

import (
	"context"
	"database/sql"
	"smile-service/entities"

	"go.uber.org/zap"
)

const execSPCustomerPICRow = `exec [sp_smile_customer_pic_getrow] $1`

func GetRowNewCustomerPic(ctx context.Context, db *sql.DB, pic_id string, log *zap.Logger) ([]*entities.DataNewCustomerPicDetail, error) {
	rows, err := db.QueryContext(
		ctx,
		execSPCustomerPICRow,
		pic_id,
	)

	if err != nil {
		log.Info(
			"Exec DB",
			zap.Any("Exec DB", err),
		)
		return nil, err
	}
	defer rows.Close()

	row_data_customer_pics := make([]*entities.DataNewCustomerPicDetail, 0)
	for rows.Next() {
		row_data_customer_pic := &entities.DataNewCustomerPicDetail{}

		if err := rows.Scan(
			&row_data_customer_pic.PicId,
			&row_data_customer_pic.Honorific,
			&row_data_customer_pic.PicName,
			&row_data_customer_pic.Position,
			&row_data_customer_pic.PhoneNo,
			&row_data_customer_pic.Email,
			&row_data_customer_pic.Status,
			&row_data_customer_pic.OutletId,
			&row_data_customer_pic.CustomerAddressId,
		); err != nil {
			log.Info("scan rows",
				zap.Any("rows", err),
			)
			return nil, err
		}

		row_data_customer_pics = append(row_data_customer_pics, row_data_customer_pic)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return row_data_customer_pics, nil
}
