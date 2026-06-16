package service

import (
	"backend-api-commerce/model"
	"backend-api-commerce/repository"
)

func CalculateCheckout(
	req model.CalculateRequest,
) (model.CalculateResponse, error) {

	var result model.CalculateResponse

	user, err := repository.GetUserByID(req.UserID)
	if err != nil {
		return result, err
	}

	member, err := repository.GetMembershipByID(user.MembershipID)
	if err != nil {
		return result, err
	}

	result.DiskonHppPercent = member.DiskonHppPercent

	for _, item := range req.Items {

		product, err := repository.GetProductByID(item.ProductID)
		if err != nil {
			return result, err
		}

		qty := float64(item.Qty)

		hargaModal := float64(product.HargaModal)

		// diskon HPP membership
		diskonHpp := hargaModal * (member.DiskonHppPercent / 100)

		hppAktual := hargaModal - diskonHpp

		totalModal := hppAktual * qty

		totalJual := item.HargaJual * qty

		result.TotalModal += totalModal
		result.TotalJual += totalJual
		result.DiskonHppAmount += diskonHpp * qty
	}

	result.FeeTransaksi =
		result.TotalJual *
			(member.FeeTransaksiPercent / 100)

	result.ProfitMember =
		result.TotalJual -
			result.TotalModal -
			result.FeeTransaksi

	return result, nil
}