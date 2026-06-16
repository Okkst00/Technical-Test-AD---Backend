package repository

import (
	"backend-api-commerce/config"
	"backend-api-commerce/model"
)

type Membership struct {
	ID                  int
	Name                string
	FeeTransaksiPercent float64
	DiskonHppPercent    float64
}

func GetMembershipByID(id int)(model.Membership,error){

	var member model.Membership

	err := config.DBLocal.QueryRow(`
		SELECT
			id,
			name,
			fee_transaksi_percent,
			diskon_hpp_percent,
			description,
			is_active
		FROM memberships
		WHERE id = ?
	`, id).Scan(
		&member.ID,
		&member.Name,
		&member.FeeTransaksiPercent,
		&member.DiskonHppPercent,
		&member.Description,
		&member.IsActive,
	)

	return member, err
}