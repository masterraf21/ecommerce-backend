package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Buyer represents buyer
type Buyer struct {
	ID              uint32 `bson:"id_buyer" json:"id_buyer"`
	Email           string `bson:"email" json:"email"`
	Name            string `bson:"name" json:"name"`
	Password        string `bson:"password" json:"password"`
	DeliveryAddress string `bson:"delivery_address" json:"delivery_address"`
}

// BuyerRepository represents repo functions for Buyer
type BuyerRepository interface {
	Store(buyer *Buyer) (primitive.ObjectID, error)
	GetByID(id uint32) (*Buyer, error)
	GetByOID(oid primitive.ObjectID) (*Buyer, error)
	UpdateArbitrary(id uint32, key string, value interface{}) error
}
