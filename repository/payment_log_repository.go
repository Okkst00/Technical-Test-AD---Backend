package repository

import (
	"backend-api-commerce/config"
)

func IsWebhookProcessed(orderCode string, status string) bool {

	var count int

	config.DBLocal.QueryRow(`
		SELECT COUNT(*)
		FROM payment_logs
		WHERE order_code = ? AND status = ?
	`, orderCode, status).Scan(&count)

	return count > 0
}