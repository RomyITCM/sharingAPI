package notification

import (
	"context"
	"database/sql"
	"smile-service/entities"

	"go.uber.org/zap"
)

const execSPDataUserNotification = `exec [sp_smile_user_notification_getrows]$1`

func GetDataUserNotification(ctx context.Context, db *sql.DB, user_id string, log *zap.Logger) ([]*entities.DataUserNotification, error) {
	rows, err := db.QueryContext(
		ctx,
		execSPDataUserNotification,
		user_id)

	if err != nil {
		log.Info("Exec DB",
			zap.Any("Exec DB", err),
		)
		return nil, err
	}
	defer rows.Close()

	data_user_notifications := make([]*entities.DataUserNotification, 0)
	for rows.Next() {
		data_user_notification := &entities.DataUserNotification{}

		if err := rows.Scan(
			&data_user_notification.TransNo,
			&data_user_notification.MessageTitle,
			&data_user_notification.MessageBody,
			&data_user_notification.IsRead,
			&data_user_notification.Type,
			&data_user_notification.Salesman,
		); err != nil {
			log.Info("scan rows",
				zap.Any("rows", err),
			)
			return nil, err
		}

		data_user_notifications = append(data_user_notifications, data_user_notification)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return data_user_notifications, nil

}
