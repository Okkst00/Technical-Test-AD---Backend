package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"backend-api-commerce/model"
	"backend-api-commerce/service"
	"backend-api-commerce/utils"

	"github.com/gorilla/mux"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {

	data, err := service.GetProducts()

	if err != nil {
		utils.Error(w, 500, err.Error())
		return
	}

	utils.JSON(w, 200, data, "success")
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {

	var req model.ProductRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		utils.Error(w, http.StatusBadRequest, "invalid request")
		return
	}

	// ambil user id dari JWT
	userID := r.Context().Value("user_id")

	if userID == nil {
		utils.Error(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	req.CreatedBy = int(userID.(float64))

	err = service.CreateProduct(req)
	if err != nil {
		utils.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.JSON(w, http.StatusCreated, nil, "product created")
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	var req model.ProductRequest

	_ = json.NewDecoder(r.Body).Decode(&req)

	err := service.UpdateProduct(id, req)
	if err != nil {
		utils.Error(w, 400, err.Error())
		return
	}

	utils.JSON(w, 200, nil, "updated")
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	err := service.DeleteProduct(id)
	if err != nil {
		utils.Error(w, 400, err.Error())
		return
	}

	utils.JSON(w, 200, nil, "deleted")
}

func GetProductByID(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		utils.Error(w, 400, "invalid product id")
		return
	}

	product, err := service.GetProductByID(id)
	if err != nil {
		utils.Error(w, 500, err.Error())
		return
	}

	utils.JSON(w, 200, product, "success")
}