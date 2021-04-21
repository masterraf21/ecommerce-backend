package mongodb

import (
	"context"
	"time"

	"github.com/masterraf21/ecommerce-backend/configs"
	"github.com/masterraf21/ecommerce-backend/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

type orderRepo struct {
	Instance    *mongo.Database
	CounterRepo models.CounterRepository
}

// NewOrderRepo will initiate order repository object
func NewOrderRepo(instance *mongo.Database, ctr models.CounterRepository) models.OrderRepository {
	return &orderRepo{Instance: instance, CounterRepo: ctr}
}

func (r *orderRepo) Store(order *models.Order) (oid primitive.ObjectID, err error) {
	collectionName := "order"
	identifier := "id_order"

	ctx, cancel := context.WithTimeout(context.Background(), configs.Constant.TimeoutOnSeconds*time.Second)
	defer cancel()

	id, err := r.CounterRepo.Get(collectionName, identifier)
	if err != nil {
		return
	}

	collection := r.Instance.Collection(collectionName)
	order.ID = id

	result, err := collection.InsertOne(ctx, order)
	if err != nil {
		return
	}

	_id := result.InsertedID
	oid = _id.(primitive.ObjectID)

	return
}

func (r *orderRepo) GetAll() (res []models.Order, err error) {
	collection := r.Instance.Collection("order")

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

func (r *orderRepo) GetByID(id uint32) (res *models.Order, err error) {
	collection := r.Instance.Collection("order")

	ctx, cancel := context.WithTimeout(context.Background(), configs.Constant.TimeoutOnSeconds*time.Second)
	defer cancel()

	err = collection.FindOne(ctx, bson.M{"id_order": id}).Decode(&res)
	if err != nil {
		return
	}

	return
}

func (r *orderRepo) GetByOID(oid primitive.ObjectID) (res *models.Order, err error) {
	collection := r.Instance.Collection("order")

	ctx, cancel := context.WithTimeout(context.Background(), configs.Constant.TimeoutOnSeconds*time.Second)
	defer cancel()

	err = collection.FindOne(ctx, bson.M{"_id": oid}).Decode(&res)
	if err != nil {
		return
	}

	return
}

func (r *orderRepo) UpdateArbitrary(id uint32, key string, value interface{}) error {
	collection := r.Instance.Collection("order")

	ctx, cancel := context.WithTimeout(context.Background(), configs.Constant.TimeoutOnSeconds*time.Second)
	defer cancel()

	_, err := collection.UpdateOne(
		ctx,
		bson.M{"id_order": id},
		bson.M{"$set": bson.M{key: value}},
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *orderRepo) GetBySellerID(sellerID uint32) (res []models.Order, err error) {
	collection := r.Instance.Collection("order")

	ctx, cancel := context.WithTimeout(context.Background(), configs.Constant.TimeoutOnSeconds*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{"id_seller": sellerID})
	if err != nil {
		return
	}

	if err = cursor.All(ctx, &res); err != nil {
		return
	}

	return
}

func (r *orderRepo) GetByBuyerID(buyerID uint32) (res []models.Order, err error) {
	collection := r.Instance.Collection("order")

	ctx, cancel := context.WithTimeout(context.Background(), configs.Constant.TimeoutOnSeconds*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{"id_buyer": buyerID})
	if err != nil {
		return
	}

	if err = cursor.All(ctx, &res); err != nil {
		return
	}

	return
}

func (r *orderRepo) GetByBuyerIDAndStatus(buyerID uint32, status string) (res []models.Order, err error) {
	collection := r.Instance.Collection("order")

	ctx, cancel := context.WithTimeout(context.Background(), configs.Constant.TimeoutOnSeconds*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{"id_buyer": buyerID, "status": status})
	if err != nil {
		return
	}

	if err = cursor.All(ctx, &res); err != nil {
		return
	}

	return
}

func (r *orderRepo) GetBySellerIDAndStatus(sellerID uint32, status string) (res []models.Order, err error) {
	collection := r.Instance.Collection("order")

	ctx, cancel := context.WithTimeout(context.Background(), configs.Constant.TimeoutOnSeconds*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{"id_seller": sellerID, "status": status})
	if err != nil {
		return
	}

	if err = cursor.All(ctx, &res); err != nil {
		return
	}

	return
}

func (r *orderRepo) GetByStatus(status string) (res []models.Order, err error) {
	collection := r.Instance.Collection("order")

	ctx, cancel := context.WithTimeout(context.Background(), configs.Constant.TimeoutOnSeconds*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{"status": status})
	if err != nil {
		return
	}

	if err = cursor.All(ctx, &res); err != nil {
		return
	}

	return
}
