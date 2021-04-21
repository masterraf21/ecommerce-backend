package mongodb

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/masterraf21/ecommerce-backend/configs"
	"github.com/masterraf21/ecommerce-backend/models"
	"github.com/masterraf21/ecommerce-backend/utils/mongodb"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/mongo"
)

type productRepoTestSuite struct {
	suite.Suite
	Instance    *mongo.Database
	ProductRepo models.ProductRepository
}

func TestProductRepository(t *testing.T) {
	suite.Run(t, new(productRepoTestSuite))
}

func (s *productRepoTestSuite) SetupSuite() {
	instance := mongodb.ConfigureMongo()
	s.Instance = instance
	counterRepo := NewCounterRepo(instance)
	s.ProductRepo = NewProductRepo(instance, counterRepo)
}

func (s *productRepoTestSuite) TearDownTest() {
	ctx, cancel := context.WithTimeout(context.Background(), configs.Constant.TimeoutOnSeconds*time.Second)
	defer cancel()
	productCollection := s.Instance.Collection("product")

	err := productCollection.Drop(ctx)
	handleError(err)
}

func (s *productRepoTestSuite) TearDownSuite() {
	ctx, cancel := context.WithTimeout(context.Background(), configs.Constant.TimeoutOnSeconds*time.Second)
	defer cancel()
	productCollection := s.Instance.Collection("product")

	err := productCollection.Drop(ctx)
	handleError(err)
}

func (s *productRepoTestSuite) TestStore() {
	s.Run("Store a single Product data", func() {
		seller := models.Seller{
			ID:            1,
			Email:         "test",
			Name:          "test",
			Password:      "test",
			PickupAddress: "test",
		}

		product := models.Product{
			ProductName: "test",
			Description: "test",
			Price:       10.11,
			SellerID:    seller.ID,
			Seller:      &seller,
		}

		oid, err := s.ProductRepo.Store(&product)
		handleError(err)

		result, err := s.ProductRepo.GetByOID(oid)
		s.Require().NoError(err)
		s.Assert().EqualValues(1, result.ID)
		s.Assert().Equal(product.ProductName, result.ProductName)
		s.Assert().True(reflect.DeepEqual(seller, *result.Seller))
	})

	s.Run("Store a single Product data after existing data stored", func() {
		seller := models.Seller{
			ID:            1,
			Email:         "test",
			Name:          "test",
			Password:      "test",
			PickupAddress: "test",
		}

		product := models.Product{
			ProductName: "test",
			Description: "test",
			Price:       10.11,
			SellerID:    seller.ID,
			Seller:      &seller,
		}

		oid, err := s.ProductRepo.Store(&product)
		handleError(err)

		result, err := s.ProductRepo.GetByOID(oid)
		s.Require().NoError(err)
		s.Assert().EqualValues(2, result.ID)
		s.Assert().Equal(product.ProductName, result.ProductName)
		s.Assert().True(reflect.DeepEqual(seller, *result.Seller))
	})
}

func (s *productRepoTestSuite) TestGet() {
	s.Run("Get a Product data by ID", func() {
		seller := models.Seller{
			ID:            1,
			Email:         "test",
			Name:          "test",
			Password:      "test",
			PickupAddress: "test",
		}

		product := models.Product{
			ProductName: "test",
			Description: "test",
			Price:       10.11,
			SellerID:    seller.ID,
			Seller:      &seller,
		}

		_, err := s.ProductRepo.Store(&product)
		handleError(err)

		result, err := s.ProductRepo.GetByID(uint32(1))
		s.Require().NoError(err)
		s.Assert().EqualValues(1, result.ID)
		s.Assert().Equal(product.ProductName, result.ProductName)
		s.Assert().True(reflect.DeepEqual(seller, *result.Seller))
	})

	s.Run("Get all Product data", func() {
		seller := models.Seller{
			ID:            1,
			Email:         "test",
			Name:          "test",
			Password:      "test",
			PickupAddress: "test",
		}

		product := models.Product{
			ProductName: "test",
			Description: "test",
			Price:       10.11,
			SellerID:    seller.ID,
			Seller:      &seller,
		}

		_, err := s.ProductRepo.Store(&product)
		handleError(err)

		_, err = s.ProductRepo.Store(&product)
		handleError(err)

		_, err = s.ProductRepo.Store(&product)
		handleError(err)

		result, err := s.ProductRepo.GetAll()
		handleError(err)

		s.Assert().Equal(4, len(result))
	})
}

func (s *productRepoTestSuite) TestGet2() {
	s.Run("Get By Seller ID", func() {
		seller1 := models.Seller{
			ID:            1,
			Email:         "test",
			Name:          "test",
			Password:      "test",
			PickupAddress: "test",
		}

		seller2 := models.Seller{
			ID:            2,
			Email:         "test",
			Name:          "test",
			Password:      "test",
			PickupAddress: "test",
		}

		product1 := models.Product{
			ProductName: "test",
			Description: "test",
			Price:       10.11,
			SellerID:    seller1.ID,
			Seller:      &seller1,
		}

		product2 := models.Product{
			ProductName: "test",
			Description: "test",
			Price:       10.11,
			SellerID:    seller2.ID,
			Seller:      &seller2,
		}

		_, err := s.ProductRepo.Store(&product1)
		handleError(err)

		_, err = s.ProductRepo.Store(&product1)
		handleError(err)

		_, err = s.ProductRepo.Store(&product1)
		handleError(err)

		_, err = s.ProductRepo.Store(&product2)
		handleError(err)

		_, err = s.ProductRepo.Store(&product2)
		handleError(err)

		result, err := s.ProductRepo.GetBySellerID(uint32(2))
		handleError(err)

		s.Assert().Equal(2, len(result))
	})
}

func (s *productRepoTestSuite) TestUpdate() {
	s.Run("Update a product data arbitrarily", func() {
		seller := models.Seller{
			ID:            1,
			Email:         "test",
			Name:          "test",
			Password:      "test",
			PickupAddress: "test",
		}

		product := models.Product{
			ProductName: "test",
			Description: "test",
			Price:       10.11,
			SellerID:    seller.ID,
			Seller:      &seller,
		}

		oid, err := s.ProductRepo.Store(&product)
		handleError(err)

		err = s.ProductRepo.UpdateArbitrary(uint32(1), "price", float32(99.9))
		handleError(err)

		result, err := s.ProductRepo.GetByOID(oid)
		handleError(err)

		s.Require().Equal(float32(99.9), result.Price)
	})
}
