package store_visit

import (
	"context"
	"database/sql"
	"smile-service/entities"
)

const (
	execSPFindingInsert = `exec [sp_smile_finding_insert]$1,$2,$3,$4,$5`
)

func DataFindingInsert(ctx context.Context, db *sql.DB,
	EmailTo string,
	Subject string,
	Remark string,
	CreatedBy string,
	CreatedByIp string) error {
	rows, err := db.QueryContext(ctx, execSPFindingInsert,
		EmailTo,
		Subject,
		Remark,
		CreatedBy,
		CreatedByIp,
	)
	if err != nil {
		return err
	}
	defer rows.Close()

	return nil
}

func InsertFinding(ctx context.Context, db *sql.DB, dataFinding *entities.DataFindingInsert) error {

	err := DataFindingInsert(ctx, db,
		dataFinding.EmailTo,
		dataFinding.Subject,
		dataFinding.Remark,
		dataFinding.CreatedBy,
		dataFinding.CreatedByIP,
	)

	if err != nil {
		return err
	}

	return nil
}
