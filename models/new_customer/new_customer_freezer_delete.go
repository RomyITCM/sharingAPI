package new_customer

import (
	"context"
	"database/sql"
	"smile-service/entities"
)

const (
	execSPOutletFreezerDelete = `exec [sp_smile_outlet_freezer_delete] $1`
)

func DataOutletFreezerDelete(ctx context.Context, db *sql.DB,
	FreezerId string) error {

	rows, err := db.QueryContext(ctx, execSPOutletFreezerDelete, FreezerId)

	if err != nil {
		return err
	}
	defer rows.Close()

	return nil
}

func DeleteOutletFreezer(ctx context.Context, db *sql.DB, freezerId *entities.FreezerId) error {
	err := DataOutletFreezerDelete(ctx, db, freezerId.FreezerId)

	if err != nil {
		return err
	}

	return nil
}
