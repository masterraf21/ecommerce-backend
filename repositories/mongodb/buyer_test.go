package mongodb

import (
	"context"
	"testing"
	"time"

	"github.com/masterraf21/ecommerce-backend/configs"
	"github.com/masterraf21/ecommerce-backend/models"
	"github.com/masterraf21/ecommerce-backend/utils/mongodb"
	testUtil "github.com/masterraf21/ecommerce-backend/utils/test"

	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/mongo"
)

type buyerRepoTestSuite struct {
	suite.Suite
	Instance  *mongo.Database
	BuyerRepo models.BuyerRepository
}

func TestBuyerRepository(t *testing.T) {
	suite.Run(t, new(buyerRepoTestSuite))
}

func (s *buyerRepoTestSuite) SetupSuite() {
	instance := mongodb.ConfigureMongo()
	s.Instance = instance
	counterRepo := NewCounterRepo(instance)
	s.BuyerRepo = NewBuyerRepo(instance, counterRepo)
}

func (s *buyerRepoTestSuite) TearDownTest() {
	ctx, cancel := context.WithTimeout(context.Background(), configs.Constant.TimeoutOnSeconds*time.Second)
	defer cancel()

	err := testUtil.DropBuyer(ctx, s.Instance)
	testUtil.HandleError(err)
	err = testUtil.DropCounter(ctx, s.Instance)
	testUtil.HandleError(err)
}

func (s *buyerRepoTestSuite) TearDownSuite() {
	ctx, cancel := context.WithTimeout(context.Background(), configs.Constant.TimeoutOnSeconds*time.Second)
	defer cancel()

	err := testUtil.DropBuyer(ctx, s.Instance)
	testUtil.HandleError(err)
	err = testUtil.DropCounter(ctx, s.Instance)
	testUtil.HandleError(err)
}

func (s *buyerRepoTestSuite) TestStore1() {
	s.Run("Store a single Buyer data", func() {
		buyer := models.Buyer{
			Email:           "test",
			Name:            "test",
			Password:        "test",
			DeliveryAddress: "test",
		}

		oid, err := s.BuyerRepo.Store(&buyer)
		testUtil.HandleError(err)

		result, err := s.BuyerRepo.GetByOID(oid)
		s.Require().NoError(err)
		s.Assert().EqualValues(1, result.ID)
		s.Assert().Equal(buyer.Email, result.Email)
		s.Assert().Equal(buyer.Name, result.Name)
		s.Assert().Equal(buyer.Password, result.Password)
		s.Assert().Equal(buyer.DeliveryAddress, result.DeliveryAddress)
	})

	s.Run("Store a single Buyer data after existing data stored", func() {
		buyer := models.Buyer{
			Email:           "test",
			Name:            "test",
			Password:        "test",
			DeliveryAddress: "test",
		}

		oid, err := s.BuyerRepo.Store(&buyer)
		testUtil.HandleError(err)

		result, err := s.BuyerRepo.GetByOID(oid)
		testUtil.HandleError(err)
		s.Require().NoError(err)
		s.Assert().EqualValues(2, result.ID)
		s.Assert().Equal(buyer.Email, result.Email)
		s.Assert().Equal(buyer.Name, result.Name)
		s.Assert().Equal(buyer.Password, result.Password)
		s.Assert().Equal(buyer.DeliveryAddress, result.DeliveryAddress)
	})
}

func (s *buyerRepoTestSuite) TestStore2() {
	s.Run("Store a single Buyer data after previous stored data updated", func() {
		buyer := models.Buyer{
			Email:           "test",
			Name:            "test",
			Password:        "test",
			DeliveryAddress: "test",
		}
		_, err := s.BuyerRepo.Store(&buyer)
		testUtil.HandleError(err)

		_, err = s.BuyerRepo.Store(&buyer)
		testUtil.HandleError(err)

		err = s.BuyerRepo.UpdateArbitrary(uint32(1), "name", "update")
		testUtil.HandleError(err)

		oid, err := s.BuyerRepo.Store(&buyer)
		testUtil.HandleError(err)

		result, err := s.BuyerRepo.GetByOID(oid)
		testUtil.HandleError(err)
		s.Require().NoError(err)
		s.Assert().EqualValues(3, result.ID)
		s.Assert().Equal(buyer.Email, result.Email)
		s.Assert().Equal(buyer.Name, result.Name)
		s.Assert().Equal(buyer.Password, result.Password)
		s.Assert().Equal(buyer.DeliveryAddress, result.DeliveryAddress)
	})
}

func (s *buyerRepoTestSuite) TestGet1() {
	s.Run("Get a Buyer Data by ID", func() {
		buyer := models.Buyer{
			Email:           "test",
			Name:            "test",
			Password:        "test",
			DeliveryAddress: "test",
		}
		_, err := s.BuyerRepo.Store(&buyer)
		testUtil.HandleError(err)

		result, err := s.BuyerRepo.GetByID(uint32(1))
		testUtil.HandleError(err)
		s.Require().NoError(err)
		s.Assert().EqualValues(1, result.ID)
		s.Assert().Equal(buyer.Email, result.Email)
		s.Assert().Equal(buyer.Name, result.Name)
		s.Assert().Equal(buyer.Password, result.Password)
		s.Assert().Equal(buyer.DeliveryAddress, result.DeliveryAddress)
	})

	s.Run("Get all Buyer data", func() {
		buyer := models.Buyer{
			Email:           "test",
			Name:            "test",
			Password:        "test",
			DeliveryAddress: "test",
		}
		_, err := s.BuyerRepo.Store(&buyer)
		testUtil.HandleError(err)

		_, err = s.BuyerRepo.Store(&buyer)
		testUtil.HandleError(err)

		_, err = s.BuyerRepo.Store(&buyer)
		testUtil.HandleError(err)

		result, err := s.BuyerRepo.GetAll()
		testUtil.HandleError(err)

		s.Assert().Equal(4, len(result))
	})
}

func (s *buyerRepoTestSuite) TestGet2() {
	s.Run("Get Empty Data", func() {
		result, err := s.BuyerRepo.GetByID(uint32(1))
		testUtil.HandleError(err)
		s.Assert().Nil(result)
	})
}

func (s *buyerRepoTestSuite) TestUpdate() {
	s.Run("Update a buyer Data arbitrarily", func() {
		buyer := models.Buyer{
			Email:           "test",
			Name:            "test",
			Password:        "test",
			DeliveryAddress: "test",
		}
		oid, err := s.BuyerRepo.Store(&buyer)
		testUtil.HandleError(err)

		err = s.BuyerRepo.UpdateArbitrary(uint32(1), "name", "update")
		testUtil.HandleError(err)

		result, err := s.BuyerRepo.GetByOID(oid)
		testUtil.HandleError(err)
		s.Require().NoError(err)
		s.Require().Equal("update", result.Name)
	})
}
