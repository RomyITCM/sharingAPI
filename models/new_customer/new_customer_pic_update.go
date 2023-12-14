package new_customer

import (
	"context"
	"database/sql"
	"smile-service/entities"
)

const (
	execSPCustomerPicUpdate = `exec [sp_smile_customer_pic_update]
	$1,$2,$3,$4,$5,$6,$7,$8,$9,$10`
)

func CustomerPicUpdate(ctx context.Context, db *sql.DB,
	PicId string,
	Honorific string,
	PicName string,
	Position string,
	PhoneNo string,
	Email string,
	OutletId string,
	CustomerAddressId string,
	CreatedBy string,
	CreatedByIp string) error {

	rows, err := db.QueryContext(ctx, execSPCustomerPicUpdate,
		PicId,
		Honorific,
		PicName,
		Position,
		PhoneNo,
		Email,
		OutletId,
		CustomerAddressId,
		CreatedBy,
		CreatedByIp,
	)
	if err != nil {
		return err
	}
	defer rows.Close()

	return nil
}

func UpdateCustomerPic(ctx context.Context, db *sql.DB,
	data_new_customer_pic *entities.InfoNewCustomerPicUpdate) error {

	err := CustomerPicUpdate(ctx, db,
		data_new_customer_pic.PicId,
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
		return err
	}

	return nil
}
