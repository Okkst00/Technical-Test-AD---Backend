package service

import (
	"backend-api-commerce/repository"
)


func ConfirmPayment(orderID int, status string) error {
	// validasi status biar aman
	if status != "paid" && status != "cancelled" {
		return nil
	}

	return repository.UpdateOrderStatus(orderID, status)
}