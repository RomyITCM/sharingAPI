package new_customer

import (
	"context"
	"database/sql"
	"smile-service/entities"
)

const (
	execSPCustomerOutletDelete = `exec [sp_smile_customer_outlet_delete] $1`
)

func DataNewCustomerOutletDelete(ctx context.Context, db *sql.DB,
	OutletId string) error {

	rows, err := db.QueryContext(ctx, execSPCustomerOutletDelete, OutletId)

	if err != nil {
		return err
	}
	defer rows.Close()

	return nil
}

func DeleteCustomerOutlet(ctx context.Context, db *sql.DB, outletId *entities.OutletId) error {
	err := DataNewCustomerOutletDelete(ctx, db, outletId.OutletId)

	if err != nil {
		return err
	}

	return nil
}
