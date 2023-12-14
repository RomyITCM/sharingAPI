package new_customer

import (
	"context"
	"database/sql"
	"encoding/base64"
	"smile-service/entities"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"go.uber.org/zap"
)

const (
	execSPNewCustomerDocumentInsert = `exec [sp_smile_new_customer_document_insert]
	$1, $2, $3, $4, $5, $6, $7, $8, $9, $10`
)

func DataNewCustomerDocumentInsert(ctx context.Context, db *sql.DB,
	data_new_customer_doc *entities.InfoNewCustomerDocument,
	log *zap.Logger) (*entities.DocId, error) {

	sess, _ := session.NewSession(
		&aws.Config{
			Region: aws.String("ap-southeast-1"),
			Credentials: credentials.NewStaticCredentials(
				*aws.String("AKIA3WB2X4HJGJTDWPGL"),
				*aws.String("VVmHtwoOR6xUosj9fi0NW4hRY9i2KVIvMFb7KX+j"),
				"", // a token will be created when the session it's used.
			),
		})

	dec := base64.NewDecoder(base64.StdEncoding, strings.NewReader(data_new_customer_doc.DocImg))

	uploader := s3manager.NewUploader(sess)
	MyBucket := "upload.file"

	dt := time.Now()
	filename := "Benfarm/Smile/Customer/Documents/" + data_new_customer_doc.CustomerRequestNo + "/" + data_new_customer_doc.DocumentNo + "_" + dt.Format("2006-01-02_15_04_05") + ".png"

	_, _ = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(MyBucket),
		ACL:    aws.String("public-read"),
		Key:    aws.String(filename),
		Body:   dec,
	})

	rows, err := db.QueryContext(ctx, execSPNewCustomerDocumentInsert,
		data_new_customer_doc.DocumentType,
		data_new_customer_doc.DocumentNo,
		data_new_customer_doc.DocumentName,
		data_new_customer_doc.DocumentAddress,
		filename,
		data_new_customer_doc.CustomerRequestNo,
		data_new_customer_doc.CustomerNo,
		data_new_customer_doc.BankCode,
		data_new_customer_doc.CreatedBy,
		data_new_customer_doc.CreatedByIp,
	)

	if err != nil {
		log.Info(
			"Exec DB", zap.Any("Exec Db", err),
		)

		return nil, err
	}
	defer rows.Close()

	doc_id := &entities.DocId{}
	for rows.Next() {
		if err := rows.Scan(
			&doc_id.DocId,
		); err != nil {
			log.Info("scan rows",
				zap.Any("rows", err),
			)
			return nil, err
		}
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return doc_id, nil
}
