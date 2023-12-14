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
	execSPNewCustomerRequestInsert = `exec [sp_smile_new_customer_request_insert]
		$1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18,$19,$20,
		$21,$22,$23,$24,$25,$26,$27,$28,$29,$30,$31,$32,$33,$34,$35,$36`
)

func DataNewCustomerRequestInsert(ctx context.Context, db *sql.DB,
	data_new_customer *entities.InfoNewCustomerRequest,
	log *zap.Logger) (*entities.CustomerRequestNo, error) {

	sess, _ := session.NewSession(
		&aws.Config{
			Region: aws.String("ap-southeast-1"),
			Credentials: credentials.NewStaticCredentials(
				*aws.String("AKIA3WB2X4HJGJTDWPGL"),
				*aws.String("VVmHtwoOR6xUosj9fi0NW4hRY9i2KVIvMFb7KX+j"),
				"", // a token will be created when the session it's used.
			),
		})

	dec := base64.NewDecoder(base64.StdEncoding, strings.NewReader(data_new_customer.VatImg))

	uploader := s3manager.NewUploader(sess)
	MyBucket := "upload.file"

	dt := time.Now()
	filename := "Benfarm/Smile/Customer/VatImg/" + data_new_customer.NpwpName + "_" + dt.Format("2006-01-02_15_04_05") + ".png"

	_, _ = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(MyBucket),
		ACL:    aws.String("public-read"),
		Key:    aws.String(filename),
		Body:   dec,
	})

	rows, err := db.QueryContext(ctx, execSPNewCustomerRequestInsert,
		data_new_customer.CustomerName,
		data_new_customer.CustomerType,
		data_new_customer.YearEstablished,
		data_new_customer.EmployeeAmount,
		data_new_customer.AnnualSales,
		data_new_customer.PhoneNo,
		data_new_customer.Fax,
		data_new_customer.Email,
		data_new_customer.Website,
		data_new_customer.StreetAddress,
		data_new_customer.ProvinceId,
		data_new_customer.CityId,
		data_new_customer.DistrictId,
		data_new_customer.SubdistrictId,
		data_new_customer.Postal,
		data_new_customer.TaxMandatory,
		data_new_customer.TaxStatus,
		data_new_customer.NpwpNo,
		data_new_customer.NpwpName,
		data_new_customer.NpwpAddress,
		filename,
		data_new_customer.CreateBy,
		data_new_customer.CreateByIp,
		data_new_customer.BillingCode,
		data_new_customer.BillingName,
		data_new_customer.BillingMethod,
		data_new_customer.FactureSchedule,
		data_new_customer.BillingDoc,
		data_new_customer.ReturAvailable,
		data_new_customer.BillingAddress,
		data_new_customer.BillingProvince,
		data_new_customer.BillingCity,
		data_new_customer.BillingDistrict,
		data_new_customer.BillingSubdistrict,
		data_new_customer.BillingPostal,
		data_new_customer.PaymentTerm,
	)

	if err != nil {
		log.Info(
			"Exec DB", zap.Any("Exec Db", err),
		)

		return nil, err
	}
	defer rows.Close()

	customer_request_no := &entities.CustomerRequestNo{}
	for rows.Next() {
		if err := rows.Scan(
			&customer_request_no.CustomerRequestNo,
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

	return customer_request_no, nil
}
