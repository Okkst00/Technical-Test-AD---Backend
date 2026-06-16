package model

import (
	"database/sql"
	"time"
)


type Product struct {
	ID int `json:"id"`
	NamaProduk string `json:"nama_produk"`
	Kategori string `json:"kategori"`
	HargaModal int `json:"harga_modal"`
	Stok int `json:"stok"`
	Status bool `json:"status"`
	CreatedBy int `json:"created_by"`
	UpdatedBy sql.NullInt64 `json:"updated_by"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

}

type ProductRequest struct {
	ID int `json:"id"`
	NamaProduk string `json:"nama_produk"`
	Kategori string `json:"kategori"`
	HargaModal int `json:"harga_modal"`
	Stok int `json:"stok"`
	Status bool `json:"status"`
	CreatedBy int `json:"created_by"`
	UpdatedBy sql.NullInt64 `json:"updated_by"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

}