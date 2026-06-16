package service

import (
	"errors"

	"backend-api-commerce/model"
	"backend-api-commerce/repository"
)



func GetProducts()([]model.Product,error){
	return repository.GetProducts()

}


func CreateProduct(
	req model.ProductRequest,
)error{


	if req.NamaProduk == ""{
		return errors.New("nama produk wajib diisi")
	}


	if req.HargaModal <=0{
		return errors.New("harga modal harus lebih dari 0")
	}


	if req.Stok <0{
		return errors.New("stok tidak boleh negatif")
	}



	return repository.CreateProduct(req)

}


func UpdateProduct(
	id int,
	req model.ProductRequest,
)error{


	return repository.UpdateProduct(id,req)

}


func DeleteProduct(id int)error{

	return repository.DeleteProduct(id)

}

func GetProductByID(id int) (model.Product, error) {
	return repository.GetProductByID(id)
}