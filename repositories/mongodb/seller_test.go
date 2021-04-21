package mongodb

import (
	"context"
	"testing"
	"time"

	"github.com/masterraf21/ecommerce-backend/configs"
	"github.com/masterraf21/ecommerce-backend/models"
	"github.com/masterraf21/ecommerce-backend/utils/mongodb"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/mongo"
)

type sellerRepoTestSuite struct {
	suite.Suite
	Instance   *mongo.Database
	SellerRepo models.SellerRepository
}

func TestSellerRepository(t *testing.T) {
	suite.Run(t, new(sellerRepoTestSuite))
}

func (s *sellerRepoTestSuite) SetupSuite() {
	instance := mongodb.ConfigureMongo()
	s.Instance = instance
	counterRepo := NewCounterRepo(instance)
	s.SellerRepo = NewSellerRepo(instance, counterRepo)
}

func (s *sellerRepoTestSuite) TearDownTest() {
	ctx, cancel := context.WithTimeout(context.Background(), configs.Constant.TimeoutOnSeconds*time.Second)
	defer cancel()
	sellerCollection := s.Instance.Collection("seller")

	err := sellerCollection.Drop(ctx)
	handleError(err)
}

func (s *sellerRepoTestSuite) TearDownSuite() {
	ctx, cancel := context.WithTimeout(context.Background(), configs.Constant.TimeoutOnSeconds*time.Second)
	defer cancel()
	sellerCollection := s.Instance.Collection("seller")

	err := sellerCollection.Drop(ctx)
	handleError(err)
}

func (s *sellerRepoTestSuite) TestStore() {
	s.Run("Store a single Seller Data", func() {
		seller := models.Seller{
			Email:         "test",
			Name:          "test",
			Password:      "test",
			PickupAddress: "test",
		}

		oid, err := s.SellerRepo.Store(&seller)
		handleError(err)

		result, err := s.SellerRepo.GetByOID(oid)
		s.Require().NoError(err)
		s.Assert().EqualValues(1, result.ID)
		s.Assert().Equal(seller.Email, result.Email)
		s.Assert().Equal(seller.Name, result.Name)
		s.Assert().Equal(seller.Password, result.Password)
		s.Assert().Equal(seller.PickupAddress, result.PickupAddress)
	})

	s.Run("Store a single Seller Data after existing data stored", func() {
		seller := models.Seller{
			Email:         "test",
			Name:          "test",
			Password:      "test",
			PickupAddress: "test",
		}

		oid, err := s.SellerRepo.Store(&seller)
		handleError(err)

		result, err := s.SellerRepo.GetByOID(oid)
		s.Require().NoError(err)
		s.Assert().EqualValues(2, result.ID)
		s.Assert().Equal(seller.Email, result.Email)
		s.Assert().Equal(seller.Name, result.Name)
		s.Assert().Equal(seller.Password, result.Password)
		s.Assert().Equal(seller.PickupAddress, result.PickupAddress)
	})
}

func (s *sellerRepoTestSuite) TestGet1() {
	s.Run("Get a Seller Data by ID", func() {
		seller := models.Seller{
			Email:         "test",
			Name:          "test",
			Password:      "test",
			PickupAddress: "test",
		}

		_, err := s.SellerRepo.Store(&seller)
		handleError(err)

		result, err := s.SellerRepo.GetByID(uint32(1))
		s.Require().NoError(err)
		s.Assert().EqualValues(1, result.ID)
		s.Assert().Equal(seller.Email, result.Email)
		s.Assert().Equal(seller.Name, result.Name)
		s.Assert().Equal(seller.Password, result.Password)
		s.Assert().Equal(seller.PickupAddress, result.PickupAddress)
	})

	s.Run("Get all Seller Data", func() {
		seller := models.Seller{
			Email:         "test",
			Name:          "test",
			Password:      "test",
			PickupAddress: "test",
		}

		_, err := s.SellerRepo.Store(&seller)
		handleError(err)

		_, err = s.SellerRepo.Store(&seller)
		handleError(err)

		_, err = s.SellerRepo.Store(&seller)
		handleError(err)

		result, err := s.SellerRepo.GetAll()
		handleError(err)

		s.Assert().Equal(4, len(result))
	})
}

func (s *sellerRepoTestSuite) TestUpdate() {
	s.Run("Update a seller data arbitrarily", func() {
		seller := models.Seller{
			Email:         "test",
			Name:          "test",
			Password:      "test",
			PickupAddress: "test",
		}

		oid, err := s.SellerRepo.Store(&seller)
		handleError(err)

		err = s.SellerRepo.UpdateArbitrary(uint32(1), "name", "update")
		handleError(err)

		result, err := s.SellerRepo.GetByOID(oid)
		handleError(err)
		s.Require().NoError(err)
		s.Require().Equal("update", result.Name)
	})
}
