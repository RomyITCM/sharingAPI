package notification

import (
	"context"
	"database/sql"
	"smile-service/entities"

	"go.uber.org/zap"
)

const execSPNotificationRead = `exec [sp_smile_user_notification_read_update]$1,$2,$3`

func UserNotificationRead(
	ctx context.Context,
	db *sql.DB,
	notification *entities.NotificationRead,
	log *zap.Logger) error {

	rows, err := db.QueryContext(
		ctx,
		execSPNotificationRead,
		notification.UserId,
		notification.TransNo,
		notification.ReaderIp,
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
