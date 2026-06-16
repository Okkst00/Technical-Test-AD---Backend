package service

import (
	"backend-api-commerce/config"
	"backend-api-commerce/model"
	"backend-api-commerce/repository"
	"fmt"
	"time"
)

func CreateOrder(
	req model.CreateOrderRequest,
	userID float64,
) (int, error) {

	tx, err := config.DBLocal.Begin()
	if err != nil {
		return 0, err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	// =========================
	// HITUNG TOTAL CHECKOUT
	// =========================
	var calculateItems []model.CheckoutItem

	for _, item := range req.Items {
		calculateItems = append(calculateItems, model.CheckoutItem{
			ProductID: item.ProductID,
			Qty:       item.Qty,
			HargaJual: item.HargaJual,
		})
	}

	calculateReq := model.CalculateRequest{
		UserID: int(userID),
		Items:  calculateItems,
	}

	calculate, err := CalculateCheckout(calculateReq)
	if err != nil {
		return 0, err
	}

	// =========================
	// CREATE ORDER
	// =========================
	order := model.Order{
		OrderCode: GenerateOrderCode(),

		UserID: int(userID),

		CustomerName:    req.Order.CustomerName,
		CustomerPhone:   req.Order.CustomerPhone,
		CustomerEmail:   req.Order.CustomerEmail,
		ShippingAddress: req.Order.ShippingAddress,

		PaymentMethodID: req.Order.PaymentMethodID,
		PaymentOptionID: req.Order.PaymentOptionID,

		Status: "pending",

		// pakai hasil perhitungan modal
		TotalAmount: calculate.TotalModal,
	}

	orderID, err := repository.CreateOrderTx(tx, order)
	if err != nil {
		return 0, err
	}

	// =========================
	// INSERT ORDER ITEMS
	// =========================
	for _, item := range req.Items {

		var stock int

		err = tx.QueryRow(
			"SELECT stok FROM products WHERE id = ? FOR UPDATE",
			item.ProductID,
		).Scan(&stock)

		if err != nil {
			return 0, err
		}

		if stock < item.Qty {
			return 0, fmt.Errorf("stok produk %d tidak cukup", item.ProductID)
		}

		price := item.HargaJual
		subtotal := price * float64(item.Qty)

		_, err = tx.Exec(`
			INSERT INTO order_items (
				order_id,
				product_id,
				price,
				qty,
				subtotal
			) VALUES (?, ?, ?, ?, ?)
		`,
			orderID,
			item.ProductID,
			price,
			item.Qty,
			subtotal,
		)
		if err != nil {
			return 0, err
		}

		_, err = tx.Exec(`
			UPDATE products
			SET stok = stok - ?
			WHERE id = ?
		`,
			item.Qty,
			item.ProductID,
		)
		if err != nil {
			return 0, err
		}
	}

	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	return orderID, nil
}

func GetOrders() ([]model.Order, error) {
	return repository.GetOrders()
}

func GetOrderHistory(role string, userID int) ([]model.Order, error) {
	if role == "admin" {
		return repository.GetAllOrders()
	}
	return repository.GetOrdersByUser(userID)
}

func GetOrderByID(id int) (model.Order, error) {
	return repository.GetOrderByID(id)
}

func GenerateOrderCode() string {
	return fmt.Sprintf("ORD-%d", time.Now().Unix())
}