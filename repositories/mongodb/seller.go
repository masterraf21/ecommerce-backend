package mongodb

import (
	"context"
	"strings"
	"time"

	"github.com/masterraf21/ecommerce-backend/configs"
	"github.com/masterraf21/ecommerce-backend/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

type sellerRepo struct {
	Instance    *mongo.Database
	CounterRepo models.CounterRepository
}

// NewSellerRepo will initiate object representing SellerRepository
func NewSellerRepo(instance *mongo.Database, ctr models.CounterRepository) models.SellerRepository {
	return &sellerRepo{Instance: instance, CounterRepo: ctr}
}

func (r *sellerRepo) Store(seller *models.Seller) (oid primitive.ObjectID, err error) {
	collectionName := "seller"
	identifier := "id_seller"

	ctx, cancel := context.WithTimeout(context.Background(), configs.Constant.TimeoutOnSeconds*time.Second)
	defer cancel()

	id, err := r.CounterRepo.Get(collectionName, identifier)
	if err != nil {
		return
	}

	collection := r.Instance.Collection(collectionName)
	seller.ID = id

	result, err := collection.InsertOne(ctx, seller)
	if err != nil {
		return
	}

	_id := result.InsertedID
	oid = _id.(primitive.ObjectID)

	return
}

func (r *sellerRepo) GetAll() (res []models.Seller, err error) {
	collection := r.Instance.Collection("seller")

	ctx, cancel := context.WithTimeout(context.Background(), configs.Constant.TimeoutOnSeconds*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return
	}

	if err = cursor.All(ctx, &res); err != nil {
		return
	}

	return
}

func (r *sellerRepo) GetByID(id uint32) (res *models.Seller, err error) {
	collection := r.Instance.Collection("seller")

	ctx, cancel := context.WithTimeout(context.Background(), configs.Constant.TimeoutOnSeconds*time.Second)
	defer cancel()

	err = collection.FindOne(ctx, bson.M{"id_seller": id}).Decode(&res)
	if err != nil {
		if strings.Contains(err.Error(), "mongo: no documents") {
			err = nil
			return
		}

		return
	}

	return
}

func (r *sellerRepo) GetByOID(oid primitive.ObjectID) (res *models.Seller, err error) {
	collection := r.Instance.Collection("seller")

	ctx, cancel := context.WithTimeout(context.Background(), configs.Constant.TimeoutOnSeconds*time.Second)
	defer cancel()

	err = collection.FindOne(ctx, bson.M{"_id": oid}).Decode(&res)
	if err != nil {
		if strings.Contains(err.Error(), "mongo: no documents") {
			err = nil
			return
		}

		return
	}

	return
}

func (r *sellerRepo) UpdateArbitrary(id uint32, key string, value interface{}) error {
	collection := r.Instance.Collection("seller")

	ctx, cancel := context.WithTimeout(context.Background(), configs.Constant.TimeoutOnSeconds*time.Second)
	defer cancel()

	_, err := collection.UpdateOne(
		ctx,
		bson.M{"id_seller": id},
		bson.M{"$set": bson.M{key: value}},
	)
	if err != nil {
		return err
	}

	return nil
}
