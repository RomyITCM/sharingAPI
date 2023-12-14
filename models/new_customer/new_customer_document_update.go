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
)

const (
	execSPNewCustomerDocumentUpdate = `exec [sp_smile_new_customer_document_update]
	$1, $2, $3, $4, $5, $6, $7, $8, $9`
)

func CustomerDocumentUpdate(ctx context.Context, db *sql.DB,
	CustomerRequestNo string,
	DocumentType string,
	DocumentNo string,
	DocumentName string,
	DocumentAddress string,
	DocImg string,
	BankCode string,
	CreatedBy string,
	CreatedByIp string) error {

	rows, err := db.QueryContext(ctx, execSPNewCustomerDocumentUpdate,
		CustomerRequestNo,
		DocumentType,
		DocumentNo,
		DocumentName,
		DocumentAddress,
		DocImg,
		BankCode,
		CreatedBy,
		CreatedByIp,
	)
	if err != nil {
		return err
	}
	defer rows.Close()

	return nil
}

func DataNewCustomerDocumentUpdate(ctx context.Context, db *sql.DB,
	data_new_doc *entities.InfoNewCustomerDocumentUpdate) error {

	filename := ""

	if data_new_doc.DocImg != "" {

		sess, _ := session.NewSession(
			&aws.Config{
				Region: aws.String("ap-southeast-1"),
				Credentials: credentials.NewStaticCredentials(
					*aws.String("AKIA3WB2X4HJGJTDWPGL"),
					*aws.String("VVmHtwoOR6xUosj9fi0NW4hRY9i2KVIvMFb7KX+j"),
					"", // a token will be created when the session it's used.
				),
			})

		dec := base64.NewDecoder(base64.StdEncoding, strings.NewReader(data_new_doc.DocImg))

		uploader := s3manager.NewUploader(sess)
		MyBucket := "upload.file"

		dt := time.Now()
		filename = "Benfarm/Smile/Customer/Documents/" + data_new_doc.CustomerRequestNo + "/" + data_new_doc.DocumentNo + "_" + dt.Format("2006-01-02_15_04_05") + ".png"

		_, _ = uploader.Upload(&s3manager.UploadInput{
			Bucket: aws.String(MyBucket),
			ACL:    aws.String("public-read"),
			Key:    aws.String(filename),
			Body:   dec,
		})
	} else {
		filename = ""
	}

	err := CustomerDocumentUpdate(ctx, db,
		data_new_doc.CustomerRequestNo,
		data_new_doc.DocumentType,
		data_new_doc.DocumentNo,
		data_new_doc.DocumentName,
		data_new_doc.DocumentAddress,
		filename,
		data_new_doc.BankCode,
		data_new_doc.CreatedBy,
		data_new_doc.CreatedByIp,
	)

	if err != nil {
		return err
	}

	return nil
}
