package need_approve

import (
	"context"
	"database/sql"
	"smile-service/entities"

	"go.uber.org/zap"
)

const (
	execSPCustomerDocumentInsert = `exec [sp_smile_customer_documents_insert] $1,$2,$3,$4,$5`
)

func DataCustomerDocumentInsert(ctx context.Context, db *sql.DB,
	data_cust_doc *entities.InfoCustomerDocument, log *zap.Logger,
) error {

	rows, err := db.QueryContext(ctx, execSPCustomerDocumentInsert,
		data_cust_doc.DocumentNo,
		data_cust_doc.CustomerNo,
		data_cust_doc.DocImg,
		// data_cust_doc.DocPath,
		data_cust_doc.CreatedBy,
		data_cust_doc.CreatedByIp,
	)

	if err != nil {
		log.Info(
			"Exec DB", zap.Any("Exec Db", err),
		)

		return nil
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(); err != nil {
			log.Info("scan rows",
				zap.Any("rows", err),
			)
			return nil
		}
	}
	if err = rows.Err(); err != nil {
		return nil
	}

	return nil
}
