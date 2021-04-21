package usecases

import "github.com/masterraf21/ecommerce-backend/models"

type productUsecase struct {
	Repo       models.ProductRepository
	SellerRepo models.SellerRepository
}

// NewProductUsecase will initiate usecase
func NewProductUsecase(prr models.ProductRepository, ssr models.SellerRepository) models.ProductUsecase {
	return &productUsecase{Repo: prr, SellerRepo: ssr}
}

func (u *productUsecase) CreateProduct(body models.ProductBody) (id uint32, err error) {
	sellerPtr, err := u.SellerRepo.GetByID(body.SellerID)
	if err != nil {
		return
	}

	product := models.Product{
		ProductName: body.ProductName,
		Description: body.Description,
		Price:       body.Price,
		SellerID:    body.SellerID,
		Seller:      sellerPtr,
	}

	oid, err := u.Repo.Store(&product)
	if err != nil {
		return
	}
	result, err := u.Repo.GetByOID(oid)
	if err != nil {
		return
	}

	id = result.ID

	return
}

func (u *productUsecase) GetAll() (res []models.Product, err error) {
	res, err = u.Repo.GetAll()
	return
}

func (u *productUsecase) GetBySellerID(sellerID uint32) (res []models.Product, err error) {
	res, err = u.Repo.GetBySellerID(sellerID)
	return
}

func (u *productUsecase) GetByID(id uint32) (res *models.Product, err error) {
	res, err = u.Repo.GetByID(id)
	return
}
