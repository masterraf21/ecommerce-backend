package apis

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/masterraf21/ecommerce-backend/models"
	httpUtil "github.com/masterraf21/ecommerce-backend/utils/http"
)

type buyerAPI struct {
	BuyerUsecase models.BuyerUsecase
}

// NewBuyerAPI will create api for buyer
func NewBuyerAPI(r *mux.Router, buc models.BuyerUsecase) {
	buyerAPI := &buyerAPI{
		BuyerUsecase: buc,
	}

	r.HandleFunc("/buyer", buyerAPI.Create).Methods("POST")
	r.HandleFunc("/buyer", buyerAPI.GetAll).Methods("GET")
	r.HandleFunc("/buyer/{id_buyer}", buyerAPI.GetByID).Methods("GET")
}

func (b *buyerAPI) Create(w http.ResponseWriter, r *http.Request) {
	var body models.BuyerBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		httpUtil.HandleError(w, r, err, "bad request body", http.StatusBadRequest)
	}
	defer r.Body.Close()

	id, err := b.BuyerUsecase.CreateBuyer(body)
	if err != nil {
		httpUtil.HandleError(w, r, err, "failed to creata buyer", http.StatusInternalServerError)
		return
	}

	var response struct {
		ID uint32 `json:"id_buyer"`
	}
	response.ID = id

	httpUtil.HandleJSONResponse(w, r, response)
}

func (b *buyerAPI) GetAll(w http.ResponseWriter, r *http.Request) {
	result, err := b.BuyerUsecase.GetAll()
	if err != nil {
		httpUtil.HandleError(w, r, err, "failed to get buyer data", http.StatusInternalServerError)
		return
	}

	var data struct {
		Data []models.Buyer `json:"data"`
	}
	data.Data = result
	httpUtil.HandleJSONResponse(w, r, data)
}

func (b *buyerAPI) GetByID(w http.ResponseWriter, r *http.Request) {
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

	result, err := b.BuyerUsecase.GetByID(uint32(buyerID))
	if err != nil {
		httpUtil.HandleError(w, r, err, "failed to get buyer data by id", http.StatusInternalServerError)
		return
	}

	var data struct {
		Data *models.Buyer `json:"data"`
	}
	data.Data = result

	httpUtil.HandleJSONResponse(w, r, data)
}
