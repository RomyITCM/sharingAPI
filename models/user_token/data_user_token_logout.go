package usertoken

import (
	"context"
	"database/sql"
	"smile-service/entities"

	"go.uber.org/zap"
)

const execSPUserTokenLogout = `exec [sp_smile_user_token_logout_update]$1,$2`

func UserTokenLogout(
	ctx context.Context,
	db *sql.DB,
	userToken *entities.UserTokenLogout,
	log *zap.Logger) error {

	rows, err := db.QueryContext(
		ctx,
		execSPUserTokenLogout,
		userToken.UserId,
		userToken.Token,
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
