package apis

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/masterraf21/ecommerce-backend/models"
	httpUtil "github.com/masterraf21/ecommerce-backend/utils/http"
)

type productAPI struct {
	ProductUsecase models.ProductUsecase
}

// NewProductAPI will create api for product
func NewProductAPI(r *mux.Router, pru models.ProductUsecase) {
	productAPI := &productAPI{
		ProductUsecase: pru,
	}

	r.HandleFunc("/product", productAPI.Create).Methods("POST")
	r.HandleFunc("/product", productAPI.GetAll).Methods("GET")
	r.HandleFunc("/product/{id_product}", productAPI.GetByID).Methods("GET")
	r.HandleFunc("/product/seller/{id_seller}", productAPI.GetBySellerID).Methods("GET")
}

func (p *productAPI) Create(w http.ResponseWriter, r *http.Request) {
	var body models.ProductBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		httpUtil.HandleError(w, r, err, "bad request body", http.StatusBadRequest)
	}
	defer r.Body.Close()

	id, err := p.ProductUsecase.CreateProduct(body)
	if err != nil {
		httpUtil.HandleError(w, r, err, "failed to creata product", http.StatusInternalServerError)
		return
	}

	var response struct {
		ID uint32 `json:"id_product"`
	}
	response.ID = id

	httpUtil.HandleJSONResponse(w, r, response)
}

func (p *productAPI) GetAll(w http.ResponseWriter, r *http.Request) {
	result, err := p.ProductUsecase.GetAll()
	if err != nil {
		httpUtil.HandleError(w, r, err, "failed to get product data", http.StatusInternalServerError)
		return
	}

	var data struct {
		Data []models.Product `json:"data"`
	}
	data.Data = result
	httpUtil.HandleJSONResponse(w, r, data)
}

func (p *productAPI) GetByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	productID, err := strconv.ParseInt(params["id_product"], 10, 64)
	if err != nil {
		httpUtil.HandleError(
			w,
			r,
			err,
			params["id_product"]+" is not integer",
			http.StatusBadRequest,
		)
		return
	}

	result, err := p.ProductUsecase.GetByID(uint32(productID))
	if err != nil {
		httpUtil.HandleError(w, r, err, "failed to get product data by id", http.StatusInternalServerError)
		return
	}

	var data struct {
		Data *models.Product `json:"data"`
	}
	data.Data = result

	httpUtil.HandleJSONResponse(w, r, data)
}

func (p *productAPI) GetBySellerID(w http.ResponseWriter, r *http.Request) {
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

	result, err := p.ProductUsecase.GetBySellerID(uint32(sellerID))
	if err != nil {
		httpUtil.HandleError(w, r, err, "failed to get product data by seller id", http.StatusInternalServerError)
		return
	}

	var data struct {
		Data []models.Product `json:"data"`
	}
	data.Data = result

	httpUtil.HandleJSONResponse(w, r, data)
}
