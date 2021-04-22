package apis

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/masterraf21/ecommerce-backend/models"
	httpUtil "github.com/masterraf21/ecommerce-backend/utils/http"
)

type sellerAPI struct {
	SellerUsecase models.SellerUsecase
}

// NewSellerAPI will create api for seller
func NewSellerAPI(r *mux.Router, suc models.SellerUsecase) {
	sellerAPI := &sellerAPI{
		SellerUsecase: suc,
	}

	r.HandleFunc("/seller", sellerAPI.Create).Methods("POST")
	r.HandleFunc("/seller", sellerAPI.GetAll).Methods("GET")
	r.HandleFunc("/seller/{id_seller}", sellerAPI.GetByID).Methods("GET")
}

func (s *sellerAPI) Create(w http.ResponseWriter, r *http.Request) {
	var body models.SellerBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		httpUtil.HandleError(w, r, err, "bad request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	id, err := s.SellerUsecase.CreateSeller(body)
	if err != nil {
		httpUtil.HandleError(w, r, err, "failed to create seller", http.StatusInternalServerError)
		return
	}

	var response struct {
		ID uint32 `json:"id_seller"`
	}
	response.ID = id

	httpUtil.HandleJSONResponse(w, r, response)
}

func (s *sellerAPI) GetAll(w http.ResponseWriter, r *http.Request) {
	result, err := s.SellerUsecase.GetAll()
	if err != nil {
		httpUtil.HandleError(w, r, err, "failed to get seller data", http.StatusInternalServerError)
		return
	}

	var data struct {
		Data []models.Seller `json:"data"`
	}
	data.Data = result

	httpUtil.HandleJSONResponse(w, r, data)
}

func (s *sellerAPI) GetByID(w http.ResponseWriter, r *http.Request) {
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

	result, err := s.SellerUsecase.GetByID(uint32(sellerID))
	if err != nil {
		httpUtil.HandleError(w, r, err, "failed to get seller data by id", http.StatusInternalServerError)
		return
	}

	var data struct {
		Data *models.Seller `json:"data"`
	}
	data.Data = result

	httpUtil.HandleJSONResponse(w, r, data)
}
