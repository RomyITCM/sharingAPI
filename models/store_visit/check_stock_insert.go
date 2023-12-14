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
	execSPDataChecklistStock = `exec [sp_smile_check_stock_insert]$1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14`
)

const insertImageDetailStock = `exec [sp_smile_check_stock_image_insert]$1,$2,$3,$4,$5,$6,$7`

const insertDetailProduct = `exec [sp_smile_check_stock_product_insert]$1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11`

func InsertDataCheckStock(ctx context.Context, db *sql.DB,
	CustomerNo string,
	BillTo string,
	ShipTo string,
	StockAvailable int,
	CreatedBy string,
	CreatedByIP string,
	ArticleNo string,
	Type string,
	Price string,
	PromoAvailable int,
	MediaPromoAvailable int,
	NotePromo string,
	StartPromoDate string,
	EndPromoDate string,
) (string, string, error) {
	rows, err := db.QueryContext(ctx, execSPDataChecklistStock,
		CustomerNo,
		BillTo,
		ShipTo,
		StockAvailable,
		CreatedBy,
		CreatedByIP,
		ArticleNo,
		Type,
		Price,
		PromoAvailable,
		MediaPromoAvailable,
		NotePromo,
		StartPromoDate,
		EndPromoDate,
	)
	if err != nil {
		return "", "", err
	}
	defer rows.Close()

	outputs := &entities.DataCheckStockOutputInsert{}

	for rows.Next() {
		if err := rows.Scan(
			&outputs.TransNo,
			&outputs.ReffNo,
		); err != nil {
			return "", "", err
		}
	}

	return outputs.TransNo, outputs.ReffNo, nil
}

func InsertDataProductCheckStock(ctx context.Context, db *sql.DB,
	ReffNo string,
	TransNo string,
	ArticleNo string,
	ExpiredDate string,
	QtyStock string,
	Condition string,
	NoteStock string,
	SalesReturnNo string,
	NoteSalesReturn string,
	CreatedBy string,
	CreatedByIP string) (string, error) {
	rows, err := db.QueryContext(ctx, insertDetailProduct,
		ReffNo,
		TransNo,
		ArticleNo,
		ExpiredDate,
		QtyStock,
		Condition,
		NoteStock,
		SalesReturnNo,
		NoteSalesReturn,
		CreatedBy,
		CreatedByIP,
	)
	if err != nil {
		return "", err
	}
	defer rows.Close()
	outputs := &entities.DataCheckStockDetailOutputInsert{}

	for rows.Next() {
		if err := rows.Scan(
			&outputs.DetailID,
		); err != nil {
			return "", err
		}
	}

	return outputs.DetailID, nil
}

func InsertCheckStock(ctx context.Context, db *sql.DB, checklistStock *entities.DataCheckStockInsert) error {

	trans_no, reff_no, err := InsertDataCheckStock(ctx, db,
		checklistStock.CustomerNo, checklistStock.BillTo, checklistStock.ShipTo, checklistStock.StockAvailable,
		checklistStock.CreatedBy, checklistStock.CreatedByIP, checklistStock.ArticleNo,
		checklistStock.Type, checklistStock.Price, checklistStock.PromoAvailable, checklistStock.MediaPromoAvailable,
		checklistStock.NotePromo, checklistStock.StartDatePromo, checklistStock.EndDatePromo,
	)

	if err != nil {
		return err
	}

	//Upload and Insert Price Images
	countSucceed := 0
	for _, image_base64 := range checklistStock.PriceImages {
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
			fileName := "Benfarm/Smile/CheckStock/" + checklistStock.CreatedBy + "/" + dt.Format("2006-01-02") + "/" + trans_no + "/" + checklistStock.Type + "_Price_" + checklistStock.CreatedBy + "_" + dt.Format("2006-01-02_15_04_05") + "_" + strconv.Itoa(countSucceed) + ".png"

			_, err = uploader.Upload(&s3manager.UploadInput{
				Bucket: aws.String(MyBucket),
				ACL:    aws.String("public-read"),
				Key:    aws.String(fileName),
				Body:   imageCheckInOut,
			})

			_, err = db.QueryContext(ctx, insertImageDetailStock,
				trans_no,
				reff_no,
				"0",
				"Price",
				fileName,
				checklistStock.CreatedBy,
				checklistStock.CreatedByIP)

		}
	}

	if checklistStock.PromoAvailable == 1 && checklistStock.MediaPromoAvailable == 1 { //If ada promo yang aktif, upload photo
		//Upload and Insert Promo Images
		countSucceed = 0
		for _, image_base64 := range checklistStock.PromoImages {
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
				fileName := "Benfarm/Smile/CheckStock/" + checklistStock.CreatedBy + "/" + dt.Format("2006-01-02") + "/" + trans_no + "/" + checklistStock.Type + "_MediaPromo_" + checklistStock.CreatedBy + "_" + dt.Format("2006-01-02_15_04_05") + "_" + strconv.Itoa(countSucceed) + ".png"

				_, err = uploader.Upload(&s3manager.UploadInput{
					Bucket: aws.String(MyBucket),
					ACL:    aws.String("public-read"),
					Key:    aws.String(fileName),
					Body:   imageCheckInOut,
				})

				_, err = db.QueryContext(ctx, insertImageDetailStock,
					trans_no,
					reff_no,
					"0",
					"MediaPromo",
					fileName,
					checklistStock.CreatedBy,
					checklistStock.CreatedByIP)

			}
		}
	}

	//Insert Data Check Stock Product
	for _, detailSKU := range checklistStock.DetailSKU {

		detail_id, err := InsertDataProductCheckStock(ctx, db,
			reff_no, trans_no, detailSKU.ArticleNo,
			detailSKU.ExpDate, detailSKU.Stock, detailSKU.Condition,
			detailSKU.NoteStock, detailSKU.SalesReturnNo, detailSKU.NoteSalesReturn,
			checklistStock.CreatedBy, checklistStock.CreatedByIP,
		)

		if err != nil {
			return err
		}

		if detailSKU.Condition == "Bad Stock" { //If bad stock, upload photo of the bad stock
			//Upload bad stock photo
			countSucceed = 0
			for _, image_base64 := range detailSKU.BadStockImages {
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
					fileName := "Benfarm/Smile/CheckStock/" + checklistStock.CreatedBy + "/" + dt.Format("2006-01-02") + "/" + trans_no + "/" + checklistStock.Type + "_BadStock_" + detailSKU.ArticleNo + "_" + checklistStock.CreatedBy + "_detail_id_" + detail_id + "_" + dt.Format("2006-01-02_15_04_05") + "_" + strconv.Itoa(countSucceed) + ".png"

					_, err = uploader.Upload(&s3manager.UploadInput{
						Bucket: aws.String(MyBucket),
						ACL:    aws.String("public-read"),
						Key:    aws.String(fileName),
						Body:   imageCheckInOut,
					})

					_, err = db.QueryContext(ctx, insertImageDetailStock,
						trans_no,
						reff_no,
						detail_id,
						"BadStock",
						fileName,
						checklistStock.CreatedBy,
						checklistStock.CreatedByIP)

				}
			}
		}

	}

	return nil
}
