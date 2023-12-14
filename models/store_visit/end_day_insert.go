package store_visit

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
	execSPEndDayInsert = `exec [sp_smile_start_day_insert]$1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11`
)

func DataEndDayInsert(ctx context.Context, db *sql.DB,
	Pic string,
	VisitDate string,
	Kilometer int,
	PicKilometer string,
	Vehicle string,
	PicVehiclePlate string,
	Address string,
	AddressLatitude string,
	AddressLongitude string,
	CreatedBy string,
	CreatedByIp string) error {
	rows, err := db.QueryContext(ctx, execSPEndDayInsert,
		Pic,
		VisitDate,
		Kilometer,
		PicKilometer,
		Vehicle,
		PicVehiclePlate,
		Address,
		AddressLatitude,
		AddressLongitude,
		CreatedBy,
		CreatedByIp,
	)
	if err != nil {
		return err
	}
	defer rows.Close()

	return nil
}

func InsertEndDay(ctx context.Context, db *sql.DB, endDay *entities.DataEndDay) error {

	sess, err := session.NewSession(
		&aws.Config{
			Region: aws.String("ap-southeast-1"),
			Credentials: credentials.NewStaticCredentials(
				*aws.String("AKIA3WB2X4HJGJTDWPGL"),
				*aws.String("VVmHtwoOR6xUosj9fi0NW4hRY9i2KVIvMFb7KX+j"),
				"", // a token will be created when the session it's used.
			),
		})

	decKM := base64.NewDecoder(base64.StdEncoding, strings.NewReader(endDay.PicKilometer))       //Convert Base64 To Image
	decPlate := base64.NewDecoder(base64.StdEncoding, strings.NewReader(endDay.PicVehiclePlate)) //Convert Base64 To Image

	uploader := s3manager.NewUploader(sess)
	MyBucket := "upload.file"

	dt := time.Now()
	filenameKm := "Benfarm/Smile/EndDay/" + endDay.Pic + "/" + dt.Format("2006-01-02") + "/KmPhoto_" + endDay.Pic + "_" + dt.Format("2006-01-02_15_04_05") + ".png"

	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(MyBucket),
		ACL:    aws.String("public-read"),
		Key:    aws.String(filenameKm),
		Body:   decKM,
	})

	filenamePlate := "Benfarm/Smile/EndDay/" + endDay.Pic + "/" + dt.Format("2006-01-02") + "/PlatePhoto_" + endDay.Pic + "_" + dt.Format("2006-01-02_15_04_05") + ".png"

	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(MyBucket),
		ACL:    aws.String("public-read"),
		Key:    aws.String(filenamePlate),
		Body:   decPlate,
	})

	err2 := DataEndDayInsert(ctx, db,
		endDay.Pic,
		endDay.VisitDate,
		endDay.Kilometer,
		filenameKm,
		endDay.Vehicle,
		filenamePlate,
		endDay.Address,
		endDay.AddressLatitude,
		endDay.AddressLongitude,
		endDay.CreatedBy,
		endDay.CreatedByIp,
	)

	if err != nil {
		return err
	} else if err2 != nil {
		return err2
	}

	return nil
}
