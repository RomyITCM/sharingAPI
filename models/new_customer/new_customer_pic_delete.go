package new_customer

import (
	"context"
	"database/sql"
	"smile-service/entities"
)

const (
	execSPCustomerPicDelete = `exec [sp_smile_customer_pic_delete] $1`
)

func DataCustomerPicDelete(
	ctx context.Context, db *sql.DB,
	PicId string,
) error {

	rows, err := db.QueryContext(
		ctx, execSPCustomerPicDelete,
		PicId)

	if err != nil {
		return err
	}
	defer rows.Close()

	return nil
}

func DeleteCustomerPic(ctx context.Context, db *sql.DB, picId *entities.PicId) error {
	err := DataCustomerPicDelete(ctx, db, picId.PicId)

	if err != nil {
		return err
	}

	return nil
}
