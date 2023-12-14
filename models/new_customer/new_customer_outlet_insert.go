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
	execSPCustomerOutletInsert = `exec [sp_smile_customer_outlet_insert]
	$1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18,$19,$20,$21,$22,$23,$24`
)

func DataNewCustomerOutletInsert(ctx context.Context, db *sql.DB,
	data_new_customer_outlet *entities.InfoNewCustomerOutlet,
	log *zap.Logger) (*entities.OutletId, error) {
	sess, _ := session.NewSession(
		&aws.Config{
			Region: aws.String("ap-southeast-1"),
			Credentials: credentials.NewStaticCredentials(
				*aws.String("AKIA3WB2X4HJGJTDWPGL"),
				*aws.String("VVmHtwoOR6xUosj9fi0NW4hRY9i2KVIvMFb7KX+j"),
				"", // a token will be created when the session it's used.
			),
		})

	dec := base64.NewDecoder(base64.StdEncoding, strings.NewReader(data_new_customer_outlet.StoreImg))

	uploader := s3manager.NewUploader(sess)
	MyBucket := "upload.file"

	dt := time.Now()
	filename := "Benfarm/Smile/Customer/OutletImg/" + data_new_customer_outlet.OutletName + "_" + dt.Format("2006-01-02_15_04_05") + ".png"

	_, _ = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(MyBucket),
		ACL:    aws.String("public-read"),
		Key:    aws.String(filename),
		Body:   dec,
	})

	rows, err := db.QueryContext(ctx, execSPCustomerOutletInsert,
		data_new_customer_outlet.OutletCode,
		data_new_customer_outlet.OutletName,
		data_new_customer_outlet.CustomerId,
		data_new_customer_outlet.StoreArea,
		data_new_customer_outlet.ShipSchedule,
		filename,
		data_new_customer_outlet.Address,
		data_new_customer_outlet.Province,
		data_new_customer_outlet.City,
		data_new_customer_outlet.District,
		data_new_customer_outlet.Subdistrict,
		data_new_customer_outlet.Zipcode,
		data_new_customer_outlet.Latitude,
		data_new_customer_outlet.Longitude,
		data_new_customer_outlet.PhoneNo,
		data_new_customer_outlet.CellphoneNo,
		data_new_customer_outlet.Email,
		data_new_customer_outlet.Salesman,
		data_new_customer_outlet.RegionId,
		data_new_customer_outlet.AreaId,
		data_new_customer_outlet.ZoneId,
		data_new_customer_outlet.BillTo,
		data_new_customer_outlet.CreatedBy,
		data_new_customer_outlet.CreatedByIp,
	)

	if err != nil {
		log.Info(
			"Exec DB", zap.Any("Exec Db", err),
		)

		return nil, err
	}
	defer rows.Close()

	outlet_id := &entities.OutletId{}
	for rows.Next() {
		if err := rows.Scan(
			&outlet_id.OutletId,
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

	return outlet_id, nil
}
