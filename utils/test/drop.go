package test

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

// DropCounter will drop counter table
func DropCounter(ctx context.Context, db *mongo.Database) (err error) {
	collection := db.Collection("counter")
	err = collection.Drop(ctx)
	return
}

// DropBuyer will drop buyer table
func DropBuyer(ctx context.Context, db *mongo.Database) (err error) {
	collection := db.Collection("buyer")
	err = collection.Drop(ctx)
	return
}

// DropSeller will drop seller table
func DropSeller(ctx context.Context, db *mongo.Database) (err error) {
	collection := db.Collection("seller")
	err = collection.Drop(ctx)
	return
}

// DropProduct will drop product table
func DropProduct(ctx context.Context, db *mongo.Database) (err error) {
	collection := db.Collection("product")
	err = collection.Drop(ctx)
	return
}

// DropOrder will drop order table
func DropOrder(ctx context.Context, db *mongo.Database) (err error) {
	collection := db.Collection("order")
	err = collection.Drop(ctx)
	return
}
