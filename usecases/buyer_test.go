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

type buyerUsecaseTestSuite struct {
	suite.Suite
	Instance     *mongo.Database
	BuyerUsecase models.BuyerUsecase
}

func TestBuyerUsecase(t *testing.T) {
	suite.Run(t, new(buyerUsecaseTestSuite))
}

func (s *buyerUsecaseTestSuite) SetupSuite() {
	instance := mongodb.ConfigureMongo()
	counterRepo := repoMongo.NewCounterRepo(instance)
	buyerRepo := repoMongo.NewBuyerRepo(instance, counterRepo)
	s.BuyerUsecase = NewBuyerUsecase(buyerRepo)
	s.Instance = instance
}

func (s *buyerUsecaseTestSuite) TearDownSuite() {
	ctx, cancel := context.WithTimeout(context.Background(), configs.Constant.TimeoutOnSeconds*time.Second)
	defer cancel()

	err := testUtil.DropBuyer(ctx, s.Instance)
	testUtil.HandleError(err)
	err = testUtil.DropCounter(ctx, s.Instance)
	testUtil.HandleError(err)
}

func (s *buyerUsecaseTestSuite) TearDownTest() {
	ctx, cancel := context.WithTimeout(context.Background(), configs.Constant.TimeoutOnSeconds*time.Second)
	defer cancel()

	err := testUtil.DropBuyer(ctx, s.Instance)
	testUtil.HandleError(err)
	err = testUtil.DropCounter(ctx, s.Instance)
	testUtil.HandleError(err)
}

func (s *buyerUsecaseTestSuite) TestCreate() {
	s.Run("Create Buyer from body", func() {
		body := models.BuyerBody{
			Email:           "test",
			Name:            "test",
			Password:        "test",
			DeliveryAddress: "test",
		}

		id, err := s.BuyerUsecase.CreateBuyer(body)
		testUtil.HandleError(err)

		result, err := s.BuyerUsecase.GetByID(id)
		testUtil.HandleError(err)

		s.Assert().EqualValues(1, id)
		s.Assert().Equal(body.Email, result.Email)
		s.Assert().Equal(body.Name, result.Name)
		s.Assert().Equal(body.DeliveryAddress, result.DeliveryAddress)
		s.Assert().NotEmpty(result.Password)
	})
}
