package notification

import (
	"context"
	"database/sql"
	"smile-service/entities"

	"go.uber.org/zap"
)

const execSPNotificationDelete = `exec [sp_smile_user_notification_delete]$1,$2`

func UserNotificationDelete(
	ctx context.Context,
	db *sql.DB,
	notification *entities.NotificationDelete,
	log *zap.Logger) error {

	rows, err := db.QueryContext(
		ctx,
		execSPNotificationDelete,
		notification.UserId,
		notification.TransNo,
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
