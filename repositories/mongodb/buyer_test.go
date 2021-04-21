package mongodb

import (
	"context"
	"testing"
	"time"

	"github.com/masterraf21/ecommerce-backend/configs"
	"github.com/masterraf21/ecommerce-backend/models"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/mongo"
)

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

type buyerRepoTestSuite struct {
	suite.Suite
	Instance  *mongo.Database
	BuyerRepo models.BuyerRepository
}

func TestBuyerRepository(t *testing.T) {
	suite.Run(t, new(buyerRepoTestSuite))
}

func (s *buyerRepoTestSuite) SetupSuite() {
	instance := configureMongo()
	s.Instance = instance
	counterRepo := NewCounterRepo(instance)
	s.BuyerRepo = NewBuyerRepo(instance, counterRepo)
}

func (s *buyerRepoTestSuite) TearDownTest() {
	ctx, cancel := context.WithTimeout(context.Background(), configs.Constant.TimeoutOnSeconds*time.Second)
	defer cancel()
	buyerCollection := s.Instance.Collection("buyer")

	err := buyerCollection.Drop(ctx)
	handleError(err)
}

func (s *buyerRepoTestSuite) TearDownSuite() {
	ctx, cancel := context.WithTimeout(context.Background(), configs.Constant.TimeoutOnSeconds*time.Second)
	defer cancel()
	buyerCollection := s.Instance.Collection("buyer")

	err := buyerCollection.Drop(ctx)
	handleError(err)
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
		handleError(err)

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
		handleError(err)

		result, err := s.BuyerRepo.GetByOID(oid)
		handleError(err)
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
		handleError(err)

		_, err = s.BuyerRepo.Store(&buyer)
		handleError(err)

		err = s.BuyerRepo.UpdateArbitrary(uint32(1), "name", "update")
		handleError(err)

		oid, err := s.BuyerRepo.Store(&buyer)
		handleError(err)

		result, err := s.BuyerRepo.GetByOID(oid)
		handleError(err)
		s.Require().NoError(err)
		s.Assert().EqualValues(3, result.ID)
		s.Assert().Equal(buyer.Email, result.Email)
		s.Assert().Equal(buyer.Name, result.Name)
		s.Assert().Equal(buyer.Password, result.Password)
		s.Assert().Equal(buyer.DeliveryAddress, result.DeliveryAddress)
	})
}

func (s *buyerRepoTestSuite) TestGet() {
	s.Run("Get a Buyer Data by ID", func() {
		buyer := models.Buyer{
			Email:           "test",
			Name:            "test",
			Password:        "test",
			DeliveryAddress: "test",
		}
		_, err := s.BuyerRepo.Store(&buyer)
		handleError(err)

		result, err := s.BuyerRepo.GetByID(uint32(1))
		handleError(err)
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
		handleError(err)

		_, err = s.BuyerRepo.Store(&buyer)
		handleError(err)

		_, err = s.BuyerRepo.Store(&buyer)
		handleError(err)

		result, err := s.BuyerRepo.GetAll()
		handleError(err)

		s.Assert().Equal(4, len(result))
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
		handleError(err)

		err = s.BuyerRepo.UpdateArbitrary(uint32(1), "name", "update")
		handleError(err)

		result, err := s.BuyerRepo.GetByOID(oid)
		handleError(err)
		s.Require().NoError(err)
		s.Require().Equal("update", result.Name)
	})
}
