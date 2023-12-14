package login

import (
	"context"
	"database/sql"
	"smile-service/entities"

	"go.uber.org/zap"
)

const execSPChangePassword = `exec [sp_smile_change_password]$1,$2,$3`

func ChangePassword(
	ctx context.Context,
	db *sql.DB,
	changePassword *entities.ChangePassword,
	log *zap.Logger) (*entities.DataChangePassword, error) {

	newPassword := GetMD5Hash(changePassword.NewPassword)
	oldPassword := GetMD5Hash(changePassword.OldPassword)
	rows, err := db.QueryContext(
		ctx,
		execSPChangePassword,
		changePassword.UserId,
		oldPassword,
		newPassword,
	)

	if err != nil {
		log.Info(
			"Exec DB", zap.Any("Exec Db", err),
		)

		return nil, err
	}
	defer rows.Close()

	data_change := &entities.DataChangePassword{}
	for rows.Next() {
		if err := rows.Scan(
			&data_change.Status,
			&data_change.Message,
		); err != nil {
			log.Info("scan rows",
				zap.Any("rows", err),
			)
			return nil, err
		}
	}

	return data_change, nil
}
