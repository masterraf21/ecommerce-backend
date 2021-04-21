package mongodb

import (
	"context"
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

func (r *counterRepo) Get(collectionName string, identifier string) (id uint32, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), configs.Constant.TimeoutOnSeconds*time.Second)
	defer cancel()

	names, err := r.Instance.ListCollectionNames(ctx, bson.M{})
	if err != nil {
		return
	}

	if !contains(names, collectionName) {
		id = uint32(1)
		return
	}

	collection := r.Instance.Collection(collectionName)

	var result bson.M
	myOptions := options.FindOne()
	myOptions.SetSort(bson.M{"$natural": -1})

	err = collection.FindOne(ctx, bson.M{}, myOptions).Decode(&result)
	if err != nil {
		return
	}

	lastID, ok := result[identifier].(int64)
	if !ok {
		panic(ok)
	}
	id = uint32(lastID) + 1

	return
}
