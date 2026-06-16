package model

type CheckoutItem struct {
	ProductID  int     `json:"product_id"`
	Qty        int     `json:"qty"`
	HargaJual  float64 `json:"harga_jual"`
	HargaModal float64 `json:"harga_modal"`
}

type CalculateRequest struct {
	UserID int            `json:"user_id"`
	Items  []CheckoutItem `json:"items"`
}

type CalculateResponse struct {
	HargaModalAwal   float64 `json:"harga_modal_awal"`
	DiskonHppPercent float64 `json:"diskon_hpp_percent"`
	DiskonHppAmount  float64 `json:"diskon_hpp_amount"`
	TotalModal       float64 `json:"total_modal"`
	TotalJual        float64 `json:"total_jual"`
	FeeTransaksi     float64 `json:"fee_transaksi"`
	ProfitMember     float64 `json:"profit_member"`
}