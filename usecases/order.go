package usecases

import (
	"github.com/masterraf21/ecommerce-backend/models"
)

type orderUsecase struct {
	Repo        models.OrderRepository
	BuyerRepo   models.BuyerRepository
	SellerRepo  models.SellerRepository
	ProductRepo models.ProductRepository
}

// NewOrderUsecase will nititate usecase for order
func NewOrderUsecase(
	orr models.OrderRepository,
	brr models.BuyerRepository,
	slr models.SellerRepository,
	prr models.ProductRepository,
) models.OrderUsecase {
	return &orderUsecase{
		Repo:        orr,
		BuyerRepo:   brr,
		SellerRepo:  slr,
		ProductRepo: prr,
	}
}

func (u *orderUsecase) CreateOrder(body models.OrderBody) (id uint32, err error) {
	sellerPtr, err := u.SellerRepo.GetByID(body.SellerID)
	if err != nil {
		return
	}

	buyerPtr, err := u.BuyerRepo.GetByID(body.BuyerID)
	if err != nil {
		return
	}

	orderDetails := make([]models.OrderDetail, 0)

	var productPtr *models.Product
	for _, product := range body.Products {
		productPtr, err = u.ProductRepo.GetByID(product.ProductID)
		if err != nil {
			return
		}

		var orderDetail models.OrderDetail

		if productPtr != nil {
			orderDetail = models.OrderDetail{
				ProductID:  product.ProductID,
				Quantity:   product.Quantity,
				Product:    productPtr,
				TotalPrice: float32(product.Quantity) * productPtr.Price,
			}
		} else {
			orderDetail = models.OrderDetail{
				ProductID: product.ProductID,
				Quantity:  product.Quantity,
			}
		}

		orderDetails = append(orderDetails, orderDetail)
	}

	totalPrice := float32(0)
	for _, orderDetail := range orderDetails {
		totalPrice += orderDetail.TotalPrice
	}

	order := models.Order{
		BuyerID:    body.BuyerID,
		Buyer:      buyerPtr,
		SellerID:   body.SellerID,
		Seller:     sellerPtr,
		Products:   orderDetails,
		TotalPrice: totalPrice,
		Status:     "Pending",
	}
	if sellerPtr != nil {
		order.SourceAddress = sellerPtr.PickupAddress
	}
	if buyerPtr != nil {
		order.DeliveryAddress = buyerPtr.DeliveryAddress
	}

	oid, err := u.Repo.Store(&order)
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

func (u *orderUsecase) AcceptOrder(id uint32) (err error) {
	err = u.Repo.UpdateArbitrary(id, "status", "Accepted")
	return
}

func (u *orderUsecase) GetAll() (res []models.Order, err error) {
	res, err = u.Repo.GetAll()
	return
}

func (u *orderUsecase) GetByID(id uint32) (res *models.Order, err error) {
	res, err = u.Repo.GetByID(id)
	return
}

func (u *orderUsecase) GetBySellerID(sellerID uint32) (res []models.Order, err error) {
	res, err = u.Repo.GetBySellerID(sellerID)
	return
}

func (u *orderUsecase) GetByBuyerID(buyerID uint32) (res []models.Order, err error) {
	res, err = u.Repo.GetByBuyerID(buyerID)
	return
}

func (u *orderUsecase) GetByBuyerIDAndStatus(buyerID uint32, status string) (res []models.Order, err error) {
	res, err = u.Repo.GetByBuyerIDAndStatus(buyerID, status)
	return
}

func (u *orderUsecase) GetBySellerIDAndStatus(sellerID uint32, status string) (res []models.Order, err error) {
	res, err = u.Repo.GetByBuyerIDAndStatus(sellerID, status)
	return
}

func (u *orderUsecase) GetByStatus(status string) (res []models.Order, err error) {
	res, err = u.Repo.GetByStatus(status)
	return
}
