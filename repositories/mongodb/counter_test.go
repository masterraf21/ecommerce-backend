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
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

type counterRepoTestSuite struct {
	suite.Suite
	Instance    *mongo.Database
	CounterRepo models.CounterRepository
}

func TestCounterRepository(t *testing.T) {
	suite.Run(t, new(counterRepoTestSuite))
}

func (s *counterRepoTestSuite) SetupSuite() {
	instance := mongodb.ConfigureMongo()
	s.Instance = instance
	s.CounterRepo = NewCounterRepo(instance)
}

func (s *counterRepoTestSuite) TearDownTest() {
	ctx, cancel := context.WithTimeout(context.Background(), configs.Constant.TimeoutOnSeconds*time.Second)
	defer cancel()
	buyerCollection := s.Instance.Collection("buyer")

	err := buyerCollection.Drop(ctx)
	if err != nil {
		panic(err)
	}
}

func (s *counterRepoTestSuite) TestGetEmpty() {
	s.Run("Get id with empty document", func() {
		collectionName := "buyer"
		identifier := "id_buyer"

		id, err := s.CounterRepo.Get(collectionName, identifier)
		handleError(err)

		s.Equal(uint32(1), id)
	})
}

func (s *counterRepoTestSuite) TestGetExisting() {
	s.Run("Get id with existing document", func() {
		ctx, cancel := context.WithTimeout(context.Background(), configs.Constant.TimeoutOnSeconds*time.Second)
		defer cancel()

		collectionName := "buyer"
		identifier := "id_buyer"

		initialID := uint32(2)
		collection := s.Instance.Collection(collectionName)
		buyer := models.Buyer{
			ID:              initialID,
			Email:           "test",
			Name:            "test",
			Password:        "test",
			DeliveryAddress: "test",
		}
		_, err := collection.UpdateOne(
			ctx,
			bson.M{identifier: buyer.ID},
			bson.M{"$set": buyer},
			options.Update().SetUpsert(true),
		)
		handleError(err)

		id, err := s.CounterRepo.Get(collectionName, identifier)
		handleError(err)

		s.Equal(initialID+1, id)
	})
}
