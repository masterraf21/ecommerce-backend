package mongodb

import (
	"context"
	"errors"
	"reflect"
	"time"

	"github.com/masterraf21/ecommerce-backend/configs"
	"github.com/masterraf21/ecommerce-backend/models"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

type counterRepo struct {
	Instance *mongo.Database
}

// NewCounterRepo will initiate new repo for counter
func NewCounterRepo(instance *mongo.Database) models.CounterRepository {
	return &counterRepo{instance}
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func dropCounter(ctx context.Context, db *mongo.Database) (err error) {
	collection := db.Collection("counter")
	err = collection.Drop(ctx)
	return
}

func initCollection(ctx context.Context, db *mongo.Database) error {
	counter := models.Counter{
		BuyerID:   0,
		ProductID: 0,
		SellerID:  0,
		OrderID:   0,
	}

	collection := db.Collection("counter")
	_, err := collection.InsertOne(ctx, counter)
	if err != nil {
		return err
	}

	return nil
}

func getLatestCounter(ctx context.Context, db *mongo.Database) (res *models.Counter, err error) {
	collection := db.Collection("counter")
	myOptions := options.FindOne()
	myOptions.SetSort(bson.M{"$natural": -1})

	err = collection.FindOne(ctx, bson.M{}, myOptions).Decode(&res)
	if err != nil {
		return
	}

	return
}

func incrementCounter(ctx context.Context, db *mongo.Database, identifier string, latestCounter models.Counter) error {
	collection := db.Collection("counter")

	counterMap := structToMap(latestCounter)
	latestID, ok := counterMap[identifier]

	if !ok {
		return errors.New("Identifier not found")
	}
	id := latestID.(uint32)

	counterMap[identifier] = id + 1

	dataBson := bson.M{}
	for k, v := range counterMap {
		dataBson[k] = v
	}

	_, err := collection.InsertOne(ctx, dataBson)
	if err != nil {
		return err
	}

	return nil
}

func structToMap(item interface{}) map[string]interface{} {
	res := map[string]interface{}{}
	if item == nil {
		return res
	}
	v := reflect.TypeOf(item)
	reflectValue := reflect.ValueOf(item)
	reflectValue = reflect.Indirect(reflectValue)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	for i := 0; i < v.NumField(); i++ {
		tag := v.Field(i).Tag.Get("json")
		field := reflectValue.Field(i).Interface()
		if tag != "" && tag != "-" {
			if v.Field(i).Type.Kind() == reflect.Struct {
				res[tag] = structToMap(field)
			} else {
				res[tag] = field
			}
		}
	}
	return res
}

func (r *counterRepo) Get(collectionName string, identifier string) (id uint32, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), configs.Constant.TimeoutOnSeconds*time.Second)
	defer cancel()

	names, err := r.Instance.ListCollectionNames(ctx, bson.M{})
	if err != nil {
		return
	}

	if !contains(names, "counter") {
		initCollection(ctx, r.Instance)
	}

	latest, err := getLatestCounter(ctx, r.Instance)
	if err != nil {
		return
	}
	latestMap := structToMap(latest)

	latestIDRaw, ok := latestMap[identifier]
	if !ok {
		err = errors.New("Identifier not found")
		return
	}

	latestID := latestIDRaw.(uint32)

	id = latestID + 1

	err = incrementCounter(ctx, r.Instance, identifier, *latest)
	if err != nil {
		return
	}

	return
}
