package store_visit

import (
	"context"
	"database/sql"
	"encoding/base64"
	"smile-service/entities"
	"strconv"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

const (
	execSPCheckInInsert = `exec [sp_smile_check_in_out_insert]$1,$2,$3,$4,$5,$6,$7,$8,$9,$10`
)

const insertImageDetail = `exec [sp_smile_check_in_out_image_insert]$1,$2,$3,$4,$5,$6`

func DataCheckInOutInsert(ctx context.Context, db *sql.DB,
	VisitID string,
	Type string,
	CheckTime string,
	CustomerNo string,
	ShipTo string,
	Address string,
	AddressLatitude string,
	AddressLongitude string,
	CreatedBy string,
	CreatedByIp string,

) (*entities.DataCheckInOutID, error) {
	rows, err := db.QueryContext(ctx, execSPCheckInInsert,
		VisitID,
		Type,
		CheckTime,
		CustomerNo,
		ShipTo,
		Address,
		AddressLatitude,
		AddressLongitude,
		CreatedBy,
		CreatedByIp,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	detail_ids := &entities.DataCheckInOutID{}

	for rows.Next() {
		if err := rows.Scan(
			&detail_ids.DetailID,
		); err != nil {
			return nil, err
		}
	}

	return detail_ids, nil
}

func InsertCheckInOut(ctx context.Context, db *sql.DB, checkInOut *entities.DataCheckInOut) (*entities.DataCheckInOutID, error) {

	detail_id, err2 := DataCheckInOutInsert(ctx, db,
		checkInOut.VisitID,
		checkInOut.Type,
		checkInOut.CheckTime,
		checkInOut.CustomerNo,
		checkInOut.ShipTo,
		checkInOut.Address,
		checkInOut.AddressLatitude,
		checkInOut.AddressLongitude,
		checkInOut.CreatedBy,
		checkInOut.CreatedByIp,
	)

	if err2 != nil {
		return nil, err2
	}

	countSucceed := 0
	for _, image_base64 := range checkInOut.Images {
		countSucceed++
		sess, err := session.NewSession(
			&aws.Config{
				Region: aws.String("ap-southeast-1"),
				Credentials: credentials.NewStaticCredentials(
					*aws.String("AKIA3WB2X4HJGJTDWPGL"),
					*aws.String("VVmHtwoOR6xUosj9fi0NW4hRY9i2KVIvMFb7KX+j"),
					"", // a token will be created when the session it's used.
				),
			})

		uploader := s3manager.NewUploader(sess)
		MyBucket := "upload.file"
		if err != nil {
			return nil, err
		} else {

			imageCheckInOut := base64.NewDecoder(base64.StdEncoding, strings.NewReader(image_base64.Base64Image)) //Convert Base64 To Image

			dt := time.Now()
			fileName := "Benfarm/Smile/CheckInOut/" + checkInOut.CreatedBy + "/" + dt.Format("2006-01-02") + "/Check" + checkInOut.Type + "_" + checkInOut.CreatedBy + "_" + dt.Format("2006-01-02_15_04_05") + "_" + strconv.Itoa(countSucceed) + ".png"

			_, err = uploader.Upload(&s3manager.UploadInput{
				Bucket: aws.String(MyBucket),
				ACL:    aws.String("public-read"),
				Key:    aws.String(fileName),
				Body:   imageCheckInOut,
			})

			_, err = db.QueryContext(ctx, insertImageDetail,
				detail_id.DetailID,
				checkInOut.VisitID,
				checkInOut.Type,
				fileName,
				checkInOut.CreatedBy,
				checkInOut.CreatedByIp)

		}
	}
	return detail_id, nil

}
