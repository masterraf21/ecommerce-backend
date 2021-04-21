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

type buyerRepo struct {
	Instance    *mongo.Database
	CounterRepo models.CounterRepository
}

// NewBuyerRepo will create an object representing BuyerRepository
func NewBuyerRepo(instance *mongo.Database, ctr models.CounterRepository) models.BuyerRepository {
	return &buyerRepo{Instance: instance, CounterRepo: ctr}
}

func (r *buyerRepo) Store(buyer *models.Buyer) (uid primitive.ObjectID, err error) {
	collectionName := "buyer"
	identifier := "id_buyer"

	ctx, cancel := context.WithTimeout(context.Background(), configs.Constant.TimeoutOnSeconds*time.Second)
	defer cancel()

	id, err := r.CounterRepo.Get(collectionName, identifier)
	if err != nil {
		return
	}

	collection := r.Instance.Collection(collectionName)
	buyer.ID = id

	result, err := collection.InsertOne(ctx, buyer)
	if err != nil {
		return
	}

	_id := result.InsertedID
	uid = _id.(primitive.ObjectID)

	return
}

func (r *buyerRepo) GetAll() (res []models.Buyer, err error) {
	collection := r.Instance.Collection("buyer")

	ctx, cancel := context.WithTimeout(context.Background(), configs.Constant.TimeoutOnSeconds*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		if strings.Contains(err.Error(), "mongo: no documents") {
			res = make([]models.Buyer, 0)
			err = nil
			return
		}

		return
	}

	if err = cursor.All(ctx, &res); err != nil {
		return
	}

	return
}

func (r *buyerRepo) GetByOID(oid primitive.ObjectID) (res *models.Buyer, err error) {
	collection := r.Instance.Collection("buyer")

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

func (r *buyerRepo) GetByID(id uint32) (res *models.Buyer, err error) {
	collection := r.Instance.Collection("buyer")

	ctx, cancel := context.WithTimeout(context.Background(), configs.Constant.TimeoutOnSeconds*time.Second)
	defer cancel()

	err = collection.FindOne(ctx, bson.M{"id_buyer": id}).Decode(&res)
	if err != nil {
		if strings.Contains(err.Error(), "mongo: no documents") {
			err = nil
			return
		}

		return
	}

	return
}

func (r *buyerRepo) UpdateArbitrary(id uint32, key string, value interface{}) error {
	collection := r.Instance.Collection("buyer")

	ctx, cancel := context.WithTimeout(context.Background(), configs.Constant.TimeoutOnSeconds*time.Second)
	defer cancel()

	_, err := collection.UpdateOne(
		ctx,
		bson.M{"id_buyer": id},
		bson.M{"$set": bson.M{key: value}},
	)
	if err != nil {
		return err
	}

	return nil
}
