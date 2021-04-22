package apis

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/masterraf21/ecommerce-backend/models"
	httpUtil "github.com/masterraf21/ecommerce-backend/utils/http"
)

type orderAPI struct {
	OrderUsecase models.OrderUsecase
}

// NewOrderAPI will create api for order
func NewOrderAPI(r *mux.Router, oru models.OrderUsecase) {
	orderAPI := &orderAPI{
		OrderUsecase: oru,
	}

	r.HandleFunc("/order", orderAPI.Create).Methods("POST")
	r.HandleFunc("/order", orderAPI.GetAll).Methods("GET")
	r.HandleFunc("/order/{id_order}", orderAPI.GetByID).Methods("GET")
	r.HandleFunc("/order/seller/{id_seller}", orderAPI.GetBySellerID).Methods("GET")
	r.HandleFunc("/order/buyer/{id_buyer}", orderAPI.GetByBuyerID).Methods("GET")
	r.HandleFunc("/order/seller/{id_seller}/accepted", orderAPI.GetAcceptedOrderBySellerID).Methods("GET")
	r.HandleFunc("/order/buyer/{id_buyer}/accepted", orderAPI.GetAcceptedOrderByBuyerID).Methods("GET")
	r.HandleFunc("/order/seller/{id_seller}/pending", orderAPI.GetPendingOrderBySellerID).Methods("GET")
	r.HandleFunc("/order/buyer/{id_buyer}/pending", orderAPI.GetPendingOrderByBuyerID).Methods("GET")
	r.HandleFunc("/order/{id_order}/accept", orderAPI.AcceptOrder).Methods("POST")
}

func (o *orderAPI) Create(w http.ResponseWriter, r *http.Request) {
	var body models.OrderBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		httpUtil.HandleError(w, r, err, "bad request body", http.StatusBadRequest)
	}
	defer r.Body.Close()

	id, err := o.OrderUsecase.CreateOrder(body)
	if err != nil {
		httpUtil.HandleError(w, r, err, "failed to creata order", http.StatusInternalServerError)
		return
	}

	var response struct {
		ID uint32 `json:"id_order"`
	}
	response.ID = id

	httpUtil.HandleJSONResponse(w, r, response)
}

func (o *orderAPI) GetAll(w http.ResponseWriter, r *http.Request) {
	result, err := o.OrderUsecase.GetAll()
	if err != nil {
		httpUtil.HandleError(w, r, err, "failed to get order data", http.StatusInternalServerError)
		return
	}

	var data struct {
		Data []models.Order `json:"data"`
	}
	data.Data = result
	httpUtil.HandleJSONResponse(w, r, data)
}

func (o *orderAPI) GetByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	orderID, err := strconv.ParseInt(params["id_order"], 10, 64)
	if err != nil {
		httpUtil.HandleError(
			w,
			r,
			err,
			params["id_order"]+" is not integer",
			http.StatusBadRequest,
		)
		return
	}

	result, err := o.OrderUsecase.GetByID(uint32(orderID))
	if err != nil {
		httpUtil.HandleError(w, r, err, "failed to get order data by id", http.StatusInternalServerError)
		return
	}

	var data struct {
		Data *models.Order `json:"data"`
	}
	data.Data = result

	httpUtil.HandleJSONResponse(w, r, data)
}

func (o *orderAPI) GetBySellerID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	sellerID, err := strconv.ParseInt(params["id_seller"], 10, 64)
	if err != nil {
		httpUtil.HandleError(
			w,
			r,
			err,
			params["id_seller"]+" is not integer",
			http.StatusBadRequest,
		)
		return
	}

	result, err := o.OrderUsecase.GetBySellerID(uint32(sellerID))
	if err != nil {
		httpUtil.HandleError(w, r, err, "failed to get order data by id seller", http.StatusInternalServerError)
		return
	}

	var data struct {
		Data []models.Order `json:"data"`
	}
	data.Data = result

	httpUtil.HandleJSONResponse(w, r, data)
}

