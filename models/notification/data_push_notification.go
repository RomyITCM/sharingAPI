package notification

import (
	"context"
	"database/sql"
	"smile-service/entities"

	"go.uber.org/zap"
)

const execSPDataPushNotification = `exec [sp_smile_user_notification_getrow]$1`

func SendNotificationSO(
	ctx context.Context,
	db *sql.DB,
	notification *entities.NotificationSend,
	log *zap.Logger,
) error {
	rows, err := db.QueryContext(
		ctx,
		execSPDataPushNotification,
		notification.TransNo,
	)

	if err != nil {
		log.Info(
			"Exec DB", zap.Any("Exec Db", err),
		)
	}
	defer rows.Close()

	for rows.Next() {
		data_notif := &entities.DataNotification{}
		if err := rows.Scan(
			&data_notif.MessageTitle,
			&data_notif.MessageBody,
			&data_notif.TransNo,
			&data_notif.Token,
		); err != nil {
			log.Info("scan rows",
				zap.Any("rows", err),
			)
			return err
		}

		_, err := SendNotificationFromHTTPRequest(data_notif.Token, data_notif.MessageTitle, data_notif.MessageBody, data_notif.TransNo)
		if err != nil {
			log.Info("send notif",
				zap.Any("notif", err),
			)
			return err
		}
	}

	return nil
}
