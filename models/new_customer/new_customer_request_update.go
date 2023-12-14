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
	execSPNewCustomerRequestUpdate = `exec [sp_smile_new_customer_request_update]
		$1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18,$19,$20,
		$21,$22,$23,$24,$25,$26,$27,$28,$29,$30,$31,$32,$33,$34,$35,$36,$37,$38`
)

func UpdateCustomerRequest(ctx context.Context, db *sql.DB,
	CustomerRequestNo string,
	CustomerName string,
	CustomerType string,
	YearEstablished string,
	EmployeeAmount string,
	AnnualSales string,
	PhoneNo string,
	Fax string,
	Email string,
	Website string,
	StreetAddress string,
	ProvinceId string,
	CityId string,
	DistrictId string,
	SubdistrictId string,
	Postal string,
	TaxMandatory string,
	TaxStatus string,
	NpwpNo string,
	NpwpName string,
	NpwpAddress string,
	VatImg string,
	CreateBy string,
	CreateByIp string,
	BillingCode string,
	BillingName string,
	BillingMethod string,
	FactureSchedule string,
	BillingDoc string,
	ReturAvailable string,
	BillingAddress string,
	BillingProvince string,
	BillingCity string,
	BillingDistrict string,
	BillingSubdistrict string,
	BillingPostal string,
	PaymentTerm string,
	Status string) error {
	rows, err := db.QueryContext(ctx, execSPNewCustomerRequestUpdate,
		CustomerRequestNo,
		CustomerName,
		CustomerType,
		YearEstablished,
		EmployeeAmount,
		AnnualSales,
		PhoneNo,
		Fax,
		Email,
		Website,
		StreetAddress,
		ProvinceId,
		CityId,
		DistrictId,
		SubdistrictId,
		Postal,
		TaxMandatory,
		TaxStatus,
		NpwpNo,
		NpwpName,
		NpwpAddress,
		VatImg,
		CreateBy,
		CreateByIp,
		BillingCode,
		BillingName,
		BillingMethod,
		FactureSchedule,
		BillingDoc,
		ReturAvailable,
		BillingAddress,
		BillingProvince,
		BillingCity,
		BillingDistrict,
		BillingSubdistrict,
		BillingPostal,
		PaymentTerm,
		Status,
	)
	if err != nil {
		return err
	}
	defer rows.Close()

	return nil
}

func CustomerReqestUpdate(ctx context.Context, db *sql.DB,
	dataCustomerRequest *entities.InfoNewCustomerRequestUpdate) error {
	filename := ""

	if dataCustomerRequest.VatImg == "" {
		filename = ""

	} else {
		sess, _ := session.NewSession(
			&aws.Config{
				Region: aws.String("ap-southeast-1"),
				Credentials: credentials.NewStaticCredentials(
					*aws.String("AKIA3WB2X4HJGJTDWPGL"),
					*aws.String("VVmHtwoOR6xUosj9fi0NW4hRY9i2KVIvMFb7KX+j"),
					"", // a token will be created when the session it's used.
				),
			})

		dec := base64.NewDecoder(base64.StdEncoding, strings.NewReader(dataCustomerRequest.VatImg))

		uploader := s3manager.NewUploader(sess)
		MyBucket := "upload.file"

		dt := time.Now()
		filename = "Benfarm/Smile/Customer/VatImg/" + dataCustomerRequest.NpwpName + "_" + dt.Format("2006-01-02_15_04_05") + ".png"

		_, _ = uploader.Upload(&s3manager.UploadInput{
			Bucket: aws.String(MyBucket),
			ACL:    aws.String("public-read"),
			Key:    aws.String(filename),
			Body:   dec,
		})
	}

	err := UpdateCustomerRequest(ctx, db,
		dataCustomerRequest.CustomerRequestNo,
		dataCustomerRequest.CustomerName,
		dataCustomerRequest.CustomerType,
		dataCustomerRequest.YearEstablished,
		dataCustomerRequest.EmployeeAmount,
		dataCustomerRequest.AnnualSales,
		dataCustomerRequest.PhoneNo,
		dataCustomerRequest.Fax,
		dataCustomerRequest.Email,
		dataCustomerRequest.Website,
		dataCustomerRequest.StreetAddress,
		dataCustomerRequest.ProvinceId,
		dataCustomerRequest.CityId,
		dataCustomerRequest.DistrictId,
		dataCustomerRequest.SubdistrictId,
		dataCustomerRequest.Postal,
		dataCustomerRequest.TaxMandatory,
		dataCustomerRequest.TaxStatus,
		dataCustomerRequest.NpwpNo,
		dataCustomerRequest.NpwpName,
		dataCustomerRequest.NpwpAddress,
		filename,
		dataCustomerRequest.CreateBy,
		dataCustomerRequest.CreateByIp,
		dataCustomerRequest.BillingCode,
		dataCustomerRequest.BillingName,
		dataCustomerRequest.BillingMethod,
		dataCustomerRequest.FactureSchedule,
		dataCustomerRequest.BillingDoc,
		dataCustomerRequest.ReturAvailable,
		dataCustomerRequest.BillingAddress,
		dataCustomerRequest.BillingProvince,
		dataCustomerRequest.BillingCity,
		dataCustomerRequest.BillingDistrict,
		dataCustomerRequest.BillingSubdistrict,
		dataCustomerRequest.BillingPostal,
		dataCustomerRequest.PaymentTerm,
		dataCustomerRequest.Status,
	)

	if err != nil {
		return err
	}

	return nil
}
