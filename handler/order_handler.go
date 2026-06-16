package handler

import (
	"backend-api-commerce/model"
	"backend-api-commerce/service"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req model.CreateOrderRequest

	// decode body
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// ambil user_id dari context
	ctxUserID := r.Context().Value("user_id")
	if ctxUserID == nil {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	// JWT default float64
	userIDFloat, ok := ctxUserID.(float64)
	if !ok {
		http.Error(w, "invalid user_id type", http.StatusUnauthorized)
		return
	}

	// SERVICE: sekarang return orderID
	orderID, err := service.CreateOrder(req, userIDFloat)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// response
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message":   "order created",
		"order_id":  orderID,
		"status":    "pending",
	})
}


func GetOrders(w http.ResponseWriter,r *http.Request){
	data,err := service.GetOrders()
	if err != nil {
		http.Error(
			w,
			err.Error(),
			400,
		)
		return
	}

	json.NewEncoder(w).Encode(data)

}

func GetOrderHistory(w http.ResponseWriter, r *http.Request) {

	role, ok := r.Context().Value("role").(string)
	if !ok {
		role = ""
	}

	uid, ok := r.Context().Value("user_id").(float64)
	if !ok {
		http.Error(w, "invalid user id", http.StatusUnauthorized)
		return
	}

	userID := int(uid)

	orders, err := service.GetOrderHistory(role, userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"data": orders,
	})
}

func GetOrderByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid order id", http.StatusBadRequest)
		return
	}

	order, err := service.GetOrderByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(order)
}