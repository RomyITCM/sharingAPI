package store_visit

import (
	"context"
	"database/sql"
	"smile-service/entities"
	"strings"
	"time"
)

const (
	execSPDataChecklistStock2 = `exec [sp_smile_check_stock_insert]$1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14`
)

const insertImageDetailStock2 = `exec [sp_smile_check_stock_image_insert]$1,$2,$3,$4,$5,$6,$7`

const insertDetailProduct2 = `exec [sp_smile_check_stock_product_insert]$1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11`

func InsertDataCheckStock2(ctx context.Context, db *sql.DB,
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
	rows, err := db.QueryContext(ctx, execSPDataChecklistStock2,
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

func InsertDataProductCheckStock2(ctx context.Context, db *sql.DB,
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
	rows, err := db.QueryContext(ctx, insertDetailProduct2,
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

func InsertCheckStock2(ctx context.Context, db *sql.DB, checklistStock *entities.DataCheckStockInsert) (*entities.DataCheckStockInsertOutput, error) {

	trans_no, reff_no, err := InsertDataCheckStock2(ctx, db,
		checklistStock.CustomerNo, checklistStock.BillTo, checklistStock.ShipTo, checklistStock.StockAvailable,
		checklistStock.CreatedBy, checklistStock.CreatedByIP, checklistStock.ArticleNo,
		checklistStock.Type, checklistStock.Price, checklistStock.PromoAvailable, checklistStock.MediaPromoAvailable,
		checklistStock.NotePromo, checklistStock.StartDatePromo, checklistStock.EndDatePromo,
	)

	if err != nil {
		return nil, err
	}

	preSignedURLs := &entities.DataCheckStockInsertOutput{}

	presignedURLPrices := make([]*entities.DataCheckStockInsertOutputURL, 0)
	dt := time.Now()
	filePath := "Benfarm/Smile/CheckStock/" + checklistStock.CreatedBy + "/" + dt.Format("2006-01-02") + "/" + trans_no + "/"
	//Upload and Insert Price Images
	countSucceed := 0
	for _, image_base64 := range checklistStock.PriceImages {
		countSucceed++
		presignedURLPrice := &entities.DataCheckStockInsertOutputURL{}

		fileName := checklistStock.Type + "_Price_" + checklistStock.CreatedBy + "_" + dt.Format("2006-01-02_15_04_05") + "_" + image_base64.Base64Image + ".png"

		url, err := getPresignedURL2(filePath + fileName)

		presignedURLPrice.FileName = fileName
		presignedURLPrice.PreSignURL = url

		presignedURLPrices = append(presignedURLPrices, presignedURLPrice)

		if err != nil {
			return nil, err
		} else {

			_, err = db.QueryContext(ctx, insertImageDetailStock2,
				trans_no,
				reff_no,
				"0",
				"Price",
				filePath+fileName,
				checklistStock.CreatedBy,
				checklistStock.CreatedByIP)

		}
		preSignedURLs.PriceImages = presignedURLPrices

	}

	if checklistStock.PromoAvailable == 1 && checklistStock.MediaPromoAvailable == 1 { //If ada promo yang aktif, upload photo
		//Upload and Insert Promo Images
		presignedURLPromos := make([]*entities.DataCheckStockInsertOutputURL, 0)
		dt := time.Now()
		countSucceed = 0
		for _, image_base64 := range checklistStock.PromoImages {
			countSucceed++
			presignedURLPromo := &entities.DataCheckStockInsertOutputURL{}

			fileName := checklistStock.Type + "_MediaPromo_" + checklistStock.CreatedBy + "_" + dt.Format("2006-01-02_15_04_05") + "_" + image_base64.Base64Image + ".png"

			url, err := getPresignedURL2(filePath + fileName)

			presignedURLPromo.FileName = fileName
			presignedURLPromo.PreSignURL = url

			presignedURLPromos = append(presignedURLPromos, presignedURLPromo)

			if err != nil {
				return nil, err
			} else {

				_, err = db.QueryContext(ctx, insertImageDetailStock2,
					trans_no,
					reff_no,
					"0",
					"MediaPromo",
					filePath+fileName,
					checklistStock.CreatedBy,
					checklistStock.CreatedByIP)

			}
		}
		preSignedURLs.MediaPromoImages = presignedURLPromos
	}

	presignedURLConditions := make([]*entities.DataCheckStockInsertOutputURL, 0)

	//Insert Data Check Stock Product
	for _, detailSKU := range checklistStock.DetailSKU {

		detail_id, err := InsertDataProductCheckStock2(ctx, db,
			reff_no, trans_no, detailSKU.ArticleNo,
			detailSKU.ExpDate, detailSKU.Stock, detailSKU.Condition,
			detailSKU.NoteStock, detailSKU.SalesReturnNo, detailSKU.NoteSalesReturn,
			checklistStock.CreatedBy, checklistStock.CreatedByIP,
		)

		if err != nil {
			return nil, err
		}

		if detailSKU.Condition == "Bad Stock" || detailSKU.Condition == "Expired Stock" || detailSKU.Condition == "Expiring Stock" { //If bad stock, upload photo of the bad stock
			//Upload bad stock photo
			dt := time.Now()
			countSucceed = 0
			for _, image_base64 := range detailSKU.BadStockImages {
				countSucceed++
				presignedURLCondition := &entities.DataCheckStockInsertOutputURL{}

				fileName := checklistStock.Type + "_" + strings.ReplaceAll(detailSKU.Condition, " ", "") + "_" + detailSKU.ArticleNo + "_" + checklistStock.CreatedBy + "_detail_id_" + detail_id + "_" + dt.Format("2006-01-02_15_04_05") + "_" + image_base64.Base64Image + ".png"

				url, err := getPresignedURL2(filePath + fileName)

				presignedURLCondition.FileName = fileName
				presignedURLCondition.PreSignURL = url

				presignedURLConditions = append(presignedURLConditions, presignedURLCondition)

				if err != nil {
					return nil, err
				} else {

					_, err = db.QueryContext(ctx, insertImageDetailStock2,
						trans_no,
						reff_no,
						detail_id,
						strings.ReplaceAll(detailSKU.Condition, " ", ""),
						filePath+fileName,
						checklistStock.CreatedBy,
						checklistStock.CreatedByIP)

				}
			}
		}
		preSignedURLs.ConditionImages = presignedURLConditions

	}

	return preSignedURLs, nil
}
