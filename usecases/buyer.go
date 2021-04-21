package usecases

import (
	"github.com/masterraf21/ecommerce-backend/utils/auth"

	"github.com/masterraf21/ecommerce-backend/models"
)

type buyerUsecase struct {
	Repo models.BuyerRepository
}

// NewBuyerUsecase will initiate usecase
func NewBuyerUsecase(br models.BuyerRepository) models.BuyerUsecase {
	return &buyerUsecase{Repo: br}
}

func (u *buyerUsecase) CreateBuyer(body models.BuyerBody) (id uint32, err error) {
	hash, err := auth.GeneratePassword(body.Password)
	if err != nil {
		return
	}
	buyer := models.Buyer{
		Email:           body.Email,
		Name:            body.Name,
		Password:        hash,
		DeliveryAddress: body.DeliveryAddress,
	}

	oid, err := u.Repo.Store(&buyer)
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

func (u *buyerUsecase) GetAll() (res []models.Buyer, err error) {
	res, err = u.Repo.GetAll()
	return
}

func (u *buyerUsecase) GetByID(id uint32) (res *models.Buyer, err error) {
	res, err = u.Repo.GetByID(id)
	return
}
