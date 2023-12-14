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
	execSPDataChecklistFreezer = `exec [sp_smile_check_freezer_insert]$1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18,$19,$20,$21,$22`
)

const insertImageDetailFreezer = `exec [sp_smile_check_freezer_image_insert]$1,$2,$3,$4,$5`

//Bedanya sama check freezer insert 2 itu cara uploadnya, yang check freezer 2 pake metode presign URL
func InsertDataCheckFreezer(ctx context.Context, db *sql.DB,
	CustomerNo string,
	BillTo string,
	ShipTo string,
	FrezerAvailable int,
	CreatedBy string,
	CreatedByIP string,
	AvailableNote string,
	Mode string,
	ArticleDescription string,
	SerialNo string,
	Brand int,
	NoteBrand string,
	Type string,
	Capacity int,
	Location int,
	NoteLocation string,
	Status int,
	NoteStatus string,
	IceThickness int,
	Temperature int,
	FreezerUse string,
	NoteFreezerUse string) (string, error) {
	rows, err := db.QueryContext(ctx, execSPDataChecklistFreezer,
		CustomerNo,
		BillTo,
		ShipTo,
		FrezerAvailable,
		CreatedBy,
		CreatedByIP,
		AvailableNote,
		Mode,
		ArticleDescription,
		SerialNo,
		Brand,
		NoteBrand,
		Type,
		Capacity,
		Location,
		NoteLocation,
		Status,
		NoteStatus,
		IceThickness,
		Temperature,
		FreezerUse,
		NoteFreezerUse,
	)
	if err != nil {
		return "", err
	}
	defer rows.Close()

	outputs := &entities.DataCheckFreezerOutputInsert{}

	for rows.Next() {
		if err := rows.Scan(
			&outputs.TransNo,
		); err != nil {
			return "", err
		}
	}

	return outputs.TransNo, nil
}

