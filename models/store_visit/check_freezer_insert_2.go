package store_visit

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"smile-service/entities"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

const (
	execSPDataChecklistFreezer2 = `exec [sp_smile_check_freezer_insert]$1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18,$19,$20,$21,$22`
)

const insertImageDetailFreezer2 = `exec [sp_smile_check_freezer_image_insert]$1,$2,$3,$4,$5`

//Bedanya sama check freezer insert doang itu cara uploadnya, yang check freezer 2 ini pake metode presign URL
func InsertDataCheckFreezer2(ctx context.Context, db *sql.DB,
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
	rows, err := db.QueryContext(ctx, execSPDataChecklistFreezer2,
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

func InsertCheckFreezer2(ctx context.Context, db *sql.DB, checklistFreezer *entities.DataCheckFreezerInsert) (*entities.DataCheckFreezerInsertOutput, error) {

	trans_no, err := InsertDataCheckFreezer2(ctx, db,
		checklistFreezer.CustomerNo, checklistFreezer.BillTo, checklistFreezer.ShipTo, checklistFreezer.FrezerAvailable,
		checklistFreezer.CreatedBy, checklistFreezer.CreatedByIP, checklistFreezer.AvailableNote, checklistFreezer.Mode, checklistFreezer.ArticleDescription, checklistFreezer.SerialNo,
		checklistFreezer.Brand, checklistFreezer.NoteBrand, checklistFreezer.Type, checklistFreezer.Capacity,
		checklistFreezer.Location, checklistFreezer.NoteLocation, checklistFreezer.Status, checklistFreezer.NoteStatus, checklistFreezer.IceThickness,
		checklistFreezer.Temperature, checklistFreezer.FreezerUse, checklistFreezer.NoteFreezerUse,
	)

	if err != nil {
		return nil, err
	}
	// preSignedURLs := make([]*entities.DataCheckFreezerInsertOutput, 0)
	preSignedURLs := &entities.DataCheckFreezerInsertOutput{}
	filePath := "Benfarm/Smile/CheckFreezer/" + trans_no + "/"

	if checklistFreezer.FrezerAvailable == 1 { //If freezer tersedia, upload photo

		//Upload and Insert SerialNo Images
		countSucceed := 0

		presignedURLSerialNos := make([]*entities.DataCheckFreezerInsertOutputURL, 0)

		for _, image_base64 := range checklistFreezer.SerialNoImages {
			countSucceed++
			presignedURLSerialNo := &entities.DataCheckFreezerInsertOutputURL{}

			dt := time.Now()
			fileName := "_SerialNo_" + checklistFreezer.CreatedBy + "_" + dt.Format("2006-01-02_15_04_05") + "_" + image_base64.Base64Image + ".png"

			url, err := getPresignedURL2(filePath + fileName)

			presignedURLSerialNo.FileName = fileName
			presignedURLSerialNo.PreSignURL = url

			presignedURLSerialNos = append(presignedURLSerialNos, presignedURLSerialNo)

			if err != nil {
				return nil, err
			} else {
				_, err = db.QueryContext(ctx, insertImageDetailFreezer,
					trans_no,
					"SerialNo",
					filePath+fileName,
					checklistFreezer.CreatedBy,
					checklistFreezer.CreatedByIP)

			}

		}
		preSignedURLs.SerialNoImages = presignedURLSerialNos

		//Upload and Insert Status Images
		countSucceed = 0
		presignedURLStatuses := make([]*entities.DataCheckFreezerInsertOutputURL, 0)

		for _, image_base64 := range checklistFreezer.StatusImages {
			countSucceed++
			presignedURLStatus := &entities.DataCheckFreezerInsertOutputURL{}

			dt := time.Now()
			fileName := "_Status_" + checklistFreezer.CreatedBy + "_" + dt.Format("2006-01-02_15_04_05") + "_" + image_base64.Base64Image + ".png"

			url, err := getPresignedURL2(filePath + fileName)

			presignedURLStatus.FileName = fileName
			presignedURLStatus.PreSignURL = url

			presignedURLStatuses = append(presignedURLStatuses, presignedURLStatus)

			if err != nil {
				return nil, err
			} else {
				_, err = db.QueryContext(ctx, insertImageDetailFreezer,
					trans_no,
					"Status",
					filePath+fileName,
					checklistFreezer.CreatedBy,
					checklistFreezer.CreatedByIP)

			}

		}
		preSignedURLs.StatusImages = presignedURLStatuses

		//Upload and Insert Location Images
		countSucceed = 0
		presignedURLLocations := make([]*entities.DataCheckFreezerInsertOutputURL, 0)
		for _, image_base64 := range checklistFreezer.LocationImages {
			countSucceed++
			presignedURLLocation := &entities.DataCheckFreezerInsertOutputURL{}

			dt := time.Now()
			fileName := "_Location_" + checklistFreezer.CreatedBy + "_" + dt.Format("2006-01-02_15_04_05") + "_" + image_base64.Base64Image + ".png"

			url, err := getPresignedURL2(filePath + fileName)

			presignedURLLocation.FileName = fileName
			presignedURLLocation.PreSignURL = url

			presignedURLLocations = append(presignedURLLocations, presignedURLLocation)

			if err != nil {
				return nil, err
			} else {
				_, err = db.QueryContext(ctx, insertImageDetailFreezer,
					trans_no,
					"Location",
					filePath+fileName,
					checklistFreezer.CreatedBy,
					checklistFreezer.CreatedByIP)

			}
		}
		preSignedURLs.LocationImages = presignedURLLocations

		//Upload and Insert Temperature Images
		countSucceed = 0
		presignedURLTemps := make([]*entities.DataCheckFreezerInsertOutputURL, 0)
		for _, image_base64 := range checklistFreezer.TemperatureImages {
			countSucceed++
			presignedURLTemp := &entities.DataCheckFreezerInsertOutputURL{}

			dt := time.Now()
			fileName := "_Temperature_" + checklistFreezer.CreatedBy + "_" + dt.Format("2006-01-02_15_04_05") + "_" + image_base64.Base64Image + ".png"

			url, err := getPresignedURL2(filePath + fileName)

			presignedURLTemp.FileName = fileName
			presignedURLTemp.PreSignURL = url

			presignedURLTemps = append(presignedURLTemps, presignedURLTemp)

			if err != nil {
				return nil, err
			} else {
				_, err = db.QueryContext(ctx, insertImageDetailFreezer,
					trans_no,
					"Temperature",
					filePath+fileName,
					checklistFreezer.CreatedBy,
					checklistFreezer.CreatedByIP)

			}
		}
		preSignedURLs.TemperatureImages = presignedURLTemps

		//Upload and Insert Usage Images
		countSucceed = 0
		presignedURLUsages := make([]*entities.DataCheckFreezerInsertOutputURL, 0)
		for _, image_base64 := range checklistFreezer.UsageImages {
			countSucceed++
			presignedURLUsage := &entities.DataCheckFreezerInsertOutputURL{}

			dt := time.Now()
			fileName := "_Usage_" + checklistFreezer.CreatedBy + "_" + dt.Format("2006-01-02_15_04_05") + "_" + image_base64.Base64Image + ".png"

			url, err := getPresignedURL2(filePath + fileName)

			presignedURLUsage.FileName = fileName
			presignedURLUsage.PreSignURL = url

			presignedURLUsages = append(presignedURLUsages, presignedURLUsage)

			if err != nil {
				return nil, err
			} else {
				_, err = db.QueryContext(ctx, insertImageDetailFreezer,
					trans_no,
					"Usage",
					filePath+fileName,
					checklistFreezer.CreatedBy,
					checklistFreezer.CreatedByIP)

			}
		}
		preSignedURLs.UsageImages = presignedURLUsages

	}
	return preSignedURLs, nil
}

func getPresignedURL2(filename string) (string, error) {

	// Load env vars
	// err := godotenv.Load(".env")
	// if err != nil {
	// 	log.Fatalf("Error loading .env file: %v", err)
	// }

	// Load the bucket name
	s3Bucket := os.Getenv("S3_BUCKET")
	if s3Bucket == "" {
		log.Fatal("an s3 bucket was unable to be loaded from env vars")
	}

	// Prepare the S3 request so a signature can be generated
	// svc := s3.New(session.NewSession(&aws.Config{
	// 	Region: aws.String("ap-southeast-1"),
	// 	Credentials: credentials.NewStaticCredentials(
	// 		*aws.String("AKIA3WB2X4HJGJTDWPGL"),
	// 		*aws.String("VVmHtwoOR6xUosj9fi0NW4hRY9i2KVIvMFb7KX+j"),
	// 		"", // a token will be created when the session it's used.
	// 	),
	// }))

	// snippet-start:[s3.go.generate_presigned_url.session]
	sess, err := session.NewSession(
		&aws.Config{
			Region: aws.String("ap-southeast-1"),
			Credentials: credentials.NewStaticCredentials(
				*aws.String("AKIA3WB2X4HJGJTDWPGL"),
				*aws.String("VVmHtwoOR6xUosj9fi0NW4hRY9i2KVIvMFb7KX+j"),
				"", // a token will be created when the session it's used.
			),
		})
	// snippet-end:[s3.go.generate_presigned_url.session]

	// snippet-start:[s3.go.generate_presigned_url.call]
	svc := s3.New(sess)

	r, _ := svc.PutObjectRequest(&s3.PutObjectInput{
		Bucket: aws.String(s3Bucket),
		Key:    aws.String(filename),
		ACL:    aws.String("public-read"),
	})

	// Create the pre-signed url with an expiry
	url, err := r.Presign(15 * time.Minute)
	if err != nil {
		fmt.Println("Failed to generate a pre-signed url: ", err)
		return "", err
	}

	// Display the pre-signed url
	fmt.Println("Pre-signed URL", url)
	return url, nil
}
