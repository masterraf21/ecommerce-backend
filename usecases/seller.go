package usecases

import (
	"github.com/masterraf21/ecommerce-backend/models"
	"github.com/masterraf21/ecommerce-backend/utils/auth"
)

type sellerUsecase struct {
	Repo models.SellerRepository
}

// NewSellerUsecase will initiate usecase
func NewSellerUsecase(srr models.SellerRepository) models.SellerUsecase {
	return &sellerUsecase{Repo: srr}
}

func (u *sellerUsecase) CreateSeller(body models.SellerBody) (id uint32, err error) {
	hash, err := auth.GeneratePassword(body.Password)
	if err != nil {
		return
	}
	seller := models.Seller{
		Email:         body.Email,
		Name:          body.Name,
		Password:      hash,
		PickupAddress: body.PickupAddress,
	}

	oid, err := u.Repo.Store(&seller)
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

func (u *sellerUsecase) GetAll() (res []models.Seller, err error) {
	res, err = u.Repo.GetAll()
	return
}

func (u *sellerUsecase) GetByID(id uint32) (res *models.Seller, err error) {
	res, err = u.Repo.GetByID(id)
	return
}
