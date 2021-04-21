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

type productRepo struct {
	Instance    *mongo.Database
	CounterRepo models.CounterRepository
}

// NewProductRepo will initiate product repo
func NewProductRepo(instance *mongo.Database, ctr models.CounterRepository) models.ProductRepository {
	return &productRepo{Instance: instance, CounterRepo: ctr}
}

func (r *productRepo) Store(product *models.Product) (oid primitive.ObjectID, err error) {
	collectionName := "product"
	identifier := "id_product"

	ctx, cancel := context.WithTimeout(context.Background(), configs.Constant.TimeoutOnSeconds*time.Second)
	defer cancel()

	id, err := r.CounterRepo.Get(collectionName, identifier)
	if err != nil {
		return
	}

	collection := r.Instance.Collection(collectionName)
	product.ID = id

	result, err := collection.InsertOne(ctx, product)
	if err != nil {
		return
	}

	_id := result.InsertedID
	oid = _id.(primitive.ObjectID)

	return
}

func (r *productRepo) GetAll() (res []models.Product, error error) {
	collection := r.Instance.Collection("product")

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

func (r *productRepo) GetBySellerID(sellerID uint32) (res []models.Product, err error) {
	collection := r.Instance.Collection("product")

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

func (r *productRepo) GetByID(id uint32) (res *models.Product, err error) {
	collection := r.Instance.Collection("product")

	ctx, cancel := context.WithTimeout(context.Background(), configs.Constant.TimeoutOnSeconds*time.Second)
	defer cancel()

	err = collection.FindOne(ctx, bson.M{"id_product": id}).Decode(&res)
	if err != nil {
		return
	}

	return
}

func (r *productRepo) GetByOID(oid primitive.ObjectID) (res *models.Product, err error) {
	collection := r.Instance.Collection("product")

	ctx, cancel := context.WithTimeout(context.Background(), configs.Constant.TimeoutOnSeconds*time.Second)
	defer cancel()

	err = collection.FindOne(ctx, bson.M{"_id": oid}).Decode(&res)
	if err != nil {
		return
	}

	return
}

func (r *productRepo) UpdateArbitrary(id uint32, key string, value interface{}) error {
	collection := r.Instance.Collection("product")

	ctx, cancel := context.WithTimeout(context.Background(), configs.Constant.TimeoutOnSeconds*time.Second)
	defer cancel()

	_, err := collection.UpdateOne(
		ctx,
		bson.M{"id_product": id},
		bson.M{"$set": bson.M{key: value}},
	)
	if err != nil {
		return err
	}

	return nil
}