func (o *orderAPI) GetByBuyerID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	buyerID, err := strconv.ParseInt(params["id_buyer"], 10, 64)
	if err != nil {
		httpUtil.HandleError(
			w,
			r,
			err,
			params["id_buyer"]+" is not integer",
			http.StatusBadRequest,
		)
		return
	}

	result, err := o.OrderUsecase.GetByBuyerID(uint32(buyerID))
	if err != nil {
		httpUtil.HandleError(w, r, err, "failed to get order data by id buyer", http.StatusInternalServerError)
		return
	}

	var data struct {
		Data []models.Order `json:"data"`
	}
	data.Data = result

	httpUtil.HandleJSONResponse(w, r, data)
}

func (o *orderAPI) GetAcceptedOrderBySellerID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	sellerID, err := strconv.ParseInt(params["id_seller"], 10, 64)
	if err != nil {
		httpUtil.HandleError(
			w,
			r,
			err,
			params["id_seller"]+" is not integer",
			http.StatusBadRequest,
		)
		return
	}

	result, err := o.OrderUsecase.GetBySellerIDAndStatus(uint32(sellerID), "Accepted")
	if err != nil {
		httpUtil.HandleError(w, r, err, "failed to get order data by id seller", http.StatusInternalServerError)
		return
	}

	var data struct {
		Data []models.Order `json:"data"`
	}
	data.Data = result

	httpUtil.HandleJSONResponse(w, r, data)
}

func (o *orderAPI) GetPendingOrderBySellerID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	sellerID, err := strconv.ParseInt(params["id_seller"], 10, 64)
	if err != nil {
		httpUtil.HandleError(
			w,
			r,
			err,
			params["id_seller"]+" is not integer",
			http.StatusBadRequest,
		)
		return
	}

	result, err := o.OrderUsecase.GetBySellerIDAndStatus(uint32(sellerID), "Pending")
	if err != nil {
		httpUtil.HandleError(w, r, err, "failed to get order data by id seller", http.StatusInternalServerError)
		return
	}

	var data struct {
		Data []models.Order `json:"data"`
	}
	data.Data = result

	httpUtil.HandleJSONResponse(w, r, data)
}

func (o *orderAPI) GetAcceptedOrderByBuyerID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	buyerID, err := strconv.ParseInt(params["id_buyer"], 10, 64)
	if err != nil {
		httpUtil.HandleError(
			w,
			r,
			err,
			params["id_buyer"]+" is not integer",
			http.StatusBadRequest,
		)
		return
	}

	result, err := o.OrderUsecase.GetByBuyerIDAndStatus(uint32(buyerID), "Accepted")
	if err != nil {
		httpUtil.HandleError(w, r, err, "failed to get order data by id buyer", http.StatusInternalServerError)
		return
	}

	var data struct {
		Data []models.Order `json:"data"`
	}
	data.Data = result

	httpUtil.HandleJSONResponse(w, r, data)
}

func (o *orderAPI) GetPendingOrderByBuyerID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	buyerID, err := strconv.ParseInt(params["id_buyer"], 10, 64)
	if err != nil {
		httpUtil.HandleError(
			w,
			r,
			err,
			params["id_buyer"]+" is not integer",
			http.StatusBadRequest,
		)
		return
	}

	result, err := o.OrderUsecase.GetByBuyerIDAndStatus(uint32(buyerID), "Pending")
	if err != nil {
		httpUtil.HandleError(w, r, err, "failed to get order data by id buyer", http.StatusInternalServerError)
		return
	}

	var data struct {
		Data []models.Order `json:"data"`
	}
	data.Data = result

	httpUtil.HandleJSONResponse(w, r, data)
}

func (o *orderAPI) AcceptOrder(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	orderID, err := strconv.ParseInt(params["id_order"], 10, 64)
	if err != nil {
		httpUtil.HandleError(
			w,
			r,
			err,
			params["id_order"]+" is not integer",
			http.StatusBadRequest,
		)
		return
	}

	err = o.OrderUsecase.AcceptOrder(uint32(orderID))
	if err != nil {
		httpUtil.HandleError(w, r, err, "error accepting order", http.StatusInternalServerError)
		return
	}

	httpUtil.HandleNoJSONResponse(w)
}
