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

type productUsecaseTestSuite struct {
	suite.Suite
	Instance       *mongo.Database
	ProductUsecase models.ProductUsecase
	SellerRepo     models.SellerRepository
}

func TestProductUsecase(t *testing.T) {
	suite.Run(t, new(productUsecaseTestSuite))
}

func (s *productUsecaseTestSuite) SetupSuite() {
	instance := mongodb.ConfigureMongo()
	counterRepo := repoMongo.NewCounterRepo(instance)
	productRepo := repoMongo.NewProductRepo(instance, counterRepo)
	sellerRepo := repoMongo.NewSellerRepo(instance, counterRepo)
	s.ProductUsecase = NewProductUsecase(productRepo, sellerRepo)
	s.SellerRepo = sellerRepo
	s.Instance = instance
}

func (s *productUsecaseTestSuite) TearDownSuite() {
	ctx, cancel := context.WithTimeout(context.Background(), configs.Constant.TimeoutOnSeconds*time.Second)
	defer cancel()

	err := testUtil.DropProduct(ctx, s.Instance)
	testUtil.HandleError(err)
	err = testUtil.DropSeller(ctx, s.Instance)
	testUtil.HandleError(err)
	err = testUtil.DropCounter(ctx, s.Instance)
	testUtil.HandleError(err)
}

func (s *productUsecaseTestSuite) TearDownTest() {
	ctx, cancel := context.WithTimeout(context.Background(), configs.Constant.TimeoutOnSeconds*time.Second)
	defer cancel()

	err := testUtil.DropBuyer(ctx, s.Instance)
	testUtil.HandleError(err)
	err = testUtil.DropSeller(ctx, s.Instance)
	testUtil.HandleError(err)
	err = testUtil.DropCounter(ctx, s.Instance)
	testUtil.HandleError(err)
}

func (s *productUsecaseTestSuite) TestStore() {
	s.Run("Test Store Product from Body", func() {
		seller := models.Seller{
			Email:         "test",
			Name:          "test",
			Password:      "test",
			PickupAddress: "test",
		}

		body := models.ProductBody{
			ProductName: "test",
			Description: "test",
			Price:       1.2,
			SellerID:    uint32(1),
		}

		_, err := s.SellerRepo.Store(&seller)
		testUtil.HandleError(err)

		id, err := s.ProductUsecase.CreateProduct(body)
		testUtil.HandleError(err)

		result, err := s.ProductUsecase.GetByID(id)
		testUtil.HandleError(err)

		s.Assert().EqualValues(1, id)
		s.Assert().Equal(&seller, result.Seller)
		s.Assert().Equal(body.ProductName, result.ProductName)
		s.Assert().Equal(body.Description, result.Description)
		s.Assert().Equal(body.Price, result.Price)
	})
}
