package usecases

import (
	"context"
	"testing"
	"time"

	"github.com/masterraf21/ecommerce-backend/configs"
	"github.com/masterraf21/ecommerce-backend/models"
	repoMongo "github.com/masterraf21/ecommerce-backend/repositories/mongodb"
	"github.com/masterraf21/ecommerce-backend/utils/mongodb"
	testUtil "github.com/masterraf21/ecommerce-backend/utils/test"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/stretchr/testify/suite"
)

type orderUsecaseTestSuite struct {
	suite.Suite
	Instance     *mongo.Database
	OrderUsecase models.OrderUsecase
	SellerRepo   models.SellerRepository
	ProductRepo  models.ProductRepository
	BuyerRepo    models.BuyerRepository
}

func TestOrderUsecase(t *testing.T) {
	suite.Run(t, new(orderUsecaseTestSuite))
}

func (s *orderUsecaseTestSuite) SetupSuite() {
	instance := mongodb.ConfigureMongo()
	counterRepo := repoMongo.NewCounterRepo(instance)
	productRepo := repoMongo.NewProductRepo(instance, counterRepo)
	sellerRepo := repoMongo.NewSellerRepo(instance, counterRepo)
	orderRepo := repoMongo.NewOrderRepo(instance, counterRepo)
	buyerRepo := repoMongo.NewBuyerRepo(instance, counterRepo)

	s.OrderUsecase = NewOrderUsecase(
		orderRepo,
		buyerRepo,
		sellerRepo,
		productRepo,
	)
	s.SellerRepo = sellerRepo
	s.ProductRepo = productRepo
	s.BuyerRepo = buyerRepo
	s.Instance = instance
}

func (s *orderUsecaseTestSuite) TearDownSuite() {
	ctx, cancel := context.WithTimeout(context.Background(), configs.Constant.TimeoutOnSeconds*time.Second)
	defer cancel()

	err := testUtil.DropProduct(ctx, s.Instance)
	testUtil.HandleError(err)
	err = testUtil.DropSeller(ctx, s.Instance)
	testUtil.HandleError(err)
	err = testUtil.DropBuyer(ctx, s.Instance)
	testUtil.HandleError(err)
	err = testUtil.DropOrder(ctx, s.Instance)
	testUtil.HandleError(err)
	err = testUtil.DropCounter(ctx, s.Instance)
	testUtil.HandleError(err)
}

func (s *orderUsecaseTestSuite) TearDownTest() {
	ctx, cancel := context.WithTimeout(context.Background(), configs.Constant.TimeoutOnSeconds*time.Second)
	defer cancel()

	err := testUtil.DropProduct(ctx, s.Instance)
	testUtil.HandleError(err)
	err = testUtil.DropSeller(ctx, s.Instance)
	testUtil.HandleError(err)
	err = testUtil.DropBuyer(ctx, s.Instance)
	testUtil.HandleError(err)
	err = testUtil.DropOrder(ctx, s.Instance)
	testUtil.HandleError(err)
	err = testUtil.DropCounter(ctx, s.Instance)
	testUtil.HandleError(err)
}

func (s *orderUsecaseTestSuite) TestCreate() {
	s.Run("Create Order from body", func() {
		seller := models.Seller{
			Email:         "test",
			Name:          "test",
			Password:      "test",
			PickupAddress: "test pickup",
		}

		buyer := models.Buyer{
			Email:           "test",
			Name:            "test",
			Password:        "test",
			DeliveryAddress: "test delivery",
		}

		product := models.Product{
			Seller:      &seller,
			ProductName: "test",
			Description: "test",
			Price:       125000,
			SellerID:    uint32(1),
		}

		product2 := models.Product{
			Seller:      &seller,
			ProductName: "test",
			Description: "test",
			Price:       15000,
			SellerID:    uint32(1),
		}

		_, err := s.BuyerRepo.Store(&buyer)
		testUtil.HandleError(err)
		_, err = s.SellerRepo.Store(&seller)
		testUtil.HandleError(err)
		_, err = s.ProductRepo.Store(&product)
		testUtil.HandleError(err)
		_, err = s.ProductRepo.Store(&product2)
		testUtil.HandleError(err)

		body := models.OrderBody{
			BuyerID:  uint32(1),
			SellerID: uint32(1),
			Products: []models.ProductDetail{
				{
					ProductID: uint32(1),
					Quantity:  10,
				},
				{
					ProductID: uint32(2),
					Quantity:  4,
				},
			},
		}

		totalPrice := product.Price*10 + product2.Price*4

		id, err := s.OrderUsecase.CreateOrder(body)
		testUtil.HandleError(err)

		result, err := s.OrderUsecase.GetByID(id)
		testUtil.HandleError(err)

		s.Assert().EqualValues(1, id)
		s.Assert().Equal(&seller, result.Seller)
		s.Assert().Equal(&buyer, result.Buyer)
		s.Assert().EqualValues(totalPrice, result.TotalPrice)
	})
}

func (s *orderUsecaseTestSuite) TestCreate2() {
	s.Run("Create order without other data available", func() {
		body := models.OrderBody{
			BuyerID:  uint32(1),
			SellerID: uint32(1),
			Products: []models.ProductDetail{
				{
					ProductID: uint32(1),
					Quantity:  10,
				},
				{
					ProductID: uint32(2),
					Quantity:  4,
				},
			},
		}

		totalPrice := 0

		id, err := s.OrderUsecase.CreateOrder(body)
		testUtil.HandleError(err)

		result, err := s.OrderUsecase.GetByID(id)
		testUtil.HandleError(err)

		s.Assert().EqualValues(1, id)
		s.Assert().Nil(result.Buyer)
		s.Assert().Nil(result.Seller)
		s.Assert().EqualValues(totalPrice, result.TotalPrice)
		s.Assert().Equal("", result.DeliveryAddress)
		s.Assert().Equal("", result.SourceAddress)
	})
}
