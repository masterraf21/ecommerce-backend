package mongodb

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/masterraf21/ecommerce-backend/configs"
	"github.com/masterraf21/ecommerce-backend/models"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/mongo"
)

type orderRepoTestSuite struct {
	suite.Suite
	Instance  *mongo.Database
	OrderRepo models.OrderRepository
}

func TestOrderRepository(t *testing.T) {
	suite.Run(t, new(orderRepoTestSuite))
}

func (s *orderRepoTestSuite) SetupSuite() {
	instance := configureMongo()
	s.Instance = instance
	counterRepo := NewCounterRepo(instance)
	s.OrderRepo = NewOrderRepo(instance, counterRepo)
}

func (s *orderRepoTestSuite) TearDownTest() {
	ctx, cancel := context.WithTimeout(context.Background(), configs.Constant.TimeoutOnSeconds*time.Second)
	defer cancel()
	orderCollection := s.Instance.Collection("order")

	err := orderCollection.Drop(ctx)
	handleError(err)
}

func (s *orderRepoTestSuite) TearDownSuite() {
	ctx, cancel := context.WithTimeout(context.Background(), configs.Constant.TimeoutOnSeconds*time.Second)
	defer cancel()
	orderCollection := s.Instance.Collection("order")

	err := orderCollection.Drop(ctx)
	handleError(err)
}

func (s *orderRepoTestSuite) TestStore() {
	s.Run("Store a single Order data", func() {
		buyer := models.Buyer{
			ID:              1,
			Email:           "test",
			Name:            "test",
			Password:        "test",
			DeliveryAddress: "test",
		}

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

		order := models.Order{
			BuyerID:         buyer.ID,
			SellerID:        seller.ID,
			Seller:          &seller,
			SourceAddress:   "test",
			DeliveryAddress: "test",
			Items: []*models.Product{
				&product,
				&product,
				&product,
			},
			Quantity:   3,
			Price:      25.123,
			TotalPrice: 100.123,
			Status:     "test",
		}

		oid, err := s.OrderRepo.Store(&order)
		handleError(err)

		result, err := s.OrderRepo.GetByOID(oid)
		handleError(err)

		s.Assert().EqualValues(1, result.ID)
		s.Assert().Nil(result.Buyer)
		s.Assert().True(reflect.DeepEqual(seller, *result.Seller))
	})
}

func (s *orderRepoTestSuite) TestGet() {
	s.Run("Get Order by ID", func() {
		buyer := models.Buyer{
			ID:              1,
			Email:           "test",
			Name:            "test",
			Password:        "test",
			DeliveryAddress: "test",
		}

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

		order := models.Order{
			BuyerID:         buyer.ID,
			Buyer:           &buyer,
			SellerID:        seller.ID,
			Seller:          &seller,
			SourceAddress:   "test",
			DeliveryAddress: "test",
			Items: []*models.Product{
				&product,
				&product,
				&product,
			},
			Quantity:   3,
			Price:      25.123,
			TotalPrice: 100.123,
			Status:     "test",
		}

		_, err := s.OrderRepo.Store(&order)
		handleError(err)

		result, err := s.OrderRepo.GetByID(uint32(1))
		handleError(err)

		s.Assert().EqualValues(1, result.ID)
		s.Assert().True(reflect.DeepEqual(buyer, *result.Buyer))
		s.Assert().True(reflect.DeepEqual(seller, *result.Seller))
	})

	s.Run("Get All Order data", func() {
		buyer := models.Buyer{
			ID:              1,
			Email:           "test",
			Name:            "test",
			Password:        "test",
			DeliveryAddress: "test",
		}

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

		order := models.Order{
			BuyerID:         buyer.ID,
			Buyer:           &buyer,
			SellerID:        seller.ID,
			Seller:          &seller,
			SourceAddress:   "test",
			DeliveryAddress: "test",
			Items: []*models.Product{
				&product,
				&product,
				&product,
			},
			Quantity:   3,
			Price:      25.123,
			TotalPrice: 100.123,
			Status:     "test",
		}

		_, err := s.OrderRepo.Store(&order)
		handleError(err)

		_, err = s.OrderRepo.Store(&order)
		handleError(err)

		_, err = s.OrderRepo.Store(&order)
		handleError(err)

		result, err := s.OrderRepo.GetAll()
		handleError(err)

		s.Assert().Equal(4, len(result))
	})
}

func (s *orderRepoTestSuite) TestGet2() {
	s.Run("Get Order By Buyer ID and Status", func() {
		buyer := models.Buyer{
			ID:              1,
			Email:           "test",
			Name:            "test",
			Password:        "test",
			DeliveryAddress: "test",
		}

		buyer2 := models.Buyer{
			ID:              2,
			Email:           "test",
			Name:            "test",
			Password:        "test",
			DeliveryAddress: "test",
		}

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

		order := models.Order{
			BuyerID:         buyer.ID,
			Buyer:           &buyer,
			SellerID:        seller.ID,
			Seller:          &seller,
			SourceAddress:   "test",
			DeliveryAddress: "test",
			Items: []*models.Product{
				&product,
				&product,
				&product,
			},
			Quantity:   3,
			Price:      25.123,
			TotalPrice: 100.123,
			Status:     "Pending",
		}

		order2 := models.Order{
			BuyerID:         buyer2.ID,
			Buyer:           &buyer2,
			SellerID:        seller.ID,
			Seller:          &seller,
			SourceAddress:   "test",
			DeliveryAddress: "test",
			Items: []*models.Product{
				&product,
				&product,
				&product,
			},
			Quantity:   3,
			Price:      25.123,
			TotalPrice: 100.123,
			Status:     "Completed",
		}

		order3 := models.Order{
			BuyerID:         buyer2.ID,
			Buyer:           &buyer2,
			SellerID:        seller.ID,
			Seller:          &seller,
			SourceAddress:   "test",
			DeliveryAddress: "test",
			Items: []*models.Product{
				&product,
				&product,
				&product,
			},
			Quantity:   3,
			Price:      25.123,
			TotalPrice: 100.123,
			Status:     "Pending",
		}

		_, err := s.OrderRepo.Store(&order)
		handleError(err)

		_, err = s.OrderRepo.Store(&order2)
		handleError(err)

		_, err = s.OrderRepo.Store(&order3)
		handleError(err)

		result, err := s.OrderRepo.GetByBuyerIDAndStatus(buyer2.ID, "Completed")
		handleError(err)

		s.Assert().Equal(1, len(result))
	})
}

func (s *orderRepoTestSuite) TestUpdate() {
	s.Run("Update arbitrary Order data field", func() {
		buyer := models.Buyer{
			ID:              1,
			Email:           "test",
			Name:            "test",
			Password:        "test",
			DeliveryAddress: "test",
		}

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

		order := models.Order{
			BuyerID:         buyer.ID,
			Buyer:           &buyer,
			SellerID:        seller.ID,
			Seller:          &seller,
			SourceAddress:   "test",
			DeliveryAddress: "test",
			Items: []*models.Product{
				&product,
				&product,
				&product,
			},
			Quantity:   3,
			Price:      25.123,
			TotalPrice: 100.123,
			Status:     "test",
		}

		oid, err := s.OrderRepo.Store(&order)
		handleError(err)

		err = s.OrderRepo.UpdateArbitrary(uint32(1), "source_address", "update")
		handleError(err)

		result, err := s.OrderRepo.GetByOID(oid)
		handleError(err)

		s.Assert().Equal("update", result.SourceAddress)
	})
}
