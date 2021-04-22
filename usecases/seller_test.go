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

type sellerUsecaseTestSuite struct {
	suite.Suite
	Instance      *mongo.Database
	SellerUsecase models.SellerUsecase
}

func TestSellerUsecase(t *testing.T) {
	suite.Run(t, new(sellerUsecaseTestSuite))
}

func (s *sellerUsecaseTestSuite) SetupSuite() {
	instance := mongodb.ConfigureMongo()
	counterRepo := repoMongo.NewCounterRepo(instance)
	SellerRepo := repoMongo.NewSellerRepo(instance, counterRepo)
	s.SellerUsecase = NewSellerUsecase(SellerRepo)
	s.Instance = instance
}

func (s *sellerUsecaseTestSuite) TearDownSuite() {
	ctx, cancel := context.WithTimeout(context.Background(), configs.Constant.TimeoutOnSeconds*time.Second)
	defer cancel()

	err := testUtil.DropSeller(ctx, s.Instance)
	testUtil.HandleError(err)
	err = testUtil.DropCounter(ctx, s.Instance)
	testUtil.HandleError(err)
}

func (s *sellerUsecaseTestSuite) TearDownTest() {
	ctx, cancel := context.WithTimeout(context.Background(), configs.Constant.TimeoutOnSeconds*time.Second)
	defer cancel()

	err := testUtil.DropSeller(ctx, s.Instance)
	testUtil.HandleError(err)
	err = testUtil.DropCounter(ctx, s.Instance)
	testUtil.HandleError(err)
}

func (s *sellerUsecaseTestSuite) TestCreate() {
	s.Run("Create Seller from body", func() {
		body := models.SellerBody{
			Email:         "test",
			Name:          "test",
			Password:      "test",
			PickupAddress: "test",
		}

		id, err := s.SellerUsecase.CreateSeller(body)
		testUtil.HandleError(err)

		result, err := s.SellerUsecase.GetByID(id)
		testUtil.HandleError(err)

		s.Assert().EqualValues(1, id)
		s.Assert().Equal(body.Email, result.Email)
		s.Assert().Equal(body.Name, result.Name)
		s.Assert().Equal(body.PickupAddress, result.PickupAddress)
		s.Assert().NotEmpty(result.Password)
	})
}
