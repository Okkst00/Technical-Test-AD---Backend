package repository

import (
	"backend-api-commerce/config"
	"backend-api-commerce/helper"
	"backend-api-commerce/model"
	"database/sql"
)



func CreateOrder(
	order model.Order,
	items []model.OrderItemRequest,
) (int64, error) {

	tx, err := config.DBLocal.Begin()
	if err != nil {
		return 0, err
	}

	result, err := tx.Exec(`
		INSERT INTO orders (
			order_code,
			user_id,
			customer_name,
			customer_phone,
			customer_email,
			shipping_address,
			payment_method_id,
			payment_option_id,
			total_amount,
			status
		) VALUES (?,?,?,?,?,?,?,?,?,?)
	`,
		helper.GenerateOrderCode(),
		order.UserID,
		order.CustomerName,
		order.CustomerPhone,
		order.CustomerEmail,
		order.ShippingAddress,
		order.PaymentMethodID,
		order.PaymentOptionID,
		0,
		order.Status,
	)

	if err != nil {
		tx.Rollback()
		return 0, err
	}

	orderID, err := result.LastInsertId()
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	// insert items
	for _, item := range items {
		_, err = tx.Exec(`
			INSERT INTO order_items (
				order_id,
				product_id,
				product_name,
				price,
				qty,
				subtotal
			) VALUES (?,?,?,?,?,?)
		`,
			orderID,
			item.ProductID,
			"",
			0,
			item.Qty,
			0,
		)

		if err != nil {
			tx.Rollback()
			return 0, err
		}
	}

	if err := tx.Commit(); err != nil {
		return 0, err
	}

	return orderID, nil
}

func nullIntToPointer(n sql.NullInt64) *int {
	if n.Valid {
		v := int(n.Int64)
		return &v
	}
	return nil
}


func GetOrders()([]model.Order,error){


	rows,err:=config.DBLocal.Query(`

	SELECT

	id,
	order_code,
	user_id,
	customer_name,
	customer_phone,
	customer_email,
	shipping_address,
	payment_method_id,
	payment_option_id,
	total_amount,
	status,
	created_at,
	updated_at

	FROM orders

	ORDER BY id DESC

	`)



	if err!=nil{
		return nil,err
	}


	defer rows.Close()



	var orders []model.Order



	for rows.Next(){


		var order model.Order



		err:=rows.Scan(

			&order.ID,
			&order.OrderCode,
			&order.UserID,
			&order.CustomerName,
			&order.CustomerPhone,
			&order.CustomerEmail,
			&order.ShippingAddress,
			&order.PaymentMethodID,
			&order.PaymentOptionID,
			&order.TotalAmount,
			&order.Status,
			&order.CreatedAt,
			&order.UpdatedAt,

		)



		if err!=nil{
			return nil,err
		}


		orders=append(
			orders,
			order,
		)

	}


	return orders,nil

}

func GetOrdersByUser(userID int) ([]model.Order, error) {
	var orders []model.Order

	query := `
		SELECT id, order_code, user_id, customer_name, customer_phone,
		       customer_email, shipping_address, payment_method_id,
		       payment_option_id, total_amount, status, created_at, updated_at
		FROM orders
		WHERE user_id = ?
		ORDER BY id DESC
	`

	rows, err := config.DBLocal.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var o model.Order
		var paymentOptionID sql.NullInt64

		err := rows.Scan(
			&o.ID,
			&o.OrderCode,
			&o.UserID,
			&o.CustomerName,
			&o.CustomerPhone,
			&o.CustomerEmail,
			&o.ShippingAddress,
			&o.PaymentMethodID,
			&paymentOptionID,
			&o.TotalAmount,
			&o.Status,
			&o.CreatedAt,
			&o.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		o.PaymentOptionID = nullIntToPointer(paymentOptionID)

		orders = append(orders, o)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return orders, nil
}

func GetAllOrders() ([]model.Order, error) {
	var orders []model.Order

	query := `
		SELECT id, order_code, user_id, customer_name, customer_phone,
		       customer_email, shipping_address, payment_method_id,
		       payment_option_id, total_amount, status, created_at, updated_at
		FROM orders
		ORDER BY id DESC
	`

	rows, err := config.DBLocal.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var o model.Order
		var paymentOptionID sql.NullInt64

		err := rows.Scan(
			&o.ID,
			&o.OrderCode,
			&o.UserID,
			&o.CustomerName,
			&o.CustomerPhone,
			&o.CustomerEmail,
			&o.ShippingAddress,
			&o.PaymentMethodID,
			&paymentOptionID,
			&o.TotalAmount,
			&o.Status,
			&o.CreatedAt,
			&o.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		o.PaymentOptionID = nullIntToPointer(paymentOptionID)

		orders = append(orders, o)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return orders, nil
}

func UpdateOrderStatus(orderID int, status string) error {
	_, err := config.DBLocal.Exec(`
		UPDATE orders
		SET status = ?
		WHERE id = ?
	`, status, orderID)

	return err
}

func GetOrderByID(id int) (model.Order, error) {

	var order model.Order

	err := config.DBLocal.QueryRow(`
		SELECT
			id,
			order_code,
			user_id,
			customer_name,
			customer_phone,
			customer_email,
			shipping_address,
			payment_method_id,
			payment_option_id,
			total_amount,
			status,
			created_at,
			updated_at
		FROM orders
		WHERE id = ?
	`, id).Scan(
		&order.ID,
		&order.OrderCode,
		&order.UserID,
		&order.CustomerName,
		&order.CustomerPhone,
		&order.CustomerEmail,
		&order.ShippingAddress,
		&order.PaymentMethodID,
		&order.PaymentOptionID,
		&order.TotalAmount,
		&order.Status,
		&order.CreatedAt,
		&order.UpdatedAt,
	)

	if err != nil {
		return order, err
	}


	items, err := GetOrderItems(id)

	if err != nil {
		return order, err
	}


	order.Items = items


	return order, nil
}

func GetOrderItems(orderID int) ([]model.OrderItem, error) {

	items := []model.OrderItem{}

	rows, err := config.DBLocal.Query(`
		SELECT
			oi.id,
			oi.order_id,
			oi.product_id,
			p.nama_produk,
			oi.qty
		FROM order_items oi
		JOIN products p
			ON p.id = oi.product_id
		WHERE oi.order_id = ?
	`,
		orderID,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()


	for rows.Next() {

		var item model.OrderItem

		err := rows.Scan(
			&item.ID,
			&item.OrderID,
			&item.ProductID,
			&item.ProductName,
			&item.Qty,
		)

		if err != nil {
			return nil, err
		}


		items = append(items, item)
	}


	return items, nil
}

func CreateOrderTx(tx *sql.Tx, order model.Order) (int, error) {

	res, err := tx.Exec(`
		INSERT INTO orders (
			order_code,
			user_id,
			customer_name,
			customer_phone,
			customer_email,
			shipping_address,
			payment_method_id,
			payment_option_id,
			total_amount,
			status
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`,
		order.OrderCode,
		order.UserID,
		order.CustomerName,
		order.CustomerPhone,
		order.CustomerEmail,
		order.ShippingAddress,
		order.PaymentMethodID,
		order.PaymentOptionID,
		order.TotalAmount,
		order.Status,
	)

	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}