package usertoken

import (
	"context"
	"database/sql"
	"smile-service/entities"

	"go.uber.org/zap"
)

const execSPUserTokenInsert = `exec [sp_smile_user_token_insert]$1,$2,$3,$4,$5`

func UserTokenInsert(
	ctx context.Context,
	db *sql.DB,
	userToken *entities.UserToken,
	log *zap.Logger) error {

	rows, err := db.QueryContext(
		ctx,
		execSPUserTokenInsert,
		userToken.UserId,
		userToken.Token,
		userToken.UserName,
		userToken.DeviceName,
		userToken.CreatedByIp,
	)

	if err != nil {
		log.Info(
			"Exec DB", zap.Any("Exec Db", err),
		)

		return err
	}
	defer rows.Close()

	return nil
}
