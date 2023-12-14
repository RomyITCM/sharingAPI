package login

import (
	"context"
	"database/sql"
	"smile-service/entities"

	"crypto/md5"
	"encoding/hex"

	"go.uber.org/zap"
)

const execSPLogin = `exec [sp_smile_login]$1,$2`

func Login(
	ctx context.Context,
	db *sql.DB,
	login *entities.Login,
	log *zap.Logger) (*entities.DataLogin, error) {

	password := GetMD5Hash(login.Password)
	rows, err := db.QueryContext(
		ctx,
		execSPLogin,
		login.UserName,
		password,
	)

	if err != nil {
		log.Info(
			"Exec DB", zap.Any("Exec Db", err),
		)

		return nil, err
	}
	defer rows.Close()

	data_login := &entities.DataLogin{}
	for rows.Next() {
		if err := rows.Scan(
			&data_login.UserId,
			&data_login.UserName,
			&data_login.RoleCode,
			&data_login.Division,
			&data_login.Dept,
			&data_login.DeptCode,
			&data_login.Email,
			&data_login.Phone,
			&data_login.ImageUrl,
			&data_login.Message,
		); err != nil {
			log.Info("scan rows",
				zap.Any("rows", err),
			)
			return nil, err
		}
	}

	return data_login, nil
}

func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
