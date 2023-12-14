package price_zone

import (
	"context"
	"database/sql"
	"smile-service/entities"

	"go.uber.org/zap"
)

const execSPArticlePrice = `exec [sp_smile_zone_product_getrows] $1, $2`

func GetRowsArticlePrice(ctx context.Context, db *sql.DB, zoneId string,
	search string, log *zap.Logger) ([]*entities.DataArticlePrice, error) {

	rows, err := db.QueryContext(
		ctx,
		execSPArticlePrice,
		zoneId,
		search,
	)

	if err != nil {
		log.Info("Exec DB",
			zap.Any("Exec DB", err),
		)
		return nil, err
	}
	defer rows.Close()

	articles := make([]*entities.DataArticlePrice, 0)
	for rows.Next() {
		article := &entities.DataArticlePrice{}

		if err := rows.Scan(
			&article.ZoneId,
			&article.ArticleNo,
			&article.ArticleDescription,
			&article.Uom,
			&article.SalesPrice,
		); err != nil {
			log.Info("scan rows",
				zap.Any("rows", err),
			)
			return nil, err
		}

		articles = append(articles, article)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return articles, nil
}
