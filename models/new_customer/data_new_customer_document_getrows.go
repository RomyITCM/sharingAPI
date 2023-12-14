package new_customer

import (
	"context"
	"database/sql"
	"smile-service/entities"

	"go.uber.org/zap"
)

const execSPNewCustomerDocument = `exec [sp_smile_new_customer_document_getrows] $1`

func GetRowsNewCustomerDocument(ctx context.Context, db *sql.DB,
	customer_request_no string, log *zap.Logger) ([]*entities.DataNewCustomerDocument, error) {

	rows, err := db.QueryContext(
		ctx,
		execSPNewCustomerDocument,
		customer_request_no,
	)

	if err != nil {
		log.Info(
			"Exec DB",
			zap.Any("Exec DB", err),
		)
		return nil, err
	}
	defer rows.Close()

	data_new_customer_documents := make([]*entities.DataNewCustomerDocument, 0)
	for rows.Next() {
		data_new_customer_document := &entities.DataNewCustomerDocument{}

		if err := rows.Scan(
			&data_new_customer_document.DocId,
			&data_new_customer_document.DocumentType,
			&data_new_customer_document.DocumentNo,
			&data_new_customer_document.DocumentName,
			&data_new_customer_document.DocumentAddress,
			&data_new_customer_document.DocImg,
			&data_new_customer_document.CustomerRequestNo,
			&data_new_customer_document.CustomerNo,
			&data_new_customer_document.BankCode,
			&data_new_customer_document.BankName,
		); err != nil {
			log.Info("scan rows",
				zap.Any("rows", err),
			)
			return nil, err
		}

		data_new_customer_documents = append(data_new_customer_documents, data_new_customer_document)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return data_new_customer_documents, nil
}
