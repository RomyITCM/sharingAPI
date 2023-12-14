package cart

import (
	"context"
	"database/sql"
	"smile-service/entities"
	"smile-service/shared"
	"time"

	"go.uber.org/zap"
)

const execSPDataCartProcess = `exec [sp_smile_cart_process]$1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12`
const execSPDataCartProcessEdit = `exec [sp_smile_cart_process_edit]$1,$2,$3,$4,$5,$6,$7,$8`

func CartProcess(
	ctx context.Context,
	db *sql.DB,
	data_transaction *entities.CartProcess,
	log *zap.Logger) (*entities.CartResult, error) {

	fileName := ""

	if data_transaction.Attachment == "PDF" {
		fileName = data_transaction.CustPoNo + "_" + data_transaction.ShipTo + ".pdf"

	} else {

		fileName = data_transaction.CustPoNo + "_" + data_transaction.ShipTo + ".png"
	}

	rows, err := db.QueryContext(
		ctx,
		execSPDataCartProcess,
		data_transaction.WHID,
		data_transaction.CustomerNo,
		data_transaction.BillTo,
		data_transaction.ShipTo,
		data_transaction.CustPoNo,
		data_transaction.CustPoDate,
		data_transaction.ExpPoDate,
		data_transaction.DelvDate,
		data_transaction.PaymentTerm,
		fileName,
		data_transaction.CreatedBy,
		data_transaction.CreatedByIp,
	)

	if err != nil {
		log.Info(
			"Exec DB", zap.Any("Exec Db", err),
		)

		return nil, err
	}
	defer rows.Close()

	cart_result := &entities.CartResult{}
	for rows.Next() {
		if err := rows.Scan(
			&cart_result.TransNo,
			&cart_result.MsgType,
			&cart_result.MsgError,
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

	if cart_result.TransNo != "" {

		path := "BEN/SO/" + cart_result.TransNo + "/"

		filename := path + fileName

		pre_signed_url, err := shared.GetPresignedURL(filename)
		if err != nil {
			return nil, err
		}

		cart_result.PreSignedURL = pre_signed_url
		cart_result.FileName = filename
	}

	// sess, err := session.NewSession(
	// 	&aws.Config{
	// 		Region: aws.String("ap-southeast-1"),
	// 		Credentials: credentials.NewStaticCredentials(
	// 			*aws.String("AKIA3WB2X4HJGJTDWPGL"),
	// 			*aws.String("VVmHtwoOR6xUosj9fi0NW4hRY9i2KVIvMFb7KX+j"),
	// 			"", // a token will be created when the session it's used.
	// 		),
	// 	})

	// if err != nil {
	// 	return nil, err
	// }

	// dec := base64.NewDecoder(base64.StdEncoding, strings.NewReader(data_transaction.Attachment)) //Convert Base64 To Image

	// uploader := s3manager.NewUploader(sess)
	// MyBucket := "upload.file"
	// path := "BEN/SO/" + cart_result.TransNo + "/"

	// filename := path + fileName

	// _, err = uploader.Upload(&s3manager.UploadInput{
	// 	Bucket: aws.String(MyBucket),
	// 	ACL:    aws.String("public-read"),
	// 	Key:    aws.String(filename),
	// 	Body:   dec,
	// })

	// if err != nil {
	// 	log.Info(
	// 		"Upload to S3", zap.Any("Upload to S3", err),
	// 	)

	// 	return nil, err
	// }

	return cart_result, nil
}

func CartProcessEdit(
	ctx context.Context,
	db *sql.DB,
	data_transaction *entities.CartProcessEdit,
	log *zap.Logger) (*entities.CartResult, error) {

	fileName := ""
	currentTime := time.Now()

	if data_transaction.EditFile {

		if data_transaction.Attachment == ".pdf" {
			fileName = data_transaction.CustPoNo + "_" + currentTime.Format("20060102") + ".pdf"

		} else {

			fileName = data_transaction.CustPoNo + "_" + currentTime.Format("20060102") + ".png"
		}
	}

	rows, err := db.QueryContext(
		ctx,
		execSPDataCartProcessEdit,
		data_transaction.TransNo,
		data_transaction.CustPoNo,
		data_transaction.DelvDate,
		data_transaction.ExpPoDate,
		data_transaction.CreatedBy,
		data_transaction.CreatedByIp,
		data_transaction.EditFile,
		fileName,
	)

	if err != nil {
		log.Info(
			"Exec DB", zap.Any("Exec Db", err),
		)

		return nil, err
	}
	defer rows.Close()

	cart_result := &entities.CartResult{}
	for rows.Next() {
		if err := rows.Scan(
			&cart_result.TransNo,
			&cart_result.MsgType,
			&cart_result.MsgError,
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

	if cart_result.TransNo != "" {
		if data_transaction.EditFile {
			path := "BEN/SO/" + cart_result.TransNo + "/"

			filename := path + fileName

			pre_signed_url, err := shared.GetPresignedURL(filename)
			if err != nil {
				return nil, err
			}

			cart_result.PreSignedURL = pre_signed_url
			cart_result.FileName = filename
		}
	}

	return cart_result, nil
}
