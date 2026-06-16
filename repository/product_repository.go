package repository

import (
	"backend-api-commerce/config"
	"backend-api-commerce/model"
	"database/sql"
	"fmt"
)


func GetProducts() ([]model.Product,error){

	rows,err := config.DBLocal.Query(`
		SELECT
		p.id,
		p.nama_produk,
		p.kategori,
		p.harga_modal,
		p.stok,
		p.status,
		p.created_by,
		p.updated_by,
		p.created_at,
		p.updated_at
		FROM products p
		ORDER BY p.id DESC
	`)


	if err != nil {
		return nil,err
	}


	defer rows.Close()


	var products []model.Product


	for rows.Next(){

		var product model.Product


		err := rows.Scan(
			&product.ID,
			&product.NamaProduk,
			&product.Kategori,
			&product.HargaModal,
			&product.Stok,
			&product.Status,
			&product.CreatedBy,
			&product.UpdatedBy,
			&product.CreatedAt,
			&product.UpdatedAt,
		)


		if err != nil {
			return nil,err
		}


		products = append(products,product)

	}


	return products,nil
}





func CreateProduct(
	req model.ProductRequest,
) error{


	_,err := config.DBLocal.Exec(`
		INSERT INTO products
		(
			nama_produk,
			kategori,
			harga_modal,
			stok,
			status,
			created_by
		)
		VALUES (?,?,?,?,?,?)
	`,
	req.NamaProduk,
	req.Kategori,
	req.HargaModal,
	req.Stok,
	req.Status,
	req.CreatedBy,
	)


	return err
}






func UpdateProduct(
	id int,
	req model.ProductRequest,
) error{


	_,err := config.DBLocal.Exec(`
		UPDATE products
		SET
		nama_produk=?,
		kategori=?,
		harga_modal=?,
		stok=?,
		status=?,
		updated_by=?
		WHERE id=?
	`,
	req.NamaProduk,
	req.Kategori,
	req.HargaModal,
	req.Stok,
	req.Status,
	req.UpdatedBy,
	id,
	)


	return err
}






func DeleteProduct(id int) error{


	_,err := config.DBLocal.Exec(
		`
		DELETE FROM products
		WHERE id=?
		`,
		id,
	)


	return err
}

func GetProductByID(id int) (model.Product, error) {

	var product model.Product

	err := config.DBLocal.QueryRow(`
		SELECT
			id,
			nama_produk,
			kategori,
			harga_modal,
			stok,
			status
		FROM products
		WHERE id = ?
	`, id).Scan(
		&product.ID,
		&product.NamaProduk,
		&product.Kategori,
		&product.HargaModal,
		&product.Stok,
		&product.Status,
	)

	if err == sql.ErrNoRows {
		return product, fmt.Errorf("product not found")
	}

	return product, err
}