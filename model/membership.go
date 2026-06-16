package model

type Membership struct {
	ID                  int     `json:"id"`
	Name                string  `json:"name"`
	FeeTransaksiPercent float64 `json:"fee_transaksi_percent"`
	DiskonHppPercent    float64 `json:"diskon_hpp_percent"`
	Description         string  `json:"description"`
	IsActive            bool    `json:"is_active"`
}