func InsertCheckFreezer(ctx context.Context, db *sql.DB, checklistFreezer *entities.DataCheckFreezerInsert) error {

	trans_no, err := InsertDataCheckFreezer(ctx, db,
		checklistFreezer.CustomerNo, checklistFreezer.BillTo, checklistFreezer.ShipTo, checklistFreezer.FrezerAvailable,
		checklistFreezer.CreatedBy, checklistFreezer.CreatedByIP, checklistFreezer.AvailableNote, checklistFreezer.Mode, checklistFreezer.ArticleDescription, checklistFreezer.SerialNo,
		checklistFreezer.Brand, checklistFreezer.NoteBrand, checklistFreezer.Type, checklistFreezer.Capacity,
		checklistFreezer.Location, checklistFreezer.NoteLocation, checklistFreezer.Status, checklistFreezer.NoteStatus, checklistFreezer.IceThickness,
		checklistFreezer.Temperature, checklistFreezer.FreezerUse, checklistFreezer.NoteFreezerUse,
	)

	if err != nil {
		return err
	}

	if checklistFreezer.FrezerAvailable == 1 { //If freezer tersedia, upload photo

		//Upload and Insert SerialNo Images
		countSucceed := 0
		for _, image_base64 := range checklistFreezer.SerialNoImages {
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
				return err
			} else {

				imageCheckInOut := base64.NewDecoder(base64.StdEncoding, strings.NewReader(image_base64.Base64Image)) //Convert Base64 To Image

				dt := time.Now()
				fileName := "Benfarm/Smile/CheckFreezer/" + trans_no + "/" + "_SerialNo_" + checklistFreezer.CreatedBy + "_" + dt.Format("2006-01-02_15_04_05") + "_" + strconv.Itoa(countSucceed) + ".png"

				_, err = uploader.Upload(&s3manager.UploadInput{
					Bucket: aws.String(MyBucket),
					ACL:    aws.String("public-read"),
					Key:    aws.String(fileName),
					Body:   imageCheckInOut,
				})

				_, err = db.QueryContext(ctx, insertImageDetailFreezer,
					trans_no,
					"SerialNo",
					fileName,
					checklistFreezer.CreatedBy,
					checklistFreezer.CreatedByIP)

			}
		}

		//Upload and Insert Status Images
		countSucceed = 0
		for _, image_base64 := range checklistFreezer.StatusImages {
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
				return err
			} else {

				imageCheckInOut := base64.NewDecoder(base64.StdEncoding, strings.NewReader(image_base64.Base64Image)) //Convert Base64 To Image

				dt := time.Now()
				fileName := "Benfarm/Smile/CheckFreezer/" + trans_no + "/" + "_Status_" + checklistFreezer.CreatedBy + "_" + dt.Format("2006-01-02_15_04_05") + "_" + strconv.Itoa(countSucceed) + ".png"

				_, err = uploader.Upload(&s3manager.UploadInput{
					Bucket: aws.String(MyBucket),
					ACL:    aws.String("public-read"),
					Key:    aws.String(fileName),
					Body:   imageCheckInOut,
				})

				_, err = db.QueryContext(ctx, insertImageDetailFreezer,
					trans_no,
					"Status",
					fileName,
					checklistFreezer.CreatedBy,
					checklistFreezer.CreatedByIP)

			}
		}

		//Upload and Insert Location Images
		countSucceed = 0
		for _, image_base64 := range checklistFreezer.LocationImages {
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
				return err
			} else {

				imageCheckInOut := base64.NewDecoder(base64.StdEncoding, strings.NewReader(image_base64.Base64Image)) //Convert Base64 To Image

				dt := time.Now()
				fileName := "Benfarm/Smile/CheckFreezer/" + trans_no + "/" + "_Location_" + checklistFreezer.CreatedBy + "_" + dt.Format("2006-01-02_15_04_05") + "_" + strconv.Itoa(countSucceed) + ".png"

				_, err = uploader.Upload(&s3manager.UploadInput{
					Bucket: aws.String(MyBucket),
					ACL:    aws.String("public-read"),
					Key:    aws.String(fileName),
					Body:   imageCheckInOut,
				})

				_, err = db.QueryContext(ctx, insertImageDetailFreezer,
					trans_no,
					"Location",
					fileName,
					checklistFreezer.CreatedBy,
					checklistFreezer.CreatedByIP)

			}
		}

		//Upload and Insert Temperature Images
		countSucceed = 0
		for _, image_base64 := range checklistFreezer.TemperatureImages {
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
				return err
			} else {

				imageCheckInOut := base64.NewDecoder(base64.StdEncoding, strings.NewReader(image_base64.Base64Image)) //Convert Base64 To Image

				dt := time.Now()
				fileName := "Benfarm/Smile/CheckFreezer/" + trans_no + "/" + "_Temperature_" + checklistFreezer.CreatedBy + "_" + dt.Format("2006-01-02_15_04_05") + "_" + strconv.Itoa(countSucceed) + ".png"

				_, err = uploader.Upload(&s3manager.UploadInput{
					Bucket: aws.String(MyBucket),
					ACL:    aws.String("public-read"),
					Key:    aws.String(fileName),
					Body:   imageCheckInOut,
				})

				_, err = db.QueryContext(ctx, insertImageDetailFreezer,
					trans_no,
					"Temperature",
					fileName,
					checklistFreezer.CreatedBy,
					checklistFreezer.CreatedByIP)

			}
		}

		//Upload and Insert Usage Images
		countSucceed = 0
		for _, image_base64 := range checklistFreezer.UsageImages {
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
				return err
			} else {

				imageCheckInOut := base64.NewDecoder(base64.StdEncoding, strings.NewReader(image_base64.Base64Image)) //Convert Base64 To Image

				dt := time.Now()
				fileName := "Benfarm/Smile/CheckFreezer/" + trans_no + "/" + "_Usage_" + checklistFreezer.CreatedBy + "_" + dt.Format("2006-01-02_15_04_05") + "_" + strconv.Itoa(countSucceed) + ".png"

				_, err = uploader.Upload(&s3manager.UploadInput{
					Bucket: aws.String(MyBucket),
					ACL:    aws.String("public-read"),
					Key:    aws.String(fileName),
					Body:   imageCheckInOut,
				})

				_, err = db.QueryContext(ctx, insertImageDetailFreezer,
					trans_no,
					"Usage",
					fileName,
					checklistFreezer.CreatedBy,
					checklistFreezer.CreatedByIP)

			}
		}
	}
	return nil
}
