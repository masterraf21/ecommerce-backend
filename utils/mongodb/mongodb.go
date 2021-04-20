package mongodb

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Option type
type Option struct {
	Hosts    string
	Database string
	Options  string
}

// NewWriteModelCollection function
func NewWriteModelCollection() []mongo.WriteModel {
	return []mongo.WriteModel{}
}

// NewUpdateModel function
func NewUpdateModel() *mongo.UpdateOneModel {
	return mongo.NewUpdateOneModel()
}

// BulkWrite function
func BulkWrite(collection *mongo.Collection, wmc []mongo.WriteModel, timeoutInSeconds time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeoutInSeconds*time.Second)
	defer cancel()

	_, err := collection.BulkWrite(ctx, wmc)
	if err != nil {
		return err
	}

	return nil
}

// Init function
func Init(option Option) (*mongo.Database, error) {
	if option.Hosts == "" {
		option.Hosts = "127.0.0.1:27017"
	}

	if option.Database == "" {
		return nil, fmt.Errorf("connecting to unknown database in MongoDB")
	}

	uri := "mongodb://" + option.Hosts + "/" + option.Database + "?" + option.Options

	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	if err := connectMongo(client); err != nil {
		return nil, err
	}

	if err := pingMongo(client, readpref.Primary()); err != nil {
		return nil, err
	}

	return client.Database(option.Database), nil
}

func connectMongo(c *mongo.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Second)
	defer cancel()

	return c.Connect(ctx)
}

func pingMongo(c *mongo.Client, rp *readpref.ReadPref) error {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Second)
	defer cancel()

	return c.Ping(ctx, rp)
}
