package new_customer

import (
	"context"
	"database/sql"
	"smile-service/entities"

	"go.uber.org/zap"
)

const execSPNewCustomerDocumentByType = `exec [sp_smile_new_customer_document_by_type_getrow] $1, $2`

func GetRowNewCustomerDocumentByType(ctx context.Context, db *sql.DB,
	customer_request_no string, doc_type string, log *zap.Logger) ([]*entities.DataNewCustomerDocument, error) {

	rows, err := db.QueryContext(
		ctx,
		execSPNewCustomerDocumentByType,
		customer_request_no,
		doc_type,
	)

	if err != nil {
		log.Info(
			"Exec DB",
			zap.Any("Exec DB", err),
		)
		return nil, err
	}
	defer rows.Close()

	dataCustDocTypes := make([]*entities.DataNewCustomerDocument, 0)
	for rows.Next() {
		dataCustDocType := &entities.DataNewCustomerDocument{}

		if err := rows.Scan(
			&dataCustDocType.DocId,
			&dataCustDocType.DocumentType,
			&dataCustDocType.DocumentNo,
			&dataCustDocType.DocumentName,
			&dataCustDocType.DocumentAddress,
			&dataCustDocType.DocImg,
			&dataCustDocType.CustomerRequestNo,
			&dataCustDocType.CustomerNo,
			&dataCustDocType.BankCode,
			&dataCustDocType.BankName,
		); err != nil {
			log.Info("scan rows",
				zap.Any("rows", err),
			)
			return nil, err
		}

		dataCustDocTypes = append(dataCustDocTypes, dataCustDocType)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return dataCustDocTypes, nil
}
